package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/huawei/DockerMachineDriver4OTC/otc"
)

func main() {
	plugin.RegisterDriver(otc.NewDriver("", ""))
}
