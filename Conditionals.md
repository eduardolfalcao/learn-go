# Conditionals in Go

Go's if statements need not be surrounded by parentheses ( ), but the braces { } are required.

```go
func (p person) getImcDescription() {
        imc = p.weight / (p.height * p.height)
        if imc < 16 {
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

```

Like for, the if statement can start with a short statement to execute before the condition.
Variables declared by the statement are only in scope until the end of the if.
Variables declared inside an if short statement are also available inside any of the else blocks.

```go
func (p person) getImcDescription() {
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

```

Switch-case structure in go follows the same well-known notation of java.

```go
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
```
