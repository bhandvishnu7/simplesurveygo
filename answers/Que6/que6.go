
/* main.go */

/*
func DeactivateExpiredSurvey() {
	for {
		expired_surveys := dao.GetAllExpiredSurveys()
		for _, survey := range expired_surveys {
			result := dao.UpdateStatusToFalse(survey)
			fmt.Println(result)
		}
	}
}


func main(){
	go DeactivateExpiredSurvey()
}
*/


/* dao/survey.go */

/*
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
*/
