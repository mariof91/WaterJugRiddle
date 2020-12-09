package main

import (
	"fmt"
	"strconv"
)

func main() {

	bucket1:= getUserInput("please enter the size of the first bucket:")
	bucket2:= getUserInput("please enter the size of the second bucket:")
	waterGalleons:= getUserInput("please enter the galleons of water that you need:")

	if bucket2 > bucket1 {
		temp:= bucket2
		bucket2 = bucket1
		bucket1 = temp
	}

	if waterGalleons > bucket1 {
		fmt.Println("No Solution")
		return
	}
	fmt.Println(getGreaterCommonDenominator(bucket1,bucket2))
	if waterGalleons % getGreaterCommonDenominator(bucket1,bucket2) != 0{
		fmt.Println("No Solution")
		return
	}

	pourWater(bucket1,bucket2,waterGalleons)
}




func getGreaterCommonDenominator (bucket1 int, bucket2 int)int{
	if bucket2 == 0{
		return bucket1
	}
	return getGreaterCommonDenominator(bucket2, bucket1 % bucket2)
}


func validateInput(number int) bool{
	return !(number>0)
}


func getUserInput(message string) int{
	var stringTemp string
	var inputNumber int
	fmt.Println(message)
	for {
		_, err := fmt.Scan(&stringTemp)
		inputNumber, err = strconv.Atoi(stringTemp)
		if err != nil || validateInput(inputNumber) {
			fmt.Println("Enter a valid number greater than 0")
		} else {
			return inputNumber
		}
	}
}


func MinOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}


func pourWater(bucket1 int, bucket2 int, destination int) {

	from := bucket1
	to:= 0

	fmt.Printf("filled %v galleon bucket\n\n", bucket1)

	showCurrentStatus(bucket1,bucket2,from,to)

	for from!= destination && to != destination {

		temp:= MinOf(from, bucket2 - to)

		to += temp
		from -=temp

		fmt.Printf("Poured from %v galleon bucket to %v galleon bucket\n\n", bucket1, bucket2)
		showCurrentStatus(bucket1,bucket2,from,to)


		if from == destination || to == destination {
			break
		}

		if from == 0{
			from = bucket1
			fmt.Printf("Filled %v galleon bucket\n\n", bucket1)
			showCurrentStatus(bucket1,bucket2,from,to)


		}

		if to == bucket2 {
			to = 0
			fmt.Printf("Emptied %v galleon bucket\n\n", bucket2)
			showCurrentStatus(bucket1,bucket2,from,to)
		}

	}

	fmt.Printf("Got the %v of galleons required", destination)


}

func showCurrentStatus (bucket1,bucket2,from,to int){
	fmt.Printf("Current status is : \n" )
	fmt.Printf("bucket of %v galleon has %v \n",bucket1, from )
	fmt.Printf("bucket of %v galleon has %v galleons \n\n", bucket2, to )
}