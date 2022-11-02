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

// makeOperationStorageReplaceStorageGeneralFileCmd returns a cmd to handle operation replaceStorageGeneralFile
func makeOperationStorageReplaceStorageGeneralFileCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "replaceStorageGeneralFile",
		Short: `Replaces the contents of a managed general use file on disk`,
		RunE:  runOperationStorageReplaceStorageGeneralFile,
	}

	if err := registerOperationStorageReplaceStorageGeneralFileParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationStorageReplaceStorageGeneralFile uses cmd flags to call endpoint api
func runOperationStorageReplaceStorageGeneralFile(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := storage.NewReplaceStorageGeneralFileParams()
	if err, _ := retrieveOperationStorageReplaceStorageGeneralFileDataFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStorageReplaceStorageGeneralFileForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStorageReplaceStorageGeneralFileNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStorageReplaceStorageGeneralFileSkipReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationStorageReplaceStorageGeneralFileResult(appCli.Storage.ReplaceStorageGeneralFile(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationStorageReplaceStorageGeneralFileParamFlags registers all flags needed to fill params
func registerOperationStorageReplaceStorageGeneralFileParamFlags(cmd *cobra.Command) error {
	if err := registerOperationStorageReplaceStorageGeneralFileDataParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStorageReplaceStorageGeneralFileForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStorageReplaceStorageGeneralFileNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStorageReplaceStorageGeneralFileSkipReloadParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationStorageReplaceStorageGeneralFileDataParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationStorageReplaceStorageGeneralFileForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationStorageReplaceStorageGeneralFileNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. General use file storage_name`

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
func registerOperationStorageReplaceStorageGeneralFileSkipReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationStorageReplaceStorageGeneralFileDataFlag(m *storage.ReplaceStorageGeneralFileParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationStorageReplaceStorageGeneralFileForceReloadFlag(m *storage.ReplaceStorageGeneralFileParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationStorageReplaceStorageGeneralFileNameFlag(m *storage.ReplaceStorageGeneralFileParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationStorageReplaceStorageGeneralFileSkipReloadFlag(m *storage.ReplaceStorageGeneralFileParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationStorageReplaceStorageGeneralFileResult parses request result and return the string content
func parseOperationStorageReplaceStorageGeneralFileResult(resp0 *storage.ReplaceStorageGeneralFileAccepted, resp1 *storage.ReplaceStorageGeneralFileNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*storage.ReplaceStorageGeneralFileDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning replaceStorageGeneralFileAccepted is not supported

		// Non schema case: warning replaceStorageGeneralFileNoContent is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*storage.ReplaceStorageGeneralFileBadRequest)
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
		resp3, ok := iResp3.(*storage.ReplaceStorageGeneralFileNotFound)
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

	// warning: non schema response replaceStorageGeneralFileAccepted is not supported by go-swagger cli yet.

	// warning: non schema response replaceStorageGeneralFileNoContent is not supported by go-swagger cli yet.

	return "", nil
}
