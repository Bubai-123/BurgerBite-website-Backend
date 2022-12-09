package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSignUp struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Age       int    `json:"age,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Contact   string `json:"contact,omitempty"`
	AuthCode  string `json:"authcode,omitempty"`
	Pass      string `json:"password"`
}

type UserDetails struct {
	FirstName     string `json:"firstname,omitempty"`
	LastName      string `json:"lastname,omitempty"`
	Email         string `json:"email,omitempty"`
	Age           int    `json:"age,omitempty"`
	Gender        string `json:"gender,omitempty"`
	Contact       string `json:"contact,omitempty"`
	Authorization string `json:"auth,omitempty"`
}

type UserPass struct {
	Id   string `json:"id"`
	Pass string `json:"pass"`
}

type Signin struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
}

type PostSignIn struct {
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Email         string `json:"email"`
	Age           int    `json:"age"`
	Gender        string `json:"gender"`
	Contact       string `json:"contact"`
	Authorization string `json:"auth"`
	JwtToken      string `json:"jwt"`
}

type MailChecker struct {
	Email string `json:"email"`
}

type AllUserDetails struct {
	Id            primitive.ObjectID `bson:"_id"`
	FirstName     string             `json:"firstname,omitempty"`
	LastName      string             `json:"lastname,omitempty"`
	Email         string             `json:"email,omitempty"`
	Age           int                `json:"age,omitempty"`
	Gender        string             `json:"gender,omitempty"`
	Contact       string             `json:"contact,omitempty"`
	Authorization string             `json:"authorization,omitempty"`
}
