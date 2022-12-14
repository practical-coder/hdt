// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/http_after_response_rule"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleCmd returns a cmd to handle operation deleteHttpAfterResponseRule
func makeOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteHTTPAfterResponseRule",
		Short: `Deletes a HTTP After Response Rule configuration by it's index from the specified parent.`,
		RunE:  runOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRule,
	}

	if err := registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRule uses cmd flags to call endpoint api
func runOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRule(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := http_after_response_rule.NewDeleteHTTPAfterResponseRuleParams()
	if err, _ := retrieveOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleIndexFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleParentNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleParentTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleResult(appCli.HTTPAfterResponseRule.DeleteHTTPAfterResponseRule(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleParamFlags registers all flags needed to fill params
func registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleParamFlags(cmd *cobra.Command) error {
	if err := registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleIndexParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleParentNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleParentTypeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleIndexParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	indexDescription := `Required. HTTP After Response Rule Index`

	var indexFlagName string
	if cmdPrefix == "" {
		indexFlagName = "index"
	} else {
		indexFlagName = fmt.Sprintf("%v.index", cmdPrefix)
	}

	var indexFlagDefault int64

	_ = cmd.PersistentFlags().Int64(indexFlagName, indexFlagDefault, indexDescription)

	return nil
}
func registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleParentNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	parentNameDescription := `Required. Parent name`

	var parentNameFlagName string
	if cmdPrefix == "" {
		parentNameFlagName = "parent_name"
	} else {
		parentNameFlagName = fmt.Sprintf("%v.parent_name", cmdPrefix)
	}

	var parentNameFlagDefault string

	_ = cmd.PersistentFlags().String(parentNameFlagName, parentNameFlagDefault, parentNameDescription)

	return nil
}
func registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleParentTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	parentTypeDescription := `Enum: ["frontend","backend"]. Required. Parent type`

	var parentTypeFlagName string
	if cmdPrefix == "" {
		parentTypeFlagName = "parent_type"
	} else {
		parentTypeFlagName = fmt.Sprintf("%v.parent_type", cmdPrefix)
	}

	var parentTypeFlagDefault string

	_ = cmd.PersistentFlags().String(parentTypeFlagName, parentTypeFlagDefault, parentTypeDescription)

	if err := cmd.RegisterFlagCompletionFunc(parentTypeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["frontend","backend"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}
func registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	transactionIdDescription := `ID of the transaction where we want to add the operation. Cannot be used when version is specified.`

	var transactionIdFlagName string
	if cmdPrefix == "" {
		transactionIdFlagName = "transaction_id"
	} else {
		transactionIdFlagName = fmt.Sprintf("%v.transaction_id", cmdPrefix)
	}

	var transactionIdFlagDefault string

	_ = cmd.PersistentFlags().String(transactionIdFlagName, transactionIdFlagDefault, transactionIdDescription)

	return nil
}
func registerOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	versionDescription := `Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.`

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "version"
	} else {
		versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
	}

	var versionFlagDefault int64

	_ = cmd.PersistentFlags().Int64(versionFlagName, versionFlagDefault, versionDescription)

	return nil
}

func retrieveOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleForceReloadFlag(m *http_after_response_rule.DeleteHTTPAfterResponseRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleIndexFlag(m *http_after_response_rule.DeleteHTTPAfterResponseRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("index") {

		var indexFlagName string
		if cmdPrefix == "" {
			indexFlagName = "index"
		} else {
			indexFlagName = fmt.Sprintf("%v.index", cmdPrefix)
		}

		indexFlagValue, err := cmd.Flags().GetInt64(indexFlagName)
		if err != nil {
			return err, false
		}
		m.Index = indexFlagValue

	}
	return nil, retAdded
}
func retrieveOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleParentNameFlag(m *http_after_response_rule.DeleteHTTPAfterResponseRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("parent_name") {

		var parentNameFlagName string
		if cmdPrefix == "" {
			parentNameFlagName = "parent_name"
		} else {
			parentNameFlagName = fmt.Sprintf("%v.parent_name", cmdPrefix)
		}

		parentNameFlagValue, err := cmd.Flags().GetString(parentNameFlagName)
		if err != nil {
			return err, false
		}
		m.ParentName = parentNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleParentTypeFlag(m *http_after_response_rule.DeleteHTTPAfterResponseRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("parent_type") {

		var parentTypeFlagName string
		if cmdPrefix == "" {
			parentTypeFlagName = "parent_type"
		} else {
			parentTypeFlagName = fmt.Sprintf("%v.parent_type", cmdPrefix)
		}

		parentTypeFlagValue, err := cmd.Flags().GetString(parentTypeFlagName)
		if err != nil {
			return err, false
		}
		m.ParentType = parentTypeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleTransactionIDFlag(m *http_after_response_rule.DeleteHTTPAfterResponseRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("transaction_id") {

		var transactionIdFlagName string
		if cmdPrefix == "" {
			transactionIdFlagName = "transaction_id"
		} else {
			transactionIdFlagName = fmt.Sprintf("%v.transaction_id", cmdPrefix)
		}

		transactionIdFlagValue, err := cmd.Flags().GetString(transactionIdFlagName)
		if err != nil {
			return err, false
		}
		m.TransactionID = &transactionIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleVersionFlag(m *http_after_response_rule.DeleteHTTPAfterResponseRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("version") {

		var versionFlagName string
		if cmdPrefix == "" {
			versionFlagName = "version"
		} else {
			versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
		}

		versionFlagValue, err := cmd.Flags().GetInt64(versionFlagName)
		if err != nil {
			return err, false
		}
		m.Version = &versionFlagValue

	}
	return nil, retAdded
}

// parseOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleResult parses request result and return the string content
func parseOperationHTTPAfterResponseRuleDeleteHTTPAfterResponseRuleResult(resp0 *http_after_response_rule.DeleteHTTPAfterResponseRuleAccepted, resp1 *http_after_response_rule.DeleteHTTPAfterResponseRuleNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*http_after_response_rule.DeleteHTTPAfterResponseRuleDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning deleteHttpAfterResponseRuleAccepted is not supported

		// Non schema case: warning deleteHttpAfterResponseRuleNoContent is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*http_after_response_rule.DeleteHTTPAfterResponseRuleNotFound)
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

	// warning: non schema response deleteHttpAfterResponseRuleAccepted is not supported by go-swagger cli yet.

	// warning: non schema response deleteHttpAfterResponseRuleNoContent is not supported by go-swagger cli yet.

	return "", nil
}
