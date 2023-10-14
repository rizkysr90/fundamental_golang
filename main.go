package main

import "fmt"

func fizzbuzz() {
	for i:=1; i<=100; i++ {
		if i % 3 == 0 && i % 5 != 0 {
			fmt.Printf("%v is fizz\n", i)
		} else if  i % 5 == 0 && i % 3 != 0 {
			fmt.Printf("%v is buzz\n", i)
		} else if i % 5 == 0 && i % 3 == 0  {
			fmt.Printf("%v is fizzbuzz\n", i)
		} else {
			fmt.Println(i)
		}
	}

}
func findPrimeNumber(max int) {
	for i := 2; i < max; i++ {
		if i == 2 {
			fmt.Println(i)
			continue
		}
		if i % 2 == 0 {
			continue
		}
		isPrime := true
		for j := 3; j * j < i+1; j++ {
			if i % j == 0 {
				isPrime =  false
				break
			}
		}
		if !isPrime {
			continue
		}
		fmt.Println(i)
	}
}
func whileLoop() {
	treshold := 10
	for treshold < 20 {
		treshold++
	}
	fmt.Println("While loop", treshold)
}
func main() {
	fizzbuzz()
	whileLoop()
	findPrimeNumber(10)
}