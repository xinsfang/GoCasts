/*
By default, go uses one core.

G represents goroutine. Contain stack pointer, base of stack, ID, cache, status
M represents OS thread. Contain a pointer to global Q of runnable goroutines, the current running goroutine and the reference to the Sched.
Sched is global struct. Contain Q free and waiting goroutines and threads.

On startup, go runtime starts a number of goroutines for GC, scheduler and user code.
An OS Thread is created to handled these goroutines. These threads can be at most equal to GOMAXPROCS.

- Start from the bottom.
A goroutine is created with initial only 2KB of stack size. Each function in go already has a check if more stack is needed or not.
The stack can be copied to another region in memory with twice the original size. This makes goroutine very light on resources.

- Blocking is fine
If a goroutine blocks on system call, it blocks it's running thread. But another thread is taken from the waiting queue of
Scheduler (the Sched struct) and used for other runnable goroutines.
However, if you communicate using channels in go which exists only in virtual space, the OS doesn’t block the thread.
Such goroutines simply go into the waiting state and other runnable goroutine (from the M struct) is scheduled in it’s place.

- Don't interrupt
The go runtime scheduler does cooperative scheduling, which means another goroutine will only be scheduled if the current one is blocking or done.

Some of these cases are:
Channel send and receive operations, if those operations would block.
The Go statement, although there is no guarantee that new goroutine will be scheduled immediately.
Blocking syscalls like file and network operations.
After being stopped for a garbage collection cycle.

This is better than pre-emptive scheduling which uses timely system interrupts (e.g. every 10 ms) to block and schedule
a new thread which may lead a task to take longer than needed to finish when number of threads increases or when a higher
priority tasks need to be scheduled while a lower priority task is running.
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

/*	for {
		time.Sleep(5 * time.Second)
		go checkLink(<-c, c) //similar effect to the following block but less readable
	}*/

	for l := range c { //it is a blocking channel
		go func(link string) { //function literal (lambda/anonymous function)
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan <- string) { //send channel
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
	close(c) //close channel so range over channels reach the end. You can only close a send channel.

}
