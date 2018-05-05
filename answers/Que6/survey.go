package dao

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Question struct {
	QuestionString string   `json:"questionString" bson:"questionString"`
	Options        []string `json:"options" bson:"options"`
}

type Answer struct {
	Question Question `json:"question" bson:"question"`
	Answer   string   `json:"answer" bson:"answer"`
}

type Survey struct {
	SurveyName  string     `json:"surveyName" bson:"surveyName"`
	Heading     string     `json:"heading" bson:"heading"`
	Description string     `json:"description" bson:"description"`
	Questions   []Question `json:"questions" bson:"questions"`
	Expiry      time.Time
	Status      bool `json:"status" bson:"status"`
}

type SurveyResponse struct {
	UserName string   `json:"userName" bson:"userName"`
	Survey   Survey   `json:"survey" bson:"survey"`
	Answers  []Answer `json:"answers" bson:"answers"`
}

func GetActiveSurveys() interface{} {
	session := MgoSession.Clone()
	defer session.Close()

	var response []interface{}
	clctn := session.DB("simplesurveys").C("survey")
	query := clctn.Find(bson.M{"status": true})
	err := query.All(&response)

	if err != nil {
		return nil
	}
	return response
}

func GetAllExpiredSurveys() []Survey {
	session := MgoSession.Clone()
	defer session.Close()

	var response []Survey
	clctn := session.DB("simplesurveys").C("survey")

	query := clctn.Find(bson.M{"$and": []bson.M{bson.M{"expiry": bson.M{"$gte": time.Now()}}, bson.M{"status": true}}})
	//query := clctn.Find(bson.M{"expiry": bson.M{"$gte": time.Now()}})
	err := query.All(&response)

	if err != nil {
		return nil
	}
	return response
}

func UpdateStatusToFalse(survey Survey) string {
	session := MgoSession.Clone()
	defer session.Close()

	clctn := session.DB("simplesurveys").C("survey")

	key := bson.M{"surveyName": survey.SurveyName}
	change := bson.M{"$set": bson.M{"status": false}}
	err := clctn.Update(key, change)
	if err != nil {
		return "Error While Updating Status to false"
	}
	return "Status Updated Successfully"
}

func GetSurveysForUser(userName string) interface{} {
	session := MgoSession.Clone()
	defer session.Close()

	var response interface{}
	clctn := session.DB("simplesurveys").C("survey_response")
	query := clctn.Find(bson.M{"userName": userName})
	err := query.All(&response)

	if err != nil {
		return nil
	}
	return response
}

func GetSurveyByName(surveyName string) interface{} {
	fmt.Println("GetSurveyByName:" + surveyName)
	session := MgoSession.Clone()
	defer session.Close()

	var response interface{}
	clctn := session.DB("simplesurveys").C("survey")
	query := clctn.Find(bson.M{"surveyname": surveyName})
	err := query.One(&response)

	if err != nil {
		return nil
	} else {
		return response
	}
}

func InsertUserResponse(userResponse SurveyResponse) {
	session := MgoSession.Clone()
	defer session.Close()

	clctn := session.DB("simplesurveys").C("survey_response")
	clctn.Insert(userResponse)
}
