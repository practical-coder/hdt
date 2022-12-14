// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/spf13/cobra"
)

// Schema cli for Resolver

// register flags to command
func registerModelResolverFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerResolverAcceptedPayloadSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerResolverHoldNx(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerResolverHoldObsolete(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerResolverHoldOther(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerResolverHoldRefused(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerResolverHoldTimeout(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerResolverHoldValid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerResolverName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerResolverParseResolvConf(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerResolverResolveRetries(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerResolverTimeoutResolve(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerResolverTimeoutRetry(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerResolverAcceptedPayloadSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	acceptedPayloadSizeDescription := ``

	var acceptedPayloadSizeFlagName string
	if cmdPrefix == "" {
		acceptedPayloadSizeFlagName = "accepted_payload_size"
	} else {
		acceptedPayloadSizeFlagName = fmt.Sprintf("%v.accepted_payload_size", cmdPrefix)
	}

	var acceptedPayloadSizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(acceptedPayloadSizeFlagName, acceptedPayloadSizeFlagDefault, acceptedPayloadSizeDescription)

	return nil
}

func registerResolverHoldNx(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	holdNxDescription := ``

	var holdNxFlagName string
	if cmdPrefix == "" {
		holdNxFlagName = "hold_nx"
	} else {
		holdNxFlagName = fmt.Sprintf("%v.hold_nx", cmdPrefix)
	}

	var holdNxFlagDefault int64

	_ = cmd.PersistentFlags().Int64(holdNxFlagName, holdNxFlagDefault, holdNxDescription)

	return nil
}

func registerResolverHoldObsolete(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	holdObsoleteDescription := ``

	var holdObsoleteFlagName string
	if cmdPrefix == "" {
		holdObsoleteFlagName = "hold_obsolete"
	} else {
		holdObsoleteFlagName = fmt.Sprintf("%v.hold_obsolete", cmdPrefix)
	}

	var holdObsoleteFlagDefault int64

	_ = cmd.PersistentFlags().Int64(holdObsoleteFlagName, holdObsoleteFlagDefault, holdObsoleteDescription)

	return nil
}

func registerResolverHoldOther(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	holdOtherDescription := ``

	var holdOtherFlagName string
	if cmdPrefix == "" {
		holdOtherFlagName = "hold_other"
	} else {
		holdOtherFlagName = fmt.Sprintf("%v.hold_other", cmdPrefix)
	}

	var holdOtherFlagDefault int64

	_ = cmd.PersistentFlags().Int64(holdOtherFlagName, holdOtherFlagDefault, holdOtherDescription)

	return nil
}

func registerResolverHoldRefused(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	holdRefusedDescription := ``

	var holdRefusedFlagName string
	if cmdPrefix == "" {
		holdRefusedFlagName = "hold_refused"
	} else {
		holdRefusedFlagName = fmt.Sprintf("%v.hold_refused", cmdPrefix)
	}

	var holdRefusedFlagDefault int64

	_ = cmd.PersistentFlags().Int64(holdRefusedFlagName, holdRefusedFlagDefault, holdRefusedDescription)

	return nil
}

func registerResolverHoldTimeout(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	holdTimeoutDescription := ``

	var holdTimeoutFlagName string
	if cmdPrefix == "" {
		holdTimeoutFlagName = "hold_timeout"
	} else {
		holdTimeoutFlagName = fmt.Sprintf("%v.hold_timeout", cmdPrefix)
	}

	var holdTimeoutFlagDefault int64

	_ = cmd.PersistentFlags().Int64(holdTimeoutFlagName, holdTimeoutFlagDefault, holdTimeoutDescription)

	return nil
}

func registerResolverHoldValid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	holdValidDescription := ``

	var holdValidFlagName string
	if cmdPrefix == "" {
		holdValidFlagName = "hold_valid"
	} else {
		holdValidFlagName = fmt.Sprintf("%v.hold_valid", cmdPrefix)
	}

	var holdValidFlagDefault int64

	_ = cmd.PersistentFlags().Int64(holdValidFlagName, holdValidFlagDefault, holdValidDescription)

	return nil
}

func registerResolverName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. `

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

func registerResolverParseResolvConf(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	parseResolvConfDescription := ``

	var parseResolvConfFlagName string
	if cmdPrefix == "" {
		parseResolvConfFlagName = "parse-resolv-conf"
	} else {
		parseResolvConfFlagName = fmt.Sprintf("%v.parse-resolv-conf", cmdPrefix)
	}

	var parseResolvConfFlagDefault bool

	_ = cmd.PersistentFlags().Bool(parseResolvConfFlagName, parseResolvConfFlagDefault, parseResolvConfDescription)

	return nil
}

func registerResolverResolveRetries(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	resolveRetriesDescription := ``

	var resolveRetriesFlagName string
	if cmdPrefix == "" {
		resolveRetriesFlagName = "resolve_retries"
	} else {
		resolveRetriesFlagName = fmt.Sprintf("%v.resolve_retries", cmdPrefix)
	}

	var resolveRetriesFlagDefault int64

	_ = cmd.PersistentFlags().Int64(resolveRetriesFlagName, resolveRetriesFlagDefault, resolveRetriesDescription)

	return nil
}

func registerResolverTimeoutResolve(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	timeoutResolveDescription := ``

	var timeoutResolveFlagName string
	if cmdPrefix == "" {
		timeoutResolveFlagName = "timeout_resolve"
	} else {
		timeoutResolveFlagName = fmt.Sprintf("%v.timeout_resolve", cmdPrefix)
	}

	var timeoutResolveFlagDefault int64

	_ = cmd.PersistentFlags().Int64(timeoutResolveFlagName, timeoutResolveFlagDefault, timeoutResolveDescription)

	return nil
}

func registerResolverTimeoutRetry(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	timeoutRetryDescription := ``

	var timeoutRetryFlagName string
	if cmdPrefix == "" {
		timeoutRetryFlagName = "timeout_retry"
	} else {
		timeoutRetryFlagName = fmt.Sprintf("%v.timeout_retry", cmdPrefix)
	}

	var timeoutRetryFlagDefault int64

	_ = cmd.PersistentFlags().Int64(timeoutRetryFlagName, timeoutRetryFlagDefault, timeoutRetryDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelResolverFlags(depth int, m *models.Resolver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, acceptedPayloadSizeAdded := retrieveResolverAcceptedPayloadSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || acceptedPayloadSizeAdded

	err, holdNxAdded := retrieveResolverHoldNxFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || holdNxAdded

	err, holdObsoleteAdded := retrieveResolverHoldObsoleteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || holdObsoleteAdded

	err, holdOtherAdded := retrieveResolverHoldOtherFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || holdOtherAdded

	err, holdRefusedAdded := retrieveResolverHoldRefusedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || holdRefusedAdded

	err, holdTimeoutAdded := retrieveResolverHoldTimeoutFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || holdTimeoutAdded

	err, holdValidAdded := retrieveResolverHoldValidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || holdValidAdded

	err, nameAdded := retrieveResolverNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, parseResolvConfAdded := retrieveResolverParseResolvConfFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parseResolvConfAdded

	err, resolveRetriesAdded := retrieveResolverResolveRetriesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || resolveRetriesAdded

	err, timeoutResolveAdded := retrieveResolverTimeoutResolveFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timeoutResolveAdded

	err, timeoutRetryAdded := retrieveResolverTimeoutRetryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timeoutRetryAdded

	return nil, retAdded
}

func retrieveResolverAcceptedPayloadSizeFlags(depth int, m *models.Resolver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	acceptedPayloadSizeFlagName := fmt.Sprintf("%v.accepted_payload_size", cmdPrefix)
	if cmd.Flags().Changed(acceptedPayloadSizeFlagName) {

		var acceptedPayloadSizeFlagName string
		if cmdPrefix == "" {
			acceptedPayloadSizeFlagName = "accepted_payload_size"
		} else {
			acceptedPayloadSizeFlagName = fmt.Sprintf("%v.accepted_payload_size", cmdPrefix)
		}

		acceptedPayloadSizeFlagValue, err := cmd.Flags().GetInt64(acceptedPayloadSizeFlagName)
		if err != nil {
			return err, false
		}
		m.AcceptedPayloadSize = acceptedPayloadSizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveResolverHoldNxFlags(depth int, m *models.Resolver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	holdNxFlagName := fmt.Sprintf("%v.hold_nx", cmdPrefix)
	if cmd.Flags().Changed(holdNxFlagName) {

		var holdNxFlagName string
		if cmdPrefix == "" {
			holdNxFlagName = "hold_nx"
		} else {
			holdNxFlagName = fmt.Sprintf("%v.hold_nx", cmdPrefix)
		}

		holdNxFlagValue, err := cmd.Flags().GetInt64(holdNxFlagName)
		if err != nil {
			return err, false
		}
		m.HoldNx = &holdNxFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveResolverHoldObsoleteFlags(depth int, m *models.Resolver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	holdObsoleteFlagName := fmt.Sprintf("%v.hold_obsolete", cmdPrefix)
	if cmd.Flags().Changed(holdObsoleteFlagName) {

		var holdObsoleteFlagName string
		if cmdPrefix == "" {
			holdObsoleteFlagName = "hold_obsolete"
		} else {
			holdObsoleteFlagName = fmt.Sprintf("%v.hold_obsolete", cmdPrefix)
		}

		holdObsoleteFlagValue, err := cmd.Flags().GetInt64(holdObsoleteFlagName)
		if err != nil {
			return err, false
		}
		m.HoldObsolete = &holdObsoleteFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveResolverHoldOtherFlags(depth int, m *models.Resolver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	holdOtherFlagName := fmt.Sprintf("%v.hold_other", cmdPrefix)
	if cmd.Flags().Changed(holdOtherFlagName) {

		var holdOtherFlagName string
		if cmdPrefix == "" {
			holdOtherFlagName = "hold_other"
		} else {
			holdOtherFlagName = fmt.Sprintf("%v.hold_other", cmdPrefix)
		}

		holdOtherFlagValue, err := cmd.Flags().GetInt64(holdOtherFlagName)
		if err != nil {
			return err, false
		}
		m.HoldOther = &holdOtherFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveResolverHoldRefusedFlags(depth int, m *models.Resolver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	holdRefusedFlagName := fmt.Sprintf("%v.hold_refused", cmdPrefix)
	if cmd.Flags().Changed(holdRefusedFlagName) {

		var holdRefusedFlagName string
		if cmdPrefix == "" {
			holdRefusedFlagName = "hold_refused"
		} else {
			holdRefusedFlagName = fmt.Sprintf("%v.hold_refused", cmdPrefix)
		}

		holdRefusedFlagValue, err := cmd.Flags().GetInt64(holdRefusedFlagName)
		if err != nil {
			return err, false
		}
		m.HoldRefused = &holdRefusedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveResolverHoldTimeoutFlags(depth int, m *models.Resolver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	holdTimeoutFlagName := fmt.Sprintf("%v.hold_timeout", cmdPrefix)
	if cmd.Flags().Changed(holdTimeoutFlagName) {

		var holdTimeoutFlagName string
		if cmdPrefix == "" {
			holdTimeoutFlagName = "hold_timeout"
		} else {
			holdTimeoutFlagName = fmt.Sprintf("%v.hold_timeout", cmdPrefix)
		}

		holdTimeoutFlagValue, err := cmd.Flags().GetInt64(holdTimeoutFlagName)
		if err != nil {
			return err, false
		}
		m.HoldTimeout = &holdTimeoutFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveResolverHoldValidFlags(depth int, m *models.Resolver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	holdValidFlagName := fmt.Sprintf("%v.hold_valid", cmdPrefix)
	if cmd.Flags().Changed(holdValidFlagName) {

		var holdValidFlagName string
		if cmdPrefix == "" {
			holdValidFlagName = "hold_valid"
		} else {
			holdValidFlagName = fmt.Sprintf("%v.hold_valid", cmdPrefix)
		}

		holdValidFlagValue, err := cmd.Flags().GetInt64(holdValidFlagName)
		if err != nil {
			return err, false
		}
		m.HoldValid = &holdValidFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveResolverNameFlags(depth int, m *models.Resolver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

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

		retAdded = true
	}

	return nil, retAdded
}

func retrieveResolverParseResolvConfFlags(depth int, m *models.Resolver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	parseResolvConfFlagName := fmt.Sprintf("%v.parse-resolv-conf", cmdPrefix)
	if cmd.Flags().Changed(parseResolvConfFlagName) {

		var parseResolvConfFlagName string
		if cmdPrefix == "" {
			parseResolvConfFlagName = "parse-resolv-conf"
		} else {
			parseResolvConfFlagName = fmt.Sprintf("%v.parse-resolv-conf", cmdPrefix)
		}

		parseResolvConfFlagValue, err := cmd.Flags().GetBool(parseResolvConfFlagName)
		if err != nil {
			return err, false
		}
		m.ParseResolvConf = parseResolvConfFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveResolverResolveRetriesFlags(depth int, m *models.Resolver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	resolveRetriesFlagName := fmt.Sprintf("%v.resolve_retries", cmdPrefix)
	if cmd.Flags().Changed(resolveRetriesFlagName) {

		var resolveRetriesFlagName string
		if cmdPrefix == "" {
			resolveRetriesFlagName = "resolve_retries"
		} else {
			resolveRetriesFlagName = fmt.Sprintf("%v.resolve_retries", cmdPrefix)
		}

		resolveRetriesFlagValue, err := cmd.Flags().GetInt64(resolveRetriesFlagName)
		if err != nil {
			return err, false
		}
		m.ResolveRetries = resolveRetriesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveResolverTimeoutResolveFlags(depth int, m *models.Resolver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	timeoutResolveFlagName := fmt.Sprintf("%v.timeout_resolve", cmdPrefix)
	if cmd.Flags().Changed(timeoutResolveFlagName) {

		var timeoutResolveFlagName string
		if cmdPrefix == "" {
			timeoutResolveFlagName = "timeout_resolve"
		} else {
			timeoutResolveFlagName = fmt.Sprintf("%v.timeout_resolve", cmdPrefix)
		}

		timeoutResolveFlagValue, err := cmd.Flags().GetInt64(timeoutResolveFlagName)
		if err != nil {
			return err, false
		}
		m.TimeoutResolve = timeoutResolveFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveResolverTimeoutRetryFlags(depth int, m *models.Resolver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	timeoutRetryFlagName := fmt.Sprintf("%v.timeout_retry", cmdPrefix)
	if cmd.Flags().Changed(timeoutRetryFlagName) {

		var timeoutRetryFlagName string
		if cmdPrefix == "" {
			timeoutRetryFlagName = "timeout_retry"
		} else {
			timeoutRetryFlagName = fmt.Sprintf("%v.timeout_retry", cmdPrefix)
		}

		timeoutRetryFlagValue, err := cmd.Flags().GetInt64(timeoutRetryFlagName)
		if err != nil {
			return err, false
		}
		m.TimeoutRetry = timeoutRetryFlagValue

		retAdded = true
	}

	return nil, retAdded
}
