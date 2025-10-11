new Vue({
    el: '#app',
    data() {
        return {
            showPopup: false,
            progress: 0,
            message: '',
            statusText: '',
            sse: null
        }
    },
    methods: {
        async checkout() {
            try {
                const res = await fetch('/api/order', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ user_id: 1001 })
                });
                const result = await res.json();
                if (!res.ok) throw new Error(result.message || '下单失败');

                this.showPopup = true;
                this.progress = 0;
                this.message = '正在下单...';
                this.statusText = '开始跟踪进度...';
                this.sse = new EventSource(`/api/trx/${result.trxid}`);
                this.sse.addEventListener('data', event => {
                    const data = JSON.parse(event.data) || {};
                    this.progress = data.progress || 0;
                    this.message = data.state || this.message;
                    this.statusText = data.status || this.statusText;
                    if (this.progress >= 100) {
                        this.statusText = '下单成功，跳转支付...';
                        this.sse.close();
                    }
                });
            } catch (err) {
                this.$toast(`下单失败：${err.message}`);
            }
        },
        closePopup() {
            this.showPopup = false;
            if (this.sse) this.sse.close();
        }
    }
})