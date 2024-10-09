package main

import (
	"fmt"
	"log"
	"net/http"
)

var dummyData map[int]Person
var pId int

func init() {
	pId = 0
	dummyData = make(map[int]Person)
}

type Person struct {
	Name string `json:"name,omitempty"`
	age  string `json:"age,omitempty"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Println("server is running in 4444")
	http.ListenAndServe(":4444", nil)

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("server is running"))

}

func createPerson(person Person) {
	dummyData[pId] = person
	pId++
}

func deletePerson(personId int) {
	_, ok := dummyData[personId]
	if ok {
		delete(dummyData, personId)
		fmt.Println("user deleted successfully")
	}
}

func update(person Person, personId int) {
	_, ok := dummyData[personId]
	if ok {
		dummyData[personId] = person
		fmt.Println("user updated successfully")
	}
}

func ReadPersonById(personId int) (Person, error) {
	person, ok := dummyData[personId]
	if ok {
		return person, nil
	}
	return Person{}, fmt.Errorf("some error occur")
}
