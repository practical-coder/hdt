// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/server_template"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServerTemplateDeleteServerTemplateCmd returns a cmd to handle operation deleteServerTemplate
func makeOperationServerTemplateDeleteServerTemplateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteServerTemplate",
		Short: `Deletes a server template configuration by it's prefix in the specified backend.`,
		RunE:  runOperationServerTemplateDeleteServerTemplate,
	}

	if err := registerOperationServerTemplateDeleteServerTemplateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServerTemplateDeleteServerTemplate uses cmd flags to call endpoint api
func runOperationServerTemplateDeleteServerTemplate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := server_template.NewDeleteServerTemplateParams()
	if err, _ := retrieveOperationServerTemplateDeleteServerTemplateBackendFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServerTemplateDeleteServerTemplateForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServerTemplateDeleteServerTemplatePrefixFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServerTemplateDeleteServerTemplateTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServerTemplateDeleteServerTemplateVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServerTemplateDeleteServerTemplateResult(appCli.ServerTemplate.DeleteServerTemplate(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServerTemplateDeleteServerTemplateParamFlags registers all flags needed to fill params
func registerOperationServerTemplateDeleteServerTemplateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServerTemplateDeleteServerTemplateBackendParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServerTemplateDeleteServerTemplateForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServerTemplateDeleteServerTemplatePrefixParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServerTemplateDeleteServerTemplateTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServerTemplateDeleteServerTemplateVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServerTemplateDeleteServerTemplateBackendParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	backendDescription := `Required. Parent backend name`

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
func registerOperationServerTemplateDeleteServerTemplateForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationServerTemplateDeleteServerTemplatePrefixParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	prefixDescription := `Required. Server template prefix`

	var prefixFlagName string
	if cmdPrefix == "" {
		prefixFlagName = "prefix"
	} else {
		prefixFlagName = fmt.Sprintf("%v.prefix", cmdPrefix)
	}

	var prefixFlagDefault string

	_ = cmd.PersistentFlags().String(prefixFlagName, prefixFlagDefault, prefixDescription)

	return nil
}
func registerOperationServerTemplateDeleteServerTemplateTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationServerTemplateDeleteServerTemplateVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationServerTemplateDeleteServerTemplateBackendFlag(m *server_template.DeleteServerTemplateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServerTemplateDeleteServerTemplateForceReloadFlag(m *server_template.DeleteServerTemplateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServerTemplateDeleteServerTemplatePrefixFlag(m *server_template.DeleteServerTemplateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("prefix") {

		var prefixFlagName string
		if cmdPrefix == "" {
			prefixFlagName = "prefix"
		} else {
			prefixFlagName = fmt.Sprintf("%v.prefix", cmdPrefix)
		}

		prefixFlagValue, err := cmd.Flags().GetString(prefixFlagName)
		if err != nil {
			return err, false
		}
		m.Prefix = prefixFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServerTemplateDeleteServerTemplateTransactionIDFlag(m *server_template.DeleteServerTemplateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServerTemplateDeleteServerTemplateVersionFlag(m *server_template.DeleteServerTemplateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationServerTemplateDeleteServerTemplateResult parses request result and return the string content
func parseOperationServerTemplateDeleteServerTemplateResult(resp0 *server_template.DeleteServerTemplateAccepted, resp1 *server_template.DeleteServerTemplateNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*server_template.DeleteServerTemplateDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning deleteServerTemplateAccepted is not supported

		// Non schema case: warning deleteServerTemplateNoContent is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*server_template.DeleteServerTemplateNotFound)
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

	// warning: non schema response deleteServerTemplateAccepted is not supported by go-swagger cli yet.

	// warning: non schema response deleteServerTemplateNoContent is not supported by go-swagger cli yet.

	return "", nil
}