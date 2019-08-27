package components

import (
	"context"
	"syscall/js"

	"github.com/stinkyfingers/gosx/attach"
	"github.com/stinkyfingers/gosx/element"
)

func ZipInput(ctx context.Context, body js.Value, zipChan chan string) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		zipChan <- this.Get("value").String()
		return nil
	})
	label := element.NewElement("label", "Zip Code", nil, nil, nil)
	zip := element.NewElement("input", "", nil, map[string]js.Func{"change": cb}, label)
	attach.AttachElements([]element.Element{*label, *zip}, body, nil)
	go func() {
		<-ctx.Done()
		cb.Release()
	}()
}
