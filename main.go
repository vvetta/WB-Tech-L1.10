package main

import (
	"fmt"
)

func main() {	
	data := []float32{-25.4, -27, 13, 19, 15.5, 24.5, -21, 32.5}
	
	tempMap := make(map[int][]float32)

	for _, temp := range data {
		key := int(getKey(temp))
		tempMap[key] = append(tempMap[key], temp)
	}

	fmt.Println(tempMap)
}

func getKey(number float32) float32 {
	var keyOne float32
	var keyTwo float32
	newNumber := number
	
	signFlag := true

	if number < 0 {
		signFlag = false
		newNumber *= -1
	}

	keyBool := false


	for {		
		if newNumber <= keyOne && newNumber >= keyTwo {
			if signFlag == false {
				return (keyOne - 10) * -1
			}	
			return keyOne - 10
		}

		if keyBool == true {
			keyTwo += 10
			keyOne += 10
			keyBool = false
			continue
		}

		keyOne += 10
		keyBool = true
	}
}
