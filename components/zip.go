package components

import (
	"syscall/js"

	"github.com/stinkyfingers/gosx/attach"
	"github.com/stinkyfingers/gosx/element"
)

// ZipInput is the text box for zip code
func (a *appManager) ZipInput() {
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		a.zip = this.Get("value").String()
		a.resultsChan <- semaphore{data: struct{}{}, dataType: "remove"}
		return nil
	})
	label := element.NewElement("label", "Zip Code", map[string]string{"class": "zipLabel"}, nil, nil)
	zip := element.NewElement("input", "", map[string]string{"class": "zip"}, map[string]js.Func{"change": cb}, label)
	attach.AttachElements([]element.Element{*label, *zip}, a.bindValue, nil)
	go func() {
		<-a.ctx.Done()
		cb.Release()
	}()
}
