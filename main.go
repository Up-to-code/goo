package main

import (
	"fmt"
	"time"
	"unicode/utf8"

)
func main() {
	  fmt.Print("Hello, World!\n")
	  const Pi = 3.14
	  fmt.Println("Pi = " + fmt.Sprint(Pi)  + "\n")
	  fmt.Printf("Pi = %v\n", Pi)
	  const test = "test"
	  // get the length of the string
      fmt.Print(fmt.Sprint(utf8.RuneCountInString(test)) + "\n")
	  var s string = "hello"
	  st2  := s + " world"
	  rsteing(st2)
	  fmt.Print(fmt.Sprint(intDivide(10, 2)) + "\n")

	  // test array
	  testArray()
	   textify()
	   slicetest()
	   maptest()
	   lops()
	   depstrings()
}



func rsteing( s string ) {
	fmt.Print(s + " \n")
}
func intDivide(a int, b int) int {
	return a / b
}
//  aryya  
func testArray() {	
	var arr = [5]int{260,2,3,4,5}

	fmt.Print(fmt.Sprint(arr) + "\n")
	fmt.Printf("arr size = %v, arr address = %v\n", len(arr), &arr[0]) // fmt.Printf("arr address = %v\n", &arr[0])
}
// if you want to use a variable in a string, you need to use %v
// %v is for variable
// %s is for string
// %d is for integer
// %f is for float
// %t is for bool
// %v %s %d %f %t
func textify() int8 {
   var a int8 = 5
   if a == 10 {
	   fmt.Print(fmt.Sprint(a) + "\n")
   }else {
	   fmt.Print(fmt.Sprint(a) + "\n")
   }
   

   switch a {
   case 10:
	   fmt.Print(fmt.Sprint(a) + "\n")
   default:
	   fmt.Print(fmt.Sprint(a) + "is not 10 \n")
   }

	return a
}


// slice
func slicetest(){
	var inSlice []uint32 = []uint32{1, 2, 3, 4, 5}
	inSlice = append(inSlice, 6)
	fmt.Print(fmt.Sprint(inSlice) + "\n")
	var inSlice2 []uint32 = make([]uint32, 0, 5)
	inSlice2 = append(inSlice2, 1 * 2, 2 * 2 , 3 * 2, 4 * 2, 5 * 2)
	fmt.Print(fmt.Sprint(inSlice2) + "\n" + fmt.Sprint(cap(inSlice2)) + "\n")
}

///  map 

func maptest(){
	var m = make(map[string]string)
	m["key"] = "value"
	m["key2"] = "value2"
	m["key"] = "ahmed  masnour"
	delete(m, "key2")
	var m2 = map[string]string{"key": "value", "key2": "value2"}
	fmt.Print(fmt.Sprint(m) + "map test 1 \n")
	fmt.Print(fmt.Sprint(m2) + " map test 2\n")
}

// lops  in maps and slices and  arrays
func lops() {
	var arr = [5]int{260,2,3,4,5}
	var arr2 = []int{260,2,3,4,5}
	var arr3 = make([]int, 5)
	fmt.Print(fmt.Sprint(arr) + "\n")
	fmt.Print(fmt.Sprint(arr2) + "\n")
	fmt.Print(fmt.Sprint(arr3) + "\n")
	var m= map[string]string{"key": "value", "key2": "value2"}
	for k, v := range m {
		fmt.Print(fmt.Sprint(k) + " " + v + "\n")
	}
	var n= 10
	for {
		if n == 0 {
			break
		}
		n--
		fmt.Print(fmt.Sprint(time.Now().Format("2006-01-02 15:04:05")) + "\n")
		fmt.Print("infinite loop \n")
	}
	 
}

//  dep in strings
func depstrings() {
	var s string = "helloب"
	 var in = s[0]
	 fmt.Print(fmt.Sprint(in) + "\n")
	 fmt.Print(fmt.Sprint(s[5]) + "\n")
	 fmt.Print(fmt.Sprint(len(s)) + "\n")
}

/// t
// type T struct {
// 	Name string
// 	Age int
// }
// func (t T) Strings() string {
// 	return fmt.Sprintf("Name: %v, Age: %v", t.Name, t.Age)
// }