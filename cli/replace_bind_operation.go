// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/practical-coder/hdt/client/bind"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationBindReplaceBindCmd returns a cmd to handle operation replaceBind
func makeOperationBindReplaceBindCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "replaceBind",
		Short: `Replaces a bind configuration by it's name in the specified frontend.`,
		RunE:  runOperationBindReplaceBind,
	}

	if err := registerOperationBindReplaceBindParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBindReplaceBind uses cmd flags to call endpoint api
func runOperationBindReplaceBind(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := bind.NewReplaceBindParams()
	if err, _ := retrieveOperationBindReplaceBindDataFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBindReplaceBindForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBindReplaceBindFrontendFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBindReplaceBindNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBindReplaceBindParentNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBindReplaceBindParentTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBindReplaceBindTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBindReplaceBindVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBindReplaceBindResult(appCli.Bind.ReplaceBind(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBindReplaceBindParamFlags registers all flags needed to fill params
func registerOperationBindReplaceBindParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBindReplaceBindDataParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBindReplaceBindForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBindReplaceBindFrontendParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBindReplaceBindNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBindReplaceBindParentNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBindReplaceBindParentTypeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBindReplaceBindTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBindReplaceBindVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBindReplaceBindDataParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(dataFlagName, "", "Optional json string for [data]. ")

	// add flags for body
	if err := registerModelBindFlags(0, "bind", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationBindReplaceBindForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationBindReplaceBindFrontendParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	frontendDescription := `Parent frontend name`

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
func registerOperationBindReplaceBindNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationBindReplaceBindParentNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	parentNameDescription := `Parent name`

	var parentNameFlagName string
	if cmdPrefix == "" {
		parentNameFlagName = "parent_name"
	} else {
		parentNameFlagName = fmt.Sprintf("%v.parent_name", cmdPrefix)
	}

	var parentNameFlagDefault string

	_ = cmd.PersistentFlags().String(parentNameFlagName, parentNameFlagDefault, parentNameDescription)

	return nil
}
func registerOperationBindReplaceBindParentTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	parentTypeDescription := `Enum: ["frontend","log_forward"]. Parent type`

	var parentTypeFlagName string
	if cmdPrefix == "" {
		parentTypeFlagName = "parent_type"
	} else {
		parentTypeFlagName = fmt.Sprintf("%v.parent_type", cmdPrefix)
	}

	var parentTypeFlagDefault string

	_ = cmd.PersistentFlags().String(parentTypeFlagName, parentTypeFlagDefault, parentTypeDescription)

	if err := cmd.RegisterFlagCompletionFunc(parentTypeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["frontend","log_forward"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}
func registerOperationBindReplaceBindTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationBindReplaceBindVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationBindReplaceBindDataFlag(m *bind.ReplaceBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("data") {
		// Read data string from cmd and unmarshal
		dataValueStr, err := cmd.Flags().GetString("data")
		if err != nil {
			return err, false
		}

		dataValue := models.Bind{}
		if err := json.Unmarshal([]byte(dataValueStr), &dataValue); err != nil {
			return fmt.Errorf("cannot unmarshal data string in models.Bind: %v", err), false
		}
		m.Data = &dataValue
	}
	dataValueModel := m.Data
	if swag.IsZero(dataValueModel) {
		dataValueModel = &models.Bind{}
	}
	err, added := retrieveModelBindFlags(0, dataValueModel, "bind", cmd)
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
func retrieveOperationBindReplaceBindForceReloadFlag(m *bind.ReplaceBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBindReplaceBindFrontendFlag(m *bind.ReplaceBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Frontend = &frontendFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBindReplaceBindNameFlag(m *bind.ReplaceBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBindReplaceBindParentNameFlag(m *bind.ReplaceBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("parent_name") {

		var parentNameFlagName string
		if cmdPrefix == "" {
			parentNameFlagName = "parent_name"
		} else {
			parentNameFlagName = fmt.Sprintf("%v.parent_name", cmdPrefix)
		}

		parentNameFlagValue, err := cmd.Flags().GetString(parentNameFlagName)
		if err != nil {
			return err, false
		}
		m.ParentName = &parentNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBindReplaceBindParentTypeFlag(m *bind.ReplaceBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("parent_type") {

		var parentTypeFlagName string
		if cmdPrefix == "" {
			parentTypeFlagName = "parent_type"
		} else {
			parentTypeFlagName = fmt.Sprintf("%v.parent_type", cmdPrefix)
		}

		parentTypeFlagValue, err := cmd.Flags().GetString(parentTypeFlagName)
		if err != nil {
			return err, false
		}
		m.ParentType = &parentTypeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBindReplaceBindTransactionIDFlag(m *bind.ReplaceBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBindReplaceBindVersionFlag(m *bind.ReplaceBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationBindReplaceBindResult parses request result and return the string content
func parseOperationBindReplaceBindResult(resp0 *bind.ReplaceBindOK, resp1 *bind.ReplaceBindAccepted, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*bind.ReplaceBindDefault)
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
		resp0, ok := iResp0.(*bind.ReplaceBindOK)
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
		resp1, ok := iResp1.(*bind.ReplaceBindAccepted)
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
		resp2, ok := iResp2.(*bind.ReplaceBindBadRequest)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*bind.ReplaceBindNotFound)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
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

	if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
		msgStr, err := json.Marshal(resp1.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
