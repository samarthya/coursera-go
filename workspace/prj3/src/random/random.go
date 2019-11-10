package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func swap(nums []int, j int) {
	var temp = nums[j]
	nums[j] = nums[j+1]
	nums[j+1] = temp
}
func bubblesort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j)
			}
		}
	}
}
func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("enter numbers :")
	line, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	strs := strings.Split(line[0:len(line)-1], " ")
	nums := make([]int, len(strs))
	for i, str := range strs {
		nums[i], _ = strconv.Atoi(str)
	}
	bubblesort(nums)
	fmt.Print("the sorted array is: ")
	fmt.Println(nums)
}
