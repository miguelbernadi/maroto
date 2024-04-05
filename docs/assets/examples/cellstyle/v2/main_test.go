package main

import (
	"testing"

	"github.com/miguelbernadi/maroto/v2/pkg/test"
)

func TestGetMaroto(t *testing.T) {
	// Act
	sut := GetMaroto()

	// Assert
	test.New(t).Assert(sut.GetStructure()).Equals("examples/cellstyle.json")
}
