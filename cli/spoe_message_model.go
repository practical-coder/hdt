// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/practical-coder/hdc/models"

	"github.com/spf13/cobra"
)

// Schema cli for SpoeMessage

// register flags to command
func registerModelSpoeMessageFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSpoeMessageACL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSpoeMessageArgs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSpoeMessageEvent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSpoeMessageName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSpoeMessageACL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: acl Acls array type is not supported by go-swagger cli yet

	return nil
}

func registerSpoeMessageArgs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	argsDescription := ``

	var argsFlagName string
	if cmdPrefix == "" {
		argsFlagName = "args"
	} else {
		argsFlagName = fmt.Sprintf("%v.args", cmdPrefix)
	}

	var argsFlagDefault string

	_ = cmd.PersistentFlags().String(argsFlagName, argsFlagDefault, argsDescription)

	return nil
}

func registerSpoeMessageEvent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var eventFlagName string
	if cmdPrefix == "" {
		eventFlagName = "event"
	} else {
		eventFlagName = fmt.Sprintf("%v.event", cmdPrefix)
	}

	if err := registerModelSpoeMessageEventFlags(depth+1, eventFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSpoeMessageName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSpoeMessageFlags(depth int, m *models.SpoeMessage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, aclAdded := retrieveSpoeMessageACLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || aclAdded

	err, argsAdded := retrieveSpoeMessageArgsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || argsAdded

	err, eventAdded := retrieveSpoeMessageEventFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || eventAdded

	err, nameAdded := retrieveSpoeMessageNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	return nil, retAdded
}

func retrieveSpoeMessageACLFlags(depth int, m *models.SpoeMessage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	aclFlagName := fmt.Sprintf("%v.acl", cmdPrefix)
	if cmd.Flags().Changed(aclFlagName) {
		// warning: acl array type Acls is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSpoeMessageArgsFlags(depth int, m *models.SpoeMessage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	argsFlagName := fmt.Sprintf("%v.args", cmdPrefix)
	if cmd.Flags().Changed(argsFlagName) {

		var argsFlagName string
		if cmdPrefix == "" {
			argsFlagName = "args"
		} else {
			argsFlagName = fmt.Sprintf("%v.args", cmdPrefix)
		}

		argsFlagValue, err := cmd.Flags().GetString(argsFlagName)
		if err != nil {
			return err, false
		}
		m.Args = argsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSpoeMessageEventFlags(depth int, m *models.SpoeMessage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	eventFlagName := fmt.Sprintf("%v.event", cmdPrefix)
	if cmd.Flags().Changed(eventFlagName) {
		// info: complex object event SpoeMessageEvent is retrieved outside this Changed() block
	}
	eventFlagValue := m.Event
	if swag.IsZero(eventFlagValue) {
		eventFlagValue = &models.SpoeMessageEvent{}
	}

	err, eventAdded := retrieveModelSpoeMessageEventFlags(depth+1, eventFlagValue, eventFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || eventAdded
	if eventAdded {
		m.Event = eventFlagValue
	}

	return nil, retAdded
}

func retrieveSpoeMessageNameFlags(depth int, m *models.SpoeMessage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for SpoeMessageEvent

// register flags to command
func registerModelSpoeMessageEventFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSpoeMessageEventCond(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSpoeMessageEventCondTest(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSpoeMessageEventName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSpoeMessageEventCond(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	condDescription := `Enum: ["if","unless"]. `

	var condFlagName string
	if cmdPrefix == "" {
		condFlagName = "cond"
	} else {
		condFlagName = fmt.Sprintf("%v.cond", cmdPrefix)
	}

	var condFlagDefault string

	_ = cmd.PersistentFlags().String(condFlagName, condFlagDefault, condDescription)

	if err := cmd.RegisterFlagCompletionFunc(condFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerSpoeMessageEventCondTest(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	condTestDescription := ``

	var condTestFlagName string
	if cmdPrefix == "" {
		condTestFlagName = "cond_test"
	} else {
		condTestFlagName = fmt.Sprintf("%v.cond_test", cmdPrefix)
	}

	var condTestFlagDefault string

	_ = cmd.PersistentFlags().String(condTestFlagName, condTestFlagDefault, condTestDescription)

	return nil
}

func registerSpoeMessageEventName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Enum: ["on-client-session","on-server-session","on-frontend-tcp-request","on-backend-tcp-request","on-tcp-response","on-frontend-http-request","on-backend-http-request","on-http-response"]. Required. `

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	if err := cmd.RegisterFlagCompletionFunc(nameFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["on-client-session","on-server-session","on-frontend-tcp-request","on-backend-tcp-request","on-tcp-response","on-frontend-http-request","on-backend-http-request","on-http-response"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSpoeMessageEventFlags(depth int, m *models.SpoeMessageEvent, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, condAdded := retrieveSpoeMessageEventCondFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || condAdded

	err, condTestAdded := retrieveSpoeMessageEventCondTestFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || condTestAdded

	err, nameAdded := retrieveSpoeMessageEventNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	return nil, retAdded
}

func retrieveSpoeMessageEventCondFlags(depth int, m *models.SpoeMessageEvent, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	condFlagName := fmt.Sprintf("%v.cond", cmdPrefix)
	if cmd.Flags().Changed(condFlagName) {

		var condFlagName string
		if cmdPrefix == "" {
			condFlagName = "cond"
		} else {
			condFlagName = fmt.Sprintf("%v.cond", cmdPrefix)
		}

		condFlagValue, err := cmd.Flags().GetString(condFlagName)
		if err != nil {
			return err, false
		}
		m.Cond = condFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSpoeMessageEventCondTestFlags(depth int, m *models.SpoeMessageEvent, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	condTestFlagName := fmt.Sprintf("%v.cond_test", cmdPrefix)
	if cmd.Flags().Changed(condTestFlagName) {

		var condTestFlagName string
		if cmdPrefix == "" {
			condTestFlagName = "cond_test"
		} else {
			condTestFlagName = fmt.Sprintf("%v.cond_test", cmdPrefix)
		}

		condTestFlagValue, err := cmd.Flags().GetString(condTestFlagName)
		if err != nil {
			return err, false
		}
		m.CondTest = condTestFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSpoeMessageEventNameFlags(depth int, m *models.SpoeMessageEvent, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}