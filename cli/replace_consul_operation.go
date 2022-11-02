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

// makeOperationServiceDiscoveryReplaceConsulCmd returns a cmd to handle operation replaceConsul
func makeOperationServiceDiscoveryReplaceConsulCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "replaceConsul",
		Short: `Replaces a Consul server configuration by it's id.`,
		RunE:  runOperationServiceDiscoveryReplaceConsul,
	}

	if err := registerOperationServiceDiscoveryReplaceConsulParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServiceDiscoveryReplaceConsul uses cmd flags to call endpoint api
func runOperationServiceDiscoveryReplaceConsul(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := service_discovery.NewReplaceConsulParams()
	if err, _ := retrieveOperationServiceDiscoveryReplaceConsulDataFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServiceDiscoveryReplaceConsulIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServiceDiscoveryReplaceConsulResult(appCli.ServiceDiscovery.ReplaceConsul(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServiceDiscoveryReplaceConsulParamFlags registers all flags needed to fill params
func registerOperationServiceDiscoveryReplaceConsulParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServiceDiscoveryReplaceConsulDataParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServiceDiscoveryReplaceConsulIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServiceDiscoveryReplaceConsulDataParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(dataFlagName, "", "Optional json string for [data]. ")

	// add flags for body
	if err := registerModelConsulFlags(0, "consul", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationServiceDiscoveryReplaceConsulIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. Consul Index`

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

func retrieveOperationServiceDiscoveryReplaceConsulDataFlag(m *service_discovery.ReplaceConsulParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("data") {
		// Read data string from cmd and unmarshal
		dataValueStr, err := cmd.Flags().GetString("data")
		if err != nil {
			return err, false
		}

		dataValue := models.Consul{}
		if err := json.Unmarshal([]byte(dataValueStr), &dataValue); err != nil {
			return fmt.Errorf("cannot unmarshal data string in models.Consul: %v", err), false
		}
		m.Data = &dataValue
	}
	dataValueModel := m.Data
	if swag.IsZero(dataValueModel) {
		dataValueModel = &models.Consul{}
	}
	err, added := retrieveModelConsulFlags(0, dataValueModel, "consul", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Data = dataValueModel
	}
	if dryRun && debug {

		dataValueDebugBytes, err := json.Marshal(m.Data)
		if err != nil {
			return err, false
		}
		logDebugf("Data dry-run payload: %v", string(dataValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationServiceDiscoveryReplaceConsulIDFlag(m *service_discovery.ReplaceConsulParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationServiceDiscoveryReplaceConsulResult parses request result and return the string content
func parseOperationServiceDiscoveryReplaceConsulResult(resp0 *service_discovery.ReplaceConsulOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*service_discovery.ReplaceConsulDefault)
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
		resp0, ok := iResp0.(*service_discovery.ReplaceConsulOK)
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
		resp1, ok := iResp1.(*service_discovery.ReplaceConsulBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*service_discovery.ReplaceConsulNotFound)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
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