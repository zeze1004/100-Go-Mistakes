package main

import (
	"errors"
	"fmt"
	"net/http"
)

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

// 주어진 에러 타입에 적합한 HTTP 상태 코드를 리턴하는 HTTP 핸들러를 작성해보자
func handler(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("transaction") // 트랜잭션 ID를 추출

	amount, err := getTransactionAmount(transactionID) // 모든 로직이 담긴 getTransactionAmount 함수 호출
	if err != nil {
		if errors.As(err, &transientError{}) { // transientError에 포인터를 전달하는 방식으로 errors.As를 호출
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
}

func (t transientError) Error() string {
	return fmt.Sprintf("transient error: %v", t.err)
}

func getTransactionAmount(transactionID string) (float32, error) {
	// 트랜잭션ID가 올바르지 않음녀 에러를 리턴
	if len(transactionID) != 5 {
		return 0, transientError{fmt.Errorf("invalid transaction ID: %s", transactionID)}
	}

	amount, err := getTransactionAmountFromDB1(transactionID)
	// DB 조회에 실패하면 transientError를 리턴
	if err != nil {
		return 0, fmt.Errorf("failed to get transaction: %s: %w", transactionID, err) // %w를 사용해 원본 에러를 포장
	}
	return amount, nil
}

// 발생한 에러 종류와 상관없이 항상 400만 리턴, transientError 부분은 전혀 실행되지 않음
func getTransactionAmountFromDB1(id string) (float32, error) {
	return 0, transientError{err: err} // transientError를 리턴
}

// getTransactionAmount는 포장한 에러를 리턴하도록 변경하기 때문에 transientError 케이스는 항상 false가 됨
// 이러한 이유때문에 에러를 포장하는 디렉티브가 추가됐고, errors.As로 포장된 에러의 타입을 확인하는 기능 추가
