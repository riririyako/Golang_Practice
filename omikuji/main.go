package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d回目のおみくじ結果は", i)
		number := rand.Intn(6)
		switch number {
		case 0:
			fmt.Println("凶です")
		case 1,2:
			fmt.Println("吉です")
		case 3,4:
			fmt.Println("中吉です")
		case 5:
			fmt.Println("大吉です")
		}
	}
}
