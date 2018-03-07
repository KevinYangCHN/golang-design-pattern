package facade

import "fmt"

type AModuleApi struct {
}

func (*AModuleApi) TestA() {
	fmt.Print("test A\n")
}

type BModuleApi struct {
}

func (*BModuleApi) TestB() {
	fmt.Print("test B\n")
}

type Facade struct {
}

func (*Facade) test() {
	a := &AModuleApi{}
	a.TestA()

	b := &BModuleApi{}
	b.TestB()
}
