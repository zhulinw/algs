package chapter1

import (
	"fmt"
	"math"
)

// @Feature 典型数组处理
// @需求 找出数组元素最大的元素
func MaxEleInSlice(s []int) int {
	max := s[0]
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

// @Feature 典型数组处理
// @需求 计算数组元素的平均值
func AverageValueSlice(s []int) int {
	total := 0
	for i := range s {
		total = total + s[i]
	}
	return total / len(s)
}

// @Feature 典型数组处理
// @需求 复制数组
func CopySlice(s []int) []int {
	result := make([]int, 0, len(s))
	for i := range s {
		result[i] = s[i]
	}
	return result
}

// @Feature 典型数组处理
// @需求 颠倒数组元素的顺序
func ReverseSlice(s []int) []int {
	for i := 0; i < len(s)/2; i++ {
		tmp := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = tmp
	}
	return s
}

// @Feature 典型数组处理
// @需求 矩阵相乘（方阵） todo: next
func MatrixMulDoubleDimensional(s []int) []int {
	return nil
}

// @Feature 二分查找
// @需求
func BinarySearch(s []int, ele int) int {
	low := 0
	high := len(s)
	for low < high {
		mid := (low + high) / 2
		midValue := s[mid]
		if midValue == ele {
			return ele
		}
		if midValue > ele {
			high = mid
		}
		if midValue < ele {
			low = mid
		}
	}
	return math.MaxInt32 // not find
}

// @Feature 二分查找, 递归实现
// @需求
func RecurBinarySearch(s []int, low, high, ele int) int {
	fmt.Println("low:", low, "high:", high)
	mid := low + (high - low)/2

	if low >= high {
		return -1
	}

	if s[mid] > ele {
		return RecurBinarySearch(s, low, mid, ele)
	}else if s[mid] < ele {
		return RecurBinarySearch(s, mid, high, ele)
	}else {
		return ele
	}
}


