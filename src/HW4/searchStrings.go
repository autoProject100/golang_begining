package HW4

import (
	"fmt"
)

const SEARCHINGSTRING = "ok"

func MultipleStringsInput() {
	var stringLine string
	var allStringLines []string
	fmt.Println("Write any strings till word 'stop' - after that input will be over")
	for stringLine != "stop" {
		fmt.Scan(&stringLine)
		if stringLine != "stop" {
			allStringLines = append(allStringLines, stringLine)
		}
	}

	amountOfMatches := 0
	for i := 0; i < len(allStringLines); i++ {
		amountOfMatches += searchForText(allStringLines[i], SEARCHINGSTRING)
	}
	fmt.Printf("Matches = %v \n", amountOfMatches)
	//fmt.Fscan(os.Stdin, &stringLine)
	//fmt.Scanln(&stringLine)
	//fmt.Printf("%v\n", stringLine)
	//for i, v := range allStringLines {
	//	fmt.Printf("%v = %v\n", i, v)
	//}
}

func searchForText(line string, searchingString string) int {
	counterOfMatches := 0
	firstEdgeCounter := 0
	lineLength := len(line)
	searchingStringLength := len(searchingString)
	//fmt.Printf("%v -- %v \n", lineLength, searchingStringLength)
	for i := searchingStringLength; i <= lineLength; i++ {
		//fmt.Printf("%v -- %v \n", line[firstEdgeCounter:i], searchingString)
		if line[firstEdgeCounter:i] == searchingString {
			counterOfMatches++
		}
		firstEdgeCounter++
	}
	return counterOfMatches
}
