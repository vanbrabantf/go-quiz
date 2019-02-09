package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var score = 0

func main() {
	filename := flag.String("filename", "problems", "the quiz you want to run")
	timelimit := flag.Int("limit", 30, "how long the quiz can run")
	flag.Parse()

	fileContent := readFileContents(*filename)
	records := getSliceFromCsv(fileContent)

	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
	go runQuiz(records)

	<-timer.C

	fmt.Printf("Your final score is %d out of %d", score, len(records))
}

func readFileContents(fileName string) string {
	file, err := ioutil.ReadFile(fileName + ".csv")

	if err != nil {
		fmt.Println("no such file")

		os.Exit(1)
	}

	return string(file)
}

func getSliceFromCsv(csvString string) [][]string {
	records := csv.NewReader(strings.NewReader(csvString))
	slices, err := records.ReadAll()

	if err != nil {
		fmt.Println("error reading the csv")

		os.Exit(2)
	}

	return slices
}

func runQuiz(records [][]string) {
	reader := bufio.NewReader(os.Stdin)

	for _, record := range records {

		fmt.Println(record[0] + "= ")
		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(text) == strings.TrimSpace(record[1]) {
			fmt.Println("✅")
			score++
		} else {
			fmt.Println("❌")
		}
	}
}
