// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/service_discovery"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServiceDiscoveryDeleteAWSRegionCmd returns a cmd to handle operation deleteAWSRegion
func makeOperationServiceDiscoveryDeleteAWSRegionCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteAWSRegion",
		Short: `Delete an AWS region configuration by it's id.`,
		RunE:  runOperationServiceDiscoveryDeleteAWSRegion,
	}

	if err := registerOperationServiceDiscoveryDeleteAWSRegionParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServiceDiscoveryDeleteAWSRegion uses cmd flags to call endpoint api
func runOperationServiceDiscoveryDeleteAWSRegion(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := service_discovery.NewDeleteAWSRegionParams()
	if err, _ := retrieveOperationServiceDiscoveryDeleteAWSRegionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServiceDiscoveryDeleteAWSRegionResult(appCli.ServiceDiscovery.DeleteAWSRegion(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServiceDiscoveryDeleteAWSRegionParamFlags registers all flags needed to fill params
func registerOperationServiceDiscoveryDeleteAWSRegionParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServiceDiscoveryDeleteAWSRegionIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServiceDiscoveryDeleteAWSRegionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. AWS region ID`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func retrieveOperationServiceDiscoveryDeleteAWSRegionIDFlag(m *service_discovery.DeleteAWSRegionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}

// parseOperationServiceDiscoveryDeleteAWSRegionResult parses request result and return the string content
func parseOperationServiceDiscoveryDeleteAWSRegionResult(resp0 *service_discovery.DeleteAWSRegionNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*service_discovery.DeleteAWSRegionDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning deleteAWSRegionNoContent is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*service_discovery.DeleteAWSRegionNotFound)
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

	// warning: non schema response deleteAWSRegionNoContent is not supported by go-swagger cli yet.

	return "", nil
}
