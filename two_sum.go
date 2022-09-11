package main

import(
	"fmt"
)

func main() {
	inputNums := []int{2,7,11,15}
	targetValue := 9
	results := twoSum(inputNums, targetValue)
	fmt.Println("results : ", results)
}

func twoSum(nums []int, target int) []int{
	
	testMap := make(map[int]int)

	for index := 0; index < len(nums); index++ {
		if otherIndex, ok := testMap[target - nums[index]]; ok {
			results := []int{index, otherIndex}
			return results
		}
		testMap[nums[index]] = index
	}

	results := []int{-1, -1}
	return results
}
