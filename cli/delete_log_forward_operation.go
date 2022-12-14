// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/log_forward"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationLogForwardDeleteLogForwardCmd returns a cmd to handle operation deleteLogForward
func makeOperationLogForwardDeleteLogForwardCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteLogForward",
		Short: `Deletes a log forward from the configuration by it's name.`,
		RunE:  runOperationLogForwardDeleteLogForward,
	}

	if err := registerOperationLogForwardDeleteLogForwardParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationLogForwardDeleteLogForward uses cmd flags to call endpoint api
func runOperationLogForwardDeleteLogForward(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := log_forward.NewDeleteLogForwardParams()
	if err, _ := retrieveOperationLogForwardDeleteLogForwardForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLogForwardDeleteLogForwardNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLogForwardDeleteLogForwardTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLogForwardDeleteLogForwardVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationLogForwardDeleteLogForwardResult(appCli.LogForward.DeleteLogForward(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationLogForwardDeleteLogForwardParamFlags registers all flags needed to fill params
func registerOperationLogForwardDeleteLogForwardParamFlags(cmd *cobra.Command) error {
	if err := registerOperationLogForwardDeleteLogForwardForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLogForwardDeleteLogForwardNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLogForwardDeleteLogForwardTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLogForwardDeleteLogForwardVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationLogForwardDeleteLogForwardForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationLogForwardDeleteLogForwardNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Log Forward name`

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
func registerOperationLogForwardDeleteLogForwardTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationLogForwardDeleteLogForwardVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationLogForwardDeleteLogForwardForceReloadFlag(m *log_forward.DeleteLogForwardParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationLogForwardDeleteLogForwardNameFlag(m *log_forward.DeleteLogForwardParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationLogForwardDeleteLogForwardTransactionIDFlag(m *log_forward.DeleteLogForwardParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationLogForwardDeleteLogForwardVersionFlag(m *log_forward.DeleteLogForwardParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationLogForwardDeleteLogForwardResult parses request result and return the string content
func parseOperationLogForwardDeleteLogForwardResult(resp0 *log_forward.DeleteLogForwardAccepted, resp1 *log_forward.DeleteLogForwardNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*log_forward.DeleteLogForwardDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning deleteLogForwardAccepted is not supported

		// Non schema case: warning deleteLogForwardNoContent is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*log_forward.DeleteLogForwardNotFound)
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

	// warning: non schema response deleteLogForwardAccepted is not supported by go-swagger cli yet.

	// warning: non schema response deleteLogForwardNoContent is not supported by go-swagger cli yet.

	return "", nil
}
