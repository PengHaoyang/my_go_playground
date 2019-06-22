package main

import (
	"fmt"
)

// variables
var varInt = 9223372036854775807
var varInt8 int8 = 127
var varInt16 int16 = 32767
var varInt32 int32 = 2147483647
var varInt64 int64 = 9223372036854775807

var varUint uint = 18446744073709551615
var varUint8 uint8 = 255
var varUint16 uint16 = 65535
var varUint32 uint32 = 4294967295
var varUint64 uint64 = 18446744073709551615

var varSrting = "this a string"
var varChars = 'g'

var varIntPtr = &varUint16        // aka: var varIntPtr *uint16 = &varUint16
var varIntPtrChecker = *varIntPtr // varIntPtrChecker = varUint16

// ----------------------

var intMap = make(map[string]int)
var strMap = make(map[string]string)
var strMap2 = map[string]string{"aaa": "1111", "bbb": "33r3f3", "ccc": "434g", "ddd": "rggr"}

func changeMap(m map[string]string) {
	m["hahaha"] = "okkkk"
}

// ----------------------

// Speakable
type Speakable interface {
	greet() string
	shout() string
}

// Person struct
type Person struct {
	name  string
	age   int
	likes []string // slice, a dynamic array
	mood  string
}

// Introduce a Person
func Introduce(name string, age int, likes []string) Person {
	person := Person{}
	person.name = name
	person.age = age
	person.likes = likes
	person.mood = "ok"
	return person
}

// Greet , a func
func (p1 *Person) Greet(p2 *Person) {
	p1.mood = "happy"
	p2.mood = "happy"
	fmt.Println(p1, p2)
}

// Shout , a func
func (p1 *Person) Shout(p2 *Person) {
	p1.mood = "angry"
	p2.mood = "sad"
	fmt.Println(p1, p2)
}

func main() {
	// fmt.Println(varIntPtr, ", ", reflect.TypeOf(varIntPtr))
	// fmt.Println(*varIntPtr, ", ", reflect.TypeOf(*varIntPtr))
	// fmt.Println(&varUint16, ", ", reflect.TypeOf(&varUint16))
	// fmt.Println(*&varUint16, ", ", reflect.TypeOf(*&varUint16))
	// fmt.Println(varIntPtrChecker, ", ", reflect.TypeOf(varIntPtrChecker))
	// var varUint162 uint16 = 345
	// varIntPtr = &varUint162
	// fmt.Println(varIntPtrChecker, ", ", reflect.TypeOf(varIntPtrChecker))
	// ------------------------------

	intMap["aa"] = 1324
	intMap["aa"] = 223
	intMap["bb"] = 3553
	fmt.Println(intMap)

	strMap["cc"] = "aaaa"
	strMap["cc"] = "fwewf"
	strMap["dd"] = "eweg"
	fmt.Println(strMap)
	fmt.Println(strMap2)
	changeMap(strMap2)
	fmt.Println(strMap2)

	// ------------------------------
	// likes := []string{}
	// likes = append(likes, "few")
	// likes = append(likes, "dwffe")
	// likes = append(likes, "fwf")
	// fmt.Println(likes)
	// ------------------------------
	// person1 := Introduce("Jack", 13, []string{"aaa", "bbbbbb", "ccc"})
	// person2 := Introduce("Sam", 17, []string{"wwr", "frfr", "cclckclk", "efg"})
	// person1.Greet(&person2)
	// fmt.Println(person1, person2)
	// person1.Shout(&person2)
	// fmt.Println(person1, person2)

}
