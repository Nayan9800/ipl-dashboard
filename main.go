package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Nayan9800/ipl-dashboard/pkg"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("IPL-Dashboard-backend")

	// New router
	r := mux.NewRouter()

	//handller for "/team"
	r.HandleFunc("/team", getTeamHandller).Methods("GET")

	//handller for "/team/{teamname}"
	r.HandleFunc("/team/{teamname}", getTeamDetails).Methods("GET")

	//handller for "/team/{teamname}/matches"
	r.HandleFunc("/team/{teamname}/matches", GetTeamMatchesByYearHandller).Methods("GET")

	//Handller for handling static file content
	fs := http.FileServer(http.Dir("frontend/build"))
	r.PathPrefix("/").Handler(fs)

	fmt.Println("starting server on port:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func getTeamHandller(w http.ResponseWriter, r *http.Request) {

	t := pkg.GetTeamNames()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func getTeamDetails(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	t, err := pkg.GetTeam(v["teamname"])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func GetTeamMatchesByYearHandller(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	y := r.URL.Query().Get("year")
	if y == "" {
		y = "2020"
	}
	data := pkg.GetTeamMatchesByYear(v["teamname"], y)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
