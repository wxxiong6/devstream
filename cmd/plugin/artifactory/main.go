package main

import (
	"github.com/devstream-io/devstream/internal/pkg/configmanager"
	"github.com/devstream-io/devstream/internal/pkg/plugin/artifactory"
	"github.com/devstream-io/devstream/internal/pkg/statemanager"
	"github.com/devstream-io/devstream/pkg/util/log"
)

// NAME is the name of this DevStream plugin.
const NAME = "artifactory"

// Plugin is the type used by DevStream core. It's a string.
type Plugin string

// Create implements the create of artifactory.
func (p Plugin) Create(options configmanager.RawOption) (statemanager.ResourceStatus, error) {
	return artifactory.Create(options)
}

// Update implements the update of artifactory.
func (p Plugin) Update(options configmanager.RawOption) (statemanager.ResourceStatus, error) {
	return artifactory.Update(options)
}

// Delete implements the delete of artifactory.
func (p Plugin) Delete(options configmanager.RawOption) (bool, error) {
	return artifactory.Delete(options)
}

// Read implements the read of artifactory.
func (p Plugin) Read(options configmanager.RawOption) (statemanager.ResourceStatus, error) {
	return artifactory.Read(options)
}

// DevStreamPlugin is the exported variable used by the DevStream core.
var DevStreamPlugin Plugin

func main() {
	log.Infof("%T: %s is a plugin for DevStream. Use it with DevStream.\n", DevStreamPlugin, NAME)
}
