package HW4

import (
	"fmt"
	"math/rand"
)

func createRandomMarksSet(amount int) []float32 {
	var marks []float32
	for i := 0; i < amount; i++ {
		marks = append(marks, float32(rand.Intn(5)+1))
	}
	fmt.Printf("Marks = %v \n", marks)
	return marks
}

func Mark(amount int) {
	mark := createRandomMarksSet(amount)
	var summ float32
	for i := 0; i < len(mark); i++ {
		summ = summ + mark[i]
	}
	fmt.Printf("Average mark = %v \n", summ/float32(len(mark)))
}
