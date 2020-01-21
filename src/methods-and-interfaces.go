package main

import "fmt"

// Vector2D : Represents a 2D vector
type Vector2D struct {
	x, y int
}

// Vector3D : Represents a 3D vector
type Vector3D struct {
	x, y, z int
}

// Geometry : Geometry interface
type Geometry interface {
	Scale(f int)
	Add(v interface{})
}

// Person : Person interaface
type Person struct {
	Name string
	Age  int
}

func (person Person) String() string {
	return fmt.Sprintf("Person data:\nName: %s \nAge %d", person.Name, person.Age)
}

// Add : [METHOD from Vector2D] Adds a vector to itself
func (v *Vector2D) Add(vector2 Vector2D) {
	v.x += vector2.x
	v.y += vector2.y
}

// Using a POINTER as a Receiver is most used, so that you dont need to be reasigning the variable everytime
// It mutates the receiver instead

// Scale : [METHOD from Vector2D] Scales a vector
func (v *Vector2D) Scale(f int) {
	v.x *= f
	v.y *= f
}

// Add : [METHOD from Vector3D] Adds a vector to itself
func (v *Vector3D) Add(vector3 Vector3D) {
	vector3.x += vector3.x
	vector3.y += vector3.y
	vector3.z += vector3.z
}

// Scale : [METHOD from Vector3D] Scales a vector
func (v *Vector3D) Scale(f int) {
	v.x *= f
	v.y *= f
	v.z *= f
}

func scaleGeometry(g Geometry) {
	g.Scale(2)
}

func main() {
	// Go has a few ways of allocating memory
	// &T{...}, &someLocalVar, new, make

	v := Vector2D{0, 5}
	v.Scale(2)

	//// Reason for using RECEIVER POINTERS
	// Avoid copying the value on each method call.
	// this can be more efficient if the receiver is a large struct, for example.

	fmt.Println(v)

	//
	// INTERFACES
	//

	// In GO to implement an interface just create the methods of that interface for a struct

	// We should design Interfaces with the actions the type should have, not the properties

	// It might seem dumb, but they are mostly used for GENERALIZING multiple types into one
	// Good blogpost about interfaces: https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
	//vectors := []Geometry{&Vector2D{1, 2}, &Vector3D{1, 2, 3}}

	// The interface{} as an arg means that it can be any type, since all types need to implement all
	// methods of the interface to `implement it` and  interface{} has none.

	// It's important to note that an interface value is two words wide and it contains a pointer to the underlying data.

	var dummyInterface = Vector2D{1, 2}
	// (value, type)
	// ({1 2}, main.Vector2D)
	describeInterface(dummyInterface)

	// Nil interfaces cause a run-time error
	var geo Geometry
	describeInterface(geo)
	// geo.Scale(1) <-- this will throw

	// The empty interface `interface{}`, used as an ANY type (Every type implements at least zero methods.)

	//
	// INTERFACE TYPE ASSERTION
	//
	var name interface{} = "TIAGO"
	value, ok := name.(string)

	fmt.Println(value, ok)

	do(12)
	do(Vector2D{1, 2})

	//
	// STRINGERS
	//

	// A Stringer is a type that can describe itself as a string
	// for instance, the fmt package, looks for this interface to print values

	me := Person{"Tiago Taquelim", 20}
	fmt.Println(me) // this will call the String() method, good for debuging

}

func describeInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
