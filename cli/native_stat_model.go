// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/haproxytech/models"

	"github.com/spf13/cobra"
)

// Schema cli for NativeStat

// register flags to command
func registerModelNativeStatFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerNativeStatBackendName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNativeStatName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNativeStatStats(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNativeStatType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerNativeStatBackendName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	backendNameDescription := ``

	var backendNameFlagName string
	if cmdPrefix == "" {
		backendNameFlagName = "backend_name"
	} else {
		backendNameFlagName = fmt.Sprintf("%v.backend_name", cmdPrefix)
	}

	var backendNameFlagDefault string

	_ = cmd.PersistentFlags().String(backendNameFlagName, backendNameFlagDefault, backendNameDescription)

	return nil
}

func registerNativeStatName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerNativeStatStats(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var statsFlagName string
	if cmdPrefix == "" {
		statsFlagName = "stats"
	} else {
		statsFlagName = fmt.Sprintf("%v.stats", cmdPrefix)
	}

	if err := registerModelNativeStatStatsFlags(depth+1, statsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerNativeStatType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["backend","server","frontend"]. `

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["backend","server","frontend"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelNativeStatFlags(depth int, m *models.NativeStat, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, backendNameAdded := retrieveNativeStatBackendNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || backendNameAdded

	err, nameAdded := retrieveNativeStatNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, statsAdded := retrieveNativeStatStatsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsAdded

	err, typeAdded := retrieveNativeStatTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveNativeStatBackendNameFlags(depth int, m *models.NativeStat, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	backendNameFlagName := fmt.Sprintf("%v.backend_name", cmdPrefix)
	if cmd.Flags().Changed(backendNameFlagName) {

		var backendNameFlagName string
		if cmdPrefix == "" {
			backendNameFlagName = "backend_name"
		} else {
			backendNameFlagName = fmt.Sprintf("%v.backend_name", cmdPrefix)
		}

		backendNameFlagValue, err := cmd.Flags().GetString(backendNameFlagName)
		if err != nil {
			return err, false
		}
		m.BackendName = backendNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNativeStatNameFlags(depth int, m *models.NativeStat, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNativeStatStatsFlags(depth int, m *models.NativeStat, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsFlagName := fmt.Sprintf("%v.stats", cmdPrefix)
	if cmd.Flags().Changed(statsFlagName) {
		// info: complex object stats NativeStatStats is retrieved outside this Changed() block
	}
	statsFlagValue := m.Stats
	if swag.IsZero(statsFlagValue) {
		statsFlagValue = &models.NativeStatStats{}
	}

	err, statsAdded := retrieveModelNativeStatStatsFlags(depth+1, statsFlagValue, statsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsAdded
	if statsAdded {
		m.Stats = statsFlagValue
	}

	return nil, retAdded
}

func retrieveNativeStatTypeFlags(depth int, m *models.NativeStat, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
