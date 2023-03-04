package main

import (
	"fmt"
	"reflect"
)

var programMenu = "\n1. Add a student \n2. Find a student by id  \n3. Find a student by name \n4. List students by course \n5. Add a course to a student \n6. Print the number of transactions executed \n7. Exit "

type program struct {
	database            database
	lastStudent         *student
	transactionExecuted int
}

func newProgram(database database) program {
	p := program{
		database:            database,
		lastStudent:         &(student{}),
		transactionExecuted: 0,
	}

	return p
}

func (p *program) run() {
	for {
		fmt.Println(programMenu)
		fmt.Println("input option:")
		var i int
		fmt.Scanf("%d", &i)

		if i == 1 {
			var name string
			fmt.Println("input student name:")
			fmt.Scanf("%s", &name)

			var id string
			fmt.Println("input student id:")
			fmt.Scanf("%s", &id)

			var major string
			fmt.Println("input student major:")
			fmt.Scanf("%s", &major)

			var age int
			fmt.Println("input student age:")
			fmt.Scanf("%d", &age)

			s := newStudent(name, id)
			s.major = major
			s.age = age

			// s.addCourse("opt", 4, 4) s和database都有

			p.database.students = append(p.database.students, s)

			// s.addCourse("opt", 4, 4) 只有s有，database就没有了

			// p.database.students[p.database.num_student()-1].addCourse("opt", 4, 4) //s没有，database有了

			p.lastStudent = &(p.database.students[p.database.num_student()-1])

			p.transactionExecuted += 1

			fmt.Printf("Student %v added", s)

			fmt.Printf("database:", p.database)
		}

		if i == 2 {
			fmt.Println("input student id:")
			var i string
			fmt.Scanf("%s", &i)
			index, s := p.database.findStudentById(i)
			if index == -1 {
				fmt.Println("Not found")
			} else {
				fmt.Println("Student")
				fmt.Printf("\tName: %s\n\tID: %s\n\tMajor: %s\n\tAge: %d\n\n", s.name, s.id, s.major, s.age)
				fmt.Println("\tCourses:", s.courses)
				p.lastStudent = &(p.database.students[index])
			}
			p.transactionExecuted += 1
		}

		if i == 3 {
			fmt.Println("input student name:")
			var i string
			fmt.Scanf("%s", &i)
			index, s := p.database.findStudentByName(i)
			if index == -1 {
				fmt.Println("Not found")
			} else {
				fmt.Println("Student")
				fmt.Printf("\tName: %s\n\tID: %s\n\tMajor: %s\n\tAge: %d\n\n", s.name, s.id, s.major, s.age)
				fmt.Println("\tCourses:", s.courses)
				p.lastStudent = &(p.database.students[index])
			}
			p.transactionExecuted += 1
		}

		if i == 4 {
			fmt.Println("input course name:")
			var i string
			fmt.Scanf("%s", &i)
			s := p.database.findStudentsByCourse(i)
			if len(s) == 0 {
				fmt.Println("No student found")
			} else {
				res := []string{}
				for _, v := range s {
					res = append(res, v.name)
				}
				fmt.Println(res)
			}
			p.transactionExecuted += 1
		}

		if i == 5 {
			if reflect.ValueOf(p.lastStudent).IsZero() {
				fmt.Println("Search or add student before adding a course")
			} else {
				var courseName string
				fmt.Println("input course name:")
				fmt.Scanf("%s", &courseName)

				var creditHours int
				fmt.Println("input creditHours:")
				fmt.Scanf("%d", &creditHours)

				var grade float64
				fmt.Println("input grade:")
				fmt.Scanf("%f", &grade)

				// p.database.students[p.database.num_student()-1].addCourse("opt", 4, 4)
				p.lastStudent.addCourse(courseName, creditHours, grade)
				// fmt.Println("p.laststudent:", p.lastStudent)
				// fmt.Println("database:", p.database) // 没有得到修改
				fmt.Printf("course %s added to student %s\n", courseName, p.lastStudent.name)

				p.transactionExecuted += 1
			}
		}

		if i == 6 {
			fmt.Println(p.transactionExecuted)
			p.transactionExecuted += 1
		}

		if i == 7 {
			break
		} else {
			continue
		}
	}
}
