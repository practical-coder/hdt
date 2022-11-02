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

// makeOperationServiceDiscoveryGetConsulsCmd returns a cmd to handle operation getConsuls
func makeOperationServiceDiscoveryGetConsulsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getConsuls",
		Short: `Returns all configured Consul servers.`,
		RunE:  runOperationServiceDiscoveryGetConsuls,
	}

	if err := registerOperationServiceDiscoveryGetConsulsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServiceDiscoveryGetConsuls uses cmd flags to call endpoint api
func runOperationServiceDiscoveryGetConsuls(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := service_discovery.NewGetConsulsParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServiceDiscoveryGetConsulsResult(appCli.ServiceDiscovery.GetConsuls(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServiceDiscoveryGetConsulsParamFlags registers all flags needed to fill params
func registerOperationServiceDiscoveryGetConsulsParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationServiceDiscoveryGetConsulsResult parses request result and return the string content
func parseOperationServiceDiscoveryGetConsulsResult(resp0 *service_discovery.GetConsulsOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*service_discovery.GetConsulsDefault)
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
		resp0, ok := iResp0.(*service_discovery.GetConsulsOK)
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
func registerModelGetConsulsOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetConsulsOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetConsulsOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data models.Consuls array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetConsulsOKBodyFlags(depth int, m *service_discovery.GetConsulsOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveGetConsulsOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetConsulsOKBodyDataFlags(depth int, m *service_discovery.GetConsulsOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type models.Consuls is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
