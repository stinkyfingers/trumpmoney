package components

import (
	"github.com/stinkyfingers/gosx/attach"
	"github.com/stinkyfingers/gosx/element"
)

// Header is the h2 heaader
func (a *appManager) Header() {
	title := "From whom are Trump's campaign contributions coming?"
	head := element.NewElement("h2", title, map[string]string{"class": "header"}, nil, nil)
	attach.AttachElements([]element.Element{*head}, a.bindValue, nil)
}
