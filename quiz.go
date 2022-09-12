package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func quizRun() string {
	//open the file//
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	// always close the file after//
	defer f.Close()
	csvReader := csv.NewReader(f)
	var records [][]string
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("An error occured ::", err)
		}
		records = append(records, rec)
	}

	correctAnswers := 0
	for i, rec := range records {
		fmt.Printf("%d. What is the answer to %s?\n", i+1, rec[0])
		var input string
		fmt.Scanln(&input)
		if input == rec[1] {
			correctAnswers++

		}

	}

	wrongAnswers := (len(records)) - correctAnswers
	totalQuestions := wrongAnswers + correctAnswers
	fmt.Println("You got", correctAnswers, "correct out of", totalQuestions, "total questions")
	return "Great job"
}
func main() {
	fmt.Println("Beginning quiz, press enter to start")
	fmt.Scanln()
	currentChannel := make(chan string, 1)

	timeLimit := flag.Duration("timeout", 30, "Define a time limit in ns")
	flag.Parse()
	if *timeLimit < time.Duration(0) {
		*timeLimit = time.Duration(30)
		}
	go func() {
		text := quizRun()
		currentChannel <- text
	}()

	select {
	case res := <-currentChannel:
		fmt.Println(res)
	case <-time.After(*timeLimit * time.Second):
		fmt.Println("You ran out of time! :(")
	}
}
