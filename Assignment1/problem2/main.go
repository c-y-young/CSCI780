package main

func main() {

	// myStudent := newStudent("ben", "001")
	// myStudent.addCourse("optimization", 4, 4.0)
	// gpa := myStudent.calculateGPA()
	// fmt.Println(myStudent, "gpa:", gpa)
	// s := student{}
	// fmt.Println(reflect.ValueOf(s).IsZero())

	db := newDatabase()

	p := newProgram(db)
	p.run()
}
