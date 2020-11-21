package main

// import "fmt"

// //Printable..
// type Printable interface {
// 	print()
// }

// type person struct {
// 	name string
// }

// type figure struct {
// 	name string
// }

// func (f *figure) print() {
// 	fmt.Println("[figure]", f.name)
// }

// func (p *person) print() {
// 	fmt.Println("[person]", p.name)
// }

// func invokePrint(p Printable) {
// 	p.print()
// }

// type close func()

// func ejemplo() {
// 	// p := &person{name: "Juan"}
// 	// f := &figure{name: "Cube"}
// 	// invokePrint(p)
// 	// invokePrint(f)
// 	f := func() {
// 		fmt.Println("Hello world")
// 	}

// 	invokeClose(f)
// }

// func invokeClose(c close) {
// 	c()
// }
