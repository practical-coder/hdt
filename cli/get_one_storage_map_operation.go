// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/storage"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationStorageGetOneStorageMapCmd returns a cmd to handle operation getOneStorageMap
func makeOperationStorageGetOneStorageMapCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getOneStorageMap",
		Short: `Returns the contents of one managed map file from disk`,
		RunE:  runOperationStorageGetOneStorageMap,
	}

	if err := registerOperationStorageGetOneStorageMapParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationStorageGetOneStorageMap uses cmd flags to call endpoint api
func runOperationStorageGetOneStorageMap(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := storage.NewGetOneStorageMapParams()
	if err, _ := retrieveOperationStorageGetOneStorageMapNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationStorageGetOneStorageMapResult(appCli.Storage.GetOneStorageMap(params, &bytes.Buffer{}, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationStorageGetOneStorageMapParamFlags registers all flags needed to fill params
func registerOperationStorageGetOneStorageMapParamFlags(cmd *cobra.Command) error {
	if err := registerOperationStorageGetOneStorageMapNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationStorageGetOneStorageMapNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Map file storage_name`

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

func retrieveOperationStorageGetOneStorageMapNameFlag(m *storage.GetOneStorageMapParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationStorageGetOneStorageMapResult parses request result and return the string content
func parseOperationStorageGetOneStorageMapResult(resp0 *storage.GetOneStorageMapOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*storage.GetOneStorageMapDefault)
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
		resp0, ok := iResp0.(*storage.GetOneStorageMapOK)
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
		resp1, ok := iResp1.(*storage.GetOneStorageMapNotFound)
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
		msgStr := fmt.Sprintf("%v", resp0.Payload)
		return string(msgStr), nil
	}

	return "", nil
}