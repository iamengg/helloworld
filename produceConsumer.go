/*

consumer x rate msg/sec
producer y rate msg/sec

	control producing trigger
	producer want to keep on producing
		once channel is full
			then stop until there is space for new values
		Once get signalled then start producing again
		***How i know if it's full?

			Rather than just starting to write even if available space is small
		***Signal to start after 'X' space is available

	Consumer starts consuming
		if no message
			***Sleep ________How i know how long to sleep?
				OR
			***Wait for signal from producer to start consuming
		else{
			process it
		}

*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const DATA_BUFFER = 100
const CONSUMER_SLEEP = 1 //in seconds
const PRODUCER_SLEEP = 1 //in seconds
const PRODUCER_TRIGGER_PERCENTAGE = 30

var msgCounter int

func producer(chDataBuffer chan<- int, chSignal <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	dataBufferIsFull := true
	var num int

	go func() {
		for {
			select {
			case <-chSignal:
				//fmt.Println("Started producing, as ", msg)
				dataBufferIsFull = false
			default:
				time.Sleep(time.Second * PRODUCER_SLEEP)
			}
		}
	}()

	for {
		num = rand.Intn(10000)
		chDataBuffer <- num
		msgCounter++
		if msgCounter%DATA_BUFFER == 0 {
			fmt.Println("Stopped producing as buffer is full")
			dataBufferIsFull = true
			for dataBufferIsFull {
				time.Sleep(time.Second * PRODUCER_SLEEP)
			}
		}

	}
}

func consumer(chDataBuffer <-chan int, chSignal chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	processedMsg := 0
	prevProcessedCnt := 0
	for {
		select {
		case msg := <-chDataBuffer:
			fmt.Println("Consumed msgCnt::", processedMsg, ",msg::", msg)
			processedMsg++
			if processedMsg-prevProcessedCnt > PRODUCER_TRIGGER_PERCENTAGE {
				//space is present for > PRODUCER_TRIGGER_PERCENTAGE messages
				chSignal <- "consumer: DATA_BUFFER is having space, start producing"
				prevProcessedCnt = processedMsg
			}
		default:
			prevProcessedCnt = processedMsg
			time.Sleep(time.Second * CONSUMER_SLEEP)
			chSignal <- "consumer: DATA_BUFFER is having space, start producing"

		}
	}
}

func main() {
	msgCounter = 0

	chDataBuffer := make(chan int, DATA_BUFFER)
	chSignal := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)
	go producer(chDataBuffer, chSignal, &wg)
	go consumer(chDataBuffer, chSignal, &wg)

	wg.Wait()
}
