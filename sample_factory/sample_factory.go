package main

import "fmt"

type Phoner interface {
	Show()
}

type xm struct {
	Name string
}

type zte struct {
	Name string
}

func Create(typeName int) Phoner {
	switch typeName {
	case 0:
		return &zte{}
	default:
		return &xm{}
	}
}

func (xm *xm) Show() {
	fmt.Println("this is a xm phone")
}

func (z *zte) Show() {
	fmt.Println("this is a zte phone")
}

func main() {
	x := Create(1)
	x.Show()
}
