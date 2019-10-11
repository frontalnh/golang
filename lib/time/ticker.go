package LibTime

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("time flies!")

			case <-done:
				fmt.Println("it's done!")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	done <- true
}
