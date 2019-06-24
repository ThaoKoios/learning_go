package main

import "fmt"

// EX1_HELLOWORLD
// func main() {
// 	fmt.Println("hello world")
// }

// EX2_VALUES
// func main() {

// 	fmt.Println("go" + "lang")

// 	fmt.Println("1+1=", 1+1)
// 	fmt.Println("7.0/3.0 =", 7.0/3.0)

// 	fmt.Println(true && false)
// 	fmt.Println(true || false)
// 	fmt.Println(!true)
// }

// EX3_VARIABLES
// func main() {

// 	var a = "initial"
// 	fmt.Println(a)

// 	var b, c int = 1, 2
// 	fmt.Println(b, c)

// 	var d = true
// 	fmt.Println(d)

// 	var e int
// 	fmt.Println(e)

// 	f := "apple"
// 	fmt.Println(f)
// }

// EX4_CONSTANTS
// import "math"

// const s string = "constant"

// func main() {
// 	fmt.Println(s)

// 	const n = 500000000

// 	const d = 3e20 / n
// 	fmt.Println(d)

// 	fmt.Println(int64(d))

// 	fmt.Println(math.Sin(n))
// }

// EX5_FOR
// func main() {

// 	i := 1
// 	for i <= 3 {
// 		fmt.Println(i)
// 		i = i + 1
// 	}

// 	for j := 7; j <= 9; j++ {
// 		fmt.Println(j)
// 	}

// 	for {
// 		fmt.Println("loop")
// 		break
// 	}

// 	for n := 0; n <= 5; n++ {
// 		if n%2 == 0 {
// 			continue
// 		}
// 		fmt.Println(n)
// 	}
// }

// EX6_IF/ELSE
// func main() {

// 	if 7%2 == 0 {
// 		fmt.Println("7 is even")
// 	} else {
// 		fmt.Println("7 is odd")
// 	}

// 	if 8%4 == 0 {
// 		fmt.Println("8 is divisible by 4")
// 	}

// 	if num := 9; num < 0 {
// 		fmt.Println(num, "is negative")
// 	} else if num < 10 {
// 		fmt.Println(num, "has 1 digit")
// 	} else {
// 		fmt.Println(num, "has multiple digits")
// 	}
// }

// EX7_SWITCH
// import "time"

// func main() {

// 	i := 2
// 	fmt.Print("Write ", i, " as ")
// 	switch i {
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	case 3:
// 		fmt.Println("three")
// 	}

// 	switch time.Now().Weekday() {
// 	case time.Saturday, time.Sunday:
// 		fmt.Println("It's the weekend")
// 	default:
// 		fmt.Println("It's a weekday")
// 	}

// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("It's  before noon")
// 	default:
// 		fmt.Println("It's after noon")
// 	}

// 	whatAmI := func(i interface{}) {
// 		switch t := i.(type) {
// 		case bool:
// 			fmt.Println("I'm a bool")
// 		case int:
// 			fmt.Println("I'm an init")
// 		default:
// 			fmt.Printf("Don't know type %T\n", t)
// 		}
// 	}
// 	whatAmI(true)
// 	whatAmI(1)
// 	whatAmI("hey")
// }

// EX8_ARRAYS
// func main() {

// 	var a [5]int
// 	fmt.Println("emp:", a)

// 	a[4] = 100
// 	fmt.Println("set:", a)
// 	fmt.Println("get:", a[4])

// 	fmt.Println("len:", len(a))

// 	b := [5]int{1, 2, 3, 4, 5}
// 	fmt.Println("dcl:", b)

// 	var twoD [2][3]int
// 	for i := 0; i < 2; i++ {
// 		for j := 0; j < 3; j++ {
// 			twoD[i][j] = i + j
// 		}
// 	}
// 	fmt.Println("2d: ", twoD)
// }

//EX9_SLICES
// func main() {

// 	s := make([]string, 3)
// 	fmt.Println("emp:", s)

// 	s[0] = "a"
// 	s[1] = "b"
// 	s[2] = "c"
// 	fmt.Println("set: ", s)
// 	fmt.Println("get", s[2])

// 	fmt.Println("len:", len(s))

// 	s = append(s, "d") // append slices
// 	s = append(s, "e", "f")
// 	fmt.Println("apd:", s)

// 	c := make([]string, len(s)) // create an empty slice c of same length as s
// 	copy(c, s)                  // copy into c from s
// 	fmt.Println("cpy:", c)

// 	l1 := s[2:5] //slice[low:high] Ex: gets a slice of the elements s[2], s[3], s[4]
// 	fmt.Println("sl1:", l1)

// 	l2 := s[:5] //slice up to (but excluding) s[5]
// 	fmt.Println("sl2:", l2)

// 	l3 := s[2:] //slice up from (and including) s[2]
// 	fmt.Println("sl3:", l3)

// 	t := []string{"g", "h", "i"} //declare and initialize a variable for slice in a single line
// 	fmt.Println("dcl:", t)

// 	twoD := make([][]int, 3) // create a multi-dimensional data structures
// 	for i := 0; i < 3; i++ {
// 		innerLen := i + 1
// 		twoD[i] = make([]int, innerLen)
// 		for j := 0; j < innerLen; j++ {
// 			twoD[i][j] = i + j
// 		}
// 	}
// 	fmt.Println("2d: ", twoD) // RESULT: [[0][1,2][2,3,4]]

// }

// EX10_MAPS
// func main() {
// 	m := make(map[string]int) //create an empty map use the builin make: make(map[key-type]val-type)

// 	m["k1"] = 7 //set key/value pairs using typical name [key] = val syntax
// 	m["k2"] = 13

// 	fmt.Println("map:", m)

// 	v1 := m["k1"]
// 	fmt.Println("v1: ", v1)

// 	fmt.Println("len: ", len(m))

// 	delete(m, "k2")
// 	fmt.Println("map: ", m)

// 	_, prs := m["k2"] //used to disambiguate between missing keys and keys with zero values like 0 or ""
// 	fmt.Println("prs: ", prs)

// 	n := map[string]int{"foo": 1, "bar": 2} //declare and initialize a new map in same line
// 	fmt.Println("map: ", n)
// }

//EX11_RANGE
// func main() {

// 	nums := []int{2, 3, 4}
// 	sum := 0
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	fmt.Println("sum:", sum)

// 	for i, num := range nums {
// 		if num == 3 {
// 			fmt.Println("index:", i)
// 		}
// 	}

// 	kvs := map[string]string{"a": "apple", "b": "banana"}
// 	for k, v := range kvs {
// 		fmt.Printf("%s -> %s\n", k, v)
// 	}

// 	for k := range kvs {
// 		fmt.Println("key:", k)
// 	}

// 	for i, c := range "go" {
// 		fmt.Println(i, c) //range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself
// 	}
// }

//EX12_FUNCTIONS
// func plus(a int, b int) int {

// 	return a + b
// }

// func plusPlus(a, b, c int) int {
// 	return a + b + c
// }

// func main() {
// 	res := plus(1, 2)
// 	fmt.Println("1+2 =", res)

// 	res = plusPlus(1, 2, 3)
// 	fmt.Println("1+2+3 =", res)
// }

// EX13_MULTIPLE RETURN VALUES
// func vals() (int, int) {
// 	return 3, 7
// }

// func main() {

// 	a, b := vals()
// 	fmt.Println(a)
// 	fmt.Println(b)

// 	_, c := vals()
// 	fmt.Println(c)

// 	d, _ := vals()
// 	fmt.Println(d)
// }

//EX14_VARIADIC FUNCTIONS
// func sum(nums ...int) {
// 	fmt.Print(nums, " ")
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	fmt.Println(total)
// }

// func main() {

// 	sum(1, 2)
// 	sum(1, 2, 3)

// 	nums := []int{1, 2, 3, 4}
// 	sum(nums...)
// }

//EX15_CLOSURES --------------ask Jax on Monday????
func intSeq(start int) func() int {
	i := start
	return func() int {
		i++
		return i
	}
}

func bomb(ticks int) func() string {
	t := ticks
	return func() string {
		t--
		if t > 0 {
			return "tick"
		}
		return "boom!"
	}
}
func main() {

	nextInt := intSeq(0)

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq(10)
	fmt.Println(newInts())
	fmt.Println(nextInt())

	b := bomb(5)
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

//EX16a_RECURSION
// func fact(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	return n * fact(n-1)
// }

// func main() {
// 	fmt.Println(fact(7))
// }

//EX16b_RECURSION
// func fact(n uint) uint {
// 	if n == 0 {
// 		return 1
// 	}
// 	return n * fact(n-1)
// }

// func main() {
// 	x := uint(7)
// 	calFact := fact(x)
// 	fmt.Println(calFact)
// }

//EX16c_FACTORIALS
// import "time"
// import _ "fmt"

// //LIM : defining the length of facts [LIM] array and limiting i
// const LIM = 10

// var facts [LIM]uint64

// func main() {
// 	fmt.Println("=====================FACTORIAL===========================")
// 	start := time.Now()
// 	for i := uint64(0); i < LIM; i++ {
// 		fmt.Printf("Factorial for %d is : %d \n", i, Factorial(uint64(i)))
// 	}
// 	end := time.Now()
// 	fmt.Printf("Calcualation finished in %s \n", end.Sub(start)) //Calculation finished in 2.0002ms

// 	fmt.Println("=====================FACTORIAL CLOSURE==========================")
// 	start = time.Now()
// 	fact := FactorialClosure()
// 	for i := uint64(0); i < LIM; i++ {
// 		fmt.Printf("Factorial closure for %d is : %d \n", uint64(i), fact(uint64(i)))
// 	}
// 	end = time.Now()
// 	fmt.Printf("Calcualtion finished in %s \n", end.Sub(start)) //Calculation finished in 1ms

// 	fmt.Println("=====================FACTORIAL MEMOIZATION======================")
// 	start = time.Now()
// 	var result uint64 = 0
// 	for i := uint64(0); i < LIM; i++ {
// 		result = FactorialMemoization(uint64(i))
// 		fmt.Printf("The factorial value for %d is %d\n", uint64(i), uint64(result))
// 	}

// 	end = time.Now()
// 	fmt.Printf("Calcualtion finished in %s\n", end.Sub(start)) //Calculation finished in 0ms
// }

// //Factorial : normal factorial
// func Factorial(n uint64) (result uint64) {
// 	if n > 0 {
// 		result = n * Factorial(n-1)
// 		return result
// 	}
// 	return 1
// }

// //FactorialClosure : using closure for factorial
// func FactorialClosure() func(n uint64) uint64 {
// 	var a, b uint64 = 1, 1
// 	//_ = a
// 	//_ = b
// 	return func(n uint64) uint64 {
// 		if n > 1 {
// 			a, b = b, uint64(n)*uint64(b)
// 		} else {
// 			return 1
// 		}

// 		return b
// 	}
// }

// //FactorialMemoization : fastest factorial
// func FactorialMemoization(n uint64) (res uint64) {
// 	if facts[n] != 0 {
// 		res = facts[n]
// 		return res
// 	}

// 	if n > 0 {
// 		res = n * FactorialMemoization(n-1)
// 		return res
// 	}

// 	return 1
// }

//EX17_POINTERS---------------GO supports pointers, allowing you to pass references to values and records within your program.-------------
// import "fmt"

// func zeroval(ival int) {
// 	ival = 0
// }

// func zeroptr(iptr *int) {
// 	*iptr = 0
// }

// func main() {
// 	i := 1
// 	fmt.Println("initial:", i)

// 	zeroval(i)
// 	fmt.Println("zeroval:", i)

// 	zeroptr(&i)
// 	fmt.Println("zeroptr", i)

// 	fmt.Println("pointer:", &i)
// }

//EX18_STRUCTS---------------Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.----------------
// import "fmt"

// type person struct {
// 	name string
// 	age  int
// }

// func main() {

// 	fmt.Println(person{"Bob", 20})

// 	fmt.Println(person{name: "Alice", age: 30})

// 	fmt.Println(person{name: "Fred"})

// 	fmt.Println(&person{name: "Ann", age: 40})

// 	s := person{name: "Sean", age: 50}
// 	fmt.Println(s.name)

// 	sp := &s
// 	fmt.Println(sp.age)

// 	sp.age = 51
// 	fmt.Println(sp.age)
// }

//EX19_METHODS
// import "fmt"

// type rect struct {
// 	width, height int
// }

// func (r *rect) area() int { //This area method has a receiver type of *rect
// 	return r.width * r.height
// }

// func (r rect) perim() int { //Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver
// 	return 2*r.width + 2*r.height
// }

// func main() {
// 	r := rect{width: 10, height: 5}

// 	fmt.Println("area: ", r.area()) //Here we call the 2 methods defined for our struct
// 	fmt.Println("perim:", r.perim())

// 	rp := &r                         //Go automatically handles conversion between values and pointers for method calls
// 	fmt.Println("area: ", rp.area()) //You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct
// 	fmt.Println("perim:", rp.perim())
// }

//EX20_INTERFACES--------Interfaces are named collections of method signatures-----------------
// import "fmt"
// import "math"

// type geometry interface { //Here’s a basic interface for geometric shapes
// 	area() float64
// 	perim() float64
// }

// type rect struct { //For our example we’ll implement this interface on rect and circle types
// 	width, height float64
// }
// type circle struct {
// 	radius float64
// }

// func (r rect) area() float64 { //To implement an interface in Go, we just need to implement all the methods in the interface. Here we implement geometry on rects
// 	return r.width * r.height
// }
// func (r rect) perim() float64 { //The implementation for circles
// 	return 2*r.width + 2*r.height
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func (c circle) perim() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func measure(g geometry) { //If a variable has an interface type, then we can call methods that are in the named interface.
// 	fmt.Println(g) //Here’s a generic measure function taking advantage of this to work on any geometry
// 	fmt.Println(g.area())
// 	fmt.Println(g.perim())
// }

// func main() { //The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure
// 	r := rect{width: 3, height: 4}
// 	c := circle{radius: 5}

// 	measure(r)
// 	measure(c)
// }

//EX21_ERRORS
//---In Go it’s idiomatic to communicate errors via an explicit, separate return value.
//This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C. Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for any other, non-error tasks.----
// import (
// 	"errors"
// 	"fmt"
// )

// func f1(arg int) (int, error) { //By convention, errors are the last return value and have type error, a built-in interface
// 	if arg == 42 {

// 		return -1, errors.New("can't work with 42") //errors.New constructs a basic error value with the given error message

// 	}

// 	return arg + 3, nil //A nil value in the error position indicates that there was no error
// }

// type argError struct { //It’s possible to use custom types as errors by implementing the Error() method on them.
// 	arg  int //Here’s a variant on the example above that uses a custom type to explicitly represent an argument error
// 	prob string
// }

// func (e *argError) Error() string {
// 	return fmt.Sprintf("%d - %s", e.arg, e.prob)
// }

// func f2(arg int) (int, error) {
// 	if arg == 42 {
// 		return -1, &argError{arg, "can't work with it"} //In this case we use &argError syntax to build a new struct, supplying values for the two fields arg and prob
// 	}
// 	return arg + 3, nil
// }

// func main() { //The two loops below test out each of our error-returning functions.
// 	//Note that the use of an inline error check on the if line is a common idiom in Go code

// 	for _, i := range []int{7, 42} {
// 		if r, e := f1(i); e != nil {
// 			fmt.Println("f1 failed:", e)
// 		} else {
// 			fmt.Println("f1 worked:", r)
// 		}
// 	}
// 	for _, i := range []int{7, 42} {
// 		if r, e := f2(i); e != nil {
// 			fmt.Println("f2 failed:", e)
// 		} else {
// 			fmt.Println("f2 worked:", r)
// 		}
// 	}

// 	_, e := f2(42) //If you want to programmatically use the data in a custom error, you’ll need to get the error as an instance of the custom error type via type assertion
// 	if ae, ok := e.(*argError); ok {
// 		fmt.Println(ae.arg)
// 		fmt.Println(ae.prob)
// 	}
// }

//EX22_GOROUTINES

//EX23_CHANNELS

//EX24_CHANNELS BUFFERING

//EX25_CHANNEL SYNCHRONIZATION

//EX26_CHANNEL DIRECTIONS

//EX27_SELECT

//EX28_TIMEOUTS

//EX29_NON-BLOCKING CHANNEL OPERATIONS

//EX30_CLOSING CHANNELS

//EX31_RANGE OVER CHANNELS

//EX32_TIMERS

//EX33_TICKERS

//EX34_WORKER POOLS

//EX35_WAITGROUPS

//EX36_RATE LIMITING

//EX37_ATOMIC COUNTERS

//EX38_MUTEXES

//EX39_STATEFUL GOROUTINES

//EX40_SORTING

//EX41_SORTING BY FUNCTIONS

//EX42_PANIC

//EX43_DEFER

//EX44_COLLECTION FUNCTIONS

//EX45_STRING FUNCTIONS

//EX46_STRING FORMATTING

//EX47_REGULAR EXPERESSIONS

//EX48_JSON

//EX49_TIME

//EX50_EPOCH

//EX51_TIME FORMATTING/PARSING

//EX52_RANDOM NUMBERS

//EX53_NUMBER PARSING

//EX54_URL PARSING

//EX55_SHA1 HASHES

//EX56_BASE64 ENCODING

//EX57_READING FILES

//EX58_WRITING FILES

//EX59_LINE FILTERS

//EX60_FILE PATHS

//EX61_COMMAND-LINE ARGUMENTS

//EX62_COMMAND-LINE FLAGS

//EX63_COMMAND-LINE SUBCOMMANDS

//EX64_ENVIRONMENT VARIABLES

//EX65_HTTP CLIENTS

//EX66_HTTP SERVERS

//EX67_SPAWNING PROCESSES

//EX68_EXEC'ING PROCESSES

//EX69_SIGNALS

//EX70_EXIT
