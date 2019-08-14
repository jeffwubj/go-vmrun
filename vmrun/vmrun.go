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
func (vmrun *Vmrun) Start(vmxpath string, nogui bool) (*types.VM, error) {
	vm, err := types.NewVM(vmxpath)
	if err != nil {
		return nil, err
	}

	gui := "gui"
	if nogui {
		gui = "nogui"
	}

	cmd := exec.Command(DefaultVmrunLocation, "start", vmxpath, gui)
	_, err = cmdStdout(cmd)
	if err != nil {
		return vm, err
	}
	return vm, nil
}

// Stop powers off given VM
func (vmrun *Vmrun) Stop(vmxpath string) (*types.VM, error) {
	vm, err := types.NewVM(vmxpath)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(DefaultVmrunLocation, "stop", vmxpath)
	_, err = cmdStdout(cmd)
	if err != nil {
		return vm, err
	}
	return vm, nil
}
