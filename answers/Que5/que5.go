/* servicehandlers/concurrancyHandler.go */

/*
var concurrency = make(chan bool, 10)
*/


// servicehandlers/ping.go

/*
func (p PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	concurrency <- true

	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)

	<-concurrency
}
*/

/* 
	Insert true value in chan of bool type to each api,
	this will handles our requirement of not allowing more than 10 conncurrent request to our simplesurvey service
*/
