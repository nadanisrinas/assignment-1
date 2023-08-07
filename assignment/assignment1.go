package assignment

import (
	"fmt"
	"os"
)

type Student struct {
	id      int
	name    string
	address string
	job     string
	reason  string
}

func AssignmentStruct() {
	//----- with struct -----
	var students = []struct {
		data Student
	}{
		{data: Student{id: 0, name: "nada", address: "Jl. selat sunda"}},
		{data: Student{id: 1, name: "yui", address: "Jl. kebagusan", reason: "switch career to programmer"}},
		{data: Student{id: 2, name: "", reason: "advance my career"}},
	}
	//find more than 1
	if len(os.Args) > 1 {

		var datafound []struct{ data Student }
		for i := 1; i < len(os.Args); i++ {
			for j := 0; j < len(students)-1; j++ {
				if students[j].data.name == os.Args[i] {
					datafound = append(datafound, students[j])
				}
			}
		}

		if len(datafound) > 0 {
			fmt.Println(datafound)
		} else {
			fmt.Println("data not found")
		}
		//find only 1
	} else if len(os.Args) < 1 {
		findName := os.Args[1]
		var datafound struct{ data Student }
		for _, student := range students {
			if student.data.name == findName {
				datafound = student
			}
		}
		if datafound.data.name != "" {

			fmt.Println(datafound)
		} else {
			fmt.Println("data not found")
		}
	} else {
		fmt.Println("no args")
	}

}
