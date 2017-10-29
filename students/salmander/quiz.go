package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	// Read the following flags from command line:
	// A. CSV file containing the problems
	fileNameFlag := flag.String("filename", "problems.csv", "CSV file in 'question,answer' format")

	// B. Flag for max time in seconds
	timeout := flag.Int("limit", 10, "time in seconds")
	flag.Parse()

	file := openFile(fileNameFlag)

	fmt.Printf("You have %d seconds to answer all the questions.\n", *timeout)

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

	timer := time.NewTimer(time.Duration(*timeout) * time.Second)

	var score int
	answerChan := make(chan string)
label:
	for i, problem := range problems {
		go func() {
			fmt.Printf("Question #%d: %s = ", i+1, problem.question)
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break label
		case answer := <-answerChan:
			if answer == problem.answer {
				fmt.Println("correct!")
				score++
			}
		}
	}

	fmt.Printf("You answered %d out of %d correctly.", score, len(problems))
}
func openFile(fileNameFlag *string) *os.File {
	fmt.Println("Opening file:", *fileNameFlag)
	// Open the file for reading
	file, err := os.Open(*fileNameFlag)
	if err != nil {
		fmt.Printf("error opening %s\n", *fileNameFlag)
		log.Fatal(err)
	}
	return file
}
