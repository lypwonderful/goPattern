package main

import "fmt"

// 1.一个工厂接口
// 2.一个产品接口

type DigtalFactory interface {
	CreateFac(dpType string) DigtalProduct
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

func (df *DigtalFacA) CreateFac(dpType string) DigtalProduct {
	if dpType == "phone" {
		return &Phone{}
	} else {
		return &Pc{}
	}
}

func (df *DigtalFacB) CreateFac(dpType string) DigtalProduct {
	if dpType == "phone" {
		return &Phone{}
	} else {
		return &Pc{}
	}
}

func main() {
	fa := &DigtalFacA{}
	fa.CreateFac("phone").Make()

	fb := &DigtalFacB{}
	fb.CreateFac("pc").Make()
}
