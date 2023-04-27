package harness

import (
	"fmt"
	"sync"

	"github.com/cheggaaa/pb/v3"
	resty "github.com/go-resty/resty/v2"
)

type HarnessAPI interface {
	GetAllConnectors() ([]Connectors, error)
}

type APIRequest struct {
	BaseURL string
	Client  *resty.Client
	APIKey  string
}

type EntityResult interface{}

func (a *APIRequest) GetAccountOverview(callCount int, callFuncs []func(string) (EntityResult, error), format string) (EntityResult, error) {
	type result struct {
		response EntityResult
		err      error
	}

	results := make(chan result)
	progressBar := pb.StartNew(callCount)

	var wg sync.WaitGroup
	wg.Add(callCount)

	for _, callFunc := range callFuncs {
		go func(callFunc func(string) (EntityResult, error)) {
			defer wg.Done()
			resp, err := callFunc(format)
			results <- result{response: resp, err: err}
			progressBar.Increment()
		}(callFunc)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	responses := make([]EntityResult, 0, callCount)
	for res := range results {
		if res.err != nil {
			progressBar.Finish()
			return nil, res.err
		}
		responses = append(responses, res)
	}

	progressBar.Finish()
	return responses, nil
}

func (api *APIRequest) GetAllConnectors(format string) (EntityResult, error) {
	resp, err := api.Client.R().Get(api.BaseURL + "/gateways")
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.String())

	return []Connectors{}, nil
}
