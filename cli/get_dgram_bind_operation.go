// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/practical-coder/hdt/client/dgram_bind"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDgramBindGetDgramBindCmd returns a cmd to handle operation getDgramBind
func makeOperationDgramBindGetDgramBindCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getDgramBind",
		Short: `Returns one dgram bind configuration by it's name in the specified log forward.`,
		RunE:  runOperationDgramBindGetDgramBind,
	}

	if err := registerOperationDgramBindGetDgramBindParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDgramBindGetDgramBind uses cmd flags to call endpoint api
func runOperationDgramBindGetDgramBind(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := dgram_bind.NewGetDgramBindParams()
	if err, _ := retrieveOperationDgramBindGetDgramBindLogForwardFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDgramBindGetDgramBindNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDgramBindGetDgramBindTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDgramBindGetDgramBindResult(appCli.DgramBind.GetDgramBind(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDgramBindGetDgramBindParamFlags registers all flags needed to fill params
func registerOperationDgramBindGetDgramBindParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDgramBindGetDgramBindLogForwardParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDgramBindGetDgramBindNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDgramBindGetDgramBindTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDgramBindGetDgramBindLogForwardParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDgramBindGetDgramBindNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDgramBindGetDgramBindTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationDgramBindGetDgramBindLogForwardFlag(m *dgram_bind.GetDgramBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDgramBindGetDgramBindNameFlag(m *dgram_bind.GetDgramBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDgramBindGetDgramBindTransactionIDFlag(m *dgram_bind.GetDgramBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDgramBindGetDgramBindResult parses request result and return the string content
func parseOperationDgramBindGetDgramBindResult(resp0 *dgram_bind.GetDgramBindOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*dgram_bind.GetDgramBindDefault)
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
		resp0, ok := iResp0.(*dgram_bind.GetDgramBindOK)
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
		resp1, ok := iResp1.(*dgram_bind.GetDgramBindNotFound)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
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
func registerModelGetDgramBindOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetDgramBindOKBodyVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGetDgramBindOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetDgramBindOKBodyVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerGetDgramBindOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelDgramBindFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetDgramBindOKBodyFlags(depth int, m *dgram_bind.GetDgramBindOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, versionAdded := retrieveGetDgramBindOKBodyVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	err, dataAdded := retrieveGetDgramBindOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetDgramBindOKBodyVersionFlags(depth int, m *dgram_bind.GetDgramBindOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveGetDgramBindOKBodyDataFlags(depth int, m *dgram_bind.GetDgramBindOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data models.DgramBind is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.DgramBind{}
	}

	err, dataAdded := retrieveModelDgramBindFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}
