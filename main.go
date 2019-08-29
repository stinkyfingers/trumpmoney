package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/stinkyfingers/trumpmoney/components"
)

//https://github.com/siongui/godom/tree/master/wasm
func main() {

	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGTERM|syscall.SIGQUIT|syscall.SIGINT)

	cancel := components.App()
	<-ch
	cancel()
}
