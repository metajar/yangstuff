package main

import (
	"fmt"
	"github.com/openconfig/ygot/ygot"
	"yangstuff/pkg/interfaces"
)

func main() {

	i := interfaces.Interfaces{}
	i.InterfaceConfigurations = &interfaces.Cisco_IOS_XRIfmgrCfg_InterfaceConfigurations{}
	ni, err := i.InterfaceConfigurations.NewInterfaceConfiguration("act", "Loopback333")
	if err != nil {
		fmt.Println(err)
	}
	ni.Description = ygot.String("Wow:This:Worked")
	ni.Shutdown = true
	bs, err := ygot.Marshal7951(ni)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	json, err := ygot.EmitJSON(&i, &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
		Indent: "  ",
		RFC7951Config: &ygot.RFC7951JSONConfig{
			AppendModuleName: true,
		},
	})
	fmt.Println(json)
}
