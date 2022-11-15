package main

import "fmt"

//12321

func isPalindrome(num int) bool {
	var arr []int
	tmp := num
	for {
		arr = append(arr, tmp%10)
		tmp = tmp / 10
		if tmp == 0 {
			break
		}
	}
	len := len(arr)

	//4/2= 2
	//5/2
	// 1 3 4 77
	//
	//0,1
	//1,3,
	//12321
	for index, val := range arr {

		if (index + 1) == (len/2 + 1) {
			break
		}

		if val != arr[len-index-1] {
			return false
		}
	}
	return true
}

func main() {
	//fmt.Println(isPalindrome(12321))
	//fmt.Println(isPalindrome(1232))

	fmt.Println(getFibonacci(1))
}

//n!
//5 4 3 2 1

func getFactorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * getFactorial(num-1)
}

//0 1, 1, 2, 3, 5, 8,

//nth num fib seq

func getFibonacci(num int) int {

	if num <= 1 {
		return num
	}

	return getFibonacci(num-2) + getFibonacci(num-1)
}
