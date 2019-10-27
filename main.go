package main

import (
	"os"
	"os/signal"
	"syscall"

	"syscall/js"

	"github.com/stinkyfingers/trumpmoney/components"
)

//https://github.com/siongui/godom/tree/master/wasm
func main() {
	loc := js.Global().Get("location").Get("href").String()
	console := js.Global().Get("console")
	console.Call("log", "Running application at "+loc)

	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGTERM|syscall.SIGQUIT|syscall.SIGINT)

	cancel := components.App()
	<-ch
	cancel()
}
