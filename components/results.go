package components

import (
	"context"
	"fmt"
	"syscall/js"

	"github.com/stinkyfingers/gosx/attach"
	"github.com/stinkyfingers/gosx/element"
)

func ResultsList(ctx context.Context, body js.Value, apiChan chan APIResponse) {
	for apiResp := range apiChan {
		renderResults(body, apiResp)
		// TODO - remove previous results
	}
}
func renderResults(body js.Value, apiResp APIResponse) {
	if apiResp.err != nil {
		attach.AttachElements([]element.Element{*element.NewElement("div", fmt.Sprintf("ERROR: %s", apiResp.err.Error()), nil, nil, nil)}, body, nil)
	}
	if apiResp.scheduleAResponse == nil {
		return
	}
	table := element.NewElement("table", "", nil, nil, nil)
	tbody := element.NewElement("tbody", "", nil, nil, table)
	thead := element.NewElement("thead", "", nil, nil, table)
	thr := element.NewElement("tr", "", nil, nil, thead)
	nameHead := element.NewElement("th", "Name", nil, nil, thr)
	emplHead := element.NewElement("th", "Employer", nil, nil, thr)

	tableItems := []element.Element{*table, *tbody, *thead, *thr, *nameHead, *emplHead}
	for _, res := range apiResp.scheduleAResponse.Results {
		tr := element.NewElement("tr", "", nil, nil, tbody)
		name := element.NewElement("td", res.ContributorName, nil, nil, tr)
		empl := element.NewElement("td", res.ContributorEmployer, nil, nil, tr)
		tableItems = append(tableItems, *tr, *name, *empl)
	}
	attach.AttachElements(tableItems, body, nil)
}
