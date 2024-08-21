package main

import (
	"fmt"
	"io"
	"math"
	"runtime"
	"strings"
	"time"
)

//////////////////////////////////// 가본 문법
func swap (x, y string) (string, string) {
	return y, x
}


var plus int;

/* return 값 명시 */
func plusValue(value int) (x int){
	x = value + plus

	return
}

///////////////////////////////////// 반복문
func iter()int{
	sum:=0

	for i:=1; sum<10; {
		sum += i
	}

	for sum < 20 {
		sum += sum
	}

	return sum
}

///////////////////////////////////// 제어문
func pow2(value float64) float64 {
	if v:=math.Pow(value,2); int(v) % 3 == 0{
		return v
	}else{
		fmt.Printf("NO~ %g\n", v);
	}

	return -1
}

func sqrt(x float64) float64 {
	z := 1.0
	minimum := float64(1000)

	for ;z<=10; z++ {
		if diff:= math.Abs((z*z - x)/ (2*z)); diff < minimum{
			minimum = diff
		}
	}

	return minimum
}

func getOs()string{
	defer fmt.Println("함수가 끝나면 실행됨")

	fmt.Println("함수 실행!")

	switch os:=runtime.GOOS; os{
	case "darwin": 
		return "MAC";
	case "linux":
		return "Linux"
	default:
		return os
	}
}

///////////////////////////////////// 포인터
func getPointer() {
	var p* int
	i := 10
	p = &i

	fmt.Println(p)
	fmt.Println(i)

	*p = 20
	fmt.Println(p, *p)
	fmt.Println(i)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}


type User struct {
	name string
	age int
}

///////////////////////////////////// 메서드
type Position struct {
	x int
	y int
}

func (p Position) Sum() int {
	return p.x + p.y
}

func (p Position) Abs() float64 {
	return math.Sqrt(float64(p.x*p.x) + float64(p.y*p.y))
}

// 포인터 리시버
func (r *Position) swap() {
	temp := r.x
	r.x = r.y
	r.y = temp
}

func scale (r* Position) {
	r.x *= 10
	r.y *= 10
}


///////////////////////////////////// 인터페이스
type ValI interface {
	Show()
}


type Val struct{
	str string
}

func (v *Val) Show() {
	if v == nil {
		fmt.Println("NIL")
		return
	}
	fmt.Println(v.str)
}

// 인터페이스 덮어쯔기 가능
// func (u User) String() string{
// 	return fmt.Sprintf("새로만든 인터페이스 >> %v (%v)", u.name, u.age)
// }

///////////// TEST******* 특정 메서드에 대한 프린트 함수 만들기
type IPAddr [4]byte

func (ip IPAddr) String() string{
	str := make([]string, len(ip))
	for i, v := range ip{
		str[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(str, ".")
}

///////////////////////////////////// Errors
type CustomError struct {
	time time.Time
	message string
}

///// 커스텀 error interface 만들기
type error interface {
	Error() string
}

// func (e CustomError) Error() string {
// 	return fmt.Sprintf("[ERROR] %v, %s", e.time, e.message)
// }

///// 에러를 추가한 Sqrt 함수 만들기
type ErrNegativeSqrt struct {
	inputValue float64
}

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v", e.inputValue)
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	minimum := float64(1000)
	
	if x < 0 {
		return 0, ErrNegativeSqrt{x}
	}

	for ;z<=10; z++ {
		if diff:= math.Abs((z*z - x)/ (2*z)); diff < minimum{
			minimum = diff
		}
	}

	return minimum, nil
}

func main(){
	/* 모듈 호출 */
	// log.SetPrefix(("LOG")) // 접두사 추가
	// // log.SetFlags(0) // 타임스탬프 제거

	// message, err := hello.Hello(("Yena"))

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(message)

	// fmt.Println(plusValue(10))


	///////////////////////////////////// 반복문
	// fmt.Println(iter())

	///////////////////////////////////// 제어문
	// fmt.Println(pow2(7))
	// fmt.Println(sqrt(10))
	// fmt.Println(getOs())
	
	///////////////////////////////////// 포인터
	// getPointer()

	//// 구조체
	// yena := User{"yena", 27}
	// quokka := &User{"quokka", 1}
	// unknown := User{name: "user"}

	// fmt.Println(yena.name, yena.age, yena)
	// fmt.Println(quokka.name, quokka.age, quokka)
	// fmt.Println(unknown.name, unknown.age, unknown)

	///////////////////////////////////// 배열
	// a  := [2]string{"y","n"}
	// fmt.Println(a)
	
	// var b [2]int
	// b[0] = 1
	// b[1] = 2
	// fmt.Println(b)

	///////////////////////////////////// 슬라이스
	// arr := [5]int{1,2,3,4,5}
	// var s []int = arr[4:5]

	// s[0] = -1 // (배열의 영역을 나타냄 -> 따라서 기존 배열이 같이 수정됨)

	// fmt.Println(arr)
	// fmt.Println(s)

	// literal := []struct {
	// 	i int
	// 	s string
	// }{
	// 	{2, "a"},
	// 	{3, "b"},
	// }
	// fmt.Println(literal)
	// fmt.Println(len(literal), cap(literal))

	//// slice - 길이와 용량
	// s := []int{2, 3, 5, 7, 11, 13}
	// fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// s = s[:1]
	// fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// s = s[:4]
	// fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// // capacity(용량) 줄어듬
	// s = s[3:]
	// fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	//// 동적 크기 배열 만들기
	// a := make([]int, 5, 10) // 배열, len, cap ////////////////////////////////////////////// ⭐️⭐️⭐️⭐️ make
	// printSlice("a", a)
	
	// b := a[7:cap(a)]
	// printSlice("b", b)
	
	// c := append(b, 1,2,3)
	// printSlice("c",c)

	// d := make([]int, ())

	///////////////////////////////////// range
	// var numbers = []int{1,2,3,4,5}
	// for i, v := range numbers {
	// 	fmt.Printf("numbers[%d] = %v\n", i, v)
	// }

	// const (
	// 	dx = 3
	// 	dy = 4
	// )

	// arr := [dy][dx]uint8{}

	// for i := range arr{
	// 	for j := range arr[i]{
	// 		arr[i][j] = uint8((i+j)/2)
	// 	}
	// }

	
	// ss := make([][]uint8, dy)
	// for y := 0; y < dy; y++ {
	// 	s := make([]uint8, dx)
	// 	for x := 0; x < dx; x++ {
	// 		s[x] = uint8((x + y) / 2)
	// 	}
	// 	ss[y] = s
	// }

	// fmt.Println(arr)
	// fmt.Println(ss)

	///////////////////////////////////// Map
	// var idMap map[int]User // map 선언
	// idMap = make(map[int]User) // 초기화

	// idMap[1] = User{"yena", 27}

	// keyMap := map[string]string{
	// 	"google": "usa",
	// 	"kakao": "korea",
	// }

	// // fmt.Println(idMap, idMap[1])
	// // fmt.Println(keyMap, keyMap["google"])

	// // keyMap["naver"] = "korea"
	// // fmt.Println(keyMap, keyMap["naver"])
	
	// // delete(keyMap, "naver")
	// // fmt.Println(keyMap, keyMap["naver"])
	
	// _, exists := keyMap["line"]
	// fmt.Println(exists)

	// for key, val := range keyMap {
	// 	fmt.Println(key, val)
	// }


	///////////////////////////////////// 함수
	// plus10 := func(x int) int {
	// 	return x + 10
	// }

	// // 익명함수
	// test := func (fun func(x int) int) int {
	// 	return fun(20) / 10
	// }

	// // 함수 원형
	// type calcurate func(int, int)int
	// calc := func (f calcurate, a int, b int) int{
	// 	result := f(a, b)
	// 	return result
	// }

	// fmt.Println(plus10(10))
	// fmt.Println(test(plus10))
	// fmt.Println(calc(func(a int, b int) int {
	// 	return a+b
	// }, 10, 20))

	
	// fibonacci := func() func() int{
	// 	var prev, cur int = -1, -1;
		
	// 	return func() int {
	// 		if prev == -1 {
	// 			prev = 0
	// 			return 0
	// 		}
	// 		if cur == -1 {
	// 			cur = 1
	// 			return 1
	// 		}

	// 		temp := cur
	// 		cur += prev
	// 		prev = temp
	// 		return cur
	// 	}
	// }
	
	// f := fibonacci()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }

	///////////////////////////////////// 메서드
	// pos := Position{10, 20}
	// fmt.Println(pos.Sum())
	
	// pos.swap()
	// fmt.Println(pos)
	
	// scale(&pos)
	// fmt.Println(pos)
	
	// v := &Position{3,4}
	// fmt.Println(v.Abs())

	///////////////////////////////////// 인터페이스
	// type Shape interface{
	// 	Sum() int
	// 	Abs() float64
	// }

	// var shape Shape
	// pos := Position{2,3}
	
	// shape = pos

	// // fmt.Println(shape.Sum())
	// fmt.Println(shape.Abs())


	// 인터페이스와 포인터
	// var test ValI = &Val{}
	
	// var value *Val
	// test2 := value

	// fmt.Println(test)
	// test.Show()

	// fmt.Println(test2)
	// test2.Show()
	
	// // 빈 인터페이스 -> unknown 같은 느낌
	// var i interface{}
	// printInterface(i)

	// i = 30
	// printInterface(i)

	// // type assertion (해당 타입을 가지고 있는지)
	// strValue, ok := i.(string) // ok가 없으면 panic
	// fmt.Println(strValue, ok)

	// numValue, ok := i.(int)
	// fmt.Println(numValue, ok)

	///////////////////////////////////// 타입 스위치
	// var i interface {}

	// switch v := i.(type){
	// case int:
	// 	fmt.Println("Int",v)
	// case string:
	// 	fmt.Println("String",v)
	// default:
	// 	fmt.Println("IDK", v)

	// }

	///////////////////////////////////// 자주 사용되는 인터페이스
	// yena := User{"yena", 27}
	// fmt.Println(yena)


	// hosts := map[string]IPAddr{
	// 	"loopback":  {127, 0, 0, 1},
	// 	"googleDNS": {8, 8, 8, 8},
	// }

	// for name, ip := range hosts {
	// 	fmt.Printf("%v: %v\n", name, ip)
	// }

	///////////////////////////////////// Errors
	// var errRun = func() error {
	// 	err := CustomError{
	// 		time.Now(),
	// 		"에러가 발생함",
	// 	}

	// 	return err
	// }

	// err := errRun()
	// fmt.Println(err)


	///// 에러를 추가한 Sqrt 함수 만들기
	// res, err := Sqrt(-2)
	// fmt.Println(res, err)

	///////////////////////////////////// Readers	
	r := strings.NewReader("Hello, Go!")

	b:= make([]byte, 4)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, byte = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	
}


func printInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type MyReader struct {
	// Read func([]byte)(int, error)
}


func (r MyReader) Read (b []byte)(int, error){
	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}