/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/


package main
import "fmt"

//functions defined 

func function1(name string, age int){
        fmt.Println("Hi ",name,". Your age is ",age )
    }
    
func sumfunc(a int, b int) int {
        return a+b
    }
    
func factorial(n int) int {
        if n==1{
            return 1
        }
        return (n*factorial(n-1))
        
    }    
    
//structs

type Employee struct{
    empnum int 
    fname string
    age int
    salary float64
}

func main() {
    /*
	//creating arrays and variables in two methods
	var array1 = [4]int{1, 2, 3, 4}
	array2 := [10]int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	var var1 int = 20
	var2 := 10
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println()

	//checking array length
	fmt.Println(len(array1))
	fmt.Println()

	//checking mutability//
	array2[2] = 5000
	fmt.Println(array2)
	fmt.Println()

	//array defining elements at index
	array4 := [6]int{3: 30, 5: 50.}
	fmt.Println(array4)
	fmt.Println()

	//slice
	slice1 := []int{}
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(slice1)
	fmt.Println()

	slice2 := []string{"St", "ri", "ng"}
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	fmt.Println(slice2)
	fmt.Println()

	slice3 := array2[2:7]
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println()

	slice4 := make([]int, 4, 8)
	fmt.Println(slice4)
	fmt.Println()

	slice5 := append(slice2, "More", "Strings")
	fmt.Println(slice5)
	fmt.Println()
	
	// if statements
	if 22>21 {
		fmt.Println("if statement example")
		}
	fmt.Println()
	num1:=5
    num2:=10
    if num1>num2 {
    	fmt.Println(num1)
    } else if num2>num1 {
    	fmt.Println(num2)
    } else {
    	fmt.Println("Numbers are equal")
    }
    fmt.Println()
    
    //switch cases
    day:=5;
    switch day{
        case 1:
        fmt.Println("Monday")
        case 2:
        fmt.Println("Tuesday")
        case 5:
        fmt.Println("Friday")
        default:
        fmt.Println("Not a valid day")
        
    }
    fmt.Println()
    
    //for loops
    for i:=0; i<8; i++{
        if i==3{
            break
        }
        if i==1{
            continue
        }
        fmt.Println(i)
    }
    fmt.Println()
    
    
    fmt.Println("Function examples")
    function1("Seb",22)
    function1("John",25)
    fmt.Println("Sum is ",sumfunc(10,5))
    
    
    fmt.Println(factorial(5))
    
    
    //structs
    
    var emp1 Employee
    var emp2 Employee
    var emp3 Employee
    
    emp1.empnum = 1
    emp1.fname = "Seb"
    emp1.age = 21
    emp1.salary = 100.500
    
    emp2.empnum = 2
    emp2.fname = "John"
    emp2.age = 21
    emp2.salary = 300.500
    
    emp3.empnum = 3 
    emp3.fname = "Kiri"
    emp3.age = 21
    emp3.salary = 200.400
    
    fmt.Println("Name of employee ",emp1.empnum," :",emp1.fname)
    fmt.Println("Age of employee ",emp1.empnum," :",emp1.age)
    fmt.Println("Salary of employee ",emp1.empnum," :",emp1.salary)
    fmt.Println()
    
    fmt.Println("Name of employee ",emp2.empnum," :",emp2.fname)
    fmt.Println("Age of employee ",emp2.empnum," :",emp2.age)
    fmt.Println("Salary of employee ",emp2.empnum," :",emp2.salary)
    fmt.Println()
    
    fmt.Println("Name of employee ",emp3.empnum," :",emp3.fname)
    fmt.Println("Age of employee ",emp3.empnum," :",emp3.age)
    fmt.Println("Salary of employee ",emp3.empnum," :",emp3.salary)
    fmt.Println()
    */
}
