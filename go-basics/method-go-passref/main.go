package main

import "fmt"

const (
	WHITE = iota
	RED
	BLACK
	GREEN
	BLUE
)

type box struct {
	width, height, depth float64
	color                color
}

type color byte
type boxList []box

func (b box) volume() float64 {
	return b.width * b.height * b.depth
}

func (b *box) setColor(c color) {
	b.color = c
}

func (bl boxList) biggestsColor() (c color) {
	c = color(WHITE)
	v := 0.00
	for _, box := range bl {
		if box.volume() > v {
			v = box.volume()
			c = box.color
		}
	}
	return
}

// Slice is passed by ref by default
func (bl boxList) paintItBlack() {
	for i := range bl {
		bl[i].setColor(BLACK)
	}

	// This does not work
	// for _, box := range bl {
	// 	box.setColor(BLACK)
	// }
}

// Receiver can be any type
func (c color) string() (cStr string) {
	colorStrings := []string{
		"WHITE",
		"RED",
		"BLACK",
		"GREEN",
		"BLUE"}
	cStr = colorStrings[c]
	return
}

func main() {
	boxes := boxList{
		{
			width:  10,
			height: 10,
			depth:  10,
			color:  WHITE,
		},
		{
			width:  10,
			height: 20,
			depth:  30,
			color:  GREEN,
		},
		{
			width:  10,
			height: 10,
			depth:  10,
			color:  BLUE,
		},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.string())
	fmt.Println("The biggest one is", boxes.biggestsColor().string())

	// Let's paint them all black
	boxes.paintItBlack()

	fmt.Println("The color of the second one is", boxes[1].color.string())
	fmt.Println("Obviously, now, the biggest one is", boxes.biggestsColor().string())
}
