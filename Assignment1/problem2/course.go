package main

type course struct {
	courseName  string
	creditHours int
	grade       float64
}

func newCourse(courseName string, creditHours int, grade float64) course {
	c := course{
		courseName:  courseName,
		creditHours: creditHours,
		grade:       grade,
	}

	return c
}
