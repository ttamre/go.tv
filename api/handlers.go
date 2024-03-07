package api

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to go.tv")
}


func UserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	switch r.Method {
		case "POST":
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			userJSON, _ := json.Marshal(user)
			log.Println(string(userJSON))
			fmt.Fprint(w,string(userJSON))

		case "GET":
			user = User{Username: r.FormValue("username"), Password: r.FormValue("password")}
			userJSON, _ := json.Marshal(user)
			log.Println(string(userJSON))
			fmt.Fprint(w,string(userJSON))

		case "PUT":
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			userJSON, _ := json.Marshal(user)
			log.Println(string(userJSON))
			fmt.Fprint(w,string(userJSON))

		case "DELETE":
			// delete user
			user = User{}
		}
}

func ReviewHandler(w http.ResponseWriter, r *http.Request) {
	var review Review

	switch r.Method {
		case "POST":
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&review)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			reviewJSON, _ := json.Marshal(review)
			log.Println(string(reviewJSON))
			fmt.Fprint(w,string(reviewJSON))

		case "GET":
			review = Review{Username: r.FormValue("username"), Password: r.FormValue("password")}
			reviewJSON, _ := json.Marshal(review)
			log.Println(string(reviewJSON))
			fmt.Fprint(w,string(reviewJSON))

		case "PUT":
			decoder := json.NewDecoder(r.Body)
			var review Review
			err := decoder.Decode(&review)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			reviewJSON, _ := json.Marshal(review)
			log.Println(string(reviewJSON))
			fmt.Fprint(w,string(reviewJSON))

		case "DELETE":
			review = Review{}
		}
	}
