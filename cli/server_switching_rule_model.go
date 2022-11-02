// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/haproxytech/models"
	"github.com/spf13/cobra"
)

// Schema cli for ServerSwitchingRule

// register flags to command
func registerModelServerSwitchingRuleFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServerSwitchingRuleCond(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSwitchingRuleCondTest(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSwitchingRuleIndex(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSwitchingRuleTargetServer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServerSwitchingRuleCond(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerServerSwitchingRuleCondTest(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerServerSwitchingRuleIndex(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	indexDescription := `Required. `

	var indexFlagName string
	if cmdPrefix == "" {
		indexFlagName = "index"
	} else {
		indexFlagName = fmt.Sprintf("%v.index", cmdPrefix)
	}

	var indexFlagDefault int64

	_ = cmd.PersistentFlags().Int64(indexFlagName, indexFlagDefault, indexDescription)

	return nil
}

func registerServerSwitchingRuleTargetServer(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	targetServerDescription := `Required. `

	var targetServerFlagName string
	if cmdPrefix == "" {
		targetServerFlagName = "target_server"
	} else {
		targetServerFlagName = fmt.Sprintf("%v.target_server", cmdPrefix)
	}

	var targetServerFlagDefault string

	_ = cmd.PersistentFlags().String(targetServerFlagName, targetServerFlagDefault, targetServerDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServerSwitchingRuleFlags(depth int, m *models.ServerSwitchingRule, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, condAdded := retrieveServerSwitchingRuleCondFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || condAdded

	err, condTestAdded := retrieveServerSwitchingRuleCondTestFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || condTestAdded

	err, indexAdded := retrieveServerSwitchingRuleIndexFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || indexAdded

	err, targetServerAdded := retrieveServerSwitchingRuleTargetServerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || targetServerAdded

	return nil, retAdded
}

func retrieveServerSwitchingRuleCondFlags(depth int, m *models.ServerSwitchingRule, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveServerSwitchingRuleCondTestFlags(depth int, m *models.ServerSwitchingRule, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveServerSwitchingRuleIndexFlags(depth int, m *models.ServerSwitchingRule, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	indexFlagName := fmt.Sprintf("%v.index", cmdPrefix)
	if cmd.Flags().Changed(indexFlagName) {

		var indexFlagName string
		if cmdPrefix == "" {
			indexFlagName = "index"
		} else {
			indexFlagName = fmt.Sprintf("%v.index", cmdPrefix)
		}

		indexFlagValue, err := cmd.Flags().GetInt64(indexFlagName)
		if err != nil {
			return err, false
		}
		m.Index = &indexFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSwitchingRuleTargetServerFlags(depth int, m *models.ServerSwitchingRule, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	targetServerFlagName := fmt.Sprintf("%v.target_server", cmdPrefix)
	if cmd.Flags().Changed(targetServerFlagName) {

		var targetServerFlagName string
		if cmdPrefix == "" {
			targetServerFlagName = "target_server"
		} else {
			targetServerFlagName = fmt.Sprintf("%v.target_server", cmdPrefix)
		}

		targetServerFlagValue, err := cmd.Flags().GetString(targetServerFlagName)
		if err != nil {
			return err, false
		}
		m.TargetServer = targetServerFlagValue

		retAdded = true
	}

	return nil, retAdded
}