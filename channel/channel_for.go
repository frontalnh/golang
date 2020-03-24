package main

import "time"

func main() {
	messages := make(chan int)

	go func() {
		messages <- 1
		messages <- 1
		messages <- 1
	}()
	// close(messages)

	// for o := range messages {
	// 	fmt.Println(o)
	// }
	time.Sleep(time.Second)
	<-messages
	<-messages
	<-messages
}
