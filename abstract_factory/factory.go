package main

import "fmt"

type DigtalFactory interface {
	CreatePhone(dpType string) DigtalProduct
	CreatePC(dpType string) DigtalProduct
}

type DigtalProduct interface {
	Make()
}

// 数码工厂A
type DigtalFacA struct {
}

// 数码工厂B
type DigtalFacB struct {
}

type Phone struct {
}

type Pc struct {
}

func (ph *Phone) Make() {
	fmt.Println("phone")
}
func (pc *Pc) Make() {
	fmt.Println("pc")
}

func (df *DigtalFacA) CreatePhone() DigtalProduct {
	return &Phone{}
}

func (df *DigtalFacB) CreatePc() DigtalProduct {
	return &Pc{}
}

func (df *DigtalFacA) CreatePc() DigtalProduct {
	return &Pc{}
}

func (df *DigtalFacB) CreatePhone() DigtalProduct {
	return &Phone{}
}

func main() {
	fa := &DigtalFacA{}
	fa.CreatePc().Make()
	fa.CreatePhone().Make()

	fb := &DigtalFacB{}
	fb.CreatePhone().Make()
	fb.CreatePc().Make()
}
