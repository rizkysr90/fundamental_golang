package main

import (
	"fmt"
)

/*
	error is an interface
	type error interface {
		Error() string
	}
*/
type DivideError struct {
	dividend float64
}

func (d DivideError) Error() string {
	return fmt.Sprintf("can not divide %v by Zero", d.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, DivideError{dividend : dividend}
		// return 0, errors.New("no dividing by 0")
	}
	return dividend / divisor, nil
}
func main() {
	res, err := divide(2,1);
	if err != nil {
		// do something
		fmt.Println("There is an error")
	}
	fmt.Println(res);


}