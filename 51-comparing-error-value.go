package main

// 에러 값을 정확하게 검사하라
// 센티널 에러(에러 값)에 적용하는 법을 알아보자
// 센티널 에러: 글로벌 변수로 정의한 에러

err := query()
if err != nil {
if err == sql.ErrNoRows { // 센티널 에러도 포장할 수 있음. sql.ErrNoRows를 fmt.Errorf와 %w 디렉티브로 포장하면, err == sql.ErrNoRows는 항상 false가 됨.
	 					  // 이를 해결하기 위해 에러값을 검사할 때는 errors.Is를 이용하면, %w와 fmt.Errorf로 포장한 에러도 검사 가능
	// 특정 에러 값에 대한 처리
	} else {
	// 그 외 에러 처리
	}
}
