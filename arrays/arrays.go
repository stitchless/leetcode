package arrays

import (
	"fmt"
	"math"
	"sort"
	"time"
)

// FindMaxConsecutiveOnes takes in a list and find the maximum number of consecutive
// digits within the list, and returns that number.
func FindMaxConsecutiveOnes(input []int) int {
	highestCnt := 0
	currentCnt := 0
	for _, numb := range input {
		if numb == 1 {
			currentCnt++
			if currentCnt > highestCnt {
				highestCnt = currentCnt
			}
		} else {
			currentCnt = 0
		}
	}

	return highestCnt
}

func FindNumbers(nums []int) int {
	outputCnt := 0
	for _, num := range nums {
		count := 0
		for num != 0 {
			num /= 10
			count += 1
		}

		if count%2 == 0 {
			outputCnt++
		}
	}

	return outputCnt
}

func SortedSquares(nums []int) []int {
	for i, num := range nums {
		nums[i] = num * num
	}

	sort.Ints(nums)

	return nums
}

func DuplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			arr[i] = 0
			i++
		}
	}
}

func Merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}

func RemoveElement(nums []int, val int) int {
	c := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, '_')
		} else {
			c++
		}
	}

	return c
}

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	c := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[c] = nums[i]
			c++
		}
	}

	return c
}

func CheckIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}

			if arr[i] == 2*arr[j] {
				return true
			}
		}
	}

	return false
}

func ValidMountainArray(arr []int) bool {
	if len(arr) <= 2 {
		return false
	}

	n := len(arr) - 1
	i := 0
	j := n
	for i < j {
		if arr[i] < arr[i+1] {
			i++
		} else {
			break
		}
	}
	for j > 0 {
		if arr[j] < arr[j-1] {
			j--
		} else {
			break
		}
	}

	return j == i && i != n && j != 0
}

func ReplaceElements(arr []int) []int {
	now := time.Now()
	maxI := -1
	for i := len(arr) - 1; i >= 0; i-- {
		ele := arr[i]
		arr[i] = maxI
		maxBool := math.Max(float64(maxI), float64(ele))
		maxI = int(maxBool)
	}

	fmt.Println(time.Since(now))
	return arr
}

func MoveZeroes(nums []int) {
	for i, j := 0, 1; j < len(nums); j++ {
		if nums[i] != 0 {
			i++
			continue
		}
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}

func SortArrayByParity(nums []int) []int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i]%2 == 0 {
			i++
			continue
		}

		if nums[j]%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	return nums
}

func HeightChecker(heights []int) int {
	tmp := make([]int, len(heights))
	copy(tmp, heights)
	sort.Ints(tmp)
	out := 0
	for i, num := range heights {
		if num != tmp[i] {
			out++
		}
	}

	return out
}

func ThirdMax(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	if nums[0] == nums[len(nums)-1] {
		return nums[0]
	}

	var tmp []int

	for i := len(nums) - 1; i >= 0; i-- {
		if len(tmp) == 0 {
			tmp = append(tmp, nums[i])
			continue
		}

		if tmp[len(tmp)-1] > nums[i] {
			tmp = append(tmp, nums[i])
			if len(tmp) == 3 {
				break
			}
		}
	}

	if len(tmp) == 2 {
		maxI := math.Max(float64(tmp[0]), float64(tmp[1]))
		return int(maxI)
	}

	return tmp[2]
}

func FindDisappearedNumbers(nums []int) []int {
	boolTmp := make([]bool, len(nums))
	for _, val := range nums {
		boolTmp[val-1] = true
	}

	var out []int
	for i, val := range boolTmp {
		if !val {
			out = append(out, i+1)
		}
	}

	return out
}
