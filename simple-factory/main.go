package main

import (
	"log"
)

// Factory 工厂角色
// 负责实现创建所有实例的内部逻辑
type Factory struct {
}

// Product 抽象产品角色
type Product interface {
	NameString() string
}

// ConcreteProductA 具体产品角色A
type ConcreteProductA struct {
}

// ConcreteProductB 具体产品角色A
type ConcreteProductB struct {
}

func main() {
	var productA, productB Product
	factory := NewFactory()
	productA = factory.CreateProduct("A")
	productB = factory.CreateProduct("B")
	log.Println(productA.NameString())
	log.Println(productB.NameString())
}

func NewFactory() Factory {
	return Factory{}
}

func (f *Factory) CreateProduct(productString string) Product {
	var pro Product

	// 注意，pro = &ConcreteProductA{}这里，Product是接口类型
	// &ConcreteProductA{}是结构体指针
	// 由于Product是interface，所以不需要给Product也加上*

	// 同时如果返回的是ConcreteProductA{}，而不是指针
	// 但是实现接口的时候定义的接受者是指针类型的，这个时候
	// 并不会自动转换成指针再调用
	// 可能是因为interface吧（如果是单纯的结构体类型而不是用接口类型，是可以自动转换的）
	switch productString {
	case "A":
		pro = &ConcreteProductA{}
	default:
		pro = &ConcreteProductB{}
	}

	return pro
}

func (c *ConcreteProductA) NameString() string {
	return "I am ConcreteProductA."
}

func (c *ConcreteProductB) NameString() string {
	return "I am ConcreteProductB."
}
