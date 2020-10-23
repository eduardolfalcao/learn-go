
package main

import "fmt"

// Author structure 
type person struct { 
    name      string 
    height    float32  
    weight float32
}

func (p person) getImcDescription() { 
	float32 imc = p.weight / p.height * p.height
	if imc < 16 {
		fmt.Printf("Person %V is with severe thinness\n", p.name)
	} else if imc < 17 {
		fmt.Printf("Person %V is with moderate thinness\n", p.name)
	} else if imc < 18.5 {
		fmt.Printf("Person %V is with light thinness\n", p.name)
	} else if imc < 25 {
		fmt.Printf("Person %V is healthy\n", p.name)
	} else if imc < 30 {
		fmt.Printf("Person %V is overweight\n", p.name)
	} else if imc < 35 {
		fmt.Printf("Person %V is overweight degree I\n", p.name)
	} else if imc < 40 {
		fmt.Printf("Person %V is overweight degree II\n", p.name)
	} else {
		fmt.Printf("Person %V is overweight degree III\n", p.name)
	}	
}

func main() {
	
	// Initializing the values 
	// of the person structure 
	p := person{ 
		name:      "Eduardo",
		height:    1.85,
		weight:    100.0}
	
	p.getImcDescription()
}

