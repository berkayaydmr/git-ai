package utils

import (
	"fmt"
	"time"
)

func LoadingWithDots(start string, done <-chan struct{}) {
	fmt.Print(start)
	dotCount := 0
	for {
		select {
		case <-done:
			return
		default:
			fmt.Print(".")
			dotCount++
			if dotCount == 4 {
				dotCount = 0
				fmt.Printf("\r%s", start)
			}

			time.Sleep(time.Second )
		}
	}
}
