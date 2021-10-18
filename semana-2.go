package main

import "fmt"

func Exercicio_2_1() {
	//função para descobrir o resultado de uma variável ao final
	//da operação
	var num_a, num_b int
	var greaterThan bool
	num_a = 2
	num_b = 3
	num_a += num_b
	num_b += num_a
	num_a *= num_b
	greaterThan = num_a > 3*num_b
	fmt.Printf("Final result was %#v\n", greaterThan)
}

func Exercicio_2_2(y int) {
	//função para corrigir um código
	var _num_a int
	var _num_z float64
	_num_a = 2
	y += _num_a
	_num_z = float64(_num_a) / float64(y)
	fmt.Printf("Given %d, final result was %.2f\n", y, _num_z)
}

type Student struct {
	name, email, address, home_address string
}

func Exercicio_2_3(_student Student) {
	fmt.Printf("Student %s's email is %s\n", _student.name, _student.address)
}

func Exercicio_2_4() {
	var array_x [2]int
	var array_y [2]int
	array_x[0] = 1
	array_y[0] = 2
	array_x[1] += array_y[0]
	array_y[1] += array_x[1] + array_x[0]
	fmt.Printf("Final result was %d\n", array_y)
}

func create_contacts(student_1, student_2, student_3 Student) []Student {
	_student_array := make([]Student, 3)
	_student_array[0] = student_1
	_student_array[1] = student_2
	_student_array[2] = student_3
	return _student_array
}

func Exercicio_2_5(student_array []Student) {
	fmt.Printf("Student %s's address is %s, and student %s's addresss is %s\n", student_array[0].name, student_array[0].email, student_array[1].name, student_array[1].email)
}

type PrimaryStudent struct {
	name string
	age  int
}

func main() {
	Exercicio_2_1()
	Exercicio_2_2(5)
	new_student := Student{"Marco", "@email", "address", "home address"}
	Exercicio_2_3(new_student)
	Exercicio_2_4()
	new_student_1 := Student{"Marco", "@email", "address", "home address"}
	new_student_2 := Student{"Preta", "@email", "address", "home address"}
	new_student_3 := Student{"Milka", "@email", "address", "home address"}
	lista_contatos := create_contacts(new_student_1, new_student_2, new_student_3)
	Exercicio_2_5(lista_contatos)
	//pStudent_1 := PrimaryStudent{"Name 1", 15}
	//pStudent_2 := PrimaryStudent{"Name 2", 16}
	//pStudent_3 := PrimaryStudent{"Name 3", 17}
}
