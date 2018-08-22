var vm = new Vue({
    el: "#app",
    template: "<div><h5>hello</h5></div>",
    data: { name: "WebView" },
    created: function () {
        alert("hi from vue, " + this.name);
    }
});
