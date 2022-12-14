// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/stick_table"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationStickTableGetStickTableCmd returns a cmd to handle operation getStickTable
func makeOperationStickTableGetStickTableCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getStickTable",
		Short: `Returns one stick table from runtime.`,
		RunE:  runOperationStickTableGetStickTable,
	}

	if err := registerOperationStickTableGetStickTableParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationStickTableGetStickTable uses cmd flags to call endpoint api
func runOperationStickTableGetStickTable(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := stick_table.NewGetStickTableParams()
	if err, _ := retrieveOperationStickTableGetStickTableNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStickTableGetStickTableProcessFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationStickTableGetStickTableResult(appCli.StickTable.GetStickTable(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationStickTableGetStickTableParamFlags registers all flags needed to fill params
func registerOperationStickTableGetStickTableParamFlags(cmd *cobra.Command) error {
	if err := registerOperationStickTableGetStickTableNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStickTableGetStickTableProcessParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationStickTableGetStickTableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Stick table name`

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
func registerOperationStickTableGetStickTableProcessParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	processDescription := `Required. Process number if master-worker mode, if not only first process is returned`

	var processFlagName string
	if cmdPrefix == "" {
		processFlagName = "process"
	} else {
		processFlagName = fmt.Sprintf("%v.process", cmdPrefix)
	}

	var processFlagDefault int64

	_ = cmd.PersistentFlags().Int64(processFlagName, processFlagDefault, processDescription)

	return nil
}

func retrieveOperationStickTableGetStickTableNameFlag(m *stick_table.GetStickTableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationStickTableGetStickTableProcessFlag(m *stick_table.GetStickTableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("process") {

		var processFlagName string
		if cmdPrefix == "" {
			processFlagName = "process"
		} else {
			processFlagName = fmt.Sprintf("%v.process", cmdPrefix)
		}

		processFlagValue, err := cmd.Flags().GetInt64(processFlagName)
		if err != nil {
			return err, false
		}
		m.Process = processFlagValue

	}
	return nil, retAdded
}

// parseOperationStickTableGetStickTableResult parses request result and return the string content
func parseOperationStickTableGetStickTableResult(resp0 *stick_table.GetStickTableOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*stick_table.GetStickTableDefault)
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
		resp0, ok := iResp0.(*stick_table.GetStickTableOK)
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
		resp1, ok := iResp1.(*stick_table.GetStickTableNotFound)
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
