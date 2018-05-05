package servicehandlers

import (
	"fmt"
	"simplesurveygo/dao"
	"time"
)

func Decorator(auth func(name string)) func(s string) {

	return func(name string) {
		fmt.Println("Started")
		startTime := time.Now()

		auth(name)

		endTime := time.Now()
		diff := endTime.Sub(startTime)
		fmt.Println("Time taken by request to UserValidation", diff.Seconds(), "seconds")
	}

}

func authenticate(token string) {
	user := dao.GetSessionDetails(token)
	userName := user.Username
	fmt.Println(userName)
}
