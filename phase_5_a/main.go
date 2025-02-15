package main

import "fmt"

func main() {

	// 1 - 20
	// detect all even nos
	// detect all odd nos

	// input is coming from main logic

	// slice of numbers : numbers

	// function to add all the even nos: func SumEven : return the sum

	// function to add all the odd nos: func SumOdd : return the sum

	sliceNum := []int{2, 3, 4, 5, 6, 7, 8, 12}

	// fmt.Println("Sum of even nuumbers", sumEven(sliceNum))
	// fmt.Println("Sum of odd nuumbers", sumOdd(sliceNum))

	channelEvenSum := make(chan int)
	defer close(channelEvenSum)

	go func() {
		channelEvenSum <- sumEven(sliceNum)
	}()

	fmt.Println("Sum of even numbers:", <-channelEvenSum)

	channelOddSum := make(chan int)
	defer close(channelOddSum)

	go func() {
		channelOddSum <- sumOdd(sliceNum)
	}()

	fmt.Println("Sum of odd numbers:", <-channelOddSum)

}

// O(n)

func sumEven(nums []int) (sum int) {

	for _, x := range nums {
		// pick even numbers
		if x%2 == 0 {
			sum = sum + x
		}
	}
	return sum
}

func sumOdd(nums []int) (sum int) {
	for _, x := range nums {
		// pick odd numbers
		if x%2 != 0 {
			sum = sum + x
		}
	}
	return sum
}
