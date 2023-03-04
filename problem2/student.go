package main

type student struct {
	name    string
	id      string
	major   string
	age     int
	courses []course
}

func newStudent(name string, id string) student {
	s := student{
		name:    name,
		id:      id,
		major:   "",
		age:     0,
		courses: []course{},
	}

	return s
}

// add course
func (s *student) addCourse(courseName string, creditHours int, grade float64) {
	s.courses = append(s.courses, newCourse(courseName, creditHours, grade))
	// actually (*s).courses = append(s.courses, course)
}

// calculate gpa
func (s student) calculateGPA() int {
	GPA := 0
	if len(s.courses) != 0 {
		for _, v := range s.courses {
			GPA += int(v.grade)
		}
		GPA /= len(s.courses)
	}
	return GPA
}
