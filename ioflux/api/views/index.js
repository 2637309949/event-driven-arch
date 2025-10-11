new Vue({
    el: '#app',
    data() {
        return {
            file: null,
            progress: 0,
            result: null,
            eventSource: null,
            showPopup: false
        };
    },
    methods: {
        triggerFileSelect() {
            this.$refs.fileInput.click();
        },
        handleFileChange(e) {
            const selectedFile = e.target.files[0];
            if (!selectedFile) return;
            this.file = selectedFile;

            // 文件选择后直接上传
            this.uploadFile();
        },
        async uploadFile() {
            const formData = new FormData();
            formData.append("file", this.file);

            this.progress = 0;
            this.result = null;
            this.showPopup = true;

            try {
                // POST 上传文件
                const response = await fetch("/api/upload", {
                    method: "POST",
                    body: formData
                });
                const data = await response.json();
                console.log("上传成功，开始监听进度：", data);

                // SSE 监听后台处理进度
                await this.listenSSE(data.taskId);
            } catch (err) {
                console.error("上传失败", err);
                this.showPopup = false;
                alert("上传失败");
            }
        },
        listenSSE(taskId) {
            return new Promise((resolve, reject) => {
                this.eventSource = new EventSource(`/api/progress?taskId=${taskId}`);

                this.eventSource.onmessage = e => {
                    const data = JSON.parse(e.data);

                    if (data.progress !== undefined) this.progress = data.progress;

                    if (data.done) {
                        this.progress = 100;
                        this.result = data.result;
                        this.eventSource.close();
                        resolve(data.result);
                    }
                };

                this.eventSource.onerror = e => {
                    console.error("SSE 错误", e);
                    this.eventSource.close();
                    reject(e);
                };
            });
        }
    }
});