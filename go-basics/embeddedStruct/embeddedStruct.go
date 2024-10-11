package main

import "fmt"

type skills []string

type human struct {
	age    uint8
	name   string
	weight uint8
	phone  string
}

type student struct {
	human
	skills
	int
	specialty string
	phone     string
}

func main() {
	mark := student{
		human: human{
			age:    12,
			name:   "Mark",
			weight: 32,
		},
		specialty: "CS",
	}

	// or
	// mark := student{human{12, "Mark", 30}, "AI"}

	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His specialty is ", mark.specialty)

	mark.specialty = "AI"
	fmt.Println("Mark changed his specialty")
	fmt.Println("His specialty is ", mark.specialty)

	fmt.Println("Mark become old. He is not an athlete anymore")
	mark.age = 46
	mark.weight += 60
	fmt.Println("His age is", mark.age)
	fmt.Println("His weight is", mark.weight)

	mark.skills = []string{"Ai"}
	mark.skills = append(mark.skills, "CS", "Data")
	fmt.Println("Mark's skills are", mark.skills)

	fmt.Println("Information about Mark:", mark.human)

	mark.int = 5
	fmt.Println("Mark.int =", mark.int)

	// Phone is override
	mark.human.phone = "777-444-XXXX"
	mark.phone = "333-222"
	fmt.Println("Mark's work phone is:", mark.phone)
	fmt.Println("Mark's personal phone is:", mark.human.phone)
}
