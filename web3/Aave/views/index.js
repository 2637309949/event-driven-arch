const { contractAddress, contractABI } = window.contractConfig;
new Vue({
    el: '#app',
    data: {
        totalSupply: 0,
        maxSupply: 1000,
        metaVisible: false,
        metaData: null,
        activeTab: 'home',
        mintQuantity: 1,
        mintLoading: false,
        popupMintNFTVisible: false,
        stats: [
            { label: '交易总数', value: 0 },
            { label: '活跃用户', value: 0 },
            { label: 'NFT 铸造', value: 0 },
            { label: '转账次数', value: 0 }
        ],
        account: null,
        balance: null,
        web3: null,
        contract: null,
        es: null,
        nftList: [],
        occurred: [],
    },
    computed: {
        progressPercent() {
            if (!this.maxSupply) return 0
            const percent = (this.totalSupply / this.maxSupply) * 100
            return percent.toFixed(2)
        }
    },
    watch: {
        activeTab(newVal) {
            switch (newVal) {
                case 'home':
                    this.$nextTick(() => {
                        this.stats.forEach(stat => {
                            this.runCountUp(stat.name, stat.value - 1, stat.value)
                        });
                        this.runCountUp('totalSupply', this.totalSupply - 1, this.totalSupply);
                        this.runCountUp('maxSupply', this.maxSupply - 1, this.maxSupply);
                    })
                    break;
                case 'me':
                    if (!this.nftList.length) {
                        this.loadMyNFTs();
                    }
                    break;
            }
        },
        totalSupply(newVal) {
            this.$nextTick(() => {
                this.runCountUp('totalSupply', newVal - 1, newVal);
            })
        },
        maxSupply(newVal) {
            this.$nextTick(() => {
                this.runCountUp('maxSupply', newVal - 1, newVal);
            })
        },
        stats: {
            deep: true,
            handler(newStats, oldStats) {
                let occurredName = this.occurred.pop()
                this.$nextTick(() => {
                    newStats.forEach((stat, index) => {
                        if (occurredName === undefined || occurredName === stat.name) {
                            this.runCountUp(stat.name, oldStats[index] ? oldStats[index].value : 0, stat.value);
                        }
                    });
                });
            }
        }
    },
    beforeDestroy() {
        if (this.es) {
            this.es.close();
            this.es = null;
        }
    },
    async mounted() {
        this.flushStat();
        if (!window.ethereum) {
            this.$toast('未检测到 MetaMask，请先安装');
            return
        }
        this.web3 = new Web3(window.ethereum);
        let walletConnected = localStorage.getItem('wallet_connected')
        if (walletConnected === 'true') {
            const accounts = await window.ethereum.request({ method: 'eth_accounts' });
            if (accounts.length) {
                this.setAccount(accounts[0]);
            }
        }
        window.ethereum.on('accountsChanged', (accounts) => {
            if (accounts.length) {
                this.setAccount(accounts[0]);
            } else {
                this.account = null;
                this.balance = null;
            }
        });
        window.ethereum.on('chainChanged', () => {
            window.location.reload();
        });

        this.contract = new this.web3.eth.Contract(contractABI, contractAddress);
        this.refreshProgressPercent();
    },
    methods: {
        async mintNFT(quantity) {
            try {
                this.mintLoading = true;
                const mintPrice = await this.contract.methods.mintPrice().call();
                const totalPrice = BigInt(mintPrice) * BigInt(quantity); // 总价
                const tx = await this.contract.methods.publicMint(quantity).send({
                    from: this.account,
                    value: totalPrice.toString()  // wei 单位
                });
                this.$toast('铸造成功');
            } catch (err) {
                console.error('publicMint err', err);
                this.$toast(err.message.includes('insufficient funds') ? '余额不足' : '铸造失败');
            } finally {
                this.mintLoading = false;
            }
        },
        flushStat() {
            if (this.es) {
                this.es.close();
                this.es = null;
            }
            this.es = new EventSource('/api/sse/occurred')
            this.es.addEventListener('data', event => {
                let data = JSON.parse(event.data);
                switch (data.type) {
                    case "stats":
                        this.stats = data.data;
                        break;
                    case "occurred":
                        let occurred = data.data;
                        for (let i = 0; i < this.stats.length; i++) {
                            let s = this.stats[i];
                            if (s.name === occurred.name) {
                                this.stats[i].value += occurred.value;
                                this.occurred.push(occurred.name);
                            }
                            if (occurred.name == 'Minted') {
                                this.totalSupply += 1
                            }
                        }
                        break;
                }
            });
        },
        loadTop() {
            this.flushStat();
            this.refreshProgressPercent();
            setTimeout(() => {
                this.$refs.loadmore.onTopLoaded()
            }, 700)
        },
        meLoadTop() {
            setTimeout(() => {
                this.$refs.meLoadmore.onTopLoaded()
            }, 700)
        },
        runCountUp(name, start, end, cb = () => { }) {
            const options = {
                startVal: 0,
                onCompleteCallback: cb,
                duration: 1,
            };
            let cu = new countUp.CountUp('stat-' + name, end, options);
            if (!cu.error) {
                cu.start();
            } else {
                console.error(cu.error);
            }
        },
        async connectWallet() {
            try {
                const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
                this.setAccount(accounts[0]);
                this.loadMyNFTs();
                localStorage.setItem('wallet_connected', 'true');
                this.$toast('钱包连接成功！');
            } catch (err) {
                console.error(err);
                this.$toast('连接钱包失败');
            }
        },
        disconnectWallet() {
            this.account = null;
            this.balance = '0';
            this.nftList = [];
            localStorage.removeItem('wallet_connected');
            this.$toast('已断开钱包');
        },
        async setAccount(account) {
            this.account = account;
            const balanceWei = await this.web3.eth.getBalance(account);
            this.balance = this.web3.utils.fromWei(balanceWei, 'ether');
        },
        async showMeta(tokenId) {
            try {
                if (!this.contract || !this.account) return;
                const balance = await this.contract.methods.balanceOf(this.account).call();
                let tokenURI = await this.contract.methods.tokenURI(tokenId).call();
                if (!tokenURI.endsWith('.json')) tokenURI += '.json';
                if (tokenURI.startsWith("ipfs://")) {
                    tokenURI = tokenURI.replace("ipfs://", "https://ipfs.io/ipfs/");
                }
                const res = await fetch(tokenURI)
                const meta = await res.json()

                // image URL
                let imageURL = "";
                if (meta.image) {
                    imageURL = meta.image.startsWith("ipfs://")
                        ? meta.image.replace("ipfs://", "https://ipfs.io/ipfs/")
                        : meta.image;
                }
                meta.imageURL = imageURL

                this.metaData = meta
                this.metaVisible = true
            } catch (err) {
                this.$toast(`加载 Metadata 失败: ${err}`)
            }
        },
        async loadMyNFTs() {
            if (!this.contract || !this.account) return;
            const balance = await this.contract.methods.balanceOf(this.account).call();
            const list = [];
            for (let i = 0; i < balance; i++) {
                const tokenId = await this.contract.methods.tokenOfOwnerByIndex(this.account, i).call();
                let tokenURI = await this.contract.methods.tokenURI(tokenId).call();
                if (!tokenURI.endsWith('.json')) tokenURI += '.json';
                if (tokenURI.startsWith("ipfs://")) {
                    tokenURI = tokenURI.replace("ipfs://", "https://ipfs.io/ipfs/");
                }
                // fetch JSON metadata
                let metadata = {};
                try {
                    const res = await fetch(tokenURI);
                    metadata = await res.json();
                } catch (e) {
                    console.error("获取 NFT metadata 失败", e);
                }
                // image URL
                let imageURL = "";
                if (metadata.image) {
                    imageURL = metadata.image.startsWith("ipfs://")
                        ? metadata.image.replace("ipfs://", "https://ipfs.io/ipfs/")
                        : metadata.image;
                }
                // 倒序显示最近铸造的
                // this.nftList.unshift({ tokenId, tokenURI, name: metadata.name, imageURL });
                this.nftList.push({ tokenId, tokenURI, name: metadata.name, imageURL });
            }
        },
        async refreshProgressPercent() {
            try {
                if (!this.contract || !this.account) return;
                const [supply, max] = await Promise.all([
                    this.contract.methods.totalSupply().call(),
                    this.contract.methods.maxSupply().call()
                ]);
                this.totalSupply = 0;
                this.$nextTick(() => this.totalSupply = parseInt(supply));
                this.maxSupply = parseInt(max);
            } catch (err) {
                console.error("读取合约数据失败:", err);
            }
        }
    }
});