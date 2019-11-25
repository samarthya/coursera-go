package main
import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomNumber(min, max int) int {
	return rand.Intn(max - min + 1) + min
}

func PopulateSlice(numberRange, min, max int) []int {
	var slice = make([]int, 0, 0)

	for i:= 0; i < numberRange; i++ {
		slice = append(slice, GetRandomNumber(min, max))
	}

	return slice
}

func DivideIntoChunks(originalSlice []int, chunkSize int) [][]int{
	var chunk []int

	chunks := make([][]int, 0, len(originalSlice)/chunkSize + 1)
	for len(originalSlice) >= chunkSize {
		chunk, originalSlice = originalSlice[:chunkSize], originalSlice[chunkSize:]
		chunks = append(chunks, chunk)
	}
	if len(originalSlice) > 0 {
		chunks = append(chunks, originalSlice[:len(originalSlice)])
	}
	return chunks

}

func BubbleSort(rawSlice []int) [] int{
	arrayLength := len(rawSlice)

	for i:= 0; i < arrayLength; i++ {
		for j:=0; j < arrayLength - i - 1; j++ {
			if rawSlice[j] > rawSlice[j + 1] {
				Swap(rawSlice, j)
			}
		} 
	}

	return rawSlice
}

func ConcurrentSort(rawSlice []int, channel chan []int) {
	channel <- BubbleSort(rawSlice)
}

func Swap(rawSlice []int, index int) {
	rawSlice[index], rawSlice[index+1] = rawSlice[index+1], rawSlice[index]
}

func SortChunks(chunks [][]int, chunkSize int, channel chan []int) []int {
	var sliceToSort = make([]int, 0, 0)
	for i:= 0; i < 4; i++ {	
		go ConcurrentSort(chunks[i], channel)
	}

	for i:= 0; i < 4; i++ {	
		x := <- channel
		fmt.Println("++++++++++++++++++")
		fmt.Println("Sorted chunk")
		fmt.Println(x)
		fmt.Println("-----------------")
		for j:= 0; j < len(x) - 1; j++ {
			sliceToSort = append(sliceToSort, x[j])
		}
	}

	return BubbleSort(sliceToSort)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	channel :=  make(chan []int)
	min := 12
	max := 300
	numberRange := GetRandomNumber(min, max)
	slice := PopulateSlice(numberRange, min, max)
	chunkSize := len(slice) / 4

	chunks := DivideIntoChunks(slice, chunkSize)
	fmt.Println("*****************")
	fmt.Println("FINAL RESULT")
	fmt.Println(SortChunks(chunks, chunkSize, channel))
	fmt.Println("*****************")

}