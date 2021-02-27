package main

import "fmt"

func main() {

	x := []int{3, 2, 4}
	y := 6
	fmt.Println(twoSum(x, y))
}

func twoSum(nums []int, target int) []int {
	a := []int{0, 0}

	for i := 0; i < len(nums)-1; i = i + 1 {
		for j := i + 1; j < len(nums); j = j + 1 {
			fmt.Println(i, j, a, nums, target)

			if nums[i]+nums[j] == target {
				a[0] = i
				a[1] = j

				fmt.Println(i, j, a)
				return a
			}
		}
	}

	return a
}
