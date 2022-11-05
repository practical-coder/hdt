// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/spf13/cobra"
)

// Schema cli for MapEntry

// register flags to command
func registerModelMapEntryFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMapEntryID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMapEntryKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMapEntryValue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMapEntryID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerMapEntryKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	keyDescription := ``

	var keyFlagName string
	if cmdPrefix == "" {
		keyFlagName = "key"
	} else {
		keyFlagName = fmt.Sprintf("%v.key", cmdPrefix)
	}

	var keyFlagDefault string

	_ = cmd.PersistentFlags().String(keyFlagName, keyFlagDefault, keyDescription)

	return nil
}

func registerMapEntryValue(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelMapEntryFlags(depth int, m *models.MapEntry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveMapEntryIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, keyAdded := retrieveMapEntryKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || keyAdded

	err, valueAdded := retrieveMapEntryValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || valueAdded

	return nil, retAdded
}

func retrieveMapEntryIDFlags(depth int, m *models.MapEntry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveMapEntryKeyFlags(depth int, m *models.MapEntry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	keyFlagName := fmt.Sprintf("%v.key", cmdPrefix)
	if cmd.Flags().Changed(keyFlagName) {

		var keyFlagName string
		if cmdPrefix == "" {
			keyFlagName = "key"
		} else {
			keyFlagName = fmt.Sprintf("%v.key", cmdPrefix)
		}

		keyFlagValue, err := cmd.Flags().GetString(keyFlagName)
		if err != nil {
			return err, false
		}
		m.Key = keyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMapEntryValueFlags(depth int, m *models.MapEntry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
