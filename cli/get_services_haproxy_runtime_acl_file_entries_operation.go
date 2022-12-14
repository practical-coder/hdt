// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/acl_runtime"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntriesCmd returns a cmd to handle operation getServicesHaproxyRuntimeAclFileEntries
func makeOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntriesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetServicesHaproxyRuntimeACLFileEntries",
		Short: `Returns an ACL runtime setting using the runtime socket.`,
		RunE:  runOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntries,
	}

	if err := registerOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntriesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntries uses cmd flags to call endpoint api
func runOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntries(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := acl_runtime.NewGetServicesHaproxyRuntimeACLFileEntriesParams()
	if err, _ := retrieveOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntriesACLIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntriesResult(appCli.ACLRuntime.GetServicesHaproxyRuntimeACLFileEntries(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntriesParamFlags registers all flags needed to fill params
func registerOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntriesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntriesACLIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntriesACLIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	aclIdDescription := `Required. ACL ID`

	var aclIdFlagName string
	if cmdPrefix == "" {
		aclIdFlagName = "acl_id"
	} else {
		aclIdFlagName = fmt.Sprintf("%v.acl_id", cmdPrefix)
	}

	var aclIdFlagDefault string

	_ = cmd.PersistentFlags().String(aclIdFlagName, aclIdFlagDefault, aclIdDescription)

	return nil
}

func retrieveOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntriesACLIDFlag(m *acl_runtime.GetServicesHaproxyRuntimeACLFileEntriesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("acl_id") {

		var aclIdFlagName string
		if cmdPrefix == "" {
			aclIdFlagName = "acl_id"
		} else {
			aclIdFlagName = fmt.Sprintf("%v.acl_id", cmdPrefix)
		}

		aclIdFlagValue, err := cmd.Flags().GetString(aclIdFlagName)
		if err != nil {
			return err, false
		}
		m.ACLID = aclIdFlagValue

	}
	return nil, retAdded
}

// parseOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntriesResult parses request result and return the string content
func parseOperationACLRuntimeGetServicesHaproxyRuntimeACLFileEntriesResult(resp0 *acl_runtime.GetServicesHaproxyRuntimeACLFileEntriesOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*acl_runtime.GetServicesHaproxyRuntimeACLFileEntriesDefault)
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
		resp0, ok := iResp0.(*acl_runtime.GetServicesHaproxyRuntimeACLFileEntriesOK)
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
		resp1, ok := iResp1.(*acl_runtime.GetServicesHaproxyRuntimeACLFileEntriesBadRequest)
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
		resp2, ok := iResp2.(*acl_runtime.GetServicesHaproxyRuntimeACLFileEntriesNotFound)
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
