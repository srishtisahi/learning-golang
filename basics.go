// use cmd+/ to comment/uncomment

/* for multi line comment
use option+shift+a */

package main

//--------------------------------------------------------------------------------//

/*
func main() {
	fmt.Println("Hello, World!")
}
*/

//--------------------------------------------------------------------------------//

/*
data types in go
numeric - int, float
boolean - true, false
string
derived - pointer, array, structure, map, interface
*/

//--------------------------------------------------------------------------------//

/* //DECLARING VARIABLES

func main() {
	var a int = 5
	var b float32 = 4.32
	const pi float64 = 3.141553723
	x,y := 14,15 // declaring a variable

	fmt.Println(a,b,pi,x,y)
} */

//--------------------------------------------------------------------------------//

/* //OPERATORS
func main() {
	x,y := 5,6
	fmt.Println(x+y) //arithmetic operations
	fmt.Println(x-y)
	fmt.Println(x*y)
	fmt.Println(x/y)
	fmt.Println(x%y)
	fmt.Println(x==y) //boolean operations
	fmt.Println(x!=y)
	fmt.Println(x>y)
	fmt.Println(x<y)

	isbool := true
	hate := false

	fmt.Println(isbool && hate) //and
	fmt.Println(isbool || hate) //or
	fmt.Println(!isbool) //not
} */

//--------------------------------------------------------------------------------//

//POINTERS

/*
func main() {
	x := 5

	fmt.Println(x)
	fmt.Println(&x) //memory address of x
} */

//change value of a variable using pointers
//call by reference
//set the pointer to the memory address of the variable

/* func main() {
	x := 10

	fmt.Println(x)

	changeValue(&x)
	fmt.Println(x)
}

func changeValue(x *int) {
	*x = 7;
} */

//--------------------------------------------------------------------------------//

//PRINTF

/* func main() {
	var name string = "Aryya Paul" //declare string
	const pi float64 = 3.141553723 // declare float
	win := true //declare boolean
	x := 5 //declare integer

	fmt.Println(len(name))
	fmt.Println(name + " is a chill dude")

	fmt.Printf("%f \n", pi) //%f is to print float
	fmt.Printf("%.3f \n", pi) //%.3f is to print float upto 3 decimal places (rounds up)

	fmt.Printf("%T \n", name) //%T is to print type of variable
	fmt.Printf("%t \n", win) //%t is to print boolean

	fmt.Printf("%d \n", x) //%d is to print integer

	fmt.Printf("%b \n", 25) //%b is to print binary of 25
	fmt.Printf("%c \n", 34) //%c is to print characters related to the ascii value
	fmt.Printf("%x \n", 15) //%x is to print hex code of 10

	fmt.Printf("%e \n", pi) //%e is to print exponential notation of pi
} */

//--------------------------------------------------------------------------------//

//LOOPS

/* func main() {
	for i := 1; i<=10; i++ {
		fmt.Println(i)
	}
}

func main() {
	i := 1
	for i<=10 {
		fmt.Println(i)
		i++
	}
}


//--------------------------------------------------------------------------------//


//NESTED LOOPS

func main() {
	for i:=1; i<5; i++ {
		for j :=1; j<i; j++ {
			fmt.Printf("*") //Printf: a formatted triangle
			//fmt.Println("*")
			//Println: normal pattern
		}
		fmt.Println()
	}
} */

//--------------------------------------------------------------------------------//

//DECISION MAKING

/* func main() {
	age := 18

	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You cannot vote")
	}
}

func main() {
	age := 80

	switch age {
		case 16: fmt.Println("Prepare for college")
		case 18: fmt.Println("Don't run after girls")
		case 20: fmt.Println("Get yourself a job!")
		default: fmt.Println("Are you even alive?")
	}
} */

//--------------------------------------------------------------------------------//

//ARRAYS

/* func main() {

	var EvenNum[5] int

	EvenNum[0] = 0
	EvenNum[1] = 2
	EvenNum[2] = 4
	EvenNum[3] = 6
	EvenNum[4] = 8

	fmt.Println(EvenNum[2])
} */

/* func main() {

	EvenNum := [5]int{0,2,4,6,8} //another way to initialize an array
	fmt.Println(EvenNum[2])

	for _, value := range EvenNum { //_ is used to ignore the index, instead of using i
		fmt.Println(value)
	}

	for i, value := range EvenNum { //using i
		fmt.Println(i, value)
	}
} */

//--------------------------------------------------------------------------------//

//SLICES

/* func main() {
	numSlice := []int{5,4,3,2,1}

	sliced := numSlice[2:5] //slice of numslice from index 3 to 5
	fmt.Println(sliced)

	slice2 := numSlice[:2] //slice of numslice from index 0 to 3
	fmt.Println(slice2)

	slice3 := numSlice[0:] //slice of numslice from index 0 to end
	fmt.Println(slice3)

	slice4 := make([]int, 5, 10) //make a slice of 5 elements and 10 capacity
	copy(slice4, numSlice) //copy numslice to slice4
	fmt.Println(slice4)

	slice5 := append(numSlice, 3, 0, -1) //append 3, 0, -1 to numslice
	fmt.Println(slice5)
} */

//--------------------------------------------------------------------------------//

//MAPS

//like a dictionary, provides key value pairs
//used for hash maps

/* func main() {
	StudentAge := make(map[string]int)

	StudentAge["Aryya"] = 23
	StudentAge["Saurabh"] = 27
	StudentAge["Prerna"] = 21
	StudentAge["Akrati"] = 19
	StudentAge["Sahiti"] = 42
	StudentAge["Kirti"] = 20

	fmt.Println(StudentAge) //prints all values in map

	fmt.Println(StudentAge["Kirti"]) //prints value of kirti's age

	fmt.Println(len(StudentAge)) //prints length of map
} */

//MAPS INSIDE OF MAPS
//used when coding for industries
/*
func main() {
	superhero := map[string]map[string]string{ //the third string is the return type
		"Superman" : map[string]string{
			"RealName" : "Clark Kent",
			"City" : "Metropolis",
		},
		"Batman" : map[string]string{
			"RealName" : "Bruce Wayne",
			"City" : "Gotham City",
		},
	}

	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}
}
*/

//--------------------------------------------------------------------------------//

//FUNCTIONS

/* func main() {
	x, y := 5, 6

	fmt.Println(add(x, y))
}

func add(num1, num2 int) int {
	return num1 + num2
}
*/

//--------------------------------------------------------------------------------//

// RECURSION
/* func main() {
	num := 5
	fmt.Println(factorial(num))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
*/

//--------------------------------------------------------------------------------//

//go specific keywords

//DEFER
//delays the execution of a function until the surrounding function returns
//multiple defers are pushed into a stack
//executed in LIFO order
//used for cleaning up resources - files, database connections, etc.

/* func main() {
	defer FirstRun() //delayed with defer
	SecondRun()
}

func FirstRun() {
	fmt.Println("First Run")
}

func SecondRun() {
	fmt.Println("Second Run")
}
*/

//--------------------------------------------------------------------------------//

//RECOVER
//used to recover from a panic

/* func main() {
	fmt.Println(div(3,0)) //dividing 3 by 0
	fmt.Println(div(3,5))
}

func div(num1, num2 int) int {
	defer func() {
		fmt.Println(recover()) //recover from panic
	}()
	solution := num1 / num2
	return solution
}

func devPanic() {
	defer func() {
		fmt.Println(recover()) //recover from panic
	}()
}
*/

//--------------------------------------------------------------------------------//

//PANIC
//similar to throwing exception in other languages
//panic stops the execution of the program
//deferred functions are executed even after panic

//--------------------------------------------------------------------------------//

//VARIADIC FUNCTIONS
//accepts any number of arguments

/* func main() {
	fmt.Println(addmeup(10, 20, 30, 40, 50))
}

func addmeup(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
} */

//--------------------------------------------------------------------------------//

//STRUCTURES
//used to create a new type
//used to group related data together

/* func main() {
	// rect1 := Rectangle{height: 10, width: 5}
	rect1 := Rectangle{10, 5}
	fmt.Println(rect1.height)
	fmt.Println(rect1.width)

	fmt.Println(rect1.area())
}

type Rectangle struct {
	height float64
	width  float64
}

func (rect *Rectangle) area() float64 {
	return rect.height * rect.width
}
*/

//--------------------------------------------------------------------------------//

//INTERFACES
//used to define a set of methods
//that a struct type must implement

/* func main() {
	rect := Rectangle{50, 60}
	circ := Circle{7}

	fmt.Println("Area of rectangle is", getArea(rect))
	fmt.Println("Area of circle is", getArea(circ))
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r1 Rectangle) area() float64 {
	return r1.height * r1.width
}

func (c1 Circle) area() float64 {
	return math.Pi * math.Pow(c1.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
} */

//------------------------------------------------//

//FILE I/O
//reading and writing files

/* import (
	"fmt"
	"log"
	"os"
	//"io/ioutil" deprecated in go 1.16
)

func main() {
	file, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal(err) //see any errors
	}
	file.WriteString("Hello, World!")
	file.Close()

	stream, err := os.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	s1 := string(stream)
	fmt.Println(s1)
} */

//--------------------------------------------------------------------------------//

//SIMPLE WEB SERVER
//go gives us a web server package
//with its http library

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", handler2)
	http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page\n") //w is the response writer, r is the request
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

//go to localhost:8080 in browser
//go to localhost:8080/Hello in browser

//--------------------------------------------------------------------------------//
