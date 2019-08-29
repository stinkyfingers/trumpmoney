package components

import (
	"errors"
	"net/http"
	"syscall/js"

	"github.com/stinkyfingers/gosx/attach"
	"github.com/stinkyfingers/gosx/element"
	"github.com/stinkyfingers/trumpmoney/api"
)

// APIResponse wraps the API's ScheduleAResponse and error
type APIResponse struct {
	scheduleAResponse *api.ScheduleAResponse
	err               error
}

const committeeID = "C00580100"

var (
	// ErrYearZip is the error for missing input data
	ErrYearZip = errors.New("year and zipcode are required")
)

// Submit is the submit button
func (a *appManager) Submit() {
	var zip, year, lastIndex, lastContributionReceiptDate string
	go func() {
		for s := range a.submitChan {
			switch s.dataType {
			case "year":
				year = s.data.(string)
			case "zip":
				zip = s.data.(string)
			}
		}
	}()

	cb := js.FuncOf(func(this js.Value, vals []js.Value) interface{} {
		a.fecCall(zip, year, lastIndex, lastContributionReceiptDate)
		return nil
	})

	button := element.NewElement("button", "Submit", nil, map[string]js.Func{"click": cb}, nil)
	attach.AttachElements([]element.Element{*button}, a.bindValue, nil)

	go func() {
		<-a.ctx.Done()
	}()
}

func (a *appManager) fecCall(zip, year, lastIndex, lastContributionReceiptDate string) {
	go func() {
		if zip == "" || year == "" {
			a.resultsChan <- semaphore{data: ErrYearZip, dataType: "error"}
			return
		}
		apiKey, err := api.GetAPIKey()
		if err != nil {
			a.resultsChan <- semaphore{data: err, dataType: "error"}
			return
		}

		c := &http.Client{}
		scheduleAResponse, err := api.GetContributions(c, committeeID, zip, year, lastIndex, lastContributionReceiptDate, apiKey)
		if err != nil {
			a.resultsChan <- semaphore{data: err, dataType: "error"}
			return
		}
		a.resultsChan <- semaphore{data: *scheduleAResponse, dataType: "fecResponse"}
	}()
}
