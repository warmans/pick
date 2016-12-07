package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const (
	INPUT_TYPE_CSV         = "csv"
	INPUT_TYPE_QUERYSTRING = "qs"
)

var outColumns = flag.String("d", "", "data to output. depends on the data type.")
var inputType = flag.String("t", INPUT_TYPE_CSV, "input type. depends on the data type.")

func main() {

	flag.Parse()

	writer := csv.NewWriter(os.Stdout)
	reader := csv.NewReader(os.Stdin)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("Expected argument with file path")
		}
		writer.Write(pick(*inputType, record, strings.Split(*outColumns, ",")))
		writer.Flush()
	}
}

func pick(input string, record, data []string) []string {

	newRecord := make([]string, len(data))

	switch input {
	case INPUT_TYPE_CSV:
		for newIndex, pickIndex := range data {
			pickIndexInt := mustAtoi(pickIndex)
			if pickIndexInt < len(record) {
				newRecord[newIndex] = record[pickIndexInt]
			}
		}
	case INPUT_TYPE_QUERYSTRING:
		for newIndex, pickIndex := range data {
			query, err := url.ParseQuery(strings.Join(record, ","))
			if err != nil {
				panic(err)
			}
			newRecord[newIndex] = query.Get(pickIndex)
		}
	}

	return newRecord
}

func mustAtoi(stringVal string) int {
	if stringVal == "" {
		return 0
	}
	intVal, err := strconv.Atoi(stringVal)
	if err != nil {
		log.Fatal(err)
	}
	return intVal
}
