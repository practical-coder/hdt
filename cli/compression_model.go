// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/spf13/cobra"
)

// Schema cli for Compression

// register flags to command
func registerModelCompressionFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCompressionAlgorithms(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCompressionOffload(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCompressionTypes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCompressionAlgorithms(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: algorithms []string array type is not supported by go-swagger cli yet

	return nil
}

func registerCompressionOffload(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	offloadDescription := ``

	var offloadFlagName string
	if cmdPrefix == "" {
		offloadFlagName = "offload"
	} else {
		offloadFlagName = fmt.Sprintf("%v.offload", cmdPrefix)
	}

	var offloadFlagDefault bool

	_ = cmd.PersistentFlags().Bool(offloadFlagName, offloadFlagDefault, offloadDescription)

	return nil
}

func registerCompressionTypes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: types []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCompressionFlags(depth int, m *models.Compression, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, algorithmsAdded := retrieveCompressionAlgorithmsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || algorithmsAdded

	err, offloadAdded := retrieveCompressionOffloadFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || offloadAdded

	err, typesAdded := retrieveCompressionTypesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typesAdded

	return nil, retAdded
}

func retrieveCompressionAlgorithmsFlags(depth int, m *models.Compression, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	algorithmsFlagName := fmt.Sprintf("%v.algorithms", cmdPrefix)
	if cmd.Flags().Changed(algorithmsFlagName) {
		// warning: algorithms array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveCompressionOffloadFlags(depth int, m *models.Compression, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	offloadFlagName := fmt.Sprintf("%v.offload", cmdPrefix)
	if cmd.Flags().Changed(offloadFlagName) {

		var offloadFlagName string
		if cmdPrefix == "" {
			offloadFlagName = "offload"
		} else {
			offloadFlagName = fmt.Sprintf("%v.offload", cmdPrefix)
		}

		offloadFlagValue, err := cmd.Flags().GetBool(offloadFlagName)
		if err != nil {
			return err, false
		}
		m.Offload = offloadFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCompressionTypesFlags(depth int, m *models.Compression, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typesFlagName := fmt.Sprintf("%v.types", cmdPrefix)
	if cmd.Flags().Changed(typesFlagName) {
		// warning: types array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
