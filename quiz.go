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

//var wg = &sync.WaitGroup{}
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

	/*fmt.Println("Now Scanning:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
	    fmt.Println(scanner.Text())
	}*/

	for i, rec := range records {
		fmt.Printf("%d. What is the answer to %s?\n", i+1, rec[0])
		var input string
		fmt.Scanln(&input)
		//var correctAnswers int

		correctAnswers := 0
		if input == rec[1] {
			correctAnswers = correctAnswers + 1
			fmt.Println("you got", correctAnswers, "correct")

		}

	}

}

/*
	// read the csv values//
	records := make( chan []string)
	go func() {
//		defer close(records)
		csvReader := csv.NewReader(f)
		for i := 0; ; i = i + 1 {
			rec, err := csvReader.Read()
		records <- rec
		}
	wg.Wait()
	printRecords(records)
	defer close(records)
}();
}
func printRecords( records chan []string ) {
	fmt.Println("finished print_records")
	defer wg.Done()
*/ //}
