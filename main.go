package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"syscall/js"

	"github.com/stinkyfingers/trumpmoney/components"
)

//https://github.com/siongui/godom/tree/master/wasm
func main() {

	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGTERM|syscall.SIGQUIT|syscall.SIGINT)

	body := js.Global().Get("document").Get("body")
	ctx, cf := context.WithCancel(context.Background())

	zipChan := make(chan string)
	yearChan := make(chan string)
	apiChan := make(chan components.APIResponse)
	removeChan := make(chan bool)

	components.ZipInput(ctx, body, zipChan, removeChan)
	components.YearSelect(ctx, body, yearChan, removeChan)
	components.Submit(ctx, body, yearChan, zipChan, apiChan)
	components.ResultsList(ctx, body, apiChan, removeChan)

	<-ch

	cf()
}
