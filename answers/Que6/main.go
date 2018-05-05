package main

import (
	"fmt"
	"net/http"
	"simplesurveygo/dao"
	"simplesurveygo/servicehandlers"
)

func DeactivateExpiredSurvey() {
	for {
		expired_surveys := dao.GetAllExpiredSurveys()
		for _, survey := range expired_surveys {
			result := dao.UpdateStatusToFalse(survey)
			fmt.Println(result)
		}
	}
}

func main() {

	// Serves the html pages
	http.Handle("/", http.FileServer(http.Dir("./static")))

	pingHandler := servicehandlers.PingHandler{}
	authHandler := servicehandlers.UserValidationHandler{}
	sessionHandler := servicehandlers.SessionHandler{}
	surveyHandler := servicehandlers.SurveyHandler{}
	userSurveyHandler := servicehandlers.UserSurveyHandler{}
	signupHandler := servicehandlers.SignupHandler{}

	go DeactivateExpiredSurvey()

	// Serves the API content
	http.Handle("/api/v1/ping/", pingHandler)

	http.Handle("/api/v1/signup/", signupHandler)
	http.Handle("/api/v1/authenticate/", authHandler)
	http.Handle("/api/v1/validate/", sessionHandler)

	http.Handle("/api/v1/survey/{surveyname}", surveyHandler)
	http.Handle("/api/v1/survey/", surveyHandler)

	http.Handle("/api/v1/usersurvey/", userSurveyHandler)
	// Start Server
	http.ListenAndServe(":3000", nil)
}
