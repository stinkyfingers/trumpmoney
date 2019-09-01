package components

import (
	"context"
	"syscall/js"
)

type appManager struct {
	ctx                         context.Context
	resultsChan                 chan semaphore
	bindValue                   js.Value
	year                        string
	zip                         string
	lastIndex                   string
	lastContributionReceiptDate string
}

type semaphore struct {
	data     interface{}
	dataType string
}

// App is the wrapper around the app's components
func App() context.CancelFunc {
	body := js.Global().Get("document").Get("body")
	ctx, cf := context.WithCancel(context.Background())

	a := &appManager{
		ctx:         ctx,
		resultsChan: make(chan semaphore),
		bindValue:   body,
	}

	a.Header()
	a.ZipInput()
	a.YearSelect()
	a.Submit()
	a.ResultsList()

	return cf
}
