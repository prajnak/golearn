package main

import "fmt"

// Channels are pipes that connect concurrent goroutines Send values into
// channels from one goroutine and receive them in another goroutine

// Channels are unbuffered by default. This means that a channel will only
// accept (chan <-) is there is a corresponding receiver (<-chan) ready to
// receive the sent value.

// Buffered channels accept a limited number of values without a corresponding
// receiver for those values

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total //send the total to given channel
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func main() {

	// use the 'make' keyword to create a channel of string type.
	messages := make(chan string)

	messages2 := make(chan string, 2)
	messages2 <- "buffered"
	messages2 <- "channel"

	fmt.Println(<-messages2)
	fmt.Println(<-messages2)
	// send a value to the channel using the 'channel<-' syntax. Here, we are sending a
	// message 'ping' to the channel messages we created earlier. We're doing
	// this from a new goroutine that runs concurrently with the main()
	// goroutine

	// receive the message sent from the other goroutine using the '<-channel' syntax
	go func() { messages <- "ping" }()
	msg := <-messages

	// sends and receives block until both sender and receiver are ready. This
	// is why this program waited at the end for the "ping" message without
	// having to use any other synchronization
	fmt.Println(msg)

	a := []int{9, 3, 4, -9, 2, 3}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c //receive from c

	fmt.Println(x, y, x+y)

	// use a buffered channel
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c1)

	// this range call will not stop reading data from a channel until the
	// channel is closed. The channel is closed with the close keyword at the
	// end of the fibonacci method call. We can use the
	// v, ok := <- ch call to test whether a channel is closed
	for i := range c1 {
		fmt.Println(i)
	}

}
