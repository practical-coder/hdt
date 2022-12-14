// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/spf13/cobra"
)

// Schema cli for SmtpchkParams

// register flags to command
func registerModelSmtpchkParamsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSmtpchkParamsDomain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSmtpchkParamsHello(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSmtpchkParamsDomain(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	domainDescription := ``

	var domainFlagName string
	if cmdPrefix == "" {
		domainFlagName = "domain"
	} else {
		domainFlagName = fmt.Sprintf("%v.domain", cmdPrefix)
	}

	var domainFlagDefault string

	_ = cmd.PersistentFlags().String(domainFlagName, domainFlagDefault, domainDescription)

	return nil
}

func registerSmtpchkParamsHello(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	helloDescription := ``

	var helloFlagName string
	if cmdPrefix == "" {
		helloFlagName = "hello"
	} else {
		helloFlagName = fmt.Sprintf("%v.hello", cmdPrefix)
	}

	var helloFlagDefault string

	_ = cmd.PersistentFlags().String(helloFlagName, helloFlagDefault, helloDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSmtpchkParamsFlags(depth int, m *models.SmtpchkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, domainAdded := retrieveSmtpchkParamsDomainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || domainAdded

	err, helloAdded := retrieveSmtpchkParamsHelloFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || helloAdded

	return nil, retAdded
}

func retrieveSmtpchkParamsDomainFlags(depth int, m *models.SmtpchkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	domainFlagName := fmt.Sprintf("%v.domain", cmdPrefix)
	if cmd.Flags().Changed(domainFlagName) {

		var domainFlagName string
		if cmdPrefix == "" {
			domainFlagName = "domain"
		} else {
			domainFlagName = fmt.Sprintf("%v.domain", cmdPrefix)
		}

		domainFlagValue, err := cmd.Flags().GetString(domainFlagName)
		if err != nil {
			return err, false
		}
		m.Domain = domainFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSmtpchkParamsHelloFlags(depth int, m *models.SmtpchkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	helloFlagName := fmt.Sprintf("%v.hello", cmdPrefix)
	if cmd.Flags().Changed(helloFlagName) {

		var helloFlagName string
		if cmdPrefix == "" {
			helloFlagName = "hello"
		} else {
			helloFlagName = fmt.Sprintf("%v.hello", cmdPrefix)
		}

		helloFlagValue, err := cmd.Flags().GetString(helloFlagName)
		if err != nil {
			return err, false
		}
		m.Hello = helloFlagValue

		retAdded = true
	}

	return nil, retAdded
}
