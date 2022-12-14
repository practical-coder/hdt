// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/storage"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationStorageReplaceStorageMapFileCmd returns a cmd to handle operation replaceStorageMapFile
func makeOperationStorageReplaceStorageMapFileCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "replaceStorageMapFile",
		Short: `Replaces the contents of a managed map file on disk`,
		RunE:  runOperationStorageReplaceStorageMapFile,
	}

	if err := registerOperationStorageReplaceStorageMapFileParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationStorageReplaceStorageMapFile uses cmd flags to call endpoint api
func runOperationStorageReplaceStorageMapFile(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := storage.NewReplaceStorageMapFileParams()
	if err, _ := retrieveOperationStorageReplaceStorageMapFileDataFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStorageReplaceStorageMapFileForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStorageReplaceStorageMapFileNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStorageReplaceStorageMapFileSkipReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationStorageReplaceStorageMapFileResult(appCli.Storage.ReplaceStorageMapFile(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationStorageReplaceStorageMapFileParamFlags registers all flags needed to fill params
func registerOperationStorageReplaceStorageMapFileParamFlags(cmd *cobra.Command) error {
	if err := registerOperationStorageReplaceStorageMapFileDataParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStorageReplaceStorageMapFileForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStorageReplaceStorageMapFileNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStorageReplaceStorageMapFileSkipReloadParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationStorageReplaceStorageMapFileDataParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	dataDescription := `Required. `

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	var dataFlagDefault string

	_ = cmd.PersistentFlags().String(dataFlagName, dataFlagDefault, dataDescription)

	return nil
}
func registerOperationStorageReplaceStorageMapFileForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	forceReloadDescription := `If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.`

	var forceReloadFlagName string
	if cmdPrefix == "" {
		forceReloadFlagName = "force_reload"
	} else {
		forceReloadFlagName = fmt.Sprintf("%v.force_reload", cmdPrefix)
	}

	var forceReloadFlagDefault bool

	_ = cmd.PersistentFlags().Bool(forceReloadFlagName, forceReloadFlagDefault, forceReloadDescription)

	return nil
}
func registerOperationStorageReplaceStorageMapFileNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationStorageReplaceStorageMapFileSkipReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	skipReloadDescription := `If set, no reload will be initiated after update`

	var skipReloadFlagName string
	if cmdPrefix == "" {
		skipReloadFlagName = "skip_reload"
	} else {
		skipReloadFlagName = fmt.Sprintf("%v.skip_reload", cmdPrefix)
	}

	var skipReloadFlagDefault bool

	_ = cmd.PersistentFlags().Bool(skipReloadFlagName, skipReloadFlagDefault, skipReloadDescription)

	return nil
}

func retrieveOperationStorageReplaceStorageMapFileDataFlag(m *storage.ReplaceStorageMapFileParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("data") {

		var dataFlagName string
		if cmdPrefix == "" {
			dataFlagName = "data"
		} else {
			dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
		}

		dataFlagValue, err := cmd.Flags().GetString(dataFlagName)
		if err != nil {
			return err, false
		}
		m.Data = dataFlagValue

	}
	return nil, retAdded
}
func retrieveOperationStorageReplaceStorageMapFileForceReloadFlag(m *storage.ReplaceStorageMapFileParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("force_reload") {

		var forceReloadFlagName string
		if cmdPrefix == "" {
			forceReloadFlagName = "force_reload"
		} else {
			forceReloadFlagName = fmt.Sprintf("%v.force_reload", cmdPrefix)
		}

		forceReloadFlagValue, err := cmd.Flags().GetBool(forceReloadFlagName)
		if err != nil {
			return err, false
		}
		m.ForceReload = &forceReloadFlagValue

	}
	return nil, retAdded
}
func retrieveOperationStorageReplaceStorageMapFileNameFlag(m *storage.ReplaceStorageMapFileParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationStorageReplaceStorageMapFileSkipReloadFlag(m *storage.ReplaceStorageMapFileParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("skip_reload") {

		var skipReloadFlagName string
		if cmdPrefix == "" {
			skipReloadFlagName = "skip_reload"
		} else {
			skipReloadFlagName = fmt.Sprintf("%v.skip_reload", cmdPrefix)
		}

		skipReloadFlagValue, err := cmd.Flags().GetBool(skipReloadFlagName)
		if err != nil {
			return err, false
		}
		m.SkipReload = &skipReloadFlagValue

	}
	return nil, retAdded
}

// parseOperationStorageReplaceStorageMapFileResult parses request result and return the string content
func parseOperationStorageReplaceStorageMapFileResult(resp0 *storage.ReplaceStorageMapFileAccepted, resp1 *storage.ReplaceStorageMapFileNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*storage.ReplaceStorageMapFileDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning replaceStorageMapFileAccepted is not supported

		// Non schema case: warning replaceStorageMapFileNoContent is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*storage.ReplaceStorageMapFileBadRequest)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*storage.ReplaceStorageMapFileNotFound)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response replaceStorageMapFileAccepted is not supported by go-swagger cli yet.

	// warning: non schema response replaceStorageMapFileNoContent is not supported by go-swagger cli yet.

	return "", nil
}
