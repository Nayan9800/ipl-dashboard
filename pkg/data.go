package pkg

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/hashicorp/go-memdb"
)

var db *memdb.MemDB
var teamdata map[string]*team

func init() {

	matches := []*match{}
	teamdata = make(map[string]*team)

	// Reading data from match-data.csv
	r, err := os.Open("resources\\match-data.csv")
	showerr(err)

	// Unmarshal the data
	err = gocsv.Unmarshal(r, &matches)
	showerr(err)

	// Creating new Memdb instance with matchschema
	db, err = memdb.NewMemDB(matchschema)
	showerr(err)

	// Sorting Array of matches into Aescending Order By Date
	sort.SliceStable(matches, func(i, j int) bool {
		t1, err := time.Parse("02-01-06", matches[i].Date)
		showerr(err)
		t2, err := time.Parse("02-01-06", matches[j].Date)
		showerr(err)
		return t1.Before(t2)
	})

	// Create Map of Teams
	for _, m := range matches {
		createOrUpdate(m.TeamA, m)
		createOrUpdate(m.TeamB, m)
	}

	//Inserting Data Into Db
	tx := db.Txn(true)
	for i, m := range matches {
		m.ID = i
		if err := tx.Insert("match", m); err != nil {
			log.Fatal(err.Error())
		}
	}
	tx.Commit()
}

//Get all Teams
func GetTeamNames() (t []*team) {
	i := 0
	for _, v := range teamdata {
		t = append(t, &team{Id: i, Name: v.Name, TotalMacthes: v.TotalMacthes, TotalWins: v.TotalWins})
		i++
	}
	return
}

//Get the team by given name with recent four matches
func GetTeam(name string) (t *team, err error) {

	t, ok := teamdata[name]
	if !ok {
		return nil, fmt.Errorf("%v Not found", name)
	}

	m := getMatchesByteamName(name)
	t.Matches = m[len(m)-4:]
	return

}

//Get team matches by given Team name and Year
func GetTeamMatchesByYear(name string, year string) []*match {

	m := []*match{}

	temp, err := time.Parse("02-01-06", "01-01-"+strings.TrimPrefix(year, "20"))
	showerr(err)

	for _, tm := range getMatchesByteamName(name) {

		t, err := time.Parse("02-01-06", tm.Date)
		showerr(err)
		if t.After(temp) || t.Equal(temp) {
			m = append(m, tm)
		}
	}

	return m
}

//Get all Matches of given Team name and sorted by Date
func getMatchesByteamName(name string) []*match {

	m := []*match{}
	txr := db.Txn(false)
	tA, err := txr.Get("match", "team_A", name)
	showerr(err)
	tB, err := txr.Get("match", "team_B", name)
	showerr(err)
	for v := tA.Next(); v != nil; v = tA.Next() {
		m = append(m, v.(*match))

	}
	for v := tB.Next(); v != nil; v = tB.Next() {
		m = append(m, v.(*match))

	}
	sort.SliceStable(m, func(i, j int) bool {
		t1, err := time.Parse("02-01-06", m[i].Date)
		showerr(err)
		t2, err := time.Parse("02-01-06", m[j].Date)
		showerr(err)
		return t1.Before(t2)
	})

	return m
}

//To show error
func showerr(e error) {
	if e != nil {
		log.Println(e.Error())
	}
}

//Helps to Create and Update team
func createOrUpdate(name string, m *match) {

	if v, ok := teamdata[name]; ok {

		v.TotalMacthes++
		if m.Winner == name {
			v.TotalWins++
		}
		teamdata[name] = v
	} else {

		tmp := &team{Name: name, TotalMacthes: 1, TotalWins: 0}
		if m.Winner == name {
			tmp.TotalWins++
		}
		teamdata[name] = tmp
	}
}
