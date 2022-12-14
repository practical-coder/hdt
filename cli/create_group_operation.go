// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/practical-coder/hdt/client/group"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationGroupCreateGroupCmd returns a cmd to handle operation createGroup
func makeOperationGroupCreateGroupCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "createGroup",
		Short: ``,
		RunE:  runOperationGroupCreateGroup,
	}

	if err := registerOperationGroupCreateGroupParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationGroupCreateGroup uses cmd flags to call endpoint api
func runOperationGroupCreateGroup(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := group.NewCreateGroupParams()
	if err, _ := retrieveOperationGroupCreateGroupDataFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationGroupCreateGroupForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationGroupCreateGroupTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationGroupCreateGroupUserlistFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationGroupCreateGroupVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationGroupCreateGroupResult(appCli.Group.CreateGroup(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationGroupCreateGroupParamFlags registers all flags needed to fill params
func registerOperationGroupCreateGroupParamFlags(cmd *cobra.Command) error {
	if err := registerOperationGroupCreateGroupDataParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationGroupCreateGroupForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationGroupCreateGroupTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationGroupCreateGroupUserlistParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationGroupCreateGroupVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationGroupCreateGroupDataParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(dataFlagName, "", "Optional json string for [data]. ")

	// add flags for body
	if err := registerModelGroupFlags(0, "group", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationGroupCreateGroupForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationGroupCreateGroupTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationGroupCreateGroupUserlistParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	userlistDescription := `Required. Parent userlist name`

	var userlistFlagName string
	if cmdPrefix == "" {
		userlistFlagName = "userlist"
	} else {
		userlistFlagName = fmt.Sprintf("%v.userlist", cmdPrefix)
	}

	var userlistFlagDefault string

	_ = cmd.PersistentFlags().String(userlistFlagName, userlistFlagDefault, userlistDescription)

	return nil
}
func registerOperationGroupCreateGroupVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationGroupCreateGroupDataFlag(m *group.CreateGroupParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("data") {
		// Read data string from cmd and unmarshal
		dataValueStr, err := cmd.Flags().GetString("data")
		if err != nil {
			return err, false
		}

		dataValue := models.Group{}
		if err := json.Unmarshal([]byte(dataValueStr), &dataValue); err != nil {
			return fmt.Errorf("cannot unmarshal data string in models.Group: %v", err), false
		}
		m.Data = &dataValue
	}
	dataValueModel := m.Data
	if swag.IsZero(dataValueModel) {
		dataValueModel = &models.Group{}
	}
	err, added := retrieveModelGroupFlags(0, dataValueModel, "group", cmd)
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
func retrieveOperationGroupCreateGroupForceReloadFlag(m *group.CreateGroupParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationGroupCreateGroupTransactionIDFlag(m *group.CreateGroupParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationGroupCreateGroupUserlistFlag(m *group.CreateGroupParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("userlist") {

		var userlistFlagName string
		if cmdPrefix == "" {
			userlistFlagName = "userlist"
		} else {
			userlistFlagName = fmt.Sprintf("%v.userlist", cmdPrefix)
		}

		userlistFlagValue, err := cmd.Flags().GetString(userlistFlagName)
		if err != nil {
			return err, false
		}
		m.Userlist = userlistFlagValue

	}
	return nil, retAdded
}
func retrieveOperationGroupCreateGroupVersionFlag(m *group.CreateGroupParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationGroupCreateGroupResult parses request result and return the string content
func parseOperationGroupCreateGroupResult(resp0 *group.CreateGroupCreated, resp1 *group.CreateGroupAccepted, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*group.CreateGroupDefault)
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
		resp0, ok := iResp0.(*group.CreateGroupCreated)
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
		resp1, ok := iResp1.(*group.CreateGroupAccepted)
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
		resp2, ok := iResp2.(*group.CreateGroupBadRequest)
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
		resp3, ok := iResp3.(*group.CreateGroupConflict)
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
