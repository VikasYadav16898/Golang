package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
func (c *Course) IsEmpty() bool  {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}



func main()  {
	fmt.Println("API - Yaduveera")

	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId:"2", CourseName: "React js", CoursePrice: 299, Author: &Author{
		Fullname: "yaduveera", Website: "Yaduveera.global.in",
	}})
	courses = append(courses, Course{CourseId:"4", CourseName: "Mern Stack", CoursePrice: 199, Author: &Author{
		Fullname: "yaduveera", Website: "go.dev.in",
	}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")


	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

	
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

func createOneCourse(w http.ResponseWriter , r *http.Request)  {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data!")
		return
	}

	// what if data = {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the sent JSON!")
		return
	}

	// genrate a unique id, string
	// append new course into courses

	
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter , r *http.Request){
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// first- grab id from req
	params := mux.Vars(r)

	// loop through value and once id matched and than remove that record and add the new value record...
	for index, course := range courses{
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// Todo: send a response when id is not found
	json.NewEncoder(w).Encode("No record with same id found")
	return
}

func deleteOneCourse(w http.ResponseWriter , r *http.Request)  {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove (index, index+1)
	for index, course := range courses{
		if course.CourseId == params["id"] {
			courses = append( courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("Succesfully deleted the course")
	return
}