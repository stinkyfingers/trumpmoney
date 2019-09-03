package components

import (
	"fmt"

	"github.com/stinkyfingers/gosx/attach"
	"github.com/stinkyfingers/gosx/element"
)

func (a *appManager) Error(err error) *element.Element {
	errDiv := element.NewElement("div", fmt.Sprintf("ERROR: %s", err.Error()), map[string]string{"class": "error"}, nil, nil)
	attach.AttachElements([]element.Element{*errDiv}, a.bindValue, nil)
	return errDiv
}
