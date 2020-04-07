// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package upgrades

import (
	"path/filepath"

	"github.com/juju/errors"
	"github.com/juju/utils"
	"gopkg.in/juju/names.v3"

	"github.com/juju/juju/worker/uniter/operation"
)

// stateStepsFor276 returns upgrade steps for Juju 2.7.6.
func stepsFor276() []Step {
	return []Step{
		&upgradeStep{
			description: "add remote-application key to hooks in uniter state files",
			targets:     []Target{HostMachine},
			run:         AddRemoteApplicationToRunningHooks(uniterStateGlob),
		},
	}
}

const uniterStateGlob = `/var/lib/juju/agents/unit-*/state/uniter`

// AddRemoteApplicationToRunningHooks finds any uniter state files on
// the machine with running hooks, and makes sure that they contain a
// remote-application key.
func AddRemoteApplicationToRunningHooks(pattern string) func(Context) error {
	return func(_ Context) error {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return errors.Annotate(err, "finding uniter state files")
		}
		for _, path := range matches {
			// First, check whether the file needs rewriting.
			stateFile := operation.NewStateFile(path)
			_, err := stateFile.Read()
			if err == nil {
				// This one's fine, leave it alone.
				logger.Debugf("state file valid: %q", path)
				continue
			}

			err = AddRemoteApplicationToHook(path)
			if err != nil {
				return errors.Annotatef(err, "fixing %q", path)
			}
		}
		return nil
	}
}

// AddRemoteApplicationToHook takes a the path to a uniter state file
// that doesn't validate, and sets hook.remote-application to the
// remote application so that it does. (If it doesn't validate for
// some other reason we won't change the file.)
func AddRemoteApplicationToHook(path string) error {
	var uniterState operation.State
	err := utils.ReadYaml(path, &uniterState)
	if err != nil {
		return errors.Trace(err)
	}

	if uniterState.Hook == nil {
		logger.Debugf("no hook found in %q, unable to fix", path)
		return nil
	}

	if uniterState.Hook.RemoteApplication != "" {
		logger.Debugf("remote-application in %q set to %q already", path,
			uniterState.Hook.RemoteApplication)
		return nil
	}

	appName, err := names.UnitApplication(uniterState.Hook.RemoteUnit)
	if err != nil {
		return errors.Trace(err)
	}

	logger.Debugf("setting remote-application to %q in %q", appName, path)
	uniterState.Hook.RemoteApplication = appName
	return errors.Annotatef(operation.NewStateFile(path).Write(&uniterState),
		"writing updated state to %q", path)
}
