package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func main() {
	createDBConnection()
	//defer DB.Close()
	//Data = make(map[string]User)
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type Nurses struct {
	Email string `json:"email"`
}
type Persons struct {
	Email         string `json:"email"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Date_of_birth string `json:"date_of_birth"`
	Sex           string `json:"sex"`
}
type Vaccinated_Nurses struct {
	Email string `json:"email"`
}
type Vaccinated_Persons struct {
	Email string `json:"email"`
}

func setupRoutes(r *gin.Engine) {

	r.GET("/nurses", redirectNurses)
	r.GET("/persons", redirectPersons)
	r.GET("/vaccinated_nurses", redirectVaccinatedNurses)
	r.GET("/vaccinated_persons", redirectVaccinatedPersons)

}
func redirectNurses(c *gin.Context) {
	Data := []Nurses{}

	SQL := "SELECT email from nurses"
	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return
	}
	defer rows.Close()
	nurse := Nurses{}

	for rows.Next() {
		rows.Scan(&nurse.Email)

		Data = append(Data, nurse)

	}
	res := gin.H{
		"nurses": Data,
	}
	c.JSON(http.StatusOK, res)
}
func redirectPersons(c *gin.Context) {
	Data := []Persons{}

	SQL := "SELECT email,first_name,last_name,date_of_birth,sex from persons"
	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return
	}
	defer rows.Close()
	person := Persons{}

	for rows.Next() {
		rows.Scan(&person.Email, &person.First_name, &person.Last_name, &person.Date_of_birth, &person.Sex)

		Data = append(Data, person)

	}
	res := gin.H{
		"persons": Data,
	}

	c.JSON(http.StatusOK, res)

}
func redirectVaccinatedNurses(c *gin.Context) {
	Data := []Nurses{}

	SQL := "select nurses.email from nurses, vaccinations where nurses.email=vaccinations.recipient"
	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return
	}
	defer rows.Close()
	nurse := Nurses{}

	for rows.Next() {
		rows.Scan(&nurse.Email)

		Data = append(Data, nurse)

	}
	res := gin.H{
		"vaccinated_Nurses": Data,
	}

	c.JSON(http.StatusOK, res)

}
func redirectVaccinatedPersons(c *gin.Context) {
	Data := []Persons{}

	SQL := "SELECT persons.email from persons, vaccinations WHERE persons.email=vaccinations.recipient"
	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return
	}
	defer rows.Close()
	person := Persons{}

	for rows.Next() {
		rows.Scan(&person.Email)

		Data = append(Data, person)

	}
	res := gin.H{
		"vaccinated_Persons": Data,
	}

	c.JSON(http.StatusOK, res)
}
