package main

import "fmt"

// go에서 부동 소수점 타입은 (허수 제외) float32, float64 두 가지
// 부동소수점은 정수 분수를 명확히 표현할 수 없어 근사치를 구하는 것
// 근사 연산의 영향과 어떻게 정확도를 높이는지 알아보자
func main() {
	var n float32 = 1.0001
	fmt.Println(n * n)

	var a float64
	positiveInf := 1 / a
	negativeInf := -1 / a
	nan := a / a
	fmt.Println(positiveInf, negativeInf, nan)
}

func f1(n int) float64 {
	result := 10_000.
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result
}

func f2(n int) float64 {
	result := 0.
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result + 10_000.
}
