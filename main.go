package main

import (
	"fmt"
	"strings"
)

func main() {
	var mult int = multiply(1, 2)

	// 2개를 동시에 반환하는 lenAndUpper 함수
	totalLength, upperName := lenAndUpper("puang")
	// _ 를 활용하면 변수를 ignore 할 수 있다.
	yeah, _ := lenAndUpper("ignored")

	names := repeatNames("june", "me", "dan")
	singleName := funcWithArguments("june", "done", "hande")

	fmt.Println(mult, totalLength, upperName, names, singleName, yeah)
}

// 인자가 타입을 공유하는 경우 a int, b int를 a, b int로 축약가능하다.
func multiply(a, b int) int {
	return a * b
}

// go의 함수는 많은 것을 동시에 반환할 수 있다.
// func [funcName(arg)] (return 타입1, return 타입2, return타입3, ...)
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToTitle(name)
}

// 자바스크립트 나머지 매게 변수 처럼 사용할 때, ...[타입]
func funcWithArguments(arg ...string) string {
	return strings.Join(arg, "")
}

// repeatNames names 나머지 매개 변수 ...string
func repeatNames(names ...string) []string {
	return names
}
