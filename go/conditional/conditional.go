package main

import(
	"fmt" 
	"strconv"
)

var imc float32

// Author structure
type person struct {
	name   string
	height float32
	weight float32
}

func (p person) getImcDescription() {
	//imc = p.weight / (p.height * p.height)
	if imc = p.weight / (p.height * p.height); imc < 16 {
		fmt.Printf("Person %v is with severe thinness\n", p.name)
	} else if imc < 17 {
		fmt.Printf("Person %v is with moderate thinness\n", p.name)
	} else if imc < 18.5 {
		fmt.Printf("Person %v is with light thinness\n", p.name)
	} else if imc < 25 {
		fmt.Printf("Person %v is healthy\n", p.name)
	} else if imc < 30 {
		fmt.Printf("Person %v is overweight\n", p.name)
	} else if imc < 35 {
		fmt.Printf("Person %v is overweight degree I\n", p.name)
	} else if imc < 40 {
		fmt.Printf("Person %v is overweight degree II\n", p.name)
	} else {
		fmt.Printf("Person %v is overweight degree III\n", p.name)
	}
}

func getMonthString(month int) string{
	switch(month){
		case 1: return "January"
		case 2: return "February"
		case 3: return "March"
		case 4: return "April"
		case 5: return "May"
		case 6: return "June"
		case 7: return "July"
		case 8: return "August"
		case 9: return "September"
		case 10: return "October"
		case 11: return "November"
		case 12: return "December"
		default: return "Value "+strconv.Itoa(month)+" is not within 1 and 12"
	}
}

func main() {

	// Initializing the values
	// of the person structure
	p := person{
		name:   "Eduardo",
		height: 1.85,
		weight: 100.0,
	}

	p.getImcDescription()
	fmt.Println(getMonthString(1))
	fmt.Println(getMonthString(12))
	fmt.Println(getMonthString(-1))
	
}
