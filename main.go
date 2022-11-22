// main.go는 컴파일 진입점이다.
package main

import "fmt"

func main() {
	// 상수는 const [상수이름] [상수타입] = 상수
	const name string = "danpacho"

	// 기본적으로 변수는 var [변수이름] [변수타입] = 변수 로 선언한다.
	var varName string = "my-friend"
	varName = "jkfds"

	// 축약형 변수는 [변수이름] := 변수 와 같은 형태로 선언 가능하고, 오직 "함수 func" 내부에서만 사용할 수 있다.
	// 또한 타입은 Go가 스스로 유추해주며 타입을 재지정할 수 없다.
	convienientVarName := "esasy"
	convienientVarName = "yeahs"

	fmt.Println(name, varName, convienientVarName)
}
