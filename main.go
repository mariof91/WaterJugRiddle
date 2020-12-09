package main

import (
	"fmt"
	"strconv"
)

func main() {

	bucket1:= getUserData("please enter the size of the first bucket:")
	bucket2:= getUserData("please enter the size of the second bucket:")
	waterGalleons:= getUserData("please enter the galleons of water that you need:")

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


func getUserData(message string) int{
	var s string
	var i int
	fmt.Println(message)
	for {
		_, err := fmt.Scan(&s)
		i, err = strconv.Atoi(s)
		if err != nil || validateInput(i) {
			fmt.Println("Enter a valid number greater than 0")
		} else {
			return i
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

	fmt.Println("Filled bucket 1")

	for from!= destination && to != destination {

		temp:= MinOf(from, bucket2 - to)

		to += temp
		from -=temp

		fmt.Println("Poured first bucket into second bucket")

		if bucket1 == destination || to == destination {
			break
		}

		if from == 0{
			from = bucket1
			fmt.Println("Filled bucket 1 again")
		}

		if to == bucket2 {
			to = 0
			fmt.Println("Emptied second bucket")
		}

	}

	fmt.Printf("Got the %v of galleons required", destination)


}