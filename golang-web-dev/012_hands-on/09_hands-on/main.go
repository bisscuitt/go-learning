package main

import (
	"encoding/csv"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// Open the file
	fh, err := os.Open("table.csv")
	if err != nil {
		log.Fatal("Could not open file:", err)
	}
	defer fh.Close()

	r := csv.NewReader(fh)
	records, err2 := r.ReadAll()
	if err2 != nil {
		log.Fatal("Could not parse CSV:", err)
	}

	err = tpl.Execute(os.Stdout, records)

}
