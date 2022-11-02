// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/server_switching_rule"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServerSwitchingRuleGetServerSwitchingRulesCmd returns a cmd to handle operation getServerSwitchingRules
func makeOperationServerSwitchingRuleGetServerSwitchingRulesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getServerSwitchingRules",
		Short: `Returns all Backend Switching Rules that are configured in specified backend.`,
		RunE:  runOperationServerSwitchingRuleGetServerSwitchingRules,
	}

	if err := registerOperationServerSwitchingRuleGetServerSwitchingRulesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServerSwitchingRuleGetServerSwitchingRules uses cmd flags to call endpoint api
func runOperationServerSwitchingRuleGetServerSwitchingRules(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := server_switching_rule.NewGetServerSwitchingRulesParams()
	if err, _ := retrieveOperationServerSwitchingRuleGetServerSwitchingRulesBackendFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServerSwitchingRuleGetServerSwitchingRulesTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServerSwitchingRuleGetServerSwitchingRulesResult(appCli.ServerSwitchingRule.GetServerSwitchingRules(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServerSwitchingRuleGetServerSwitchingRulesParamFlags registers all flags needed to fill params
func registerOperationServerSwitchingRuleGetServerSwitchingRulesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServerSwitchingRuleGetServerSwitchingRulesBackendParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServerSwitchingRuleGetServerSwitchingRulesTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServerSwitchingRuleGetServerSwitchingRulesBackendParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationServerSwitchingRuleGetServerSwitchingRulesTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationServerSwitchingRuleGetServerSwitchingRulesBackendFlag(m *server_switching_rule.GetServerSwitchingRulesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServerSwitchingRuleGetServerSwitchingRulesTransactionIDFlag(m *server_switching_rule.GetServerSwitchingRulesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationServerSwitchingRuleGetServerSwitchingRulesResult parses request result and return the string content
func parseOperationServerSwitchingRuleGetServerSwitchingRulesResult(resp0 *server_switching_rule.GetServerSwitchingRulesOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*server_switching_rule.GetServerSwitchingRulesDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*server_switching_rule.GetServerSwitchingRulesOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}

// register flags to command
func registerModelGetServerSwitchingRulesOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetServerSwitchingRulesOKBodyVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGetServerSwitchingRulesOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetServerSwitchingRulesOKBodyVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	versionDescription := ``

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "_version"
	} else {
		versionFlagName = fmt.Sprintf("%v._version", cmdPrefix)
	}

	var versionFlagDefault int64

	_ = cmd.PersistentFlags().Int64(versionFlagName, versionFlagDefault, versionDescription)

	return nil
}

func registerGetServerSwitchingRulesOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data models.ServerSwitchingRules array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetServerSwitchingRulesOKBodyFlags(depth int, m *server_switching_rule.GetServerSwitchingRulesOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, versionAdded := retrieveGetServerSwitchingRulesOKBodyVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	err, dataAdded := retrieveGetServerSwitchingRulesOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetServerSwitchingRulesOKBodyVersionFlags(depth int, m *server_switching_rule.GetServerSwitchingRulesOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	versionFlagName := fmt.Sprintf("%v._version", cmdPrefix)
	if cmd.Flags().Changed(versionFlagName) {

		var versionFlagName string
		if cmdPrefix == "" {
			versionFlagName = "_version"
		} else {
			versionFlagName = fmt.Sprintf("%v._version", cmdPrefix)
		}

		versionFlagValue, err := cmd.Flags().GetInt64(versionFlagName)
		if err != nil {
			return err, false
		}
		m.Version = versionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveGetServerSwitchingRulesOKBodyDataFlags(depth int, m *server_switching_rule.GetServerSwitchingRulesOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type models.ServerSwitchingRules is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
