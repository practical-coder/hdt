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

// makeOperationStickRuleGetStickRulesCmd returns a cmd to handle operation getStickRules
func makeOperationStickRuleGetStickRulesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getStickRules",
		Short: `Returns all Stick Rules that are configured in specified backend.`,
		RunE:  runOperationStickRuleGetStickRules,
	}

	if err := registerOperationStickRuleGetStickRulesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationStickRuleGetStickRules uses cmd flags to call endpoint api
func runOperationStickRuleGetStickRules(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := stick_rule.NewGetStickRulesParams()
	if err, _ := retrieveOperationStickRuleGetStickRulesBackendFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStickRuleGetStickRulesTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationStickRuleGetStickRulesResult(appCli.StickRule.GetStickRules(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationStickRuleGetStickRulesParamFlags registers all flags needed to fill params
func registerOperationStickRuleGetStickRulesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationStickRuleGetStickRulesBackendParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStickRuleGetStickRulesTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationStickRuleGetStickRulesBackendParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationStickRuleGetStickRulesTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationStickRuleGetStickRulesBackendFlag(m *stick_rule.GetStickRulesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationStickRuleGetStickRulesTransactionIDFlag(m *stick_rule.GetStickRulesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationStickRuleGetStickRulesResult parses request result and return the string content
func parseOperationStickRuleGetStickRulesResult(resp0 *stick_rule.GetStickRulesOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*stick_rule.GetStickRulesDefault)
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
		resp0, ok := iResp0.(*stick_rule.GetStickRulesOK)
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
func registerModelGetStickRulesOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetStickRulesOKBodyVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGetStickRulesOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetStickRulesOKBodyVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerGetStickRulesOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data models.StickRules array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetStickRulesOKBodyFlags(depth int, m *stick_rule.GetStickRulesOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, versionAdded := retrieveGetStickRulesOKBodyVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	err, dataAdded := retrieveGetStickRulesOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetStickRulesOKBodyVersionFlags(depth int, m *stick_rule.GetStickRulesOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveGetStickRulesOKBodyDataFlags(depth int, m *stick_rule.GetStickRulesOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type models.StickRules is not supported by go-swagger cli yet
	}

	return nil, retAdded
}