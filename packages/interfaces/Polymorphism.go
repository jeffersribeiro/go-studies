package interfaces

import "fmt"

/*
Polymorphism is the ability to write code that can take on different behavior through the implementation of type.

 We have the declaration of strcts named Pentagon, Hexagon, Octagon and Decagon with the
 implementation of the Geometru interface.
*/

type Germetry interface {
	Edges() int
}

type Pengaton struct{}
type Hexagon struct{}
type Octagon struct{}
type Decagon struct{}

func (p Pengaton) Edges() int {
	return 5
}

func (h Hexagon) Edges() int {
	return 6
}

func (o Octagon) Edges() int {
	return 8
}

func (d Decagon) Edges() int {
	return 10
}

func Parameter(geo Germetry, value int) int {
	num := geo.Edges()

	calculation := num * value
	return calculation
}

func Polymorphism() {

	p := new(Pengaton)

	h := new(Hexagon)

	o := new(Octagon)

	d := new(Decagon)

	g := [...]Germetry{p, h, o, d}

	for _, i := range g {
		fmt.Println(Parameter(i, 5))
	}
}

/*
  We have our polymorphic Edges functions that accepts values thta implement the Geometry interface.
  Using polymorphic approach the created here Parameter is used bu each concrete type value
  thats's passed in.
*/
