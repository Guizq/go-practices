package main

import "fmt"

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	nums := make([]int, 10)
	tgt := 8
	for i, _ := range nums {
		nums[i] = i * 2
	}
	fmt.Println(search(nums, tgt))
}

func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	for left <= right {

		mid := left + (right-left)>>1

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
