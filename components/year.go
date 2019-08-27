package components

import (
	"context"
	"syscall/js"

	"github.com/stinkyfingers/gosx/attach"
	"github.com/stinkyfingers/gosx/element"
)

type year struct {
	value     string
	elementID string
	cb        js.Func
}

var yearOptions = []string{"", "2020", "2019", "2018", "2017", "2016"}

func YearSelect(ctx context.Context, body js.Value, yearChan chan string) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		yearChan <- this.Get("value").String()
		return nil
	})

	sel := element.NewElement("select", "", nil, map[string]js.Func{"change": cb}, nil)
	elements := []element.Element{*sel}
	for _, option := range yearOptions {
		elements = append(elements, *element.NewElement("option", option, map[string]string{"value": option}, nil, sel))
	}

	attach.AttachElements(elements, body, nil)
	go func() {
		<-ctx.Done()
		cb.Release()
	}()
}
