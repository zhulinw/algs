package leetcode

import (
	"testing"

	"github.com/gruntpig/algs/testdata"
)

// 剑指 Offer 10- I. 斐波那契数列
// 写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
// F(0) = 0,   F(1) = 1
// F(N) = F(N - 1) + F(N - 2), 其中 N > 1.

// 斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
func fibRemainder(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return (fibRemainder(n-1) + fibRemainder(n-2)) % 1000000007
}

func fibRemainderDp(n int) int {
	dp := make([]int, n+1)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}

func Test_FibRemainder(t *testing.T) {
	input, except := testdata.ClimbFibTestData()
	for i := range input {
		got := fibRemainderDp(input[i])
		if got != except[i] {
			t.Fatalf("want: %v, but got: %v, test data: %v", except[i], got, input[i])
		}
	}
}

// 509. 斐波那契数
// 斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。
// 该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
// F(0) = 0，F(1) = 1
// F(n) = F(n - 1) + F(n - 2)，其中 n > 1
// 给你 n ，请计算 F(n) 。
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
