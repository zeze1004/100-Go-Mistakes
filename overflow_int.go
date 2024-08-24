package main

import (
	"fmt"
	"math"
	"strconv"
)

/* go에서 제공하는 정수타입은 모두 10가지
부호있는 정수: int8, int16, int32, int64
부호없는 정수: uint8, uint16, uint32, uint64
나머지 정수 타입은 흔하게 사용하는 int와 uint -> 두 타입은 시스템에 따라 32/64비트 크기인지 정해짐
*/

/*
	오버플로우에 대해 알아보자

int32를 최댓값으로 초기화한 뒤, 값을 증가시키면 어떤 결과가 나올까?
int32 최댓값 -> 0b01111111111111111111111111111111 (양수일 때는 부호를 표시하는 왼쪽 끝 비트가 0, 31비트가 1) -> 2147483647
int32 최댓값 + 1 -> 0b10000000000000000000000000000000 (32비트로 표현할 수 있는 부호있는 정수 중에서 가장 작은 값이 됨) -> -2147483648
*/
func main() {
	var counter int32 = math.MaxInt32    // int32 최댓값: 2147483647
	counter++                            // -2147483648
	fmt.Printf("counter: %d\n", counter) // 오버플로우를 출력해도 에러가 아님!!!!

	// 그럼 오버플로우 체크를 어떻게할까?
	var counter2 int32 = math.MaxInt32 - 1
	if counter2 == math.MaxInt32 { // 오버플로우 체크
		panic("int overflow")
	}
	counter2++

	// 덧셈 연산에서 발생하는 정수 오버플로우 검사
	var counter3 = math.MaxInt - 1
	tmp := 1
	if counter3 > math.MaxInt-tmp {
		panic("int overflow")
	}
	fmt.Printf(strconv.Itoa(counter3 + tmp))
}

// MultiplyInt 곱셈에서 발생하는 정수 오버플로우 검사: 최솟값(math.MinInt)에 대해 검사
func MultiplyInt(a, b int) int {
	if a == 0 || b == 0 { // 둘 중 하나가 0이면 0을 리턴
		return 0
	}

	result := a * b
	if a == 1 || b == 1 { // 둘 중 하나가 1이면 a*b를 리턴
		return result
	}
	if a == math.MinInt || b == math.MinInt { // 둘 중 하나가 math.MinInt인지 검사
		panic("integer overflow")
	}
	if result/b != a { // a*b의 결과가 a로 나누어도 b가 되는지 검사
		panic("integer overflow")
	}

	return result
}
