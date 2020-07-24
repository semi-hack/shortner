const app = new Vue({
    el: '#app',
    data: {
        url: ''
    },
    methods: {
        createUrl() {
            console.log(this.url);
        }
    }
})