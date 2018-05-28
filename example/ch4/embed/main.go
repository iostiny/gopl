// Embed demonstrates basic struct embeding.
package main

import "fmt"

type Point struct{ X, Y int }

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w = Wheel{Circle{Point{3, 4}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point: Point{
				X: 3, Y: 4,
			},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w)
	// Output:
	// main.Wheel{Circle:main.Circle{Point:main.Point{X:3, Y:4}, Radius:5}, Spokes:20}

	w.X = 42

	fmt.Printf("%v\n", w)
	// Output:
	// {{{42 4} 5} 20}
}

// Be careful the diff between `%#v`` and `%v``, `#` adverb causes
// dispaly value in a form to Go syntax.