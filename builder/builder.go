package builder

import (
	"context"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	// TODO: Replace with your own config
	"github.com/hashicorp/packer-plugin-sdk/template/config"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	MockOption          string `mapstructure:"mock"`
}

type Builder struct {
	config Config
}

func (b *Builder) ConfigSpec() hcldec.ObjectSpec {
	// TODO: Replace with your own config spec
	// return b.config.FlatMapstructure().HCL2Spec()
	return nil
}

/**
 * Prepares the builder for building.
 * This is where you can validate the configuration, set defaults, etc.
 */
func (b *Builder) Prepare(raws ...interface{}) (generatedVars []string, warnings []string, err error) {
	err = config.Decode(&b.config, &config.DecodeOpts{
		// TODO: Replace with your own PluginType
		PluginType:  "packer.builder.scaffolding",
		Interpolate: true,
	}, raws...)

	if err != nil {
		return nil, nil, err
	}
	// Return the placeholder for the generated data that will become available to provisioners and post-processors.
	// If the builder doesn't generate any data, just return an empty slice of string: []string{}
	buildGeneratedData := []string{"GeneratedMockData"}
	return buildGeneratedData, nil, nil
}

/**
 * Run is the main function that will be called by Packer to run the builder.
 * This is where the actual work of the builder is done.
 */
func (b *Builder) Run(ctx context.Context, ui packer.Ui, hook packer.Hook) (packer.Artifact, error) {
	// TODO: Replace with your own builder logic ...
	return nil, nil
}
