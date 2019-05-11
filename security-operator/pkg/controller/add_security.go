package controller

import (
	"github.com/Shehanir/k8s-securityOperator/security-operator/pkg/controller/security"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, security.Add)
}
