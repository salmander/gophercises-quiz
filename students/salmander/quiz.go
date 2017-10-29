package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	// Read the following flags from command line:
	// A. CSV file containing the problems
	fileNameFlag := flag.String("filename", "problems.csv", "-filename problems.csv")

	flag.Parse()
	fmt.Println("Opening file:", *fileNameFlag)

	// Open the file for reading
	file, err := os.Open(*fileNameFlag)
	if err != nil {
		fmt.Printf("error opening %s\n", *fileNameFlag)
		log.Fatal(err)
	}

	// Parse the csv file
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error reading csv file")
		log.Fatal(err)
	}

	// Make a slice of problems
	problems := make([]Problem, len(lines))
	for i, line := range lines {
		problems[i].question = line[0]
		problems[i].answer = line[1]
	}

	var answer string
	var score int
	for i, problem := range problems {
		fmt.Printf("Question #%d: %s =\n", i+1, problem.question)
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			fmt.Println("correct!")
			score++
		}
	}

	fmt.Printf("You answered %d out of %d correctly.", score, len(problems))
}
