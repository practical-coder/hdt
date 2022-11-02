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

// makeOperationStickTableGetStickTablesCmd returns a cmd to handle operation getStickTables
func makeOperationStickTableGetStickTablesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getStickTables",
		Short: `Returns an array of all stick tables.`,
		RunE:  runOperationStickTableGetStickTables,
	}

	if err := registerOperationStickTableGetStickTablesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationStickTableGetStickTables uses cmd flags to call endpoint api
func runOperationStickTableGetStickTables(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := stick_table.NewGetStickTablesParams()
	if err, _ := retrieveOperationStickTableGetStickTablesProcessFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationStickTableGetStickTablesResult(appCli.StickTable.GetStickTables(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationStickTableGetStickTablesParamFlags registers all flags needed to fill params
func registerOperationStickTableGetStickTablesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationStickTableGetStickTablesProcessParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationStickTableGetStickTablesProcessParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	processDescription := `Process number if master-worker mode, if not all processes are returned`

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

func retrieveOperationStickTableGetStickTablesProcessFlag(m *stick_table.GetStickTablesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Process = &processFlagValue

	}
	return nil, retAdded
}

// parseOperationStickTableGetStickTablesResult parses request result and return the string content
func parseOperationStickTableGetStickTablesResult(resp0 *stick_table.GetStickTablesOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*stick_table.GetStickTablesDefault)
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
		resp0, ok := iResp0.(*stick_table.GetStickTablesOK)
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
