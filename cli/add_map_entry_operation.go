// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/practical-coder/hdt/client/maps"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationMapsAddMapEntryCmd returns a cmd to handle operation addMapEntry
func makeOperationMapsAddMapEntryCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "addMapEntry",
		Short: `Adds an entry into the map file.`,
		RunE:  runOperationMapsAddMapEntry,
	}

	if err := registerOperationMapsAddMapEntryParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationMapsAddMapEntry uses cmd flags to call endpoint api
func runOperationMapsAddMapEntry(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := maps.NewAddMapEntryParams()
	if err, _ := retrieveOperationMapsAddMapEntryDataFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationMapsAddMapEntryForceSyncFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationMapsAddMapEntryMapFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationMapsAddMapEntryResult(appCli.Maps.AddMapEntry(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationMapsAddMapEntryParamFlags registers all flags needed to fill params
func registerOperationMapsAddMapEntryParamFlags(cmd *cobra.Command) error {
	if err := registerOperationMapsAddMapEntryDataParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationMapsAddMapEntryForceSyncParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationMapsAddMapEntryMapParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationMapsAddMapEntryDataParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(dataFlagName, "", "Optional json string for [data]. ")

	// add flags for body
	if err := registerModelMapEntryFlags(0, "mapEntry", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationMapsAddMapEntryForceSyncParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	forceSyncDescription := `If true, immediately syncs changes to disk`

	var forceSyncFlagName string
	if cmdPrefix == "" {
		forceSyncFlagName = "force_sync"
	} else {
		forceSyncFlagName = fmt.Sprintf("%v.force_sync", cmdPrefix)
	}

	var forceSyncFlagDefault bool

	_ = cmd.PersistentFlags().Bool(forceSyncFlagName, forceSyncFlagDefault, forceSyncDescription)

	return nil
}
func registerOperationMapsAddMapEntryMapParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	mapDescription := `Required. Mapfile attribute storage_name`

	var mapFlagName string
	if cmdPrefix == "" {
		mapFlagName = "map"
	} else {
		mapFlagName = fmt.Sprintf("%v.map", cmdPrefix)
	}

	var mapFlagDefault string

	_ = cmd.PersistentFlags().String(mapFlagName, mapFlagDefault, mapDescription)

	return nil
}

func retrieveOperationMapsAddMapEntryDataFlag(m *maps.AddMapEntryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("data") {
		// Read data string from cmd and unmarshal
		dataValueStr, err := cmd.Flags().GetString("data")
		if err != nil {
			return err, false
		}

		dataValue := models.MapEntry{}
		if err := json.Unmarshal([]byte(dataValueStr), &dataValue); err != nil {
			return fmt.Errorf("cannot unmarshal data string in models.MapEntry: %v", err), false
		}
		m.Data = &dataValue
	}
	dataValueModel := m.Data
	if swag.IsZero(dataValueModel) {
		dataValueModel = &models.MapEntry{}
	}
	err, added := retrieveModelMapEntryFlags(0, dataValueModel, "mapEntry", cmd)
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
func retrieveOperationMapsAddMapEntryForceSyncFlag(m *maps.AddMapEntryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("force_sync") {

		var forceSyncFlagName string
		if cmdPrefix == "" {
			forceSyncFlagName = "force_sync"
		} else {
			forceSyncFlagName = fmt.Sprintf("%v.force_sync", cmdPrefix)
		}

		forceSyncFlagValue, err := cmd.Flags().GetBool(forceSyncFlagName)
		if err != nil {
			return err, false
		}
		m.ForceSync = &forceSyncFlagValue

	}
	return nil, retAdded
}
func retrieveOperationMapsAddMapEntryMapFlag(m *maps.AddMapEntryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("map") {

		var mapFlagName string
		if cmdPrefix == "" {
			mapFlagName = "map"
		} else {
			mapFlagName = fmt.Sprintf("%v.map", cmdPrefix)
		}

		mapFlagValue, err := cmd.Flags().GetString(mapFlagName)
		if err != nil {
			return err, false
		}
		m.Map = mapFlagValue

	}
	return nil, retAdded
}

// parseOperationMapsAddMapEntryResult parses request result and return the string content
func parseOperationMapsAddMapEntryResult(resp0 *maps.AddMapEntryCreated, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*maps.AddMapEntryDefault)
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
		resp0, ok := iResp0.(*maps.AddMapEntryCreated)
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
		resp1, ok := iResp1.(*maps.AddMapEntryBadRequest)
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
		resp2, ok := iResp2.(*maps.AddMapEntryConflict)
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