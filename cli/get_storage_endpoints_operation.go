// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/discovery"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDiscoveryGetStorageEndpointsCmd returns a cmd to handle operation getStorageEndpoints
func makeOperationDiscoveryGetStorageEndpointsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getStorageEndpoints",
		Short: `Returns a list of endpoints that use HAProxy storage for persistency, e.g. maps, ssl certificates...`,
		RunE:  runOperationDiscoveryGetStorageEndpoints,
	}

	if err := registerOperationDiscoveryGetStorageEndpointsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDiscoveryGetStorageEndpoints uses cmd flags to call endpoint api
func runOperationDiscoveryGetStorageEndpoints(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := discovery.NewGetStorageEndpointsParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDiscoveryGetStorageEndpointsResult(appCli.Discovery.GetStorageEndpoints(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDiscoveryGetStorageEndpointsParamFlags registers all flags needed to fill params
func registerOperationDiscoveryGetStorageEndpointsParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationDiscoveryGetStorageEndpointsResult parses request result and return the string content
func parseOperationDiscoveryGetStorageEndpointsResult(resp0 *discovery.GetStorageEndpointsOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*discovery.GetStorageEndpointsDefault)
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
		resp0, ok := iResp0.(*discovery.GetStorageEndpointsOK)
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