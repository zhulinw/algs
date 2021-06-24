package leetcode

import (
	"testing"

	"github.com/gruntpig/algs/leetcode/testdata"
)

// 70. 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 注意：给定 n 是一个正整数。

// 输入： 2
// 输出： 2
// 解释： 有两种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶
// 2.  2 阶

// 知识：斐波那契数列 指的是这样一个数列：0、1、1、2、3、5、8、13、21、34

// 思路：利用上次的求和计算的结果
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	n1 := 1
	n2 := 2
	for i := 3; i <= n; i++ {
		temp := n1 + n2
		n1 = n2
		n2 = temp
	}
	return n2
}

// 思路：递归思路
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs2(n-1) + climbStairs2(n-2)
}

// 思路：DP，动态规划
func climbStairs3(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

func Test_ClimbStairs(t *testing.T) {
	input, except := testdata.ClimbStairsTestData()
	for i := range input {
		got := climbStairs3(input[i])
		if got != except[i] {
			t.Fatalf("want: %v, but got: %v, test data: %v", except[i], got, input[i])
		}
	}
}
