package leetcode

import (
	"testing"

	"github.com/gruntpig/algs/testdata"
	"github.com/gruntpig/algs/utils"
)

// 55.跳跃游戏
// 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个下标。

// 输入：nums = [2,3,1,1,4]
// 输出：true
// 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

// 思路：每次更新能跳到的最远距离，
// 当前位置超过最远能跳到的距离时，返回false
func canJump(nums []int) bool {
	max := 0
	for i := range nums {
		if i > max { // 当前位置超过最远能跳到的位置
			return false
		}
		// 更新能跳到的最大的位置
		max = utils.MaxInt(max, i+nums[i])
	}

	return true
}

func Test_CanJump(t *testing.T) {
	input, except := testdata.JumpGameTestData()
	for i := range input {
		canJump := canJump(input[i])
		if canJump != except[i] {
			t.Fatalf("can't jump, got: %v, test data: %v", canJump, input[i])
		}
	}
}
