// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/information"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationInformationGetInfoCmd returns a cmd to handle operation getInfo
func makeOperationInformationGetInfoCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getInfo",
		Short: `Return API, hardware and OS information`,
		RunE:  runOperationInformationGetInfo,
	}

	if err := registerOperationInformationGetInfoParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationInformationGetInfo uses cmd flags to call endpoint api
func runOperationInformationGetInfo(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := information.NewGetInfoParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationInformationGetInfoResult(appCli.Information.GetInfo(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationInformationGetInfoParamFlags registers all flags needed to fill params
func registerOperationInformationGetInfoParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationInformationGetInfoResult parses request result and return the string content
func parseOperationInformationGetInfoResult(resp0 *information.GetInfoOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*information.GetInfoDefault)
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
		resp0, ok := iResp0.(*information.GetInfoOK)
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