package components

import (
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

// YearSelect is the select dropdown for year
func (a *appManager) YearSelect() {
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		a.year = this.Get("value").String()
		a.resultsChan <- semaphore{data: struct{}{}, dataType: "remove"}
		return nil
	})
	label := element.NewElement("label", "Year", map[string]string{"class": "yearLabel"}, nil, nil)
	sel := element.NewElement("select", "", map[string]string{"class": "year"}, map[string]js.Func{"change": cb}, label)
	elements := []element.Element{*label, *sel}
	for _, option := range yearOptions {
		elements = append(elements, *element.NewElement("option", option, map[string]string{"value": option}, nil, sel))
	}
	attach.AttachElements(elements, a.bindValue, nil)
	go func() {
		<-a.ctx.Done()
		cb.Release()
	}()
}
