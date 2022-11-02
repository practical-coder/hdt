// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/transactions"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationTransactionsStartTransactionCmd returns a cmd to handle operation startTransaction
func makeOperationTransactionsStartTransactionCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "startTransaction",
		Short: `Starts a new transaction and returns it's id`,
		RunE:  runOperationTransactionsStartTransaction,
	}

	if err := registerOperationTransactionsStartTransactionParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTransactionsStartTransaction uses cmd flags to call endpoint api
func runOperationTransactionsStartTransaction(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := transactions.NewStartTransactionParams()
	if err, _ := retrieveOperationTransactionsStartTransactionVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTransactionsStartTransactionResult(appCli.Transactions.StartTransaction(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTransactionsStartTransactionParamFlags registers all flags needed to fill params
func registerOperationTransactionsStartTransactionParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTransactionsStartTransactionVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTransactionsStartTransactionVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	versionDescription := `Required. Configuration version on which to work on`

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

func retrieveOperationTransactionsStartTransactionVersionFlag(m *transactions.StartTransactionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Version = versionFlagValue

	}
	return nil, retAdded
}

// parseOperationTransactionsStartTransactionResult parses request result and return the string content
func parseOperationTransactionsStartTransactionResult(resp0 *transactions.StartTransactionCreated, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*transactions.StartTransactionDefault)
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
		resp0, ok := iResp0.(*transactions.StartTransactionCreated)
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
		resp1, ok := iResp1.(*transactions.StartTransactionTooManyRequests)
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
func registerModelStartTransactionTooManyRequestsBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerStartTransactionTooManyRequestsBodyCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStartTransactionTooManyRequestsBodyMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerStartTransactionTooManyRequestsBodyCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	codeDescription := ``

	var codeFlagName string
	if cmdPrefix == "" {
		codeFlagName = "code"
	} else {
		codeFlagName = fmt.Sprintf("%v.code", cmdPrefix)
	}

	var codeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(codeFlagName, codeFlagDefault, codeDescription)

	return nil
}

func registerStartTransactionTooManyRequestsBodyMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	messageDescription := ``

	var messageFlagName string
	if cmdPrefix == "" {
		messageFlagName = "message"
	} else {
		messageFlagName = fmt.Sprintf("%v.message", cmdPrefix)
	}

	var messageFlagDefault string

	_ = cmd.PersistentFlags().String(messageFlagName, messageFlagDefault, messageDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelStartTransactionTooManyRequestsBodyFlags(depth int, m *transactions.StartTransactionTooManyRequestsBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, codeAdded := retrieveStartTransactionTooManyRequestsBodyCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || codeAdded

	err, messageAdded := retrieveStartTransactionTooManyRequestsBodyMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageAdded

	return nil, retAdded
}

func retrieveStartTransactionTooManyRequestsBodyCodeFlags(depth int, m *transactions.StartTransactionTooManyRequestsBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	codeFlagName := fmt.Sprintf("%v.code", cmdPrefix)
	if cmd.Flags().Changed(codeFlagName) {

		var codeFlagName string
		if cmdPrefix == "" {
			codeFlagName = "code"
		} else {
			codeFlagName = fmt.Sprintf("%v.code", cmdPrefix)
		}

		codeFlagValue, err := cmd.Flags().GetInt64(codeFlagName)
		if err != nil {
			return err, false
		}
		m.Code = codeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStartTransactionTooManyRequestsBodyMessageFlags(depth int, m *transactions.StartTransactionTooManyRequestsBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	messageFlagName := fmt.Sprintf("%v.message", cmdPrefix)
	if cmd.Flags().Changed(messageFlagName) {

		var messageFlagName string
		if cmdPrefix == "" {
			messageFlagName = "message"
		} else {
			messageFlagName = fmt.Sprintf("%v.message", cmdPrefix)
		}

		messageFlagValue, err := cmd.Flags().GetString(messageFlagName)
		if err != nil {
			return err, false
		}
		m.Message = messageFlagValue

		retAdded = true
	}

	return nil, retAdded
}
