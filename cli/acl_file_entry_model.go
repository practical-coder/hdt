// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/spf13/cobra"
)

// Schema cli for ACLFileEntry

// register flags to command
func registerModelACLFileEntryFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerACLFileEntryID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerACLFileEntryValue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerACLFileEntryID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerACLFileEntryValue(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	valueDescription := ``

	var valueFlagName string
	if cmdPrefix == "" {
		valueFlagName = "value"
	} else {
		valueFlagName = fmt.Sprintf("%v.value", cmdPrefix)
	}

	var valueFlagDefault string

	_ = cmd.PersistentFlags().String(valueFlagName, valueFlagDefault, valueDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelACLFileEntryFlags(depth int, m *models.ACLFileEntry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveACLFileEntryIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, valueAdded := retrieveACLFileEntryValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || valueAdded

	return nil, retAdded
}

func retrieveACLFileEntryIDFlags(depth int, m *models.ACLFileEntry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveACLFileEntryValueFlags(depth int, m *models.ACLFileEntry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	valueFlagName := fmt.Sprintf("%v.value", cmdPrefix)
	if cmd.Flags().Changed(valueFlagName) {

		var valueFlagName string
		if cmdPrefix == "" {
			valueFlagName = "value"
		} else {
			valueFlagName = fmt.Sprintf("%v.value", cmdPrefix)
		}

		valueFlagValue, err := cmd.Flags().GetString(valueFlagName)
		if err != nil {
			return err, false
		}
		m.Value = valueFlagValue

		retAdded = true
	}

	return nil, retAdded
}
