// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/practical-coder/hdt/client/http_check"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationHTTPCheckCreateHTTPCheckCmd returns a cmd to handle operation createHttpCheck
func makeOperationHTTPCheckCreateHTTPCheckCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "createHTTPCheck",
		Short: `Adds a new HTTP check of the specified type in the specified parent.`,
		RunE:  runOperationHTTPCheckCreateHTTPCheck,
	}

	if err := registerOperationHTTPCheckCreateHTTPCheckParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationHTTPCheckCreateHTTPCheck uses cmd flags to call endpoint api
func runOperationHTTPCheckCreateHTTPCheck(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := http_check.NewCreateHTTPCheckParams()
	if err, _ := retrieveOperationHTTPCheckCreateHTTPCheckDataFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationHTTPCheckCreateHTTPCheckForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationHTTPCheckCreateHTTPCheckParentNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationHTTPCheckCreateHTTPCheckParentTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationHTTPCheckCreateHTTPCheckTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationHTTPCheckCreateHTTPCheckVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationHTTPCheckCreateHTTPCheckResult(appCli.HTTPCheck.CreateHTTPCheck(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationHTTPCheckCreateHTTPCheckParamFlags registers all flags needed to fill params
func registerOperationHTTPCheckCreateHTTPCheckParamFlags(cmd *cobra.Command) error {
	if err := registerOperationHTTPCheckCreateHTTPCheckDataParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationHTTPCheckCreateHTTPCheckForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationHTTPCheckCreateHTTPCheckParentNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationHTTPCheckCreateHTTPCheckParentTypeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationHTTPCheckCreateHTTPCheckTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationHTTPCheckCreateHTTPCheckVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationHTTPCheckCreateHTTPCheckDataParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(dataFlagName, "", "Optional json string for [data]. ")

	// add flags for body
	if err := registerModelHTTPCheckFlags(0, "httpCheck", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationHTTPCheckCreateHTTPCheckForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationHTTPCheckCreateHTTPCheckParentNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationHTTPCheckCreateHTTPCheckParentTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	parentTypeDescription := `Enum: ["backend","defaults"]. Required. Parent type`

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
			if err := json.Unmarshal([]byte(`["backend","defaults"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}
func registerOperationHTTPCheckCreateHTTPCheckTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationHTTPCheckCreateHTTPCheckVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationHTTPCheckCreateHTTPCheckDataFlag(m *http_check.CreateHTTPCheckParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("data") {
		// Read data string from cmd and unmarshal
		dataValueStr, err := cmd.Flags().GetString("data")
		if err != nil {
			return err, false
		}

		dataValue := models.HTTPCheck{}
		if err := json.Unmarshal([]byte(dataValueStr), &dataValue); err != nil {
			return fmt.Errorf("cannot unmarshal data string in models.HTTPCheck: %v", err), false
		}
		m.Data = &dataValue
	}
	dataValueModel := m.Data
	if swag.IsZero(dataValueModel) {
		dataValueModel = &models.HTTPCheck{}
	}
	err, added := retrieveModelHTTPCheckFlags(0, dataValueModel, "httpCheck", cmd)
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
func retrieveOperationHTTPCheckCreateHTTPCheckForceReloadFlag(m *http_check.CreateHTTPCheckParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationHTTPCheckCreateHTTPCheckParentNameFlag(m *http_check.CreateHTTPCheckParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationHTTPCheckCreateHTTPCheckParentTypeFlag(m *http_check.CreateHTTPCheckParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.ParentType = parentTypeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationHTTPCheckCreateHTTPCheckTransactionIDFlag(m *http_check.CreateHTTPCheckParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationHTTPCheckCreateHTTPCheckVersionFlag(m *http_check.CreateHTTPCheckParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationHTTPCheckCreateHTTPCheckResult parses request result and return the string content
func parseOperationHTTPCheckCreateHTTPCheckResult(resp0 *http_check.CreateHTTPCheckCreated, resp1 *http_check.CreateHTTPCheckAccepted, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*http_check.CreateHTTPCheckDefault)
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
		resp0, ok := iResp0.(*http_check.CreateHTTPCheckCreated)
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
		resp1, ok := iResp1.(*http_check.CreateHTTPCheckAccepted)
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
		resp2, ok := iResp2.(*http_check.CreateHTTPCheckBadRequest)
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
		resp3, ok := iResp3.(*http_check.CreateHTTPCheckConflict)
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