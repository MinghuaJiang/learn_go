package main

import (
	"fmt"
	"go-functions/simplemath"
)

func main() {
	answer, err := simplemath.Divide(2, 0)
	if err != nil {
	    fmt.Printf("An error occured %s\n", err.Error())
	}else{
		fmt.Printf("%f\n", answer)
	}

	answer2, err2 := simplemath.Divide2(2, 0)
	if err != nil {
		fmt.Printf("An error occured %s\n", err2.Error())
	}else{
		fmt.Printf("%f\n", answer2)
	}

	total := simplemath.Sum(1.0, 5.0, 9.0)
	fmt.Printf("total of sum %f\n", total)
	numbers := []float64{2.0, 3.0, 4.0}
	total2 := simplemath.Sum(numbers...)
	fmt.Printf("total of sum %f\n", total2)

	sv := simplemath.NewSemanticVersion(1,2,3)
	sv.IncrementMajor()
	fmt.Println(sv.String())
}


