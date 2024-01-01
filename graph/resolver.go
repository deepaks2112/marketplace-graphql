package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"fmt"

	"deepaks2112.github.io/marketplace/backend"
)

type Resolver struct{
	repo backend.Repository
}

func (r *Resolver) Init () {
	err := r.repo.Connect()
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
}
