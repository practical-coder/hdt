// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/practical-coder/hdt/client/acl"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationACLCreateACLCmd returns a cmd to handle operation createAcl
func makeOperationACLCreateACLCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "createAcl",
		Short: `Adds a new ACL line of the specified type in the specified parent.`,
		RunE:  runOperationACLCreateACL,
	}

	if err := registerOperationACLCreateACLParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationACLCreateACL uses cmd flags to call endpoint api
func runOperationACLCreateACL(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := acl.NewCreateACLParams()
	if err, _ := retrieveOperationACLCreateACLDataFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationACLCreateACLForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationACLCreateACLParentNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationACLCreateACLParentTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationACLCreateACLTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationACLCreateACLVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationACLCreateACLResult(appCli.ACL.CreateACL(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationACLCreateACLParamFlags registers all flags needed to fill params
func registerOperationACLCreateACLParamFlags(cmd *cobra.Command) error {
	if err := registerOperationACLCreateACLDataParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationACLCreateACLForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationACLCreateACLParentNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationACLCreateACLParentTypeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationACLCreateACLTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationACLCreateACLVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationACLCreateACLDataParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(dataFlagName, "", "Optional json string for [data]. ")

	// add flags for body
	if err := registerModelACLFlags(0, "acl", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationACLCreateACLForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationACLCreateACLParentNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	parentNameDescription := `Required. Parent name`

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
func registerOperationACLCreateACLParentTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	parentTypeDescription := `Enum: ["frontend","backend"]. Required. Parent type`

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
			if err := json.Unmarshal([]byte(`["frontend","backend"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}
func registerOperationACLCreateACLTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationACLCreateACLVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationACLCreateACLDataFlag(m *acl.CreateACLParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("data") {
		// Read data string from cmd and unmarshal
		dataValueStr, err := cmd.Flags().GetString("data")
		if err != nil {
			return err, false
		}

		dataValue := models.ACL{}
		if err := json.Unmarshal([]byte(dataValueStr), &dataValue); err != nil {
			return fmt.Errorf("cannot unmarshal data string in models.ACL: %v", err), false
		}
		m.Data = &dataValue
	}
	dataValueModel := m.Data
	if swag.IsZero(dataValueModel) {
		dataValueModel = &models.ACL{}
	}
	err, added := retrieveModelACLFlags(0, dataValueModel, "acl", cmd)
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
func retrieveOperationACLCreateACLForceReloadFlag(m *acl.CreateACLParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationACLCreateACLParentNameFlag(m *acl.CreateACLParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.ParentName = parentNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationACLCreateACLParentTypeFlag(m *acl.CreateACLParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationACLCreateACLTransactionIDFlag(m *acl.CreateACLParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationACLCreateACLVersionFlag(m *acl.CreateACLParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationACLCreateACLResult parses request result and return the string content
func parseOperationACLCreateACLResult(resp0 *acl.CreateACLCreated, resp1 *acl.CreateACLAccepted, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*acl.CreateACLDefault)
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
		resp0, ok := iResp0.(*acl.CreateACLCreated)
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
		resp1, ok := iResp1.(*acl.CreateACLAccepted)
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
		resp2, ok := iResp2.(*acl.CreateACLBadRequest)
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
		resp3, ok := iResp3.(*acl.CreateACLConflict)
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
