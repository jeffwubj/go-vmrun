package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/jeffwubj/go-vmrun/vmrun"
)

func main() {
	vmx, err := vmrun.List()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	spew.Dump(vmx)
}
