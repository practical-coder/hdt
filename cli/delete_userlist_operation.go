// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/userlist"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationUserlistDeleteUserlistCmd returns a cmd to handle operation deleteUserlist
func makeOperationUserlistDeleteUserlistCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteUserlist",
		Short: `Deletes a userlist configuration by it's name.`,
		RunE:  runOperationUserlistDeleteUserlist,
	}

	if err := registerOperationUserlistDeleteUserlistParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUserlistDeleteUserlist uses cmd flags to call endpoint api
func runOperationUserlistDeleteUserlist(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := userlist.NewDeleteUserlistParams()
	if err, _ := retrieveOperationUserlistDeleteUserlistForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserlistDeleteUserlistNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserlistDeleteUserlistTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserlistDeleteUserlistVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUserlistDeleteUserlistResult(appCli.Userlist.DeleteUserlist(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationUserlistDeleteUserlistParamFlags registers all flags needed to fill params
func registerOperationUserlistDeleteUserlistParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUserlistDeleteUserlistForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserlistDeleteUserlistNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserlistDeleteUserlistTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserlistDeleteUserlistVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUserlistDeleteUserlistForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationUserlistDeleteUserlistNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Userlist name`

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
func registerOperationUserlistDeleteUserlistTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationUserlistDeleteUserlistVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationUserlistDeleteUserlistForceReloadFlag(m *userlist.DeleteUserlistParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationUserlistDeleteUserlistNameFlag(m *userlist.DeleteUserlistParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationUserlistDeleteUserlistTransactionIDFlag(m *userlist.DeleteUserlistParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationUserlistDeleteUserlistVersionFlag(m *userlist.DeleteUserlistParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationUserlistDeleteUserlistResult parses request result and return the string content
func parseOperationUserlistDeleteUserlistResult(resp0 *userlist.DeleteUserlistAccepted, resp1 *userlist.DeleteUserlistNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*userlist.DeleteUserlistDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning deleteUserlistAccepted is not supported

		// Non schema case: warning deleteUserlistNoContent is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*userlist.DeleteUserlistNotFound)
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

	// warning: non schema response deleteUserlistAccepted is not supported by go-swagger cli yet.

	// warning: non schema response deleteUserlistNoContent is not supported by go-swagger cli yet.

	return "", nil
}
