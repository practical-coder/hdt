// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/stats"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationStatsGetStatsCmd returns a cmd to handle operation getStats
func makeOperationStatsGetStatsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getStats",
		Short: `Getting stats from the HAProxy.`,
		RunE:  runOperationStatsGetStats,
	}

	if err := registerOperationStatsGetStatsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationStatsGetStats uses cmd flags to call endpoint api
func runOperationStatsGetStats(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := stats.NewGetStatsParams()
	if err, _ := retrieveOperationStatsGetStatsNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStatsGetStatsParentFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStatsGetStatsTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationStatsGetStatsResult(appCli.Stats.GetStats(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationStatsGetStatsParamFlags registers all flags needed to fill params
func registerOperationStatsGetStatsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationStatsGetStatsNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStatsGetStatsParentParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStatsGetStatsTypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationStatsGetStatsNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Object name to get stats for`

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
func registerOperationStatsGetStatsParentParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	parentDescription := `Object parent name to get stats for, in case the object is a server`

	var parentFlagName string
	if cmdPrefix == "" {
		parentFlagName = "parent"
	} else {
		parentFlagName = fmt.Sprintf("%v.parent", cmdPrefix)
	}

	var parentFlagDefault string

	_ = cmd.PersistentFlags().String(parentFlagName, parentFlagDefault, parentDescription)

	return nil
}
func registerOperationStatsGetStatsTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	typeDescription := `Enum: ["frontend","backend","server"]. Object type to get stats for (one of frontend, backend, server)`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["frontend","backend","server"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func retrieveOperationStatsGetStatsNameFlag(m *stats.GetStatsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = &nameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationStatsGetStatsParentFlag(m *stats.GetStatsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("parent") {

		var parentFlagName string
		if cmdPrefix == "" {
			parentFlagName = "parent"
		} else {
			parentFlagName = fmt.Sprintf("%v.parent", cmdPrefix)
		}

		parentFlagValue, err := cmd.Flags().GetString(parentFlagName)
		if err != nil {
			return err, false
		}
		m.Parent = &parentFlagValue

	}
	return nil, retAdded
}
func retrieveOperationStatsGetStatsTypeFlag(m *stats.GetStatsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("type") {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = &typeFlagValue

	}
	return nil, retAdded
}

// parseOperationStatsGetStatsResult parses request result and return the string content
func parseOperationStatsGetStatsResult(resp0 *stats.GetStatsOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*stats.GetStatsDefault)
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
		resp0, ok := iResp0.(*stats.GetStatsOK)
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
		resp1, ok := iResp1.(*stats.GetStatsInternalServerError)
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
