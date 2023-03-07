package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Models for course - file
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// Fake DB
var courses []Course

// middleware, helper - file
func isEmpty(c *Course) bool  {
	return c.CourseId == "" && c.CourseName == ""
}



func main()  {
	
}

// Controllers - file
// serve home route
func serveHome(w http.ResponseWriter , r *http.Request)  {
	w.Write([]byte("<h1>Welcome to API by Yaduveera...</h1>"))
}

func getAllCourses(w http.ResponseWriter , r *http.Request)  {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	// following line is similar to res.json() of node
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter , r *http.Request)  {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from req and is similar to req.params in node
	params := mux.Vars(r)

	// loop through courses and find matching id and return the response
	for _, course := range courses{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with the given id")
	return

}