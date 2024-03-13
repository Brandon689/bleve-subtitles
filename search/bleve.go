package search

import (
	"fmt"
	"github.com/Brandon689/bleve-subtitles/types"
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/google/uuid"
	"log"
	"os"
)

var index bleve.Index

func InitBleve() {

	indexMapping := bleve.NewIndexMapping()

	// Create a new text field mapping and set the analyzer to the custom one
	fieldMapping := bleve.NewTextFieldMapping()
	fieldMapping.Analyzer = "kagome"

	// Add the field mapping to the index mapping
	indexMapping.DefaultMapping.AddFieldMappingsAt("Lines", fieldMapping)

	// Register the custom analyzer
	//bleve.Config.DefaultIndexType = "upside_down"
	//bleve.Config.DefaultKVStore = "boltdb"
	registry.RegisterAnalyzer("kagome", AnalyzerConstructor)

	//indexMapping := bleve.NewIndexMapping()
	//
	//fieldMapping := bleve.NewTextFieldMapping()
	//fieldMapping.Analyzer = "ja"
	////mapping.DefaultMapping.AddFieldMappingsAt("Line", fieldMapping)
	//
	//indexMapping.DefaultMapping.AddFieldMappingsAt("Lines", fieldMapping)
	indexPath := "subtitles.bleve"

	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		index, err = bleve.New(indexPath, indexMapping)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Open an existing index
		index, err = bleve.Open(indexPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Terms() {
	fieldDict, err := index.FieldDict("Lines")
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the terms
	term, err := fieldDict.Next()
	for err == nil && term != nil {
		fmt.Println("Term:|" + term.Term + "|")
		term, err = fieldDict.Next()
	}
	if err != nil {
		log.Fatal(err)
	}
}

func IndexItems(nodes []types.Subtitle) {

	batch := index.NewBatch()
	for _, node := range nodes {
		batch.Index(uuid.NewString(), node)
	}
	err := index.Batch(batch)
	fmt.Println(err)
	//return b.BIndex.Batch(batch)

	//for _, item := range data {
	//	err := index.Index(uuid.NewString(), item)
	//	//fmt.Println('a')
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
}

func Inspect() {
	query := bleve.NewMatchAllQuery()

	// Create a search request
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	for _, hit := range searchResult.Hits {
		// Retrieve the document by ID
		doc, err := index.Document(hit.ID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(doc)
		// Print the document ID
		fmt.Println("Document ID:", hit.ID)

		// Iterate through fields and print terms
		//for _, field := range doc.Fields {
		//	// Print field name
		//	fmt.Println("Field:", field.Name())
		//
		//	// Iterate through terms in the field
		//	terms := doc.FieldTerms(field.Name())
		//	for _, term := range terms {
		//		fmt.Println("  Term:", term)
		//	}
		//}
	}
}

func Query(query1 string) {
	//queryString := query1
	//query := bleve.NewQueryStringQuery("Lines:" + queryString)
	//query := bleve.NewMatchQuery(query1)
	query := bleve.NewMatchQuery(query1)
	//query.SetField("Lines")
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"Lines", "Episode", "StartAt", "EndAt"}
	searchResult, _ := index.Search(searchRequest)
	fmt.Println(searchResult)
}
