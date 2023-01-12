package main

import (
	"fmt"
	xr "github.com/nleiva/xrgrpc"
	"github.com/openconfig/ygot/ygot"
	"log"
	"yangstuff/pkg/interfaces"
)

func main() {
	router, err := xr.BuildRouter(
		xr.WithUsername("grpc"),
		xr.WithPassword("53cret"),
		xr.WithHost("192.168.88.3:57344"),
		xr.WithTimeout(50),
	)

	conn, ctx, err := xr.Connect(*router)
	if err != nil {
		log.Fatalf("could not setup a client connection to %s, %v", router.Host, err)
	}
	defer conn.Close()
	var output string

	var yp = `{	
    "Cisco-IOS-XR-ifmgr-cfg:interface-configurations": [null]
}
`

	output, err = xr.GetConfig(ctx, conn, yp, 1001)
	if err != nil {
		log.Fatalf("could not get the config from %s, %v", router.Host, err)
	}
	fmt.Printf("\nconfig from %s\n %s\n", router.Host, output)
	i := interfaces.Interfaces{}
	i.InterfaceConfigurations = &interfaces.Cisco_IOS_XRIfmgrCfg_InterfaceConfigurations{}
	ni, err := i.InterfaceConfigurations.NewInterfaceConfiguration("act", "Loopback333")
	if err != nil {
		fmt.Println(err)
	}
	ni.Description = ygot.String("Wow:This:Worked")
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

	io, err := xr.MergeConfig(ctx, conn, json, 100023)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(io)
	xr.CommitConfig(ctx, conn, uint32(io), 3333)

	deleteLspInterfaces := interfaces.Interfaces{}
	deleteLspInterfaces.InterfaceConfigurations = &interfaces.Cisco_IOS_XRIfmgrCfg_InterfaceConfigurations{}
	_, err = deleteLspInterfaces.InterfaceConfigurations.NewInterfaceConfiguration("act", "tunnel-te29")
	if err != nil {
		fmt.Println(err)
	}
	json, err = ygot.EmitJSON(&deleteLspInterfaces, &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
		Indent: "  ",
		RFC7951Config: &ygot.RFC7951JSONConfig{
			AppendModuleName: true,
		},
	})
	print(json)
	io, err = xr.DeleteConfig(ctx, conn, json, 100023333)
	if err != nil {
		fmt.Println(err)
	}
}
