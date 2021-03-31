package config

import (
	"encoding/csv"
	"log"
	"os"
)

var path = "./config/serverData.csv"

type ServerData struct {
	ServerID     string
	CanvasToken  string
	AlertChannel string
}

func UpdateData(serverItem *ServerData) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data [][]string
	data = append(data, []string{serverItem.ServerID, serverItem.CanvasToken, serverItem.AlertChannel})

	w := csv.NewWriter(f)
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	log.Println("Appended successfully")
}

func ReadData() []ServerData {
	var data []ServerData

	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Unable to read input file "+path, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+path, err)
	}

	for _, recs := range records {
		data = append(data, *&ServerData{recs[0], recs[1], recs[2]})

	}
	return data
}
