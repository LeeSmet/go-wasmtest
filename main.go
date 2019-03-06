// +build js

package main

import (
	"fmt"

	"syscall/js"
)

var wallets = []string{"wallet1", "wallet2", "wallet3"}

func main() {
	fmt.Println("Hello wasm")

	window := js.Global()
	if window == js.Null() {
		fmt.Println("Failed to get a handle on the window")
		return
	}

	document := window.Get("document")
	if document == js.Null() {
		fmt.Println("Failed to get a handle on the document")
		return
	}

	title := document.Call("getElementById", "title")
	if title == js.Null() {
		fmt.Println("Failed to get a reference to the title")
		return
	}

	title.Set("innerText", "Hello from go wasm")

	body := document.Get("body")
	if body == js.Null() {
		fmt.Println("Failed to get body")
	}

	currentBody := body.Get("innerHTML").String()

	walletList := "<ul>\n"
	for _, walletName := range wallets {
		walletList += fmt.Sprintf("<li>%v</li>\n", walletName)
	}
	walletList += "</ul>\n"

	currentBody += walletList
	body.Set("innerHTML", currentBody)
}
