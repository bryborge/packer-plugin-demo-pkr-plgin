package version

import "github.com/hashicorp/packer-plugin-sdk/version"

var (
	Version           = "0.2.4"
	VersionPrerelease = ""
	VersionMetadata   = ""
	PluginVersion     = version.NewPluginVersion(Version, VersionPrerelease, VersionMetadata)
)
