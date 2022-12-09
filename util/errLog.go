package util

import (
	entity "bargerbites/Entity"
	database "bargerbites/mongodb"
	"context"
)

func LogError(errMsg string, user string) { //LogError(err_msg string , user string)
	var responses entity.Response
	UserPassDB, err := database.GetMongoDbCollection("BurgerBites", "error_logs")
	if err != nil {
		responses.Status = 401
		responses.Response = "Failed to save User Detail "
		return
	}
	var errlog entity.Errlog

	errlog.Massage = errMsg
	errlog.Time = TimeNow()
	errlog.User = user

	resPass, err := UserPassDB.InsertOne(context.Background(), errlog)
	if err != nil {
		responses.Status = 401
		responses.Response = "Failed to save User Password"
		return
	}
	id := ObjectIdToString(resPass.InsertedID)
	responses.Status = 200
	responses.Response = "Error Logged , Log Id:" + " " + id
}
