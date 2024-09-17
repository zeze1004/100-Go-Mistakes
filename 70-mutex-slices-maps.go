package main

import (
	"fmt"
	"sync"
)

// 슬라이스와 맵에 뮤텍스를 잘 적용하라
func main() {
	c := Cache{
		balances: make(map[string]float64),
	}
	c.AddBalance("1", 1.0)
	c.AddBalance("2", 3.0)
	c.AddBalance("3", 2.0)
	fmt.Println(c.AverageBalance1())
}

// 고객의 잔액을 캐시에 저장하기 위한 Cache 구조체
type Cache struct {
	mu       sync.RWMutex       // RWMutex: 쓰려는 고루틴이 없으면 여러 고루틴이 동시에 읽을 수 있음
	balances map[string]float64 // 고객 ID를 키로 하는 잔액을 저장하는 맵
}

// 잔액을 캐시에 저장하는 함수
func (c *Cache) AddBalance(id string, balance float64) {
	c.mu.Lock()              // 잔액을 추가할 때는 뮤텍스를 잠그고
	c.balances[id] = balance // 함수명이 Add면 +=가 맞지 않나? 이럴거면 SetBalance가 맞는 듯?
	c.mu.Unlock()            // 잔액 추가가 끝나면 뮤텍스를 해제
}

// 모든 고객의 평균 잔액을 계산
// 함수가 끝날 때 락이 풀려서, 함수전체가 크리티커컬 섹션
func (c *Cache) AverageBalance1() float64 {
	c.mu.RLock()         // 잔액을 읽을 때는 읽기 뮤텍스를 잠그고
	defer c.mu.RUnlock() // 함수가 끝나면 뮤텍스를 해제: 함수 전체를 보호

	sum := 0.
	for _, balance := range c.balances {
		sum += balance
	}
	return sum / float64(len(c.balances))
}

func (c *Cache) AverageBalance2() float64 {
	c.mu.RLock()
	m := make(map[string]float64, len(c.balances)) // 맵 복제

	for k, v := range c.balances {
		m[k] = v
	}
	c.mu.RUnlock()

	sum := 0.
	// 반복연산은 크리티컬 섹션 밖에서 본제본에 대해 수행
	for _, balance := range m {
		sum += balance
	}
	return sum / float64(len(m))
}
