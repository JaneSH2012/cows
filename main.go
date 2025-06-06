package main

import (
	"fmt"
)

func main() {
	foodPerCow := 2

	for cowNumber := 10; cowNumber <= 39; cowNumber++ {
		food := cowNumber * foodPerCow
		cowWord := "коров"
		if cowNumber%10 == 1 && cowNumber != 11 {
			cowWord = "коровы"
		}
		fmt.Println("Стадо из", cowNumber, cowWord, "съедает", food, "кг еды")

	}
}
