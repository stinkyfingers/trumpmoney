package main

import (
	"context"
	"log"
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

	components.ZipInput(ctx, body, zipChan)
	components.YearSelect(ctx, body, yearChan)
	components.Submit(ctx, body, yearChan, zipChan, apiChan)
	components.ResultsList(ctx, body, apiChan)
	<-ch
	log.Print("PRE DONE")

	cf()
	log.Print("DONE")
}
