// Program to generate TypeScript definition files for Golang structs that are
// serialized to JSON for the web UI.
//
//go:generate bazelisk run --config=mayberemote //:go -- run . -o ../../../modules/json/index.ts
package main

import (
	"flag"
	"io"

	"go.skia.org/infra/go/go2ts"
	"go.skia.org/infra/go/sklog"
	"go.skia.org/infra/go/util"
	"go.skia.org/infra/machine/go/machine"
	"go.skia.org/infra/machine/go/machineserver/rpc"
)

func main() {
	var outputPath = flag.String("o", "", "Path to the output TypeScript file.")
	flag.Parse()

	generator := go2ts.New()
	generator.AddMultiple(
		rpc.SetNoteRequest{},
		rpc.SupplyChromeOSRequest{},
		rpc.SetAttachedDevice{},
	)
	generator.AddIgnoreNil(rpc.ListMachinesResponse{})
	generator.AddUnion(machine.AllAttachedDevices)
	generator.AddUnion(machine.AllPowerCycleStates)
	generator.AddUnion(machine.AllTaskRequestorStates)

	err := util.WithWriteFile(*outputPath, func(w io.Writer) error {
		return generator.Render(w)
	})
	if err != nil {
		sklog.Fatal(err)
	}
}
