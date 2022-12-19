/*
  Define Type thta Satisfies Multiple Interfaces

  interfaces allow any user-defined type to satisfy interface types at once.

  Using *Type Assertion* you can get a value of a concrete type back and you can call methods on it thta are
  defined on ohter interface, but aren't part of the interface satisfying.
*/

package interfaces

import "fmt"

type Polygons interface {
	Perimeter()
}

type Object interface {
	NumberOfSide()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("Perimeter of Pentagon ", 5*p)
}

func (p Pentagon) NumberOfSide() {
	fmt.Println("Pengaton has 5 sides")
}

func DefineTypeSatisfiesMultipleInterfaces() {
	var p Polygons = Pentagon(50)
	p.Perimeter()

	var o Pentagon = p.(Pentagon)
	o.NumberOfSide()

	var obj Object = Pentagon(50)

	obj.NumberOfSide()

	var pent Pentagon = obj.(Pentagon)

	pent.Perimeter()

}

/*
  When a user-defined type implements the set of methods declared by an interface type, values of the user-defined
  type can be assigned to values of the interface type. this assignment stores the value of the user-defined
  type into the interface value, When a  method call is made against an interface value, the equivalent
  method for the stored user-defined value will be executed. Since any user-defined type can implement any
  interface, method calls against an interface value are polymorphic in nature. The user-defined type in this
  relationship is often reffered as *concrete type*.
*/
