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

var (
	// ErrYearZip is the error for missing input data
	ErrYearZip = errors.New("year and zipcode are required")
)

// Submit is the submit button
func (a *appManager) Submit() {
	cb := js.FuncOf(func(this js.Value, vals []js.Value) interface{} {
		a.fecCallAll()
		return nil
	})

	button := element.NewElement("button", "Submit", map[string]string{"class": "submit"}, map[string]js.Func{"click": cb}, nil)
	attach.AttachElements([]element.Element{*button}, a.bindValue, nil)

	go func() {
		<-a.ctx.Done()
	}()
}

func (a *appManager) fecCall() {
	go func() {
		if a.zip == "" || a.year == "" {
			a.resultsChan <- semaphore{data: ErrYearZip, dataType: "error"}
			return
		}
		apiKey, err := api.GetAPIKey()
		if err != nil {
			a.resultsChan <- semaphore{data: err, dataType: "error"}
			return
		}

		c := &http.Client{}
		scheduleAResponse, err := api.GetContributions(c, a.zip, a.year, a.lastIndex, a.lastContributionReceiptDate, apiKey)
		if err != nil {
			a.resultsChan <- semaphore{data: err, dataType: "error"}
			return
		}
		a.resultsChan <- semaphore{data: *scheduleAResponse, dataType: "fecResponse"}
	}()
}

func (a *appManager) fecCallAll() {
	go func() {
		if a.zip == "" || a.year == "" {
			a.resultsChan <- semaphore{data: ErrYearZip, dataType: "error"}
			return
		}
		apiKey, err := api.GetAPIKey()
		if err != nil {
			a.resultsChan <- semaphore{data: err, dataType: "error"}
			return
		}

		c := &http.Client{}
		results, err := api.GetContributionsPaged(c, a.zip, a.year, apiKey)
		if len(results) == 0 {
			a.resultsChan <- semaphore{data: api.ErrEOR, dataType: "error"}
			return
		}
		if err != nil {
			a.resultsChan <- semaphore{data: err, dataType: "error"}
			return
		}
		a.resultsChan <- semaphore{data: results, dataType: "fecResponse"}
	}()
}
