// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/user"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationUserGetUsersCmd returns a cmd to handle operation getUsers
func makeOperationUserGetUsersCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getUsers",
		Short: ``,
		RunE:  runOperationUserGetUsers,
	}

	if err := registerOperationUserGetUsersParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUserGetUsers uses cmd flags to call endpoint api
func runOperationUserGetUsers(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := user.NewGetUsersParams()
	if err, _ := retrieveOperationUserGetUsersTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserGetUsersUserlistFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUserGetUsersResult(appCli.User.GetUsers(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationUserGetUsersParamFlags registers all flags needed to fill params
func registerOperationUserGetUsersParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUserGetUsersTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserGetUsersUserlistParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUserGetUsersTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationUserGetUsersUserlistParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationUserGetUsersTransactionIDFlag(m *user.GetUsersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationUserGetUsersUserlistFlag(m *user.GetUsersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationUserGetUsersResult parses request result and return the string content
func parseOperationUserGetUsersResult(resp0 *user.GetUsersOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*user.GetUsersDefault)
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
		resp0, ok := iResp0.(*user.GetUsersOK)
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
func registerModelGetUsersOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetUsersOKBodyVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGetUsersOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetUsersOKBodyVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerGetUsersOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data models.Users array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetUsersOKBodyFlags(depth int, m *user.GetUsersOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, versionAdded := retrieveGetUsersOKBodyVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	err, dataAdded := retrieveGetUsersOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetUsersOKBodyVersionFlags(depth int, m *user.GetUsersOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveGetUsersOKBodyDataFlags(depth int, m *user.GetUsersOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type models.Users is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
