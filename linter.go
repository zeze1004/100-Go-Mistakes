package main

import "fmt"

// linter를 이용하면 가려진 변수를 찾을 수 있음
/* vet이 가려진 변수를 찾아냄
# go vet -vettool=$(which shadow)
./linter.go:9:3: declaration of "i" shadows declaration at line 7
*/
func main() {
	i := 0
	if true {
		i := 1 // 가려진 변수
		fmt.Println(i)
	}
	fmt.Println(i)
}
