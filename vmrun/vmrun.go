package vmrun

import (
	"os/exec"
	"strings"

	"github.com/jeffwubj/go-vmrun/types"
)

// Vmrun is the client to the vmrun cli
type Vmrun struct {
}

// List list all running VMs
func (vmrun *Vmrun) List() ([]*types.VM, error) {
	var vms []*types.VM
	cmd := exec.Command(DefaultVmrunLocation, "list")
	buf, err := cmdStdout(cmd)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		if strings.Contains(line, "vmx") {
			vm, err := types.NewVM(line)
			if err != nil {
				return vms, err
			}
			vms = append(vms, vm)
		}
	}
	return vms, nil
}

// Start powers on given VM
func (vmrun *Vmrun) Start() error {
	return nil
}

// Stop powers off given VM
func (vmrun *Vmrun) Stop() error {
	return nil
}
