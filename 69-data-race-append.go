package main

import "fmt"

// append를 이용해 슬라이스에 원소를 추가할 때 데이터 경쟁이 발생할 수 있음
func main() {
	// # 예제1: 슬라이스 하나를 초기화한 뒤 고루티 두 개를 생성하고 고루틴은 각각 append로 슬라이스를 새로 만들어서 원소 추가함
	slice1 := make([]int, 1) // 길이가 1이고 용량이 1인 슬라이스 생성

	go func() {
		s1 := append(slice1, 1)
		fmt.Println(s1)
	}()

	go func() {
		s2 := append(slice1, 1) // 새로 만든 고루틴에서 슬라이스에 원소 추가
		fmt.Println(s2)
	}()

	// 위 코드에서는 데이터 경쟁이 발생하지 않음
	// 슬라이스가 가득 찾다면 고 런타임은 내부 배열을 새로 만든 뒤 원소를 추가하고, 그렇지 않으면 기존 내부 배열에 원소를 추가함
	// 고루틴마다 append를 호출하면, 슬라이스는 이미 꽉 찬 상태이므로 새로 생성된 내부 배열 기반의 슬라이스를 리턴
	// 배열을 새로 만들어서 기존 배열을 변경하지 않기 때문에 데이터 경쟁이 발생하지 않는 것

	// # 예제2: 길이가 1인 슬라이스 대신, 용량은 1이고 길이는 0인 슬라이스
	slice2 := make([]int, 0, 1)

	//go func() {
	//	s3 := append(slice2, 1)
	//	fmt.Println(s3)
	//}()
	//
	//go func() {
	//	s4 := append(slice2, 1)
	//	fmt.Println(s4)
	//}()

	// 위 코드에서는 데이터 경쟁이 발생할 수 있음
	// 슬라이스의 길이가 0이므로, 두 고루틴이 슬라이스에 원소를 추가할 때마다 내부 배열의 동일한 지점을 업데이트하려고 하기에 경쟁 상태가 발생
	// 경쟁이 발생하지 않게 하려면 어떻게 해야할까? => slice의 복제본을 만들면 됨

	go func() {
		sCopy := make([]int, len(slice2), cap(slice2))
		copy(sCopy, slice2) // 슬라이스 복제

		s3 := append(sCopy, 1) // 복제한 슬라이스에 원소 추가
		fmt.Println(s3)
	}()

	go func() {
		sCopy := make([]int, len(slice2), cap(slice2))
		copy(sCopy, slice2)

		s4 := append(sCopy, 1)
		fmt.Println(s4)
	}()

	// 두 고루틴 모두 슬라이스를 복제, 그리고 append를 원본 슬라이스가 아닌 복제한 슬라이스에 적용
	// => 고루틴이 서로 다른 데이터를 다루기에 데이터 경쟁이 발생하지 않음
}
