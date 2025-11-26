package Utils

import (
	"fmt"
	"time"
)

func PrintWithDelay(text string) {
	var delay time.Duration = 20 * time.Millisecond
	for _, char := range text {
		fmt.Printf(string(char))
		time.Sleep(delay)
	}
	fmt.Println()
}

func Separate_block() {
	fmt.Println("_______")
}
