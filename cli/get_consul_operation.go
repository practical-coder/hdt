// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/practical-coder/hdt/client/service_discovery"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServiceDiscoveryGetConsulCmd returns a cmd to handle operation getConsul
func makeOperationServiceDiscoveryGetConsulCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getConsul",
		Short: `Returns one Consul server configuration by it's id.`,
		RunE:  runOperationServiceDiscoveryGetConsul,
	}

	if err := registerOperationServiceDiscoveryGetConsulParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServiceDiscoveryGetConsul uses cmd flags to call endpoint api
func runOperationServiceDiscoveryGetConsul(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := service_discovery.NewGetConsulParams()
	if err, _ := retrieveOperationServiceDiscoveryGetConsulIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServiceDiscoveryGetConsulResult(appCli.ServiceDiscovery.GetConsul(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServiceDiscoveryGetConsulParamFlags registers all flags needed to fill params
func registerOperationServiceDiscoveryGetConsulParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServiceDiscoveryGetConsulIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServiceDiscoveryGetConsulIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. Consul server id`

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

func retrieveOperationServiceDiscoveryGetConsulIDFlag(m *service_discovery.GetConsulParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationServiceDiscoveryGetConsulResult parses request result and return the string content
func parseOperationServiceDiscoveryGetConsulResult(resp0 *service_discovery.GetConsulOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*service_discovery.GetConsulDefault)
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
		resp0, ok := iResp0.(*service_discovery.GetConsulOK)
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
		resp1, ok := iResp1.(*service_discovery.GetConsulNotFound)
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
func registerModelGetConsulOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetConsulOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetConsulOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelConsulFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetConsulOKBodyFlags(depth int, m *service_discovery.GetConsulOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveGetConsulOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetConsulOKBodyDataFlags(depth int, m *service_discovery.GetConsulOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data models.Consul is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.Consul{}
	}

	err, dataAdded := retrieveModelConsulFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}
