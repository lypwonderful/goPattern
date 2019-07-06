package main

import "fmt"

// 货物
type goods struct {
	Name  string
	Price float32
	Saler saler
}

func NewGoodsCtx(name string, price float32, s saler) *goods {
	return &goods{
		Name:  name,
		Price: price,
		Saler: s,
	}
}

type saler interface {
	sale(p float32) float32
}

func (g *goods) sale() {

}

func (g *goods) saleStrategy() float32 {
	return g.Saler.sale(g.Price)
}

type normal struct {
}

func (n *normal) sale(p float32) float32 {
	return p
}

type halfPrice struct {
}

func (halfPrice) sale(p float32) float32 {
	return p / 2
}

func main() {
	gCtx := NewGoodsCtx("phone", 10000, &normal{})
	fmt.Println(gCtx.saleStrategy())

	gCtx = NewGoodsCtx("phone", 10000, &halfPrice{})
	fmt.Println(gCtx.saleStrategy())
}
