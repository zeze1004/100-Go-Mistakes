package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//cond1()
	cron2()
}

// 뮤텍스를 이용해 업데이터 고루틴이 매초 잔액을증가시키고 리스너 고루틴은 목표 금액에 도달할 때까지 루프를 돔
func cond1() {
	type Donation struct { // 현재 잔액과 뮤텍스로 구성된 Donation 구조체를 생성하고 초기화
		mu      sync.RWMutex
		balance int
	}
	donation := &Donation{}

	// 리스너 고루틴
	f := func(goal int) {
		donation.mu.RLock()
		for donation.balance < goal { // 목표 금액에 도달했는지 확인
			donation.mu.RUnlock()
			donation.mu.RLock()
		}
		fmt.Printf("목표액에 도달했습니다. 현재 잔액: %d\n", donation.balance)
		donation.mu.RUnlock()
	}
	go f(10)
	go f(15)

	// 업데이트 고루틴
	go func() {
		for { // 잔액을 계속 증가시킴
			time.Sleep(time.Second)
			donation.mu.Lock()
			donation.balance++
			donation.mu.Unlock()
		}
	}()
}

func cron2() {
	type Donation struct {
		ch      chan int // 채널 하나를 갖도록 Donation 수정
		balance int
	}
	donation := &Donation{ch: make(chan int)}

	// 리스너 고루틴
	f := func(goal int) {
		for balance := range donation.ch { // 채널 업데이트를 수신
			if balance >= goal {
				fmt.Printf("목표액에 도달했습니다. 현재 잔액: %d\n", balance)
				return
			}
		}
	}
	go f(10)
	go f(15)

	// 업데이트 고루틴
	for {
		time.Sleep(time.Second)
		donation.balance++
		donation.ch <- donation.balance // 잔액이 변경될 때마다 채널에 메세지를 보냄
	}
}

func cron3() {
	type Donation struct {
		cond    *sync.Cond
		balance int
	}

	donation := &Donation{cond: sync.NewCond(&sync.Mutex{})}

	// 리스너 고루틴
	f := func(goal int) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			donation.cond.Wait() // 잠금/잠금 해제 사이에서 (잔액이 변경되는) 상태를 기다림
		}
		fmt.Printf("목표액에 도달했습니다. 현재 잔액: %d\n", donation.balance)
		donation.cond.L.Unlock()
	}
	go f(10)
	go f(15)

	// 업데이트 고루틴
	for {
		time.Sleep(time.Second)
		donation.cond.L.Lock()
		donation.balance++        // 잠금/잠금 해제 사이에서 잔액을 증가시킴
		donation.cond.Broadcast() // (잔액이 변경되는) 상태가 됐음을 브로드캐스팅함
		donation.cond.L.Unlock()
	}
}
