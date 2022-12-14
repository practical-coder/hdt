// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/tcp_response_rule"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationTCPResponseRuleDeleteTCPResponseRuleCmd returns a cmd to handle operation deleteTcpResponseRule
func makeOperationTCPResponseRuleDeleteTCPResponseRuleCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteTCPResponseRule",
		Short: `Deletes a TCP Response Rule configuration by it's index from the specified backend.`,
		RunE:  runOperationTCPResponseRuleDeleteTCPResponseRule,
	}

	if err := registerOperationTCPResponseRuleDeleteTCPResponseRuleParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTCPResponseRuleDeleteTCPResponseRule uses cmd flags to call endpoint api
func runOperationTCPResponseRuleDeleteTCPResponseRule(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := tcp_response_rule.NewDeleteTCPResponseRuleParams()
	if err, _ := retrieveOperationTCPResponseRuleDeleteTCPResponseRuleBackendFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTCPResponseRuleDeleteTCPResponseRuleForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTCPResponseRuleDeleteTCPResponseRuleIndexFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTCPResponseRuleDeleteTCPResponseRuleTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTCPResponseRuleDeleteTCPResponseRuleVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTCPResponseRuleDeleteTCPResponseRuleResult(appCli.TCPResponseRule.DeleteTCPResponseRule(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTCPResponseRuleDeleteTCPResponseRuleParamFlags registers all flags needed to fill params
func registerOperationTCPResponseRuleDeleteTCPResponseRuleParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTCPResponseRuleDeleteTCPResponseRuleBackendParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTCPResponseRuleDeleteTCPResponseRuleForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTCPResponseRuleDeleteTCPResponseRuleIndexParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTCPResponseRuleDeleteTCPResponseRuleTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTCPResponseRuleDeleteTCPResponseRuleVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTCPResponseRuleDeleteTCPResponseRuleBackendParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	backendDescription := `Required. Parent backend name`

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
func registerOperationTCPResponseRuleDeleteTCPResponseRuleForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationTCPResponseRuleDeleteTCPResponseRuleIndexParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	indexDescription := `Required. TCP Response Rule Index`

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
func registerOperationTCPResponseRuleDeleteTCPResponseRuleTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationTCPResponseRuleDeleteTCPResponseRuleVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationTCPResponseRuleDeleteTCPResponseRuleBackendFlag(m *tcp_response_rule.DeleteTCPResponseRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationTCPResponseRuleDeleteTCPResponseRuleForceReloadFlag(m *tcp_response_rule.DeleteTCPResponseRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationTCPResponseRuleDeleteTCPResponseRuleIndexFlag(m *tcp_response_rule.DeleteTCPResponseRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationTCPResponseRuleDeleteTCPResponseRuleTransactionIDFlag(m *tcp_response_rule.DeleteTCPResponseRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationTCPResponseRuleDeleteTCPResponseRuleVersionFlag(m *tcp_response_rule.DeleteTCPResponseRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationTCPResponseRuleDeleteTCPResponseRuleResult parses request result and return the string content
func parseOperationTCPResponseRuleDeleteTCPResponseRuleResult(resp0 *tcp_response_rule.DeleteTCPResponseRuleAccepted, resp1 *tcp_response_rule.DeleteTCPResponseRuleNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*tcp_response_rule.DeleteTCPResponseRuleDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning deleteTcpResponseRuleAccepted is not supported

		// Non schema case: warning deleteTcpResponseRuleNoContent is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*tcp_response_rule.DeleteTCPResponseRuleNotFound)
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

	// warning: non schema response deleteTcpResponseRuleAccepted is not supported by go-swagger cli yet.

	// warning: non schema response deleteTcpResponseRuleNoContent is not supported by go-swagger cli yet.

	return "", nil
}
