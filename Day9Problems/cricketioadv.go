package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Cricket struct {
	Name     string
	Period   string
	Matches  int
	Runs     int
	Avgscore string
}

func main() {
	var filename string
	fmt.Println("Enter the CSV file name")
	fmt.Scanf("%s", &filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	header, err := csvReader.Read()
	rec, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	var Playerlist []Cricket
	for _, record := range rec {
		matches, _ := strconv.Atoi(record[2])
		updatedmatch := matches + 10
		runs, _ := strconv.Atoi(record[3])

		player := Cricket{
			Name:     record[0],
			Period:   record[1],
			Matches:  updatedmatch,
			Runs:     runs,
			Avgscore: record[4],
		}
		Playerlist = append(Playerlist, player)
	}
	Playerlist[0].Runs = Playerlist[0].Runs + 1300
	Playerlist[1].Runs = Playerlist[1].Runs + 10000
	Playerlist[2].Runs = Playerlist[2].Runs + 1500
	Playerlist[3].Runs = Playerlist[3].Runs + 100
	Playerlist[4].Runs = Playerlist[4].Runs + 3000
	f, err := os.Create("Cricupdated.csv")
	defer f.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(f)
	defer w.Flush()
	_ = w.Write([]string(header))

	var data [][]string
	for _, record := range Playerlist {
		row := []string{record.Name, record.Period, strconv.Itoa(record.Matches), strconv.Itoa(record.Runs), record.Avgscore}
		data = append(data, row)
	}
	w.WriteAll(data)
}
