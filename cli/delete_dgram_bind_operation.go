// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/dgram_bind"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDgramBindDeleteDgramBindCmd returns a cmd to handle operation deleteDgramBind
func makeOperationDgramBindDeleteDgramBindCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteDgramBind",
		Short: `Deletes a dgram bind configuration by it's name in the specified log forward.`,
		RunE:  runOperationDgramBindDeleteDgramBind,
	}

	if err := registerOperationDgramBindDeleteDgramBindParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDgramBindDeleteDgramBind uses cmd flags to call endpoint api
func runOperationDgramBindDeleteDgramBind(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := dgram_bind.NewDeleteDgramBindParams()
	if err, _ := retrieveOperationDgramBindDeleteDgramBindForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDgramBindDeleteDgramBindLogForwardFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDgramBindDeleteDgramBindNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDgramBindDeleteDgramBindTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDgramBindDeleteDgramBindVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDgramBindDeleteDgramBindResult(appCli.DgramBind.DeleteDgramBind(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDgramBindDeleteDgramBindParamFlags registers all flags needed to fill params
func registerOperationDgramBindDeleteDgramBindParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDgramBindDeleteDgramBindForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDgramBindDeleteDgramBindLogForwardParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDgramBindDeleteDgramBindNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDgramBindDeleteDgramBindTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDgramBindDeleteDgramBindVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDgramBindDeleteDgramBindForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDgramBindDeleteDgramBindLogForwardParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	logForwardDescription := `Required. Parent log forward name`

	var logForwardFlagName string
	if cmdPrefix == "" {
		logForwardFlagName = "log_forward"
	} else {
		logForwardFlagName = fmt.Sprintf("%v.log_forward", cmdPrefix)
	}

	var logForwardFlagDefault string

	_ = cmd.PersistentFlags().String(logForwardFlagName, logForwardFlagDefault, logForwardDescription)

	return nil
}
func registerOperationDgramBindDeleteDgramBindNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Bind name`

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
func registerOperationDgramBindDeleteDgramBindTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDgramBindDeleteDgramBindVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationDgramBindDeleteDgramBindForceReloadFlag(m *dgram_bind.DeleteDgramBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDgramBindDeleteDgramBindLogForwardFlag(m *dgram_bind.DeleteDgramBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("log_forward") {

		var logForwardFlagName string
		if cmdPrefix == "" {
			logForwardFlagName = "log_forward"
		} else {
			logForwardFlagName = fmt.Sprintf("%v.log_forward", cmdPrefix)
		}

		logForwardFlagValue, err := cmd.Flags().GetString(logForwardFlagName)
		if err != nil {
			return err, false
		}
		m.LogForward = logForwardFlagValue

	}
	return nil, retAdded
}
func retrieveOperationDgramBindDeleteDgramBindNameFlag(m *dgram_bind.DeleteDgramBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDgramBindDeleteDgramBindTransactionIDFlag(m *dgram_bind.DeleteDgramBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDgramBindDeleteDgramBindVersionFlag(m *dgram_bind.DeleteDgramBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDgramBindDeleteDgramBindResult parses request result and return the string content
func parseOperationDgramBindDeleteDgramBindResult(resp0 *dgram_bind.DeleteDgramBindAccepted, resp1 *dgram_bind.DeleteDgramBindNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*dgram_bind.DeleteDgramBindDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning deleteDgramBindAccepted is not supported

		// Non schema case: warning deleteDgramBindNoContent is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*dgram_bind.DeleteDgramBindNotFound)
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

	// warning: non schema response deleteDgramBindAccepted is not supported by go-swagger cli yet.

	// warning: non schema response deleteDgramBindNoContent is not supported by go-swagger cli yet.

	return "", nil
}
