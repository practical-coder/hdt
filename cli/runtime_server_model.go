// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/spf13/cobra"
)

// Schema cli for RuntimeServer

// register flags to command
func registerModelRuntimeServerFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRuntimeServerAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRuntimeServerAdminState(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRuntimeServerID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRuntimeServerName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRuntimeServerOperationalState(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRuntimeServerPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRuntimeServerAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	addressDescription := ``

	var addressFlagName string
	if cmdPrefix == "" {
		addressFlagName = "address"
	} else {
		addressFlagName = fmt.Sprintf("%v.address", cmdPrefix)
	}

	var addressFlagDefault string

	_ = cmd.PersistentFlags().String(addressFlagName, addressFlagDefault, addressDescription)

	return nil
}

func registerRuntimeServerAdminState(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	adminStateDescription := `Enum: ["ready","maint","drain"]. `

	var adminStateFlagName string
	if cmdPrefix == "" {
		adminStateFlagName = "admin_state"
	} else {
		adminStateFlagName = fmt.Sprintf("%v.admin_state", cmdPrefix)
	}

	var adminStateFlagDefault string

	_ = cmd.PersistentFlags().String(adminStateFlagName, adminStateFlagDefault, adminStateDescription)

	if err := cmd.RegisterFlagCompletionFunc(adminStateFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["ready","maint","drain"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerRuntimeServerID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerRuntimeServerName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerRuntimeServerOperationalState(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	operationalStateDescription := `Enum: ["up","down","stopping"]. `

	var operationalStateFlagName string
	if cmdPrefix == "" {
		operationalStateFlagName = "operational_state"
	} else {
		operationalStateFlagName = fmt.Sprintf("%v.operational_state", cmdPrefix)
	}

	var operationalStateFlagDefault string

	_ = cmd.PersistentFlags().String(operationalStateFlagName, operationalStateFlagDefault, operationalStateDescription)

	if err := cmd.RegisterFlagCompletionFunc(operationalStateFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["up","down","stopping"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerRuntimeServerPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	portDescription := ``

	var portFlagName string
	if cmdPrefix == "" {
		portFlagName = "port"
	} else {
		portFlagName = fmt.Sprintf("%v.port", cmdPrefix)
	}

	var portFlagDefault int64

	_ = cmd.PersistentFlags().Int64(portFlagName, portFlagDefault, portDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRuntimeServerFlags(depth int, m *models.RuntimeServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressAdded := retrieveRuntimeServerAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressAdded

	err, adminStateAdded := retrieveRuntimeServerAdminStateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || adminStateAdded

	err, idAdded := retrieveRuntimeServerIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, nameAdded := retrieveRuntimeServerNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, operationalStateAdded := retrieveRuntimeServerOperationalStateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || operationalStateAdded

	err, portAdded := retrieveRuntimeServerPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portAdded

	return nil, retAdded
}

func retrieveRuntimeServerAddressFlags(depth int, m *models.RuntimeServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	addressFlagName := fmt.Sprintf("%v.address", cmdPrefix)
	if cmd.Flags().Changed(addressFlagName) {

		var addressFlagName string
		if cmdPrefix == "" {
			addressFlagName = "address"
		} else {
			addressFlagName = fmt.Sprintf("%v.address", cmdPrefix)
		}

		addressFlagValue, err := cmd.Flags().GetString(addressFlagName)
		if err != nil {
			return err, false
		}
		m.Address = addressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRuntimeServerAdminStateFlags(depth int, m *models.RuntimeServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	adminStateFlagName := fmt.Sprintf("%v.admin_state", cmdPrefix)
	if cmd.Flags().Changed(adminStateFlagName) {

		var adminStateFlagName string
		if cmdPrefix == "" {
			adminStateFlagName = "admin_state"
		} else {
			adminStateFlagName = fmt.Sprintf("%v.admin_state", cmdPrefix)
		}

		adminStateFlagValue, err := cmd.Flags().GetString(adminStateFlagName)
		if err != nil {
			return err, false
		}
		m.AdminState = adminStateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRuntimeServerIDFlags(depth int, m *models.RuntimeServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRuntimeServerNameFlags(depth int, m *models.RuntimeServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRuntimeServerOperationalStateFlags(depth int, m *models.RuntimeServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	operationalStateFlagName := fmt.Sprintf("%v.operational_state", cmdPrefix)
	if cmd.Flags().Changed(operationalStateFlagName) {

		var operationalStateFlagName string
		if cmdPrefix == "" {
			operationalStateFlagName = "operational_state"
		} else {
			operationalStateFlagName = fmt.Sprintf("%v.operational_state", cmdPrefix)
		}

		operationalStateFlagValue, err := cmd.Flags().GetString(operationalStateFlagName)
		if err != nil {
			return err, false
		}
		m.OperationalState = operationalStateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRuntimeServerPortFlags(depth int, m *models.RuntimeServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portFlagName := fmt.Sprintf("%v.port", cmdPrefix)
	if cmd.Flags().Changed(portFlagName) {

		var portFlagName string
		if cmdPrefix == "" {
			portFlagName = "port"
		} else {
			portFlagName = fmt.Sprintf("%v.port", cmdPrefix)
		}

		portFlagValue, err := cmd.Flags().GetInt64(portFlagName)
		if err != nil {
			return err, false
		}
		m.Port = &portFlagValue

		retAdded = true
	}

	return nil, retAdded
}
