package api

import (
	"context"
	"errors"
	"net/http"
)

type processResponse struct {
	http *http.Response
	err  error
}

type processRequest struct {
	client       *Client
	request      *request
	responseChan chan *processResponse
}

type processor struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	requests   chan *processRequest
}

func newProcessor(ctx context.Context) *processor {
	p := &processor{
		requests: make(chan *processRequest, 100),
	}
	p.ctx, p.cancelFunc = context.WithCancel(ctx)
	go p.process()
	return p
}

func (p *processor) do(req *processRequest) error {
	select {
	case p.requests <- req:
		return nil
	default:
		return errors.New("request queue full")
	}
}

func (p *processor) cancel() {
	p.cancelFunc()
}

func (p *processor) process() {
	defer p.cancelFunc()
	for {
		select {
		case req := <-p.requests:
			var httpRequest *http.Request
			var resp *processResponse

			resp = new(processResponse)

			// Try to build http httpRequest
			httpRequest, resp.err = req.request.toHTTP()

			// if no error during http request building...
			if resp.err == nil {

				// attempt to set auth cookie, if needed and found
				if req.request.authenticated {
					if req.client.sessionCookie != nil {
						httpRequest.AddCookie(req.client.sessionCookie)
					}
				}

				// new timeout context...
				ctx, cancel := context.WithTimeout(req.request.ctx, req.client.Config().RequestTimeout)
				httpRequest.WithContext(ctx)

				// attempt request execution
				resp.http, resp.err = req.client.client.Do(httpRequest)

				// always call cancel
				cancel()
			}

			// push response to responseChan
			req.responseChan <- resp

		case <-p.ctx.Done():
			return
		}
	}
}
