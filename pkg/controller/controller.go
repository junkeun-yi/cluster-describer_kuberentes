package controller

import (
	"github.com/junkeun-yi/cluster-describer_kuberentes/pkg/functions"
)

type Controller struct {
	FunctionSet		functions.FunctionSet
}

func (c *Controller) Run() {
	f := c.FunctionSet

	f.GetAllInfo()
}

