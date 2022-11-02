// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/cluster"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationClusterGetClusterCmd returns a cmd to handle operation getCluster
func makeOperationClusterGetClusterCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getCluster",
		Short: `Returns cluster data`,
		RunE:  runOperationClusterGetCluster,
	}

	if err := registerOperationClusterGetClusterParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationClusterGetCluster uses cmd flags to call endpoint api
func runOperationClusterGetCluster(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := cluster.NewGetClusterParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationClusterGetClusterResult(appCli.Cluster.GetCluster(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationClusterGetClusterParamFlags registers all flags needed to fill params
func registerOperationClusterGetClusterParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationClusterGetClusterResult parses request result and return the string content
func parseOperationClusterGetClusterResult(resp0 *cluster.GetClusterOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*cluster.GetClusterDefault)
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
		resp0, ok := iResp0.(*cluster.GetClusterOK)
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
