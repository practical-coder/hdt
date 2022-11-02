// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/spoe"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSpoeDeleteSpoeAgentCmd returns a cmd to handle operation deleteSpoeAgent
func makeOperationSpoeDeleteSpoeAgentCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteSpoeAgent",
		Short: `Deletes a SPOE agent from the configuration in one SPOE scope.`,
		RunE:  runOperationSpoeDeleteSpoeAgent,
	}

	if err := registerOperationSpoeDeleteSpoeAgentParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSpoeDeleteSpoeAgent uses cmd flags to call endpoint api
func runOperationSpoeDeleteSpoeAgent(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := spoe.NewDeleteSpoeAgentParams()
	if err, _ := retrieveOperationSpoeDeleteSpoeAgentNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSpoeDeleteSpoeAgentScopeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSpoeDeleteSpoeAgentSpoeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSpoeDeleteSpoeAgentTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSpoeDeleteSpoeAgentVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSpoeDeleteSpoeAgentResult(appCli.Spoe.DeleteSpoeAgent(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSpoeDeleteSpoeAgentParamFlags registers all flags needed to fill params
func registerOperationSpoeDeleteSpoeAgentParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSpoeDeleteSpoeAgentNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSpoeDeleteSpoeAgentScopeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSpoeDeleteSpoeAgentSpoeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSpoeDeleteSpoeAgentTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSpoeDeleteSpoeAgentVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSpoeDeleteSpoeAgentNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Spoe agent name`

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
func registerOperationSpoeDeleteSpoeAgentScopeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	scopeDescription := `Required. Spoe scope`

	var scopeFlagName string
	if cmdPrefix == "" {
		scopeFlagName = "scope"
	} else {
		scopeFlagName = fmt.Sprintf("%v.scope", cmdPrefix)
	}

	var scopeFlagDefault string

	_ = cmd.PersistentFlags().String(scopeFlagName, scopeFlagDefault, scopeDescription)

	return nil
}
func registerOperationSpoeDeleteSpoeAgentSpoeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	spoeDescription := `Required. Spoe file name`

	var spoeFlagName string
	if cmdPrefix == "" {
		spoeFlagName = "spoe"
	} else {
		spoeFlagName = fmt.Sprintf("%v.spoe", cmdPrefix)
	}

	var spoeFlagDefault string

	_ = cmd.PersistentFlags().String(spoeFlagName, spoeFlagDefault, spoeDescription)

	return nil
}
func registerOperationSpoeDeleteSpoeAgentTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSpoeDeleteSpoeAgentVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSpoeDeleteSpoeAgentNameFlag(m *spoe.DeleteSpoeAgentParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("name") {

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
		m.Name = nameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationSpoeDeleteSpoeAgentScopeFlag(m *spoe.DeleteSpoeAgentParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("scope") {

		var scopeFlagName string
		if cmdPrefix == "" {
			scopeFlagName = "scope"
		} else {
			scopeFlagName = fmt.Sprintf("%v.scope", cmdPrefix)
		}

		scopeFlagValue, err := cmd.Flags().GetString(scopeFlagName)
		if err != nil {
			return err, false
		}
		m.Scope = scopeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationSpoeDeleteSpoeAgentSpoeFlag(m *spoe.DeleteSpoeAgentParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("spoe") {

		var spoeFlagName string
		if cmdPrefix == "" {
			spoeFlagName = "spoe"
		} else {
			spoeFlagName = fmt.Sprintf("%v.spoe", cmdPrefix)
		}

		spoeFlagValue, err := cmd.Flags().GetString(spoeFlagName)
		if err != nil {
			return err, false
		}
		m.Spoe = spoeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationSpoeDeleteSpoeAgentTransactionIDFlag(m *spoe.DeleteSpoeAgentParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSpoeDeleteSpoeAgentVersionFlag(m *spoe.DeleteSpoeAgentParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSpoeDeleteSpoeAgentResult parses request result and return the string content
func parseOperationSpoeDeleteSpoeAgentResult(resp0 *spoe.DeleteSpoeAgentNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*spoe.DeleteSpoeAgentDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning deleteSpoeAgentNoContent is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*spoe.DeleteSpoeAgentNotFound)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response deleteSpoeAgentNoContent is not supported by go-swagger cli yet.

	return "", nil
}
