// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/peer"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationPeerDeletePeerCmd returns a cmd to handle operation deletePeer
func makeOperationPeerDeletePeerCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deletePeer",
		Short: `Deletes a peer from the configuration by it's name.`,
		RunE:  runOperationPeerDeletePeer,
	}

	if err := registerOperationPeerDeletePeerParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPeerDeletePeer uses cmd flags to call endpoint api
func runOperationPeerDeletePeer(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := peer.NewDeletePeerParams()
	if err, _ := retrieveOperationPeerDeletePeerForceReloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPeerDeletePeerNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPeerDeletePeerTransactionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPeerDeletePeerVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationPeerDeletePeerResult(appCli.Peer.DeletePeer(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationPeerDeletePeerParamFlags registers all flags needed to fill params
func registerOperationPeerDeletePeerParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPeerDeletePeerForceReloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPeerDeletePeerNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPeerDeletePeerTransactionIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPeerDeletePeerVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPeerDeletePeerForceReloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationPeerDeletePeerNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Peer name`

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
func registerOperationPeerDeletePeerTransactionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationPeerDeletePeerVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationPeerDeletePeerForceReloadFlag(m *peer.DeletePeerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationPeerDeletePeerNameFlag(m *peer.DeletePeerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationPeerDeletePeerTransactionIDFlag(m *peer.DeletePeerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationPeerDeletePeerVersionFlag(m *peer.DeletePeerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationPeerDeletePeerResult parses request result and return the string content
func parseOperationPeerDeletePeerResult(resp0 *peer.DeletePeerAccepted, resp1 *peer.DeletePeerNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*peer.DeletePeerDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning deletePeerAccepted is not supported

		// Non schema case: warning deletePeerNoContent is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*peer.DeletePeerNotFound)
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

	// warning: non schema response deletePeerAccepted is not supported by go-swagger cli yet.

	// warning: non schema response deletePeerNoContent is not supported by go-swagger cli yet.

	return "", nil
}