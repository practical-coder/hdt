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

// makeOperationSpoeReplaceSpoeMessageCmd returns a cmd to handle operation replaceSpoeMessage
func makeOperationSpoeReplaceSpoeMessageCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "replaceSpoeMessage",
		Short: `Replaces a spoe message configuration in one SPOE scope.`,
		RunE:  runOperationSpoeReplaceSpoeMessage,
	}

	if err := registerOperationSpoeReplaceSpoeMessageParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSpoeReplaceSpoeMessage uses cmd flags to call endpoint api
func runOperationSpoeReplaceSpoeMessage(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := spoe.NewReplaceSpoeMessageParams()
	if err, _ := retrieveOperationSpoeReplaceSpoeMessageDataFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSpoeReplaceSpoeMessageNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSpoeReplaceSpoeMessageScopeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSpoeReplaceSpoeMessageSpoeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSpoeReplaceSpoeMessageTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSpoeReplaceSpoeMessageVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSpoeReplaceSpoeMessageResult(appCli.Spoe.ReplaceSpoeMessage(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSpoeReplaceSpoeMessageParamFlags registers all flags needed to fill params
func registerOperationSpoeReplaceSpoeMessageParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSpoeReplaceSpoeMessageDataParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSpoeReplaceSpoeMessageNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSpoeReplaceSpoeMessageScopeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSpoeReplaceSpoeMessageSpoeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSpoeReplaceSpoeMessageTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSpoeReplaceSpoeMessageVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSpoeReplaceSpoeMessageDataParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(dataFlagName, "", "Optional json string for [data]. ")

	// add flags for body
	if err := registerModelSpoeMessageFlags(0, "spoeMessage", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationSpoeReplaceSpoeMessageNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Spoe message name`

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
func registerOperationSpoeReplaceSpoeMessageScopeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSpoeReplaceSpoeMessageSpoeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSpoeReplaceSpoeMessageTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSpoeReplaceSpoeMessageVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSpoeReplaceSpoeMessageDataFlag(m *spoe.ReplaceSpoeMessageParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("data") {
		// Read data string from cmd and unmarshal
		dataValueStr, err := cmd.Flags().GetString("data")
		if err != nil {
			return err, false
		}

		dataValue := models.SpoeMessage{}
		if err := json.Unmarshal([]byte(dataValueStr), &dataValue); err != nil {
			return fmt.Errorf("cannot unmarshal data string in models.SpoeMessage: %v", err), false
		}
		m.Data = &dataValue
	}
	dataValueModel := m.Data
	if swag.IsZero(dataValueModel) {
		dataValueModel = &models.SpoeMessage{}
	}
	err, added := retrieveModelSpoeMessageFlags(0, dataValueModel, "spoeMessage", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Data = dataValueModel
	}
	if dryRun && debug {

		dataValueDebugBytes, err := json.Marshal(m.Data)
		if err != nil {
			return err, false
		}
		logDebugf("Data dry-run payload: %v", string(dataValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationSpoeReplaceSpoeMessageNameFlag(m *spoe.ReplaceSpoeMessageParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSpoeReplaceSpoeMessageScopeFlag(m *spoe.ReplaceSpoeMessageParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSpoeReplaceSpoeMessageSpoeFlag(m *spoe.ReplaceSpoeMessageParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSpoeReplaceSpoeMessageTransactionIDFlag(m *spoe.ReplaceSpoeMessageParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSpoeReplaceSpoeMessageVersionFlag(m *spoe.ReplaceSpoeMessageParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSpoeReplaceSpoeMessageResult parses request result and return the string content
func parseOperationSpoeReplaceSpoeMessageResult(resp0 *spoe.ReplaceSpoeMessageOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*spoe.ReplaceSpoeMessageDefault)
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
		resp0, ok := iResp0.(*spoe.ReplaceSpoeMessageOK)
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
		resp1, ok := iResp1.(*spoe.ReplaceSpoeMessageBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*spoe.ReplaceSpoeMessageNotFound)
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

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
