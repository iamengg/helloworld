/*
val - primitive,
	int, float, interface, wg, struct

by ref
	map, chan, slice ptr -

a:=[]int{1,2}
b:=[]int{2,3}
c:=a
a=b
appending 5 to b
a = 2,3
b 2,3,5
c = 1,2


stock [7,1,5,3,6,4]

	producer
		main routine

	consumer
		even
		odd


*/
package main

import (
	"fmt"
	"sync"
)

func produce(wg *sync.WaitGroup, data chan<- int) {
	wg.Add(1)

	for i := 1; i <= 100; i++ {
		data <- i
	}
	
	wg.Done()
	close(data)
}

func main() {

	data := make(chan int, 100)
	var wg sync.WaitGroup
	produce(&wg, data)
	
	consumeData(data)
	wg.Wait()
}

func evenPrint(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Even ", num)
}

func oddPrint(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Odd ", num)
}

func consumeData(data <-chan int) {
	var wg sync.WaitGroup

	for values := range data {
		wg.Add(1)
		if values%2 == 0 {
			go evenPrint(values, &wg)
		} else {
			go oddPrint(values, &wg)
		}
	}
	wg.Wait()

}

func mainPurchaseSale() {
	stock := []int{7, 1, 5, 3, 6, 14}
	//4 + 11
	//purchase
	//find selling price , sell
	//{}
	//14,6,

	/*7,1
	6
	5
	3
	11,2*/
	//repeat

	var purchase, profit int
	const MAX_PRICE = 999999
	purchase = MAX_PRICE
	profit = 0
	lowestPrice := 0
	//highPrice := 0
	totalDays := len(stock)
	for index, val := range stock {
		if val < purchase {
			purchase = val
			if purchase < lowestPrice {
				lowestPrice = purchase
			}
		} else {

			if val > purchase {
				for {
					if index >= totalDays-1 {
						break
					}
					if stock[index+1] > stock[index] {

						index++
					} else {
						break
					}
				}
				//highPrice = stock[index]
			}

			profit += stock[index] - purchase
			purchase = MAX_PRICE
		}
	}

	fmt.Println("Profit is ", profit)

}
