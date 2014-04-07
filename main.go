package main

import (
	"code.google.com/p/biogo.ncbi/blast"

	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	tool  = "blast-testing"
	email = "atma.ivancevic@gmail.com"
	retry = 15
)

// Example Usage: MyBlast file.fasta txid[5] blastn 10 7 4,-5
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

	// gapCostString := os.Args[7]
	// gapCostExistence, err := int(strconv.ParseInt((strings.Split(gapCostString, ","))[0], 0, 64))
	// if err != nil {
	// 	panic(err)
	// }
	// gapCostExtension, err := int(strconv.ParseInt((strings.Split(gapCostString, ","))[1], 0, 64))
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Query File Name: " + queryFileName)
	fmt.Println("Organism: " + organism)
	fmt.Println("BLAST Program: " + program)
	fmt.Printf("Expect Threshold: %f\n", expectThreshold)
	fmt.Printf("Word Size: %d\n", wordSize)
	fmt.Printf("Match Score: %d\n", matchScore)
	fmt.Printf("Mismatch Score: %d\n", mismatchScore)
	// fmt.Printf("Gap Cost - Existence: %d\n", gapCostExistence)
	// fmt.Printf("Gap Cost - Extension: %d\n", gapCostExtension)

	// Read the query sequence from file called queryFileName
	query := "gcgcaggcggctgcgacgggcgcgctaagcgcggccgagaggagccaccccacgtcgaggtcaggggcagaagccgggaggaccccatgcccgaagggcggcggccaagaggagttaccccacgtccgaggtcaggggcagcggccgagagtaccagactgcgacggcgcaggaacggccgagaggagctaccccgcgtccgaggtcagggggggcggccgagaggagatacccagcgtccgaggtcaggggcggcgacgagaggagttaccccgcgtccgaggtcaggggcggcggccgggaggagctaccccacgcccctaagcccgaggccaggggcggcggccgggaggagcaaccccacgcccgaggccaggggcggcggccgggaggaccaaccccacgtccaaggagccgtggctgcgcgggcacaggagggcctagaggagctatcccacgttgaaggtcaggaagggcggcggtgaggagatacccctcgtccaaggtaaggagcagcggctgcgctttgctggagcagccgtgaagagataccccacgcccaaggtaagagaaacccaagtaagacggtaggtgttgcaagagggcatcagagggcagacacactgaaaccatactcacagaaaactagtcaatctaatcacactaggaccacagccttgtctaactcaatgaaactaagccatgcccgtggggcaacccaagatgggcgggtcatggtggagagatctgacagaatgtggtccactggagaagggaatggcaaaccacttcagtattcttgccttgagaaccccatgaacagtatgaaaaggcaaaatgataggatactgaaagaggaactccccaggtcagtaggtgcccaatatgctactggagatcagtggagaaataactccagaaagaatgaagggatggagccaaagcaaaaacaatacccagctgtggatgtgactggtgatagaagcaaggtccgatgctgtaaagagcaatattgcataggaacctggaatgtcaggtccatgaatcaaggcaaattggaagtggtcaaacaagagatggcaagagtgaatgtcgacattctaggaatcagcgaactgaaatggactggaatgggtgaatttaactcagatgaccattatatctactactgcgggcaggaatccctcagaagaaatggagtggccatcatggtcaacaaaagagtccgaaatgcagtacttggatgcaatctcaaaaacgacagaatgatctctgttcgtttccaaggcaaaccattcaatatcacagtaatccaagtctatgccccaaccagtaacgctgaagaagctgaagttgaacggttctatgaagacctacaagaccttttagaactaacacccaaaaaagatgtccttttcattataggggactggaatgcaaaagtaggaagtcaagaaacacctggagtaacaggcaaatttggccttggaatacggaatgaagcagggcaaagactaatagagttttgccaagaaaatgcactggtcatagcaaacaccctcttccaacaacacaagagaagactctatacatggacatcaccagatgtcaacaccgaaatcagattgattatattctttgcagccaaagatggagaagctctatacagtcagcaaaaacaagaccaggagctgactgtggctcagaccatgaactccttattgccaaattcagacttaaattgaagaaagtagggaaaaccactagaccattcaggtatgacctaaatcaaatcccttatgattatacagtggaagtgagaaatagatttaagggcctagatctgatagatagagtgcctgatgaactatggaatgaggttcgtgacattgtacaggagacagggatcaagaccatccccatggaaaagaaatgcaaaaaagcaaaatggctgtctg"

	pp := &blast.PutParameters{
		Program:     program,
		Database:    "nr",
		EntrezQuery: "bos taurus[organism]",
		WordSize:    int(wordSize),
		Expect:      &expectThreshold,
		NuclReward:  int(matchScore),
		NuclPenalty: int(mismatchScore)}

	var gp *blast.GetParameters
	gp = nil

	r, err := blast.Put(query, pp, tool, email)
	if err != nil {
		panic(err)
	}

	var o *blast.Output
	for k := 0; k < retry; k++ {
		// Wait for RTOE to elapse and get search status.
		var s *blast.SearchInfo
		s, err = r.SearchInfo(tool, email)
		if err != nil {
			panic(err)
		}

		// Output search status.
		fmt.Println(s)

		switch s.Status {
		case "WAITING":
			continue
		case "FAILED":
			fmt.Printf("search: %s failed", r)
			return
		case "UNKNOWN":
			fmt.Printf("search: %s expired", r)
			return
		case "READY":
			if !s.HaveHits {
				fmt.Printf("search: %s no hits", r)
				return
			}
		default:
			fmt.Printf("An unknown error occurred.")
			return
		}

		// We have hits, so get the BLAST output.
		o, err = r.GetOutput(gp, tool, email)
		if err != nil {
			panic(err)
		}
		fmt.Println(o.Program)
		return
	}

	fmt.Printf("%s exceeded retries", r)
	return

}
