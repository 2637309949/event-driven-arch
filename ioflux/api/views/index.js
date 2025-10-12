new Vue({
    el: '#app',
    data() {
        return {
            progress: 0,
            eventSource: null,
            showPopup: false,
            message: '',
            statusText: '',
        };
    },
    methods: {
        triggerFileSelect() {
            this.$refs.fileInput.click();
        },
        handleFileChange(e) {
            const file = e.target.files[0];
            if (!file) return;
            this.uploadFile(file);
        },
        async uploadFile(file) {
            try {
                const form = new FormData();
                form.append("file", file);
                this.progress = 0;
                this.message = '正在上传...';
                this.statusText = '开始跟踪进度...';
                this.showPopup = true;
                const res = await fetch("/api/upload", {
                    method: "POST",
                    body: form
                });
                const result = await res.json();
                this.eventSource = new EventSource(`/api/trx/${result.trxid}`);
                this.eventSource.addEventListener('data', e => {
                    const data = JSON.parse(e.data) || {};
                    this.progress = data.progress || 0;
                    this.message = data.state || this.message;
                    this.statusText = data.status || this.statusText;
                    if (this.progress >= 100) {
                        this.statusText = '下单成功，跳转支付...';
                        this.eventSource.close();
                    }
                });
            } catch (err) {
                console.error("上传失败", err);
                this.showPopup = false;
                this.$toast(`上传失败：${err.message}`);
            }
        },
        closePopup() {
            this.showPopup = false;
            if (this.eventSource) this.eventSource.close();
        }
    }
});