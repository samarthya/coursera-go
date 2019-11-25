package main

import  (
	"fmt"
	"bufio"
	"strings"
	"sort"
	"os"
	"log"
	"strconv"
	"sync"
)

func sortArray(nums [] int, wg  *sync.WaitGroup) {
	fmt.Printf("> %v\n", nums)
	sort.Ints(nums)
	fmt.Printf(">> %v\n", nums)
	wg.Done()
}

/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. 
Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that 
it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
*/


func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Numbers (Press enter when done) >")
	line, err := in.ReadString('\n')
	sizeLen := 0
	var wg sync.WaitGroup
	
	if err != nil {
		log.Fatal(err)
	}

	line = strings.Trim(line, "\n")
	strs := strings.Split(line, " ")

	nums := make([]int, len(strs))

	for i, str := range strs {
		nums[i], _ = strconv.Atoi(str)
	}

	if len(nums) <= 4 {
		sort.Ints(nums)
	} else {
		fmt.Printf(" Len : %d\n", len(nums))

		sizeLen = (len(nums)/4)
		wg.Add(4)

		if sizeLen > 0 {
			x := 0
			y := sizeLen
			z := 0
			
			for  z < 4 {

				go sortArray(nums[x : y], &wg)
				
				x = y
				y += sizeLen
				z++
				
				if( y > len(nums) || z > 2) {
					y = len(nums)
				}
			}
		}

		wg.Wait()
		sort.Ints(nums)
	}
		
	

	fmt.Print("the sorted array is: ")
	fmt.Println(nums)

}