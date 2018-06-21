package design

import "fmt"

// https://stackoverflow.com/questions/2604169/visitor-patterns-purpose-with-examples
// Visitor visitor
type Visitor interface {
	VisitElementA(e *ElementA)
	VisitElementB(e *ElementB)
	VisitElementC(e *ElementC)
}

type Accepter interface {
	Accept(v Visitor)
}

type ElementA struct {
}

func (a *ElementA) Accept(v Visitor) {
	v.VisitElementA(a)
}

type ElementB struct {
}

func (b *ElementB) Accept(v Visitor) {
	v.VisitElementB(b)
}

type ElementC struct {
}

func (c *ElementC) Accept(v Visitor) {
	v.VisitElementC(c)
}

type visitorImpl struct {
	a *ElementA
	b *ElementB
	c *ElementC
}

func newVisitorImpl(a *ElementA, b *ElementB, c *ElementC) *visitorImpl {
	return &visitorImpl{
		a: a,
		b: b,
		c: c,
	}
}

func (impl *visitorImpl) VisitElementA(e *ElementA) {
	fmt.Println("visit visitorImpl VisitElementA")
}

func (impl *visitorImpl) VisitElementB(e *ElementB) {
	fmt.Println("visit visitorImpl VisitElementB")
}

func (impl *visitorImpl) VisitElementC(e *ElementC) {
	fmt.Println("visit visitorImpl VisitElementC")
}

func (impl *visitorImpl) visit() {
	impl.a.Accept(impl)
	impl.b.Accept(impl)
	impl.c.Accept(impl)
}
