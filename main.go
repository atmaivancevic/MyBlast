package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Example Usage: MyBlast file.fasta txid[5] blastn 10 7 4,-5 12,8
func main() {

	queryFileName := os.Args[1]

	organism := os.Args[2]

	program := os.Args[3]

	expectThreshold, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		panic(err)
	}

	wordSize, err := strconv.ParseInt(os.Args[5], 0, 64)
	if err != nil {
		panic(err)
	}

	matchMismatchString := os.Args[6]
	matchScore, err := strconv.ParseInt((strings.Split(matchMismatchString, ","))[0], 0, 64)
	if err != nil {
		panic(err)
	}
	mismatchScore, err := strconv.ParseInt((strings.Split(matchMismatchString, ","))[1], 0, 64)
	if err != nil {
		panic(err)
	}

	gapCostString := os.Args[7]
	gapCostExistence, err := strconv.ParseInt((strings.Split(gapCostString, ","))[0], 0, 64)
	if err != nil {
		panic(err)
	}
	gapCostExtension, err := strconv.ParseInt((strings.Split(gapCostString, ","))[1], 0, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Query File Name: " + queryFileName)
	fmt.Println("Organism: " + organism)
	fmt.Println("BLAST Program: " + program)
	fmt.Printf("Expect Threshold: %f\n", expectThreshold)
	fmt.Printf("Word Size: %d\n", wordSize)
	fmt.Printf("Match Score: %d\n", matchScore)
	fmt.Printf("Mismatch Score: %d\n", mismatchScore)
	fmt.Printf("Gap Cost - Existence: %d\n", gapCostExistence)
	fmt.Printf("Gap COst - Extension: %d\n", gapCostExtension)

}
