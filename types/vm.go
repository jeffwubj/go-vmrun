package types

import (
	"net"

	"gopkg.in/ini.v1"
)

// VM type
type VM struct {
	// vmx path
	Path        string
	DisplayName string
	CPU         int
	Memory      int
	IP          net.Addr
}

// NewVM creates VM according to vmx file
func NewVM(vmxpath string) (*VM, error) {
	vm := &VM{}
	cfg, err := ini.Load(vmxpath)
	if err != nil {
		return vm, err
	}
	vm.Path = vmxpath
	vm.DisplayName = cfg.Section("").Key("displayName").String()
	vm.Memory, _ = cfg.Section("").Key("memSize").Int()
	return vm, nil
}
