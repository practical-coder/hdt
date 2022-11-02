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

// makeOperationSpoeGetSpoeAgentsCmd returns a cmd to handle operation getSpoeAgents
func makeOperationSpoeGetSpoeAgentsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getSpoeAgents",
		Short: `Returns an array of all configured spoe agents in one scope.`,
		RunE:  runOperationSpoeGetSpoeAgents,
	}

	if err := registerOperationSpoeGetSpoeAgentsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSpoeGetSpoeAgents uses cmd flags to call endpoint api
func runOperationSpoeGetSpoeAgents(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := spoe.NewGetSpoeAgentsParams()
	if err, _ := retrieveOperationSpoeGetSpoeAgentsScopeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSpoeGetSpoeAgentsSpoeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSpoeGetSpoeAgentsTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSpoeGetSpoeAgentsResult(appCli.Spoe.GetSpoeAgents(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSpoeGetSpoeAgentsParamFlags registers all flags needed to fill params
func registerOperationSpoeGetSpoeAgentsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSpoeGetSpoeAgentsScopeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSpoeGetSpoeAgentsSpoeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSpoeGetSpoeAgentsTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSpoeGetSpoeAgentsScopeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSpoeGetSpoeAgentsSpoeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSpoeGetSpoeAgentsTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSpoeGetSpoeAgentsScopeFlag(m *spoe.GetSpoeAgentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSpoeGetSpoeAgentsSpoeFlag(m *spoe.GetSpoeAgentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSpoeGetSpoeAgentsTransactionIDFlag(m *spoe.GetSpoeAgentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSpoeGetSpoeAgentsResult parses request result and return the string content
func parseOperationSpoeGetSpoeAgentsResult(resp0 *spoe.GetSpoeAgentsOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*spoe.GetSpoeAgentsDefault)
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
		resp0, ok := iResp0.(*spoe.GetSpoeAgentsOK)
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
func registerModelGetSpoeAgentsOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetSpoeAgentsOKBodyVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGetSpoeAgentsOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetSpoeAgentsOKBodyVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerGetSpoeAgentsOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data models.SpoeAgents array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetSpoeAgentsOKBodyFlags(depth int, m *spoe.GetSpoeAgentsOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, versionAdded := retrieveGetSpoeAgentsOKBodyVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	err, dataAdded := retrieveGetSpoeAgentsOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetSpoeAgentsOKBodyVersionFlags(depth int, m *spoe.GetSpoeAgentsOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveGetSpoeAgentsOKBodyDataFlags(depth int, m *spoe.GetSpoeAgentsOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type models.SpoeAgents is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
