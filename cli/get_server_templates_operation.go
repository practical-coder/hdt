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

// makeOperationServerTemplateGetServerTemplatesCmd returns a cmd to handle operation getServerTemplates
func makeOperationServerTemplateGetServerTemplatesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getServerTemplates",
		Short: `Returns an array of all server templates that are configured in specified backend.`,
		RunE:  runOperationServerTemplateGetServerTemplates,
	}

	if err := registerOperationServerTemplateGetServerTemplatesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServerTemplateGetServerTemplates uses cmd flags to call endpoint api
func runOperationServerTemplateGetServerTemplates(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := server_template.NewGetServerTemplatesParams()
	if err, _ := retrieveOperationServerTemplateGetServerTemplatesBackendFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServerTemplateGetServerTemplatesTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServerTemplateGetServerTemplatesResult(appCli.ServerTemplate.GetServerTemplates(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServerTemplateGetServerTemplatesParamFlags registers all flags needed to fill params
func registerOperationServerTemplateGetServerTemplatesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServerTemplateGetServerTemplatesBackendParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServerTemplateGetServerTemplatesTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServerTemplateGetServerTemplatesBackendParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationServerTemplateGetServerTemplatesTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationServerTemplateGetServerTemplatesBackendFlag(m *server_template.GetServerTemplatesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServerTemplateGetServerTemplatesTransactionIDFlag(m *server_template.GetServerTemplatesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationServerTemplateGetServerTemplatesResult parses request result and return the string content
func parseOperationServerTemplateGetServerTemplatesResult(resp0 *server_template.GetServerTemplatesOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*server_template.GetServerTemplatesDefault)
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
		resp0, ok := iResp0.(*server_template.GetServerTemplatesOK)
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
func registerModelGetServerTemplatesOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetServerTemplatesOKBodyVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGetServerTemplatesOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetServerTemplatesOKBodyVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerGetServerTemplatesOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data models.ServerTemplates array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetServerTemplatesOKBodyFlags(depth int, m *server_template.GetServerTemplatesOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, versionAdded := retrieveGetServerTemplatesOKBodyVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	err, dataAdded := retrieveGetServerTemplatesOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetServerTemplatesOKBodyVersionFlags(depth int, m *server_template.GetServerTemplatesOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveGetServerTemplatesOKBodyDataFlags(depth int, m *server_template.GetServerTemplatesOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type models.ServerTemplates is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
