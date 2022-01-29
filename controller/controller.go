package controller

import (
	"encoding/json"
	"go-crud/data"
	"go-crud/service"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PersonApi struct {
	Service service.PersonService
}

type Route struct {
	Handle func(w http.ResponseWriter, r *http.Request)
	Path   string
	Method []string
}

func (p PersonApi) GetAllRoutes() []Route {
	routes := make([]Route, 1)
	routes = append(routes, Route{
		Handle: p.getPerson,
		Path:   "/v1/person/{id}",
		Method: p.createMethods("GET"),
	})
	routes = append(routes, Route{
		Handle: p.updatePerson,
		Path:   "/v1/person/{id}",
		Method: p.createMethods("PUT"),
	})
	routes = append(routes, Route{
		Handle: p.createPerson,
		Path:   "/v1/person",
		Method: p.createMethods("POST"),
	})
	return routes
}

func (p PersonApi) createMethods(methods ...string) []string {
	return append(make([]string, 1), methods...)
}

func (p PersonApi) getPerson(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return
	}

	result, _ := p.Service.GetPerson(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (p PersonApi) createPerson(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person data.Person
	json.Unmarshal(requestBody, &person)

	result, _ := p.Service.CreatePerson(&person)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (p PersonApi) updatePerson(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(400)
		return
	}
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person data.Person
	json.Unmarshal(requestBody, &person)

	result, _ := p.Service.UpdatePerson(id, &person)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
