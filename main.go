package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    int
	CourseName  string
	CoursePrice float64
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string
	Website  string
}

// fake DB
var courses []Course

// middleware, helper - file

func main() {
	fmt.Println("")
	r := mux.NewRouter()

	//seeding
	course1 := Course{CourseId: 42, CourseName: "DSA", CoursePrice: 4999, Author: &Author{Fullname: "Rahul", Website: "dsabyrahul.com"}}
	_ = Course{CourseId: 43, CourseName: "OS", CoursePrice: 9999, Author: &Author{Fullname: "Ajay", Website: "osbyajay.com"}}
	courses = append(courses, course1)

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/create-course", createOneCourse).Methods("POST")
	r.HandleFunc("/get-all", getAllCourses).Methods("GET")

	// listen to a port
	fmt.Println(courses)
	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "welcome to the Homepage")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCourse Course
	err := json.NewDecoder(r.Body).Decode(&newCourse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	courses = append(courses, newCourse)
	err = json.NewEncoder(w).Encode(newCourse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// Get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
