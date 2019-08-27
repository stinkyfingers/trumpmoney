package components

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"syscall/js"

	"github.com/stinkyfingers/gosx/attach"
	"github.com/stinkyfingers/gosx/element"
	"github.com/stinkyfingers/trumpmoney/fec"
)

type APIResponse struct {
	scheduleAResponse *fec.ScheduleAResponse
	err               error
}

const committeeID = "C00580100"
const apiKey = ""

func Submit(ctx context.Context, body js.Value, yearChan, zipChan chan string, apiChan chan APIResponse) {
	var zip, year, lastIndex, lastContributionReceiptDate string
	go func() {
		for {
			select {
			case year = <-yearChan:
			case zip = <-zipChan:
			case <-ctx.Done():
				return
			}
		}
	}()

	cb := js.FuncOf(func(this js.Value, vals []js.Value) interface{} {
		fecCall(zip, year, lastIndex, lastContributionReceiptDate, apiKey, apiChan) // TODO handle error
		return nil
	})

	button := element.NewElement("button", "Submit", nil, map[string]js.Func{"click": cb}, nil)
	attach.AttachElements([]element.Element{*button}, body, nil)

	go func() {
		<-ctx.Done()
		// cb.Release()
	}()
}

// TODO - wrap better?
func fecCall(zip, year, lastIndex, lastContributionReceiptDate, apiKey string, apiChan chan APIResponse) {
	go func() {
		apiKey, err := getAPIKey()
		if err != nil {
			log.Fatal(err) // TODO
		}

		c := &http.Client{}
		scheduleAResponse, err := fec.GetContributions(c, committeeID, zip, lastIndex, lastContributionReceiptDate, apiKey)
		apiChan <- APIResponse{
			scheduleAResponse: scheduleAResponse,
			err:               err,
		}
	}()
}

func getAPIKey() (string, error) {
	// location := js.Global().Get("location").String()
	// log.Print("--", location)
	// url := fmt.Sprintf("%s/%s", strings.TrimRight(location, "/"), "apikey")
	url := "http://trumpmoney.john-shenk.com/fioobar"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes.Trim(b, "\n\t ")), nil
}
