// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdt/client/cluster"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationClusterInitiateCertificateRefreshCmd returns a cmd to handle operation initiateCertificateRefresh
func makeOperationClusterInitiateCertificateRefreshCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "initiateCertificateRefresh",
		Short: `Initiates a certificate refresh`,
		RunE:  runOperationClusterInitiateCertificateRefresh,
	}

	if err := registerOperationClusterInitiateCertificateRefreshParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationClusterInitiateCertificateRefresh uses cmd flags to call endpoint api
func runOperationClusterInitiateCertificateRefresh(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := cluster.NewInitiateCertificateRefreshParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationClusterInitiateCertificateRefreshResult(appCli.Cluster.InitiateCertificateRefresh(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationClusterInitiateCertificateRefreshParamFlags registers all flags needed to fill params
func registerOperationClusterInitiateCertificateRefreshParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationClusterInitiateCertificateRefreshResult parses request result and return the string content
func parseOperationClusterInitiateCertificateRefreshResult(resp0 *cluster.InitiateCertificateRefreshOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*cluster.InitiateCertificateRefreshDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning initiateCertificateRefreshOK is not supported

		// Non schema case: warning initiateCertificateRefreshForbidden is not supported

		return "", respErr
	}

	// warning: non schema response initiateCertificateRefreshOK is not supported by go-swagger cli yet.

	return "", nil
}
