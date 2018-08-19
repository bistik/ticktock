package main

import (
	"fmt"
	"github.com/zserge/webview"
	// "html/template"
	"net/url"
)

func main() {
	Test()
	w := webview.New(webview.Settings{
		Title: "ticktock: ",
		URL:   `data:text/html,` + url.PathEscape(string(MustAsset("asset/index.html"))),
	})
	defer w.Exit()

	css := string(MustAsset("asset/styles.css"))
	fmt.Println(css)
	w.Dispatch(func() {
		//w.InjectCSS(css)
		/*
			w.Eval(fmt.Sprintf(`(function(css){
					var style = document.createElement('style')
					var head = document.head || document.getElementsByTagName('head')[0];
					style.setAttribute('type', 'text/css');
					if (style.styleSheet) {
						style.styleSheet.cssText = css;
					} else {
						style.appendChild(document.createTextNode(css))
					}
					alert("hello css");
				})("%s")`, template.JSEscapeString(css)))
		*/
	})
	w.Run()
}
