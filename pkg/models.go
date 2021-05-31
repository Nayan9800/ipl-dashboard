package pkg

import "github.com/hashicorp/go-memdb"

type match struct {
	ID            int    `csv:"id" json:"id"`
	City          string `csv:"city" json:"city"`
	Date          string `csv:"date" json:"date"`
	ManOfTheMatch string `csv:"man_of_the_match" json:"playerOfMatch"`
	Venue         string `csv:"venue" json:"venue"`
	TeamA         string `csv:"team_A" json:"team1"`
	TeamB         string `csv:"team_B" json:"team2"`
	TossWinner    string `csv:"toss_winner" json:"tossWinner"`
	ChoseTo       string `csv:"chose_to" json:"tossDecision"`
	Winner        string `csv:"winner" json:"matchWinner"`
	Result        string `csv:"result" json:"result"`
	ResultMargin  string `csv:"result_margin" json:"resultMargin"`
	Umpire1       string `csv:"umpire_1" json:"umpire1"`
	Umpire2       string `csv:"umpire_2" json:"umpire2"`
}
type team struct {
	Id           int      `json:"id"`
	Name         string   `json:"teamName"`
	TotalMacthes int      `json:"totalMatches"`
	TotalWins    int      `json:"totalWins"`
	Matches      []*match `json:"matches"`
}

var matchschema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"match": &memdb.TableSchema{
			Name: "match",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:         "id",
					Unique:       true,
					Indexer:      &memdb.IntFieldIndex{Field: "ID"},
					AllowMissing: false,
				},
				"city": &memdb.IndexSchema{
					Name:    "city",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "City"},
				},
				"date": &memdb.IndexSchema{
					Name:    "date",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Date"},
				},
				"man_of_the_match": &memdb.IndexSchema{
					Name:    "man_of_the_match",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "ManOfTheMatch"},
				},
				"venue": &memdb.IndexSchema{
					Name:    "venue",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Venue"},
				},
				"team_A": &memdb.IndexSchema{
					Name:    "team_A",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "TeamA"},
				},
				"team_B": &memdb.IndexSchema{
					Name:    "team_B",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "TeamB"},
				},
				"tosswinner": &memdb.IndexSchema{
					Name:    "tosswinner",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "TossWinner"},
				},
				"choseto": &memdb.IndexSchema{
					Name:    "choseto",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "ChoseTo"},
				},
				"winner": &memdb.IndexSchema{
					Name:    "winner",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Winner"},
				},
				"result": &memdb.IndexSchema{
					Name:    "result",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Result"},
				},
				"result_margin": &memdb.IndexSchema{
					Name:    "result_margin",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "ResultMargin"},
				},
				"umpire_1": &memdb.IndexSchema{
					Name:    "umpire_1",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Umpire1"},
				},
				"umpire_2": &memdb.IndexSchema{
					Name:    "umpire_2",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Umpire2"},
				},
			},
		},
	},
}
