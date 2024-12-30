package main

import (
	"fmt"
	"os"

	Builder "github.com/bryborge/demo-packer-plugin/builder"
	Version "github.com/bryborge/demo-packer-plugin/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder("my-builder", new(Builder.Builder))
	pps.SetVersion(Version.PluginVersion)

	err := pps.Run()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
