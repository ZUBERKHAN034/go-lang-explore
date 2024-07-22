package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	ID         string  `json:"id"`
	CourseName string  `json:"courseName"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// Fake DB - file
var courses []Course

// Middleware - file
func (course *Course) IsEmpty() bool {
	return course.CourseName == ""
}

func main() {
	fmt.Println("API in GO lang...")
}

// Controller - file
func serveHome(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Requesting home...", req.URL.Path)
	res.Write([]byte("<h1>Hello World!<h1>"))
}

func getCourses(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Requesting get all courses...", req.URL.Path)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(courses)
}

func getCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Requesting get course...", req.URL.Path)
	res.Header().Set("Content-Type", "application/json")

	// Get the id from the request
	params := mux.Vars(req)
	id := params["id"]

	var resCourse *Course = nil
	// Loop through the courses to find the course with the id
	for i := range courses {
		if courses[i].ID == id {
			resCourse = &courses[i]
			break
		}
	}

	// If course not found
	if resCourse == nil {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(map[string]string{"message": "Course not found"})
		return
	}

	// Return the course
	json.NewEncoder(res).Encode(resCourse)
}

func createCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Requesting create course...", req.URL.Path)
	res.Header().Set("Content-Type", "application/json")

	if req.Body == nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{"message": "Invalid request body"})
		return
	}
	defer req.Body.Close()

	var course Course
	err := json.NewDecoder(req.Body).Decode(&course)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{"message": "Invalid JSON format"})
		return
	}

	if course.IsEmpty() {
		res.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(res).Encode(map[string]string{"message": "Empty request body"})
		return
	}

	course.ID = uuid.New().String()

	courses = append(courses, course)

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(course)
}

func updateCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Requesting update course...", req.URL.Path)
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	id := params["id"]

	for index, course := range courses {
		if course.ID == id {
			defer req.Body.Close()

			var updatedCourse Course
			err := json.NewDecoder(req.Body).Decode(&updatedCourse)
			if err != nil {
				res.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(res).Encode(map[string]string{"message": "Invalid JSON format"})
				return
			}

			updatedCourse.ID = id          // Ensure the ID remains unchanged
			courses[index] = updatedCourse // Update the course in place
			json.NewEncoder(res).Encode(updatedCourse)
			return
		}
	}

	// If no course with the provided ID was found
	res.WriteHeader(http.StatusNotFound)
	json.NewEncoder(res).Encode(map[string]string{"message": "Course not found"})
}

func deleteCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Requesting delete course...", req.URL.Path)
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	id := params["id"]

	for index, course := range courses {
		if course.ID == id {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(res).Encode(map[string]string{"message": "Course deleted successfully"})
			return
		}
	}

	// If no course with the provided ID was found
	res.WriteHeader(http.StatusNotFound)
	json.NewEncoder(res).Encode(map[string]string{"message": "Course not found"})
}
