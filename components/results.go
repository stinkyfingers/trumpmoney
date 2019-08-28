package components

import (
	"context"
	"syscall/js"

	"github.com/stinkyfingers/gosx/attach"
	"github.com/stinkyfingers/gosx/element"
)

func ResultsList(ctx context.Context, body js.Value, apiChan chan APIResponse, removeChan chan bool) {
	table := element.NewElement("table", "", nil, nil, nil)
	thead := element.NewElement("thead", "", nil, nil, table)
	thr := element.NewElement("tr", "", nil, nil, thead)
	nameHead := element.NewElement("th", "Name", nil, nil, thr)
	emplHead := element.NewElement("th", "Employer", nil, nil, thr)
	tableItems := []element.Element{*table, *thead, *thr, *nameHead, *emplHead}
	attach.AttachElements(tableItems, body, nil)

	var tbody *element.Element

	go func() {
		for apiResp := range apiChan {
			if tbody == nil || tbody.Null() {
				tbody = element.NewElement("tbody", "", nil, nil, table)
				attach.AttachElements([]element.Element{*tbody}, body, nil)
			}
			elements := renderResults(tbody, apiResp)
			attach.AttachElements(elements, body, nil)
		}
	}()

	go func() {
		for range removeChan {
			if tbody == nil || tbody.Null() {
				continue
			}
			attach.Remove(*tbody)
		}
	}()
}

func renderResults(tbody *element.Element, apiResp APIResponse) []element.Element {
	if apiResp.err != nil {
		// TODO
		// attach.AttachElements([]element.Element{*element.NewElement("div", fmt.Sprintf("ERROR: %s", apiResp.err.Error()), nil, nil, nil)}, body, nil)
	}
	if apiResp.scheduleAResponse == nil {
		return nil
	}
	var elements []element.Element
	for _, res := range apiResp.scheduleAResponse.Results {
		tr := element.NewElement("tr", "", nil, nil, tbody)
		name := element.NewElement("td", res.ContributorName, nil, nil, tr)
		empl := element.NewElement("td", res.ContributorEmployer, nil, nil, tr)
		elements = append(elements, *tr, *name, *empl)
	}
	return elements
}

// func renderResults(body js.Value, apiResp APIResponse) {
// 	if apiResp.err != nil {
// 		attach.AttachElements([]element.Element{*element.NewElement("div", fmt.Sprintf("ERROR: %s", apiResp.err.Error()), nil, nil, nil)}, body, nil)
// 	}
// 	if apiResp.scheduleAResponse == nil {
// 		return
// 	}
// 	// table := element.NewElement("table", "", nil, nil, nil)
// 	// tbody := element.NewElement("tbody", "", nil, nil, table)
// 	// thead := element.NewElement("thead", "", nil, nil, table)
// 	// thr := element.NewElement("tr", "", nil, nil, thead)
// 	// nameHead := element.NewElement("th", "Name", nil, nil, thr)
// 	// emplHead := element.NewElement("th", "Employer", nil, nil, thr)
// 	//
// 	// tableItems := []element.Element{*table, *tbody, *thead, *thr, *nameHead, *emplHead}
// 	for _, res := range apiResp.scheduleAResponse.Results {
// 		tr := element.NewElement("tr", "", nil, nil, tbody)
// 		name := element.NewElement("td", res.ContributorName, nil, nil, tr)
// 		empl := element.NewElement("td", res.ContributorEmployer, nil, nil, tr)
// 		tableItems = append(tableItems, *tr, *name, *empl)
// 	}
// 	attach.AttachElements(tableItems, body, nil)
// }
