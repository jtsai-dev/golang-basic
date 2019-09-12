package main

import "testing"

/* 测试文件以 _test 结尾，测试方法以 Test 开头
 * go tool cover (run package tests) 可获取测试结果，包括是否通过、代码覆盖率等
 */

// 定义测试所需的case
var tests = []struct{ a, b, c int }{
	{1, 2, 3},
	{2, 3, 5},
	{3, 4, 7},
}

func TestSum(t *testing.T) {
	for _, tt := range tests {
		// 调用方法，进行断言
		if actual := sum(tt.a, tt.b); actual != tt.c {
			t.Errorf("%d + %d = %d, but got %d", tt.a, tt.b, tt.c, actual)
		}
	}
}

/* 性能测试：方法以 Benchmark 开头
 * 测试结果：BenchmarkSum-4     	2000000000   0.36 ns/op	        0 B/op	       		 0 allocs/op
 * 						 4个cpu	   运行次数		 平均每次耗时(纳秒)	 每次执行分配的内存		每次执行分配了0次对象
 */
func BenchmarkSum(b *testing.B) {
	numA, numB := 1, 2
	r := 3

	for i := 0; i < b.N; i++ {
		actual := sum(numA, numB)
		if actual != r {
			b.Errorf("%d + %d = %d, but got %d", numA, numB, r, actual)
		}
	}
}
