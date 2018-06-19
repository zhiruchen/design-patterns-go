package design

import "testing"

func TestVisitor(t *testing.T) {
	vt := newVisitorImpl(
		&ElementA{},
		&ElementB{},
		&ElementC{},
	)

	vt.visit()
}
