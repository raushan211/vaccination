package main

import (
	"database/sql"
	"fmt"
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
type Delete struct {
	Email string `json:"email"`
}

func setupRoutes(r *gin.Engine) {

	r.GET("/nurses", redirectNurses)
	r.GET("/persons", redirectPersons)
	r.GET("/vaccinated_nurses", redirectVaccinatedNurses)
	r.GET("/vaccinated_persons", redirectVaccinatedPersons)
	r.DELETE("/delete_nurses/", deleteNurses)
	r.DELETE("/delete_persons/", deletePersons)
	r.POST("/add_nurse/", addNurse)
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
func deleteNurses(c *gin.Context) {
	Data := Nurses{}
	c.Bind(&Data)
	SQL := "DELETE from nurses where email=$1"
	_, err := DB.Exec(SQL, Data.Email)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return
	}
	//defer rows.Close()
	//nurse := Nurses{}

	res := gin.H{
		"nurses": Data,
	}
	c.JSON(http.StatusOK, res)
}
func deletePersons(c *gin.Context) {
	Data := Persons{}
	c.Bind(&Data)
	SQL := "DELETE from persons where email=$1 "
	_, err := DB.Exec(SQL, Data.Email)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return
	}
	//defer rows.Close()
	//person := Persons{}

	res := gin.H{
		"persons": Data,
	}

	c.JSON(http.StatusOK, res)

}
func addNurse(c *gin.Context) {
	reqBody := Persons{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"Error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	result, _ := addNurses(reqBody)
	if result == false {
		res := gin.H{
			"Error": "Something is wromg",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res := gin.H{
		"message": "successfully inserted",
		"status":  200,
	}
	c.JSON(http.StatusCreated, res)
}
func addNurses(reqbody Persons) (bool, string) {
	var result = true
	var err_responce = ""

	sqlStatement := `
INSERT INTO persons(first_name, last_name,email, date_of_birth, sex)
VALUES ($1, $2, $3, $4,$5)`
	_, err2 := DB.Exec(sqlStatement, reqbody.First_name, reqbody.Last_name, reqbody.Email, reqbody.Date_of_birth, reqbody.Sex)
	fmt.Println(err2)
	if err2 != nil {

		err_responce = "Something went wrong"
		result = false
		return result, err_responce
	}

	sqlStatement2 := `
INSERT INTO nurses(email) VALUES ($1)`
	_, err3 := DB.Exec(sqlStatement2, reqbody.Email)
	fmt.Println(err3)
	//fmt.Println(user)
	if err3 != nil {

		err_responce = "Something went wrong"
		result = false
		return result, err_responce
	}
	return result, err_responce
}
