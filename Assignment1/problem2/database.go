package main

type database struct {
	students []student
}

func newDatabase() database {
	db := database{
		students: []student{},
	}

	return db
}

// add a studnet
func (db *database) addStudent(name string, id string) student {
	s := newStudent(name, id)
	db.students = append(db.students, s)
	return s
}

// find a student by the student's id
func (db database) findStudentById(id string) (int, student) {
	s := student{}
	for i, v := range db.students {
		if v.id == id {
			return i, v
			break
		}
	}
	return -1, s
}

// find a student by the student's name
func (db database) findStudentByName(name string) (int, student) {
	s := student{}
	for i, v := range db.students {
		if v.name == name {
			return i, v
		}
	}
	return -1, s
}

// fund list all students who are taking a course
func (db database) findStudentsByCourse(courseName string) []student {
	targetStudents := []student{}
	for _, v := range db.students {
		for _, c := range v.courses {
			if c.courseName == courseName {
				targetStudents = append(targetStudents, v)
				break
			}
		}
	}

	return targetStudents
}

// return the number of students in databases
func (db database) num_student() int {
	return len(db.students)
}
