package main

import (
	"fmt"
	"reflect"
)

// Person struct this is
type Person struct {
	name string
	age  int
}

// some array
var persons = [...]Person{
	Person{"Jack", 13},
	Person{"Tom", 20},
	Person{"John", 24},
	Person{"Sam", 40},
	Person{"Peter", 80},
	Person{"Steve", 90},
}

// some map
var strMap = map[string]string{
	"aaa": "1111",
	"bbb": "22222",
	"ccc": "333333",
	"ddd": "4444444",
}

func main() {
	// use for ;; and len to iterate array
	// for n := 0; n < len(persons); n++ {
	// 	fmt.Println(n, ", ", persons[n])
	// }
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~")
	// use range to iterate array
	for n, person := range persons {
		fmt.Print(n, ", ", person)
		if person.age < 50 {
			fmt.Println("   young person")
		} else {
			fmt.Println("   elder person")
		}
	}
	// use range to iterate map
	for k, v := range strMap {
		fmt.Println(k, "-->", v, reflect.TypeOf(k), reflect.TypeOf(k))
	}
}
