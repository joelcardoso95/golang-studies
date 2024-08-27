package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"strings"
)

// basic types (numbers, strings, booleans)
var myInt int
var myUint uint // just positive numbers
var myFloat float32
var myFloat64 float64

// aggregate types (array, struct)

// reference types (pointers, slices, maps, functions, channels)

// interface types
type AnimalInterface interface {
	Says() string
	HowManyLegs() int
}

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car2 struct {
	Make       string
	Model      string
	Year       int
	isElectric bool
	isHybrid   bool
	Vehicle    Vehicle
}

func (vehicle Vehicle) showDetails() {
	fmt.Println("Number of wheels", vehicle.NumberOfWheels)
	fmt.Println("Number of passengers", vehicle.NumberOfPassengers)
}

func (car2 Car2) show() {
	fmt.Println("Make", car2.Make)
	fmt.Println("Model", car2.Model)
	fmt.Println("Year", car2.Year)
	fmt.Println("Is electric", car2.isElectric)
	fmt.Println("Is hybrid", car2.isHybrid)
	car2.Vehicle.showDetails()
}

func (d *Cat) Says() string {
	return d.Sound
}

func (d *Cat) HowManyLegs() int {
	return d.NumberOfLegs
}

func (a *Animal) MakeSound() {
	fmt.Printf("%s says %s\n", a.Name, a.Sound)
}

var keyPressChan chan rune

func main() {
	myInt = 10
	myUint = 20

	myFloat = 10.2
	myFloat64 = 200.4

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Hello, World!"

	log.Println(myString)

	// Go strings is immutable
	myString = "Hello, World! " + "I'm a string!"

	var myBool = true
	myBool = false
	log.Println(myBool)

	// Arrays
	var myStrings [3]int
	myStrings[0] = 1
	myStrings[1] = 2
	myStrings[2] = 3

	fmt.Println(myStrings)

	// Structs
	myCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "BMW",
		Model:         "M3",
		Year:          2020,
	}

	fmt.Printf("my car is a %d %s %s\n", myCar.Year, myCar.Make, myCar.Model)

	// Pointers
	x := 10

	myFirstPoint := &x // memory address of x

	fmt.Println("x is ", x)
	fmt.Println("myFirstPoint is ", myFirstPoint)

	*myFirstPoint = 20 // content value of x

	fmt.Println("x is ", x)

	changeValueOfPointer(&x)

	fmt.Println("After function call x is ", x)

	// Slices
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "bird")
	animals = append(animals, "fish")

	fmt.Println(animals)

	for _, x := range animals {
		fmt.Println(x)
	}

	fmt.Println("Length of animals is ", len(animals))
	fmt.Println("The first two elements of animals are ", animals[:2])

	animals = deleteFromSlice(animals, 2)
	fmt.Println(animals)

	// Maps
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	delete(intMap, "two")

	el, ok := intMap["two"]
	if ok {
		fmt.Println("two is in the map and its value is ", el)
	} else {
		fmt.Println("two is not in the map")
	}

	// Functions
	c := addTwoNumbers(10, 20)
	fmt.Println(c)

	total := sumManyNumbers(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)

	dog := Animal{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}
	dog.MakeSound()

	// channels
	keyPressChan = make(chan rune)

	go listenForKeyPress()
	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if strings.ToLower(string(char)) == "q" {
			break
		}

		keyPressChan <- char
	}

	// interfaces
	dog2 := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	cat := Cat{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
		HasTail:      true,
	}

	riddle(&dog2)
	riddle(&cat)

	// expressions
	age := 10
	name := "John"
	rightHanded := true

	rightHanded = !rightHanded

	fmt.Printf("a is %d, name is %s, rightHanded is %t\n", age, name, rightHanded)

	ageInTenYears := age + 10
	fmt.Printf("In ten years, John will be %d\n", ageInTenYears)

	isATeenager := age >= 13
	fmt.Printf("Is John a teenager? %t\n", isATeenager)

	// boolean expressions
	// > < >= <= == != && ||
	apples := 10
	oranges := 5
	fmt.Println(apples == oranges)
	fmt.Println(apples != oranges)

	jack := Employee{
		Name:     "Jack",
		Age:      30,
		Salary:   50000,
		FullTime: true,
	}

	jill := Employee{
		Name:     "Jill",
		Age:      25,
		Salary:   60000,
		FullTime: false,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, employee := range employees {
		if employee.Age > 30 {
			fmt.Println(employee.Name, "is over 30")
		} else {
			fmt.Println(employee.Name, "is under 30")
		}

		if employee.Age > 30 && employee.Salary > 50000 {
			fmt.Println(employee.Name, "is over 30 and makes more than 50000")
		} else {
			fmt.Println(employee.Name, "is under 30 or makes less than 50000")
		}

		if employee.Age > 30 || employee.Salary > 50000 {
			fmt.Println(employee.Name, "is over 30 or makes more than 50000")
		} else {
			fmt.Println(employee.Name, "is under 30 and makes less than 50000")
		}

		if (employee.Age > 30 || employee.Salary < 5000) && employee.FullTime {
			fmt.Println(employee.Name, "matches our unclear criteria")
		}
	}

	// composition
	suv := Vehicle{
		NumberOfWheels:     4,
		NumberOfPassengers: 5,
	}

	volvoXC90 := Car2{
		Make:       "Volvo",
		Model:      "XC90",
		Year:       2020,
		isElectric: false,
		isHybrid:   true,
		Vehicle:    suv,
	}

	volvoXC90.show()

}

func changeValueOfPointer(num *int) {
	*num = 50
}

func deleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}

func addTwoNumbers(a int, b int) int {
	return a + b
}

func sumManyNumbers(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed ", string(key))
	}
}

func riddle(a AnimalInterface) {
	riddle := fmt.Sprintf(`This animal has %s legs and says %d`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
