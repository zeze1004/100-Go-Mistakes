package main

// 에러 포장하기: 에러를 레퍼 컨데이너에 담아서 원본 에러도 사용 가능하게 만드는 것을 뜻함
// 에러 포장하기는 다음 두 가지 용도로 사용됨 1) 에러에 문맥 추가(단순 에러 결과만 뱉는게 아니라 어떤 상황에서 누가 에러가 났는지 문맥 추가) 2) 에러 구체화

// 커스텀 에러 구조체로 만든 특정한 에러 타입만 반환할 수 있었는데, 1.13부터 %w 도입돼서 에러 타입 별도 생성할 필요 없이 원본 에러에 문맥 정보 추가해서 감쌀 수 있음
