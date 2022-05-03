package contoller

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq" // <------------ here
	"html"
	"log"
	"net/http"
	"web-project/dao"
	"web-project/model"
)

// contoller or handler
func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register controller, %q", html.EscapeString(r.URL.Path))

	var register model.Register

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//fmt.Fprintf(w, "Registration Data : %+v", register)

	db := dao.OpenConnection()

	sqlStatement := `INSERT INTO person (first_name, last_name, email, mobile) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(sqlStatement, register.FirstName, register.LastName, register.Email, register.Mobile)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

func GETHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get controller")
	fmt.Fprintf(w, "Get controller, %q", html.EscapeString(r.URL.Path))

	db := dao.OpenConnection()

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Fatal(err)
	}

	var people []model.Register

	for rows.Next() {
		var person model.Register
		rows.Scan(&person.FirstName, &person.LastName)
		people = append(people, person)
	}

	peopleBytes, _ := json.MarshalIndent(people, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleBytes)

	defer rows.Close()
	defer db.Close()
}
