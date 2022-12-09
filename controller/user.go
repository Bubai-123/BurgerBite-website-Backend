package controller

import (
	entity "bargerbites/Entity"
	"bargerbites/auth"
	"bargerbites/constants"
	mongodb "bargerbites/mongodb"
	"bargerbites/util"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " WELCOME TO BURGERBITES!")
}

func SignUpUser(w http.ResponseWriter, r *http.Request) {
	db := os.Getenv("DBNAME")
	var responses entity.Response

	UserDB, err := mongodb.GetMongoDbCollection(db, "user")
	if err != nil {
		responses.Status = 401
		responses.Response = "Failed to save User Detail "

		json.NewEncoder(w).Encode(responses)
		return
	}

	var person entity.UserSignUp
	a, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal([]byte(a), &person)
	var creds entity.UserDetails
	result := UserDB.FindOne(context.Background(), bson.M{"email": person.Email})
	errDecode := result.Decode(&creds)

	if errDecode == nil {
		responses.Status = 202
		responses.Response = "Hi," + creds.FirstName + " " + creds.LastName + " " + "your account already Registered with this email "

		json.NewEncoder(w).Encode(responses)
		return
	}

	var Details entity.UserDetails

	Details.Age = person.Age
	if person.AuthCode == constants.SELLER_CODE {
		Details.Authorization = "SELLER"
	} else {
		Details.Authorization = "CUSTOMER"
	}

	_, NumCheck := strconv.Atoi(person.Contact)
	if NumCheck != nil {
		responses.Status = 401
		responses.Response = "Please Insert Valid Phone Number"

		json.NewEncoder(w).Encode(responses)

		return
	}
	if len(person.Contact) > 10 || len(person.Contact) < 10 {
		responses.Status = 401
		responses.Response = "Please Insert Valid Phone Number"

		json.NewEncoder(w).Encode(responses)

		return
	}

	Details.Contact = person.Contact
	if !strings.Contains(person.Email, "@") || !strings.Contains(person.Email, ".com") {
		responses.Status = 401
		responses.Response = "Please Insert Valid Email"
		json.NewEncoder(w).Encode(responses)

		return
	}
	Details.Email = person.Email
	Details.FirstName = person.FirstName
	Details.Gender = person.Gender
	Details.LastName = person.LastName

	resUser, err := UserDB.InsertOne(context.Background(), Details)
	if err != nil {
		responses.Status = 401
		responses.Response = "Failed to save User Password"
		json.NewEncoder(w).Encode(responses)

		return
	}
	user_Id := util.ObjectIdToString(resUser.InsertedID)

	var userPass entity.UserPass
	userPass.Id = user_Id
	HashPass, _ := auth.HashPassword(person.Pass)
	userPass.Pass = HashPass

	UserPassDB, err := mongodb.GetMongoDbCollection(db, "user_pass")
	if err != nil {
		responses.Status = 401
		responses.Response = "Failed to save User Detail "

		json.NewEncoder(w).Encode(responses)

		return
	}

	resPass, err := UserPassDB.InsertOne(context.Background(), userPass)
	if err != nil {
		responses.Status = 401
		responses.Response = "Failed to save User Password"
		json.NewEncoder(w).Encode(responses)

		return
	}

	ID := util.ObjectIdToString(resPass.InsertedID)

	if ID != "" && user_Id != "" {
		responses.Status = 200
		responses.Response = "account created Successfully "
	}

	json.NewEncoder(w).Encode(responses)

}
