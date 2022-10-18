package reposcaffolding

import (
	"github.com/devstream-io/devstream/internal/pkg/configmanager"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller/reposcaffolding"
	"github.com/devstream-io/devstream/internal/pkg/statemanager"
	"github.com/devstream-io/devstream/pkg/util/log"
)

func Update(options configmanager.RawOption) (statemanager.ResourceStatus, error) {
	operator := &plugininstaller.Operator{
		PreExecuteOperations: plugininstaller.PreExecuteOperations{
			reposcaffolding.Validate,
		},
		ExecuteOperations: plugininstaller.ExecuteOperations{
			reposcaffolding.DeleteRepo,
			reposcaffolding.InstallRepo,
		},
		GetStatusOperation: reposcaffolding.GetDynamicStatus,
	}

	// Execute all Operations in Operator
	status, err := operator.Execute(plugininstaller.RawOptions(options))
	if err != nil {
		return nil, err
	}
	log.Debugf("Return map: %v", status)
	return status, nil
}
