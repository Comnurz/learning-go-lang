package main

import "fmt"

func main() {
	j := 0

	for j <= 5 {
		fmt.Println(j)
		j += 1
	}

	for i := 0; i <= 10; i++ {
		if i%2 != 0 && i < 5 {
			fmt.Println(i, " is odd and less then 5")
		} else if i%2 == 0 && i < 5 {
			fmt.Println(i, " is even and less than 5")
		} else {
			fmt.Println(i, " is greater than 5")
		}
	}
}
