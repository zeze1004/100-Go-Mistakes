package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// sync.WaitGroup을 잘 사용하라

func listing1() {
	wg := sync.WaitGroup{}
	var v uint64

	for i := 0; i < 3; i++ {
		go func() { // 고루틴을 생성하고
			wg.Add(1)               // 대기 그룹의 카운트를 증가 <= 경쟁상태가 발생하는 이유는 부모 고루틴이 아닌 새로 생성된 고루틴에서 호출돼서. 따라서 wg.Wiat()를 호출하기 전에 대기 그룹에게 세 고루틴이 끝나길 기다리도록 알리지 못할 수 있음
			atomic.AddUint64(&v, 1) // v를 원자 연산으로 증가
			wg.Done()               // 대기 그룹의 카운트를 감소
		}()
	}

	wg.Wait()      // 모든 코루틴이 v를 증가시키고 나서 출력할 때까지 기다림
	fmt.Println(v) // 3가 나와야하는데 경쟁 상태일 때 2가 출력됨
}

// 문제 해결(1): wg.Add를 호출하고 나서 3까지 반복
func listing2() {
	wg := sync.WaitGroup{}
	var v uint64

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}
}

// 문제 해결(2): wg.Add를 호출하고 나서 자식 고루틴을 구동시키는 것
func listing3() {
	wg := sync.WaitGroup{}
	var v uint64

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}
}
