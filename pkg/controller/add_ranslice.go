package controller

import (
	"github.com/ast9501/ran-operator/pkg/controller/ranslice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, ranslice.Add)
}
