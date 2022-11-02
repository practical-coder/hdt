// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/practical-coder/hdc/models"
	"github.com/spf13/cobra"
)

// Schema cli for PgsqlCheckParams

// register flags to command
func registerModelPgsqlCheckParamsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPgsqlCheckParamsUsername(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPgsqlCheckParamsUsername(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	usernameDescription := ``

	var usernameFlagName string
	if cmdPrefix == "" {
		usernameFlagName = "username"
	} else {
		usernameFlagName = fmt.Sprintf("%v.username", cmdPrefix)
	}

	var usernameFlagDefault string

	_ = cmd.PersistentFlags().String(usernameFlagName, usernameFlagDefault, usernameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPgsqlCheckParamsFlags(depth int, m *models.PgsqlCheckParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, usernameAdded := retrievePgsqlCheckParamsUsernameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || usernameAdded

	return nil, retAdded
}

func retrievePgsqlCheckParamsUsernameFlags(depth int, m *models.PgsqlCheckParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	usernameFlagName := fmt.Sprintf("%v.username", cmdPrefix)
	if cmd.Flags().Changed(usernameFlagName) {

		var usernameFlagName string
		if cmdPrefix == "" {
			usernameFlagName = "username"
		} else {
			usernameFlagName = fmt.Sprintf("%v.username", cmdPrefix)
		}

		usernameFlagValue, err := cmd.Flags().GetString(usernameFlagName)
		if err != nil {
			return err, false
		}
		m.Username = usernameFlagValue

		retAdded = true
	}

	return nil, retAdded
}
