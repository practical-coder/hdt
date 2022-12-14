// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/spoe"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSpoeGetOneSpoeFileCmd returns a cmd to handle operation getOneSpoeFile
func makeOperationSpoeGetOneSpoeFileCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getOneSpoeFile",
		Short: `Returns one SPOE file.`,
		RunE:  runOperationSpoeGetOneSpoeFile,
	}

	if err := registerOperationSpoeGetOneSpoeFileParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSpoeGetOneSpoeFile uses cmd flags to call endpoint api
func runOperationSpoeGetOneSpoeFile(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := spoe.NewGetOneSpoeFileParams()
	if err, _ := retrieveOperationSpoeGetOneSpoeFileNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSpoeGetOneSpoeFileResult(appCli.Spoe.GetOneSpoeFile(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSpoeGetOneSpoeFileParamFlags registers all flags needed to fill params
func registerOperationSpoeGetOneSpoeFileParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSpoeGetOneSpoeFileNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSpoeGetOneSpoeFileNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. SPOE file name`

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

func retrieveOperationSpoeGetOneSpoeFileNameFlag(m *spoe.GetOneSpoeFileParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSpoeGetOneSpoeFileResult parses request result and return the string content
func parseOperationSpoeGetOneSpoeFileResult(resp0 *spoe.GetOneSpoeFileOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*spoe.GetOneSpoeFileDefault)
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
		resp0, ok := iResp0.(*spoe.GetOneSpoeFileOK)
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
		resp1, ok := iResp1.(*spoe.GetOneSpoeFileNotFound)
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
func registerModelGetOneSpoeFileOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetOneSpoeFileOKBodyVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGetOneSpoeFileOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetOneSpoeFileOKBodyVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	versionDescription := ``

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "_version"
	} else {
		versionFlagName = fmt.Sprintf("%v._version", cmdPrefix)
	}

	var versionFlagDefault int64

	_ = cmd.PersistentFlags().Int64(versionFlagName, versionFlagDefault, versionDescription)

	return nil
}

func registerGetOneSpoeFileOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	dataDescription := ``

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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetOneSpoeFileOKBodyFlags(depth int, m *spoe.GetOneSpoeFileOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, versionAdded := retrieveGetOneSpoeFileOKBodyVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	err, dataAdded := retrieveGetOneSpoeFileOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetOneSpoeFileOKBodyVersionFlags(depth int, m *spoe.GetOneSpoeFileOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	versionFlagName := fmt.Sprintf("%v._version", cmdPrefix)
	if cmd.Flags().Changed(versionFlagName) {

		var versionFlagName string
		if cmdPrefix == "" {
			versionFlagName = "_version"
		} else {
			versionFlagName = fmt.Sprintf("%v._version", cmdPrefix)
		}

		versionFlagValue, err := cmd.Flags().GetInt64(versionFlagName)
		if err != nil {
			return err, false
		}
		m.Version = versionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveGetOneSpoeFileOKBodyDataFlags(depth int, m *spoe.GetOneSpoeFileOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {

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

		retAdded = true
	}

	return nil, retAdded
}
