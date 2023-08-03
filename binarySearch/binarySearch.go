package binarySearch

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	low := 0
	high := len(nums)

	for low < high {
		middle := (low + high) / 2
		switch {
		case nums[middle] == target:
			return middle
		case nums[middle] > target:
			high = middle
		case nums[middle] < target:
			low = middle + 1
		}
	}

	return -1
}
