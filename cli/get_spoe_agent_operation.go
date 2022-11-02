// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/practical-coder/hdt/client/spoe"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSpoeGetSpoeAgentCmd returns a cmd to handle operation getSpoeAgent
func makeOperationSpoeGetSpoeAgentCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getSpoeAgent",
		Short: `Returns one spoe agent configuration in one SPOE scope.`,
		RunE:  runOperationSpoeGetSpoeAgent,
	}

	if err := registerOperationSpoeGetSpoeAgentParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSpoeGetSpoeAgent uses cmd flags to call endpoint api
func runOperationSpoeGetSpoeAgent(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := spoe.NewGetSpoeAgentParams()
	if err, _ := retrieveOperationSpoeGetSpoeAgentNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSpoeGetSpoeAgentScopeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSpoeGetSpoeAgentSpoeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSpoeGetSpoeAgentTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSpoeGetSpoeAgentResult(appCli.Spoe.GetSpoeAgent(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSpoeGetSpoeAgentParamFlags registers all flags needed to fill params
func registerOperationSpoeGetSpoeAgentParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSpoeGetSpoeAgentNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSpoeGetSpoeAgentScopeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSpoeGetSpoeAgentSpoeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSpoeGetSpoeAgentTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSpoeGetSpoeAgentNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSpoeGetSpoeAgentScopeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSpoeGetSpoeAgentSpoeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSpoeGetSpoeAgentTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSpoeGetSpoeAgentNameFlag(m *spoe.GetSpoeAgentParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSpoeGetSpoeAgentScopeFlag(m *spoe.GetSpoeAgentParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSpoeGetSpoeAgentSpoeFlag(m *spoe.GetSpoeAgentParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSpoeGetSpoeAgentTransactionIDFlag(m *spoe.GetSpoeAgentParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSpoeGetSpoeAgentResult parses request result and return the string content
func parseOperationSpoeGetSpoeAgentResult(resp0 *spoe.GetSpoeAgentOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*spoe.GetSpoeAgentDefault)
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
		resp0, ok := iResp0.(*spoe.GetSpoeAgentOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*spoe.GetSpoeAgentNotFound)
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
func registerModelGetSpoeAgentOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetSpoeAgentOKBodyVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGetSpoeAgentOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetSpoeAgentOKBodyVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerGetSpoeAgentOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelSpoeAgentFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetSpoeAgentOKBodyFlags(depth int, m *spoe.GetSpoeAgentOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, versionAdded := retrieveGetSpoeAgentOKBodyVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	err, dataAdded := retrieveGetSpoeAgentOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetSpoeAgentOKBodyVersionFlags(depth int, m *spoe.GetSpoeAgentOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveGetSpoeAgentOKBodyDataFlags(depth int, m *spoe.GetSpoeAgentOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data models.SpoeAgent is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.SpoeAgent{}
	}

	err, dataAdded := retrieveModelSpoeAgentFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}