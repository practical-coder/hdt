// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/stick_rule"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationStickRuleDeleteStickRuleCmd returns a cmd to handle operation deleteStickRule
func makeOperationStickRuleDeleteStickRuleCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteStickRule",
		Short: `Deletes a Stick Rule configuration by it's index from the specified backend.`,
		RunE:  runOperationStickRuleDeleteStickRule,
	}

	if err := registerOperationStickRuleDeleteStickRuleParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationStickRuleDeleteStickRule uses cmd flags to call endpoint api
func runOperationStickRuleDeleteStickRule(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := stick_rule.NewDeleteStickRuleParams()
	if err, _ := retrieveOperationStickRuleDeleteStickRuleBackendFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStickRuleDeleteStickRuleForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStickRuleDeleteStickRuleIndexFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStickRuleDeleteStickRuleTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStickRuleDeleteStickRuleVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationStickRuleDeleteStickRuleResult(appCli.StickRule.DeleteStickRule(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationStickRuleDeleteStickRuleParamFlags registers all flags needed to fill params
func registerOperationStickRuleDeleteStickRuleParamFlags(cmd *cobra.Command) error {
	if err := registerOperationStickRuleDeleteStickRuleBackendParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStickRuleDeleteStickRuleForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStickRuleDeleteStickRuleIndexParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStickRuleDeleteStickRuleTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStickRuleDeleteStickRuleVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationStickRuleDeleteStickRuleBackendParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	backendDescription := `Required. Backend name`

	var backendFlagName string
	if cmdPrefix == "" {
		backendFlagName = "backend"
	} else {
		backendFlagName = fmt.Sprintf("%v.backend", cmdPrefix)
	}

	var backendFlagDefault string

	_ = cmd.PersistentFlags().String(backendFlagName, backendFlagDefault, backendDescription)

	return nil
}
func registerOperationStickRuleDeleteStickRuleForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	forceReloadDescription := `If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.`

	var forceReloadFlagName string
	if cmdPrefix == "" {
		forceReloadFlagName = "force_reload"
	} else {
		forceReloadFlagName = fmt.Sprintf("%v.force_reload", cmdPrefix)
	}

	var forceReloadFlagDefault bool

	_ = cmd.PersistentFlags().Bool(forceReloadFlagName, forceReloadFlagDefault, forceReloadDescription)

	return nil
}
func registerOperationStickRuleDeleteStickRuleIndexParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	indexDescription := `Required. Stick Rule Index`

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
func registerOperationStickRuleDeleteStickRuleTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	transactionIdDescription := `ID of the transaction where we want to add the operation. Cannot be used when version is specified.`

	var transactionIdFlagName string
	if cmdPrefix == "" {
		transactionIdFlagName = "transaction_id"
	} else {
		transactionIdFlagName = fmt.Sprintf("%v.transaction_id", cmdPrefix)
	}

	var transactionIdFlagDefault string

	_ = cmd.PersistentFlags().String(transactionIdFlagName, transactionIdFlagDefault, transactionIdDescription)

	return nil
}
func registerOperationStickRuleDeleteStickRuleVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	versionDescription := `Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.`

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "version"
	} else {
		versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
	}

	var versionFlagDefault int64

	_ = cmd.PersistentFlags().Int64(versionFlagName, versionFlagDefault, versionDescription)

	return nil
}

func retrieveOperationStickRuleDeleteStickRuleBackendFlag(m *stick_rule.DeleteStickRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("backend") {

		var backendFlagName string
		if cmdPrefix == "" {
			backendFlagName = "backend"
		} else {
			backendFlagName = fmt.Sprintf("%v.backend", cmdPrefix)
		}

		backendFlagValue, err := cmd.Flags().GetString(backendFlagName)
		if err != nil {
			return err, false
		}
		m.Backend = backendFlagValue

	}
	return nil, retAdded
}
func retrieveOperationStickRuleDeleteStickRuleForceReloadFlag(m *stick_rule.DeleteStickRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("force_reload") {

		var forceReloadFlagName string
		if cmdPrefix == "" {
			forceReloadFlagName = "force_reload"
		} else {
			forceReloadFlagName = fmt.Sprintf("%v.force_reload", cmdPrefix)
		}

		forceReloadFlagValue, err := cmd.Flags().GetBool(forceReloadFlagName)
		if err != nil {
			return err, false
		}
		m.ForceReload = &forceReloadFlagValue

	}
	return nil, retAdded
}
func retrieveOperationStickRuleDeleteStickRuleIndexFlag(m *stick_rule.DeleteStickRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("index") {

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
		m.Index = indexFlagValue

	}
	return nil, retAdded
}
func retrieveOperationStickRuleDeleteStickRuleTransactionIDFlag(m *stick_rule.DeleteStickRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("transaction_id") {

		var transactionIdFlagName string
		if cmdPrefix == "" {
			transactionIdFlagName = "transaction_id"
		} else {
			transactionIdFlagName = fmt.Sprintf("%v.transaction_id", cmdPrefix)
		}

		transactionIdFlagValue, err := cmd.Flags().GetString(transactionIdFlagName)
		if err != nil {
			return err, false
		}
		m.TransactionID = &transactionIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationStickRuleDeleteStickRuleVersionFlag(m *stick_rule.DeleteStickRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("version") {

		var versionFlagName string
		if cmdPrefix == "" {
			versionFlagName = "version"
		} else {
			versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
		}

		versionFlagValue, err := cmd.Flags().GetInt64(versionFlagName)
		if err != nil {
			return err, false
		}
		m.Version = &versionFlagValue

	}
	return nil, retAdded
}

// parseOperationStickRuleDeleteStickRuleResult parses request result and return the string content
func parseOperationStickRuleDeleteStickRuleResult(resp0 *stick_rule.DeleteStickRuleAccepted, resp1 *stick_rule.DeleteStickRuleNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*stick_rule.DeleteStickRuleDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning deleteStickRuleAccepted is not supported

		// Non schema case: warning deleteStickRuleNoContent is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*stick_rule.DeleteStickRuleNotFound)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response deleteStickRuleAccepted is not supported by go-swagger cli yet.

	// warning: non schema response deleteStickRuleNoContent is not supported by go-swagger cli yet.

	return "", nil
}
