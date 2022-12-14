// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/maps"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationMapsDeleteRuntimeMapEntryCmd returns a cmd to handle operation deleteRuntimeMapEntry
func makeOperationMapsDeleteRuntimeMapEntryCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteRuntimeMapEntry",
		Short: `Delete all the map entries from the map by its id.`,
		RunE:  runOperationMapsDeleteRuntimeMapEntry,
	}

	if err := registerOperationMapsDeleteRuntimeMapEntryParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationMapsDeleteRuntimeMapEntry uses cmd flags to call endpoint api
func runOperationMapsDeleteRuntimeMapEntry(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := maps.NewDeleteRuntimeMapEntryParams()
	if err, _ := retrieveOperationMapsDeleteRuntimeMapEntryForceSyncFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationMapsDeleteRuntimeMapEntryIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationMapsDeleteRuntimeMapEntryMapFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationMapsDeleteRuntimeMapEntryResult(appCli.Maps.DeleteRuntimeMapEntry(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationMapsDeleteRuntimeMapEntryParamFlags registers all flags needed to fill params
func registerOperationMapsDeleteRuntimeMapEntryParamFlags(cmd *cobra.Command) error {
	if err := registerOperationMapsDeleteRuntimeMapEntryForceSyncParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationMapsDeleteRuntimeMapEntryIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationMapsDeleteRuntimeMapEntryMapParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationMapsDeleteRuntimeMapEntryForceSyncParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationMapsDeleteRuntimeMapEntryIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. Map id`

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
func registerOperationMapsDeleteRuntimeMapEntryMapParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationMapsDeleteRuntimeMapEntryForceSyncFlag(m *maps.DeleteRuntimeMapEntryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationMapsDeleteRuntimeMapEntryIDFlag(m *maps.DeleteRuntimeMapEntryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationMapsDeleteRuntimeMapEntryMapFlag(m *maps.DeleteRuntimeMapEntryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationMapsDeleteRuntimeMapEntryResult parses request result and return the string content
func parseOperationMapsDeleteRuntimeMapEntryResult(resp0 *maps.DeleteRuntimeMapEntryNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*maps.DeleteRuntimeMapEntryDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning deleteRuntimeMapEntryNoContent is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*maps.DeleteRuntimeMapEntryNotFound)
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

	// warning: non schema response deleteRuntimeMapEntryNoContent is not supported by go-swagger cli yet.

	return "", nil
}
