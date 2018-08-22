package main

import (
	"fmt"
	"github.com/zserge/webview"
)

func loadVue(w webview.WebView) {
	w.Eval(string(MustAsset("asset/vue.min.js")))
	w.Eval(string(MustAsset("asset/app.js")))
	w.Eval(string(MustAsset("asset/test.js")))
	//w.Eval(`alert("hello")`)
	fmt.Println("load vue")
}
