package auth

import (
	entity "bargerbites/Entity"
	database "bargerbites/mongodb"
	"bargerbites/util"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

var UserFirstname string
var UserLastname string
var Email string
var Authorization string
var Contact string

func GetJwtToken(w http.ResponseWriter, c *http.Request) {

	var user entity.Signin
	var responses entity.Response
	db := os.Getenv("DBNAME")

	usr, err := ioutil.ReadAll(c.Body)
	if err != nil {
		responses.Status = 401
		responses.Response = "Server Down !"

		json.NewEncoder(w).Encode(responses)
		return
	}
	json.Unmarshal([]byte(usr), &user)

	var m entity.MailChecker
	var creds entity.AllUserDetails
	m.Email = user.Email

	UserInDB, err := database.GetMongoDbCollection(db, "user")
	if err != nil {
		responses.Status = 401
		responses.Response = "Authentication Server Down !"

		json.NewEncoder(w).Encode(responses)
		return
	}
	result := UserInDB.FindOne(context.Background(), bson.M{"email": user.Email})
	errDecode := result.Decode(&creds)

	if errDecode != nil {
		responses.Status = 401
		responses.Response = "Please SignUp with this email !"

		json.NewEncoder(w).Encode(responses)
		return
	}

	PassDB, err := database.GetMongoDbCollection(db, "user_pass")
	if err != nil {
		responses.Status = 401
		responses.Response = "Authentication Server Down !"

		json.NewEncoder(w).Encode(responses)
		return
	}

	var passW entity.UserPass

	user_Id := util.ObjectIdToString(creds.Id)

	resultPass := PassDB.FindOne(context.Background(), bson.M{"id": user_Id})

	errPass := resultPass.Decode(&passW)
	if errPass != nil {
		responses.Status = 401
		responses.Response = "Please Provide correct password !"
		json.NewEncoder(w).Encode(responses)
		return
	}

	checkPass := CheckPasswordHash(user.Pass, passW.Pass)
	if !checkPass {
		responses.Status = 401
		responses.Response = "Please Provide correct password, password mismatch !"
		json.NewEncoder(w).Encode(responses)
		return
	}
	jwt, err := GenerateJWT(creds)
	if err != nil {
		responses.Status = 401
		responses.Response = "Unauthorised"
		json.NewEncoder(w).Encode(responses)
		return
	}
	var userDetails entity.PostSignIn

	userDetails.Age = creds.Age
	userDetails.Authorization = creds.Authorization
	userDetails.Contact = creds.Contact
	userDetails.Email = creds.Email
	userDetails.FirstName = creds.FirstName
	userDetails.Gender = creds.Gender
	userDetails.JwtToken = jwt
	userDetails.LastName = creds.LastName

	json.NewEncoder(w).Encode(userDetails)

}

func GenerateJWT(creds entity.AllUserDetails) (string, error) {
	var mySigningKey = []byte(os.Getenv("SERCET_KEY"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["auth"] = creds.Authorization
	claims["email"] = creds.Email
	claims["firstname"] = creds.FirstName
	claims["lastname"] = creds.LastName
	claims["contact"] = creds.Contact
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", fmt.Errorf("something went wrong: %s", err.Error())
	}
	return tokenString, nil
}

func MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/authenticate" {
			next.ServeHTTP(w, r)
			return
		}
		if r.URL.Path == "/signup" {
			next.ServeHTTP(w, r)
			return
		}
		if r.Header["Authorization"] == nil {
			http.Error(w, "Invalid User", http.StatusForbidden)
		}
		var mySigningKey = []byte(os.Getenv("SERCET_KEY"))

		tokenString := strings.Split(r.Header["Authorization"][0], " ")

		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok && token.Valid {
				http.Error(w, "Token is of Diffrernt type !", http.StatusForbidden)
			}
			return mySigningKey, nil
		})
		if err != nil {
			http.Error(w, "Token Expired !", http.StatusForbidden)
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// obtains claims
			UserFirstname = fmt.Sprint(claims["firstname"])
			UserLastname = fmt.Sprint(claims["lastname"])
			Email = fmt.Sprint(claims["email"])
			Authorization = fmt.Sprint(claims["auth"])
			Contact = fmt.Sprint(claims["contact"])
		}

		next.ServeHTTP(w, r)
	})
}

//for Unvarified token every jwt can be parsed but need field names
// func extractToken(accessToken string) {
// 	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
// 	if err != nil {
// 		fmt.Printf("Error %s", err)
// 	}
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 		// obtains claims
// 		UserFirstname = fmt.Sprint(claims["firstname"])
// 		UserLastname = fmt.Sprint(claims["lastname"])
// 		Email = fmt.Sprint(claims["email"])
// 		Authorization = fmt.Sprint(claims["auth"])
// 	}
// }

func GetUserFirstName() string {
	return UserFirstname
}

func GetUserLasttName() string {
	return UserLastname
}

func GetEmail() string {
	return Email
}

func GetAuthorization() string {
	return Authorization
}

func GetUserContact() string {
	return Contact
}
