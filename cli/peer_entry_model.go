// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/practical-coder/hdc/models"
	"github.com/spf13/cobra"
)

// Schema cli for PeerEntry

// register flags to command
func registerModelPeerEntryFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPeerEntryAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPeerEntryName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPeerEntryPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPeerEntryAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	addressDescription := `Required. `

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

func registerPeerEntryName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. `

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

func registerPeerEntryPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	portDescription := `Required. `

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
func retrieveModelPeerEntryFlags(depth int, m *models.PeerEntry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressAdded := retrievePeerEntryAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressAdded

	err, nameAdded := retrievePeerEntryNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, portAdded := retrievePeerEntryPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portAdded

	return nil, retAdded
}

func retrievePeerEntryAddressFlags(depth int, m *models.PeerEntry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Address = &addressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePeerEntryNameFlags(depth int, m *models.PeerEntry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePeerEntryPortFlags(depth int, m *models.PeerEntry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
