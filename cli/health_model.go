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

// Schema cli for Health

// register flags to command
func registerModelHealthFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerHealthHaproxy(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerHealthHaproxy(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	haproxyDescription := `Enum: ["up","down","unknown"]. `

	var haproxyFlagName string
	if cmdPrefix == "" {
		haproxyFlagName = "haproxy"
	} else {
		haproxyFlagName = fmt.Sprintf("%v.haproxy", cmdPrefix)
	}

	var haproxyFlagDefault string

	_ = cmd.PersistentFlags().String(haproxyFlagName, haproxyFlagDefault, haproxyDescription)

	if err := cmd.RegisterFlagCompletionFunc(haproxyFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["up","down","unknown"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelHealthFlags(depth int, m *models.Health, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, haproxyAdded := retrieveHealthHaproxyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || haproxyAdded

	return nil, retAdded
}

func retrieveHealthHaproxyFlags(depth int, m *models.Health, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	haproxyFlagName := fmt.Sprintf("%v.haproxy", cmdPrefix)
	if cmd.Flags().Changed(haproxyFlagName) {

		var haproxyFlagName string
		if cmdPrefix == "" {
			haproxyFlagName = "haproxy"
		} else {
			haproxyFlagName = fmt.Sprintf("%v.haproxy", cmdPrefix)
		}

		haproxyFlagValue, err := cmd.Flags().GetString(haproxyFlagName)
		if err != nil {
			return err, false
		}
		m.Haproxy = haproxyFlagValue

		retAdded = true
	}

	return nil, retAdded
}
