package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []region{
		region{"Southern",
			[]hotel{
				{"Fake Hotel", "123 Fake Street", "Fakey Beach", 81652},
				{"Hotel Dela Fake", "4096 Fake Street", "Fakington", 11023},
			},
		},
		region{"Central",
			[]hotel{
				{"Hotel Fakeafornia", "456 Fake Street", "Fakesbury", 90211},
				{"Park Fake Hotel", "1024 Fake Street", "Upper Fakering", 23798},
			},
		},
		region{"Northern",
			[]hotel{
				{"Fake International on Fake", "789 Fake Street", "Faking-on-Fake", 90210},
			},
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
