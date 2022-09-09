package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	//	"bufio"
	//	"sync"
)

func main() {
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
			return
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
	fmt.Println("you got", correctAnswers, "correct out of", totalQuestions, "total questions")
//	fmt.Println("you got", wrongAnswers, "wrong")
}
