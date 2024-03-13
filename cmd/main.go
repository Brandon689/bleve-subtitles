//package main
//
//import (
//	"fmt"
//	"log"
//	"time"
//
//	"github.com/blevesearch/bleve/v2"
//)
//
//type Subtitle struct {
//	ID      string
//	Episode string
//	Lines   []string
//	StartAt time.Duration
//	EndAt   time.Duration
//}
//
//func main() {
//	// Open or create a new Bleve index
//	index, err := bleve.Open("my_index.bleve")
//	if err == bleve.ErrorIndexPathDoesNotExist {
//		indexMapping := bleve.NewIndexMapping()
//		indexMapping.DefaultMapping.AddFieldMappingsAt("Lines", bleve.NewTextFieldMapping())
//
//		index, err = bleve.New("my_index.bleve", indexMapping)
//		if err != nil {
//			log.Fatal(err)
//		}
//	} else if err != nil {
//		log.Fatal(err)
//	}
//
//	// Sample subtitle data
//	subtitle := Subtitle{
//		Episode: "Episode 1",
//		Lines:   []string{"Subtitle line 1", "Subtitle line 2"},
//		StartAt: 0,
//		EndAt:   10 * time.Second,
//	}
//
//	// Index the subtitle
//	err = index.Index("", subtitle)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Query for a specific line within the Lines field
//	query := bleve.NewQueryStringQuery("Lines:\"Subtitle line 1\"")
//
//	searchRequest := bleve.NewSearchRequest(query)
//	searchResult, err := index.Search(searchRequest)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Print the search result
//	fmt.Println("Search Result:", searchResult)
//
//	// Close the index when done
//	index.Close()
//}

package main

import (
	"fmt"
	"github.com/Brandon689/bleve-subtitles/data"
	"github.com/Brandon689/bleve-subtitles/search"
	"github.com/Brandon689/bleve-subtitles/types"
	_ "github.com/blevesearch/bleve/v2/analysis/lang/cjk"
	"time"
)

var Subs []types.Subtitle

type Stopwatch struct {
	startTime time.Time
}

func (s *Stopwatch) Start() {
	s.startTime = time.Now()
}

func (s *Stopwatch) Elapsed() time.Duration {
	return time.Since(s.startTime)
}

func main() {

	//n := subs[:100000]
	//for _, item := range n {
	//	id := data.GenerateID("Arakaw", item.Lines)
	//	item.ID = id
	//}

	//fmt.Println(len(n))
	search.InitBleve()
	//search.Inspect()
	stopwatch := &Stopwatch{}

	// Start the stopwatch

	root := "C:\\Users\\Brandon\\Pictures\\New folder\\"
	subs := data.GetSubtitles(root)
	subs = subs[:1000]
	stopwatch.Start()
	search.IndexItems(subs)

	// Get the elapsed time
	elapsedTime := stopwatch.Elapsed()

	// Print the elapsed time
	fmt.Println("Elapsed Time:", elapsedTime)
	search.Terms()
	//search.Query("ある")
}
