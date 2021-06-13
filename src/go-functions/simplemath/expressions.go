package simplemath

import (
	"errors"
	"math"
)

func Sum(values...float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}

	return total
}

func Add(p1,p2 float64) float64 {
	return p1 + p2
}

func Subject(p1,p2 float64) float64 {
	return p1 - p2
}

func Multiply(p1,p2 float64) float64 {
	return p1 * p2
}

func Divide(p1,p2 float64) (float64,error) {
	if p2 == 0 {
		return math.NaN(), errors.New("can't divide by 0")
	}

	return p1 / p2, nil
}

func Divide2(p1,p2 float64) (answer float64, err error) {
	if p2 == 0 {
		err = errors.New("can't divide by 0")
	}

	answer = p1 / p2
	return
}