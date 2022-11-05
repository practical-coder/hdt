// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/haproxytech/client-native/v4/models"

	"github.com/spf13/cobra"
)

// Schema cli for ProcessInfo

// register flags to command
func registerModelProcessInfoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerProcessInfoError(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerProcessInfoInfo(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerProcessInfoRuntimeAPI(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerProcessInfoError(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	errorDescription := ``

	var errorFlagName string
	if cmdPrefix == "" {
		errorFlagName = "error"
	} else {
		errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
	}

	var errorFlagDefault string

	_ = cmd.PersistentFlags().String(errorFlagName, errorFlagDefault, errorDescription)

	return nil
}

func registerProcessInfoInfo(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var infoFlagName string
	if cmdPrefix == "" {
		infoFlagName = "info"
	} else {
		infoFlagName = fmt.Sprintf("%v.info", cmdPrefix)
	}

	if err := registerModelProcessInfoItemFlags(depth+1, infoFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerProcessInfoRuntimeAPI(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	runtimeApiDescription := ``

	var runtimeApiFlagName string
	if cmdPrefix == "" {
		runtimeApiFlagName = "runtimeAPI"
	} else {
		runtimeApiFlagName = fmt.Sprintf("%v.runtimeAPI", cmdPrefix)
	}

	var runtimeApiFlagDefault string

	_ = cmd.PersistentFlags().String(runtimeApiFlagName, runtimeApiFlagDefault, runtimeApiDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelProcessInfoFlags(depth int, m *models.ProcessInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, errorAdded := retrieveProcessInfoErrorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorAdded

	err, infoAdded := retrieveProcessInfoInfoFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || infoAdded

	err, runtimeApiAdded := retrieveProcessInfoRuntimeAPIFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || runtimeApiAdded

	return nil, retAdded
}

func retrieveProcessInfoErrorFlags(depth int, m *models.ProcessInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorFlagName := fmt.Sprintf("%v.error", cmdPrefix)
	if cmd.Flags().Changed(errorFlagName) {

		var errorFlagName string
		if cmdPrefix == "" {
			errorFlagName = "error"
		} else {
			errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
		}

		errorFlagValue, err := cmd.Flags().GetString(errorFlagName)
		if err != nil {
			return err, false
		}
		m.Error = errorFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveProcessInfoInfoFlags(depth int, m *models.ProcessInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	infoFlagName := fmt.Sprintf("%v.info", cmdPrefix)
	if cmd.Flags().Changed(infoFlagName) {
		// info: complex object info ProcessInfoItem is retrieved outside this Changed() block
	}
	infoFlagValue := m.Info
	if swag.IsZero(infoFlagValue) {
		infoFlagValue = &models.ProcessInfoItem{}
	}

	err, infoAdded := retrieveModelProcessInfoItemFlags(depth+1, infoFlagValue, infoFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || infoAdded
	if infoAdded {
		m.Info = infoFlagValue
	}

	return nil, retAdded
}

func retrieveProcessInfoRuntimeAPIFlags(depth int, m *models.ProcessInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	runtimeApiFlagName := fmt.Sprintf("%v.runtimeAPI", cmdPrefix)
	if cmd.Flags().Changed(runtimeApiFlagName) {

		var runtimeApiFlagName string
		if cmdPrefix == "" {
			runtimeApiFlagName = "runtimeAPI"
		} else {
			runtimeApiFlagName = fmt.Sprintf("%v.runtimeAPI", cmdPrefix)
		}

		runtimeApiFlagValue, err := cmd.Flags().GetString(runtimeApiFlagName)
		if err != nil {
			return err, false
		}
		m.RuntimeAPI = runtimeApiFlagValue

		retAdded = true
	}

	return nil, retAdded
}
