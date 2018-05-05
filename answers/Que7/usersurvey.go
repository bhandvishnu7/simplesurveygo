package servicehandlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"simplesurveygo/dao"
)

type UserSurveyHandler struct {
}

func (p UserSurveyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p UserSurveyHandler) Get(r *http.Request) SrvcRes {
	tokens, _ := r.Header["Token"]
	token := tokens[0]

	Decorator(authenticate)(token)

	user := dao.GetSessionDetails(token)
	userName := user.Username
	return Response200OK(dao.GetSurveysForUser(userName))
}

func (p UserSurveyHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p UserSurveyHandler) Post(r *http.Request) SrvcRes {
	body, err := ioutil.ReadAll(r.Body)

	var userResponse dao.SurveyResponse
	err = json.Unmarshal(body, &userResponse)

	dao.InsertUserResponse(userResponse)
	if err == nil {
		return Simple200OK("Feedback updated successfully")
	} else {
		return InternalServerError("Something went wrong")
	}
}
