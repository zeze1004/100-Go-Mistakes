package main

import "fmt"

// 에러 타입을 정확하게 검사하라

// 49장에서 %w로 에러를 포장하는 방법을 설명했는데, 이 방법을 적용하는 순간, 에러 타입을 검사하는 방법도 바꿔야함.

// 구체적인 예시
// 주어진 ID로 DB를 조회하여 거래 금액을 리턴하는 API에 대해 두 가지 에러가 발생할 수 있음
// 1) ID가 잘못된 경우(스트링 길이가 5가 아닐 때): 400 리턴 2) DB 조회 실패: 500 리턴
// 임시 에러임을 표시하는 transientError 타입을 만들어 나머지 경우는 400에러를 리턴하게 구현함

// 커스텀 에러 타입 정의
type transientError struct {
	err error
}

func (t transientError) Error() string {
	return fmt.Sprintf("transient error: %v", t.err)
}

func getTransactionAmount(transactionID string) (float32, error) {
	// 트랜잭션ID가 올바르지 않음녀 에러를 리턴
	if len(transactionID) != 5 {
		return 0, transientError{fmt.Errorf("invalid transaction ID: %s", transactionID)}
	}

	amount, err := getTransactionAmount(transactionID)
	// DB 조회에 실패하면 transientError를 리턴
	if err != nil {
		return 0, transientError{err: err}
	}
	return amount, nil
}
