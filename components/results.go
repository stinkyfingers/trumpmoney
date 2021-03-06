package components

import (
	"github.com/stinkyfingers/gosx/attach"
	"github.com/stinkyfingers/gosx/element"
	"github.com/stinkyfingers/trumpmoney/api"
)

// ResultsList is the table of results
func (a *appManager) ResultsList() {
	table := element.NewElement("table", "", map[string]string{"class": "resultsTable"}, nil, nil)
	thead := element.NewElement("thead", "", nil, nil, table)
	thr := element.NewElement("tr", "", nil, nil, thead)
	nameHead := element.NewElement("th", "Name", nil, nil, thr)
	emplHead := element.NewElement("th", "Employer", nil, nil, thr)
	occupationHead := element.NewElement("th", "Occupation", nil, nil, thr)
	tableItems := []element.Element{*table, *thead, *thr, *nameHead, *emplHead, *occupationHead}
	var loading *element.Element
	attach.AttachElements(tableItems, a.bindValue, nil)

	var tbody *element.Element
	var errDiv *element.Element

	go func() {
		for s := range a.resultsChan {
			switch s.dataType {
			case "loading":
				if s.data.(bool) {
					loading = element.NewElement("div", "Loading...", map[string]string{"class": "loading"}, nil, nil)
					attach.AttachElements([]element.Element{*loading}, a.bindValue, nil)
				} else if !loading.Null() {
					attach.Remove(*loading)
				}
			case "error":
				errDiv = a.Error(s.data.(error))

			case "fecResponse":
				if tbody == nil || tbody.Null() {
					tbody = element.NewElement("tbody", "", nil, nil, table)
					attach.AttachElements([]element.Element{*tbody}, a.bindValue, nil)
				}
				scheduleAResponse := s.data.([]api.Result)

				elements := renderResults(tbody, scheduleAResponse)
				attach.AttachElements(elements, a.bindValue, nil)

			case "remove":
				if errDiv != nil && !errDiv.Null() {
					attach.Remove(*errDiv)
				}
				if tbody != nil && !tbody.Null() {
					attach.Remove(*tbody)
				}
			}
		}
	}()

}

func renderResults(tbody *element.Element, apiResp []api.Result) []element.Element {
	var elements []element.Element
	for _, res := range apiResp {
		tr := element.NewElement("tr", "", nil, nil, tbody)
		name := element.NewElement("td", res.ContributorName, nil, nil, tr)
		empl := element.NewElement("td", res.ContributorEmployer, nil, nil, tr)
		occupation := element.NewElement("td", res.ContributorOccupation, nil, nil, tr)
		elements = append(elements, *tr, *name, *empl, *occupation)
	}
	return elements
}
