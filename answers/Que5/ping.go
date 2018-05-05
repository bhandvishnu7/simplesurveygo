package servicehandlers

import (
	"net/http"
)

type PingHandler struct {
}

func (p PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	concurrency <- true
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
	<-concurrency
}

func (p PingHandler) Get(r *http.Request) SrvcRes {
	return Response200OK("All Good")
}

func (p PingHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p PingHandler) Post(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}
