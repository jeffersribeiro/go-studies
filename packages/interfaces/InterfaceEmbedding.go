package interfaces

import "fmt"

type Geometry2 interface {
	Edges() int
}

type Polygons2 interface {
	Geometry2 // Interface embedding another interface
}

type Pentagon2 int
type Hexagon2 int
type Octagon2 int
type Decagon2 int

func (p Pentagon2) Edges() int { return 5 }
func (h Hexagon2) Edges() int  { return 6 }
func (o Octagon2) Edges() int  { return 8 }
func (d Decagon2) Edges() int  { return 10 }

func InterfaceEmbedding() {
	p := new(Pentagon2)
	h := new(Hexagon2)
	o := new(Octagon2)
	d := new(Decagon2)

	polygons := [...]Polygons2{p, h, o, d}
	for i := range polygons {
		fmt.Println(polygons[i].Edges())
	}
}

/*
  When one type is embedded into another type, the methods
*/
