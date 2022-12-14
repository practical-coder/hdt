// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/declare_capture"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDeclareCaptureDeleteDeclareCaptureCmd returns a cmd to handle operation deleteDeclareCapture
func makeOperationDeclareCaptureDeleteDeclareCaptureCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteDeclareCapture",
		Short: `Deletes a declare capture configuration by it's index in the specified frontend.`,
		RunE:  runOperationDeclareCaptureDeleteDeclareCapture,
	}

	if err := registerOperationDeclareCaptureDeleteDeclareCaptureParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDeclareCaptureDeleteDeclareCapture uses cmd flags to call endpoint api
func runOperationDeclareCaptureDeleteDeclareCapture(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := declare_capture.NewDeleteDeclareCaptureParams()
	if err, _ := retrieveOperationDeclareCaptureDeleteDeclareCaptureForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDeclareCaptureDeleteDeclareCaptureFrontendFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDeclareCaptureDeleteDeclareCaptureIndexFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDeclareCaptureDeleteDeclareCaptureTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDeclareCaptureDeleteDeclareCaptureVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDeclareCaptureDeleteDeclareCaptureResult(appCli.DeclareCapture.DeleteDeclareCapture(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDeclareCaptureDeleteDeclareCaptureParamFlags registers all flags needed to fill params
func registerOperationDeclareCaptureDeleteDeclareCaptureParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDeclareCaptureDeleteDeclareCaptureForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDeclareCaptureDeleteDeclareCaptureFrontendParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDeclareCaptureDeleteDeclareCaptureIndexParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDeclareCaptureDeleteDeclareCaptureTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDeclareCaptureDeleteDeclareCaptureVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDeclareCaptureDeleteDeclareCaptureForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDeclareCaptureDeleteDeclareCaptureFrontendParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	frontendDescription := `Required. Parent frontend name`

	var frontendFlagName string
	if cmdPrefix == "" {
		frontendFlagName = "frontend"
	} else {
		frontendFlagName = fmt.Sprintf("%v.frontend", cmdPrefix)
	}

	var frontendFlagDefault string

	_ = cmd.PersistentFlags().String(frontendFlagName, frontendFlagDefault, frontendDescription)

	return nil
}
func registerOperationDeclareCaptureDeleteDeclareCaptureIndexParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	indexDescription := `Required. Declare Capture Index`

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
func registerOperationDeclareCaptureDeleteDeclareCaptureTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDeclareCaptureDeleteDeclareCaptureVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationDeclareCaptureDeleteDeclareCaptureForceReloadFlag(m *declare_capture.DeleteDeclareCaptureParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDeclareCaptureDeleteDeclareCaptureFrontendFlag(m *declare_capture.DeleteDeclareCaptureParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("frontend") {

		var frontendFlagName string
		if cmdPrefix == "" {
			frontendFlagName = "frontend"
		} else {
			frontendFlagName = fmt.Sprintf("%v.frontend", cmdPrefix)
		}

		frontendFlagValue, err := cmd.Flags().GetString(frontendFlagName)
		if err != nil {
			return err, false
		}
		m.Frontend = frontendFlagValue

	}
	return nil, retAdded
}
func retrieveOperationDeclareCaptureDeleteDeclareCaptureIndexFlag(m *declare_capture.DeleteDeclareCaptureParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDeclareCaptureDeleteDeclareCaptureTransactionIDFlag(m *declare_capture.DeleteDeclareCaptureParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDeclareCaptureDeleteDeclareCaptureVersionFlag(m *declare_capture.DeleteDeclareCaptureParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDeclareCaptureDeleteDeclareCaptureResult parses request result and return the string content
func parseOperationDeclareCaptureDeleteDeclareCaptureResult(resp0 *declare_capture.DeleteDeclareCaptureAccepted, resp1 *declare_capture.DeleteDeclareCaptureNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*declare_capture.DeleteDeclareCaptureDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning deleteDeclareCaptureAccepted is not supported

		// Non schema case: warning deleteDeclareCaptureNoContent is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*declare_capture.DeleteDeclareCaptureNotFound)
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

	// warning: non schema response deleteDeclareCaptureAccepted is not supported by go-swagger cli yet.

	// warning: non schema response deleteDeclareCaptureNoContent is not supported by go-swagger cli yet.

	return "", nil
}
