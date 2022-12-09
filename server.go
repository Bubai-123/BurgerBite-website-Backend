package main

import (
	"bargerbites/auth"
	database "bargerbites/mongodb"
	"bargerbites/routes"
	"log"
	"net/http"
	"os"

	_ "crypto/sha256"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	r := mux.NewRouter()
	if len(os.Args) > 1 {
		if os.Args[1] == "dev" {
			err := godotenv.Load(".dev.env")
			if err != nil {
				log.Println("error loading env")
			}
		} else if os.Args[1] == "sit" {
			err := godotenv.Load(".sit.env")
			if err != nil {
				log.Println("error loading env")
			}
		} else if os.Args[1] == "uat" {
			err := godotenv.Load(".uat.env")
			if err != nil {
				log.Println("error loading env")
			}
		} else if os.Args[1] == "prod" {
			err := godotenv.Load(".prod.env")
			if err != nil {
				log.Println("error loading env")
			}
		}
	} else {
		err := godotenv.Load(".dev.env")
		if err != nil {
			log.Println("error loading env")
		}
	}

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	database.Setup()
	r.HandleFunc("/authenticate", auth.GetJwtToken).Methods("POST")
	r.Use(auth.MiddleWare)

	routes.RouterDefination(r)
	log.Println("URL: ", host+"/:"+port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, r))
}
