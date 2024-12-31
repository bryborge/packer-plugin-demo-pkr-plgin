package main

import (
	"fmt"
	"os"

	scaffoldingBuilder "github.com/bryborge/demo-packer-plugin/builder/scaffolding"
	scaffoldingVersion "github.com/bryborge/demo-packer-plugin/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder("my-builder", new(scaffoldingBuilder.Builder))
	// Register Provisioners, Post-Processors, and Datasources here ...
	pps.SetVersion(scaffoldingVersion.PluginVersion)

	err := pps.Run()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
