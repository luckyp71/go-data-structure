package main

import "fmt"

type Employee struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// Linear Search over list of int
func intLinearSearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

// Linear Search over list of string
func stringLinearSearch(data []string, key string) bool {
	for _, datum := range data {
		if datum == key {
			return true
		}
	}
	return false
}

// Linear search over list of object
func objLinearSearch(data []Employee, key string) bool {
	var employee_names []string
	for _, employee := range data {
		employee_names = append(employee_names, employee.Name)
	}

	for _, name := range employee_names {
		if name == key {
			return true
		}
	}

	return false
}

func main() {
	items := []int{1, 2, 4, 5, 10, 22, 44, 55, 66, 33, 22, 33, 44, 55, 33, 44}
	fmt.Println(intLinearSearch(items, 30))

	data := []string{"Customer 1", "Customer 2", "Customer 3", "Customer 4"}
	fmt.Println(stringLinearSearch(data, "Customer 2"))

	emp_data := []Employee{{"Employee 1", "Address 1"}, {"Employee 2", "Address 2"}}

	fmt.Println(objLinearSearch(emp_data, "Employee 2"))
}
