// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/haproxytech/client-native/v4/models"

	"github.com/spf13/cobra"
)

// Schema cli for ClusterSettings

// register flags to command
func registerModelClusterSettingsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerClusterSettingsBootstrapKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterSettingsCluster(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterSettingsMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterSettingsStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerClusterSettingsBootstrapKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	bootstrapKeyDescription := ``

	var bootstrapKeyFlagName string
	if cmdPrefix == "" {
		bootstrapKeyFlagName = "bootstrap_key"
	} else {
		bootstrapKeyFlagName = fmt.Sprintf("%v.bootstrap_key", cmdPrefix)
	}

	var bootstrapKeyFlagDefault string

	_ = cmd.PersistentFlags().String(bootstrapKeyFlagName, bootstrapKeyFlagDefault, bootstrapKeyDescription)

	return nil
}

func registerClusterSettingsCluster(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var clusterFlagName string
	if cmdPrefix == "" {
		clusterFlagName = "cluster"
	} else {
		clusterFlagName = fmt.Sprintf("%v.cluster", cmdPrefix)
	}

	if err := registerModelClusterSettingsClusterFlags(depth+1, clusterFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerClusterSettingsMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	modeDescription := `Enum: ["single","cluster"]. `

	var modeFlagName string
	if cmdPrefix == "" {
		modeFlagName = "mode"
	} else {
		modeFlagName = fmt.Sprintf("%v.mode", cmdPrefix)
	}

	var modeFlagDefault string

	_ = cmd.PersistentFlags().String(modeFlagName, modeFlagDefault, modeDescription)

	if err := cmd.RegisterFlagCompletionFunc(modeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["single","cluster"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerClusterSettingsStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := `Enum: ["active","unreachable","waiting_approval"]. `

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "status"
	} else {
		statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	var statusFlagDefault string

	_ = cmd.PersistentFlags().String(statusFlagName, statusFlagDefault, statusDescription)

	if err := cmd.RegisterFlagCompletionFunc(statusFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["active","unreachable","waiting_approval"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelClusterSettingsFlags(depth int, m *models.ClusterSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, bootstrapKeyAdded := retrieveClusterSettingsBootstrapKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bootstrapKeyAdded

	err, clusterAdded := retrieveClusterSettingsClusterFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clusterAdded

	err, modeAdded := retrieveClusterSettingsModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || modeAdded

	err, statusAdded := retrieveClusterSettingsStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	return nil, retAdded
}

func retrieveClusterSettingsBootstrapKeyFlags(depth int, m *models.ClusterSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	bootstrapKeyFlagName := fmt.Sprintf("%v.bootstrap_key", cmdPrefix)
	if cmd.Flags().Changed(bootstrapKeyFlagName) {

		var bootstrapKeyFlagName string
		if cmdPrefix == "" {
			bootstrapKeyFlagName = "bootstrap_key"
		} else {
			bootstrapKeyFlagName = fmt.Sprintf("%v.bootstrap_key", cmdPrefix)
		}

		bootstrapKeyFlagValue, err := cmd.Flags().GetString(bootstrapKeyFlagName)
		if err != nil {
			return err, false
		}
		m.BootstrapKey = bootstrapKeyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterSettingsClusterFlags(depth int, m *models.ClusterSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	clusterFlagName := fmt.Sprintf("%v.cluster", cmdPrefix)
	if cmd.Flags().Changed(clusterFlagName) {
		// info: complex object cluster ClusterSettingsCluster is retrieved outside this Changed() block
	}
	clusterFlagValue := m.Cluster
	if swag.IsZero(clusterFlagValue) {
		clusterFlagValue = &models.ClusterSettingsCluster{}
	}

	err, clusterAdded := retrieveModelClusterSettingsClusterFlags(depth+1, clusterFlagValue, clusterFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clusterAdded
	if clusterAdded {
		m.Cluster = clusterFlagValue
	}

	return nil, retAdded
}

func retrieveClusterSettingsModeFlags(depth int, m *models.ClusterSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	modeFlagName := fmt.Sprintf("%v.mode", cmdPrefix)
	if cmd.Flags().Changed(modeFlagName) {

		var modeFlagName string
		if cmdPrefix == "" {
			modeFlagName = "mode"
		} else {
			modeFlagName = fmt.Sprintf("%v.mode", cmdPrefix)
		}

		modeFlagValue, err := cmd.Flags().GetString(modeFlagName)
		if err != nil {
			return err, false
		}
		m.Mode = modeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterSettingsStatusFlags(depth int, m *models.ClusterSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusFlagName := fmt.Sprintf("%v.status", cmdPrefix)
	if cmd.Flags().Changed(statusFlagName) {

		var statusFlagName string
		if cmdPrefix == "" {
			statusFlagName = "status"
		} else {
			statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
		}

		statusFlagValue, err := cmd.Flags().GetString(statusFlagName)
		if err != nil {
			return err, false
		}
		m.Status = statusFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for ClusterSettingsCluster

// register flags to command
func registerModelClusterSettingsClusterFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerClusterSettingsClusterClusterLogTargets(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterSettingsClusterAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterSettingsClusterAPIBasePath(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterSettingsClusterClusterID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterSettingsClusterDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterSettingsClusterName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterSettingsClusterPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerClusterSettingsClusterClusterLogTargets(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ClusterLogTargets []*ClusterLogTarget array type is not supported by go-swagger cli yet

	return nil
}

func registerClusterSettingsClusterAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	addressDescription := ``

	var addressFlagName string
	if cmdPrefix == "" {
		addressFlagName = "address"
	} else {
		addressFlagName = fmt.Sprintf("%v.address", cmdPrefix)
	}

	var addressFlagDefault string

	_ = cmd.PersistentFlags().String(addressFlagName, addressFlagDefault, addressDescription)

	return nil
}

func registerClusterSettingsClusterAPIBasePath(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	apiBasePathDescription := ``

	var apiBasePathFlagName string
	if cmdPrefix == "" {
		apiBasePathFlagName = "api_base_path"
	} else {
		apiBasePathFlagName = fmt.Sprintf("%v.api_base_path", cmdPrefix)
	}

	var apiBasePathFlagDefault string

	_ = cmd.PersistentFlags().String(apiBasePathFlagName, apiBasePathFlagDefault, apiBasePathDescription)

	return nil
}

func registerClusterSettingsClusterClusterID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	clusterIdDescription := ``

	var clusterIdFlagName string
	if cmdPrefix == "" {
		clusterIdFlagName = "cluster_id"
	} else {
		clusterIdFlagName = fmt.Sprintf("%v.cluster_id", cmdPrefix)
	}

	var clusterIdFlagDefault string

	_ = cmd.PersistentFlags().String(clusterIdFlagName, clusterIdFlagDefault, clusterIdDescription)

	return nil
}

func registerClusterSettingsClusterDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := ``

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerClusterSettingsClusterName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

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

func registerClusterSettingsClusterPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	portDescription := ``

	var portFlagName string
	if cmdPrefix == "" {
		portFlagName = "port"
	} else {
		portFlagName = fmt.Sprintf("%v.port", cmdPrefix)
	}

	var portFlagDefault int64

	_ = cmd.PersistentFlags().Int64(portFlagName, portFlagDefault, portDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelClusterSettingsClusterFlags(depth int, m *models.ClusterSettingsCluster, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, clusterLogTargetsAdded := retrieveClusterSettingsClusterClusterLogTargetsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clusterLogTargetsAdded

	err, addressAdded := retrieveClusterSettingsClusterAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressAdded

	err, apiBasePathAdded := retrieveClusterSettingsClusterAPIBasePathFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || apiBasePathAdded

	err, clusterIdAdded := retrieveClusterSettingsClusterClusterIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clusterIdAdded

	err, descriptionAdded := retrieveClusterSettingsClusterDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, nameAdded := retrieveClusterSettingsClusterNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, portAdded := retrieveClusterSettingsClusterPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portAdded

	return nil, retAdded
}

func retrieveClusterSettingsClusterClusterLogTargetsFlags(depth int, m *models.ClusterSettingsCluster, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	clusterLogTargetsFlagName := fmt.Sprintf("%v.ClusterLogTargets", cmdPrefix)
	if cmd.Flags().Changed(clusterLogTargetsFlagName) {
		// warning: ClusterLogTargets array type []*ClusterLogTarget is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveClusterSettingsClusterAddressFlags(depth int, m *models.ClusterSettingsCluster, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	addressFlagName := fmt.Sprintf("%v.address", cmdPrefix)
	if cmd.Flags().Changed(addressFlagName) {

		var addressFlagName string
		if cmdPrefix == "" {
			addressFlagName = "address"
		} else {
			addressFlagName = fmt.Sprintf("%v.address", cmdPrefix)
		}

		addressFlagValue, err := cmd.Flags().GetString(addressFlagName)
		if err != nil {
			return err, false
		}
		m.Address = addressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterSettingsClusterAPIBasePathFlags(depth int, m *models.ClusterSettingsCluster, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	apiBasePathFlagName := fmt.Sprintf("%v.api_base_path", cmdPrefix)
	if cmd.Flags().Changed(apiBasePathFlagName) {

		var apiBasePathFlagName string
		if cmdPrefix == "" {
			apiBasePathFlagName = "api_base_path"
		} else {
			apiBasePathFlagName = fmt.Sprintf("%v.api_base_path", cmdPrefix)
		}

		apiBasePathFlagValue, err := cmd.Flags().GetString(apiBasePathFlagName)
		if err != nil {
			return err, false
		}
		m.APIBasePath = apiBasePathFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterSettingsClusterClusterIDFlags(depth int, m *models.ClusterSettingsCluster, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	clusterIdFlagName := fmt.Sprintf("%v.cluster_id", cmdPrefix)
	if cmd.Flags().Changed(clusterIdFlagName) {

		var clusterIdFlagName string
		if cmdPrefix == "" {
			clusterIdFlagName = "cluster_id"
		} else {
			clusterIdFlagName = fmt.Sprintf("%v.cluster_id", cmdPrefix)
		}

		clusterIdFlagValue, err := cmd.Flags().GetString(clusterIdFlagName)
		if err != nil {
			return err, false
		}
		m.ClusterID = clusterIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterSettingsClusterDescriptionFlags(depth int, m *models.ClusterSettingsCluster, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterSettingsClusterNameFlags(depth int, m *models.ClusterSettingsCluster, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveClusterSettingsClusterPortFlags(depth int, m *models.ClusterSettingsCluster, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portFlagName := fmt.Sprintf("%v.port", cmdPrefix)
	if cmd.Flags().Changed(portFlagName) {

		var portFlagName string
		if cmdPrefix == "" {
			portFlagName = "port"
		} else {
			portFlagName = fmt.Sprintf("%v.port", cmdPrefix)
		}

		portFlagValue, err := cmd.Flags().GetInt64(portFlagName)
		if err != nil {
			return err, false
		}
		m.Port = &portFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for ClusterLogTarget

// register flags to command
func registerModelClusterLogTargetFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerClusterLogTargetAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterLogTargetLogFormat(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterLogTargetPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterLogTargetProtocol(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerClusterLogTargetAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	addressDescription := `Required. `

	var addressFlagName string
	if cmdPrefix == "" {
		addressFlagName = "address"
	} else {
		addressFlagName = fmt.Sprintf("%v.address", cmdPrefix)
	}

	var addressFlagDefault string

	_ = cmd.PersistentFlags().String(addressFlagName, addressFlagDefault, addressDescription)

	return nil
}

func registerClusterLogTargetLogFormat(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	logFormatDescription := ``

	var logFormatFlagName string
	if cmdPrefix == "" {
		logFormatFlagName = "log_format"
	} else {
		logFormatFlagName = fmt.Sprintf("%v.log_format", cmdPrefix)
	}

	var logFormatFlagDefault string

	_ = cmd.PersistentFlags().String(logFormatFlagName, logFormatFlagDefault, logFormatDescription)

	return nil
}

func registerClusterLogTargetPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	portDescription := `Required. `

	var portFlagName string
	if cmdPrefix == "" {
		portFlagName = "port"
	} else {
		portFlagName = fmt.Sprintf("%v.port", cmdPrefix)
	}

	var portFlagDefault int64

	_ = cmd.PersistentFlags().Int64(portFlagName, portFlagDefault, portDescription)

	return nil
}

func registerClusterLogTargetProtocol(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	protocolDescription := `Enum: ["tcp","udp"]. Required. `

	var protocolFlagName string
	if cmdPrefix == "" {
		protocolFlagName = "protocol"
	} else {
		protocolFlagName = fmt.Sprintf("%v.protocol", cmdPrefix)
	}

	var protocolFlagDefault string

	_ = cmd.PersistentFlags().String(protocolFlagName, protocolFlagDefault, protocolDescription)

	if err := cmd.RegisterFlagCompletionFunc(protocolFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["tcp","udp"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelClusterLogTargetFlags(depth int, m *models.ClusterLogTarget, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressAdded := retrieveClusterLogTargetAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressAdded

	err, logFormatAdded := retrieveClusterLogTargetLogFormatFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || logFormatAdded

	err, portAdded := retrieveClusterLogTargetPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portAdded

	err, protocolAdded := retrieveClusterLogTargetProtocolFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || protocolAdded

	return nil, retAdded
}

func retrieveClusterLogTargetAddressFlags(depth int, m *models.ClusterLogTarget, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	addressFlagName := fmt.Sprintf("%v.address", cmdPrefix)
	if cmd.Flags().Changed(addressFlagName) {

		var addressFlagName string
		if cmdPrefix == "" {
			addressFlagName = "address"
		} else {
			addressFlagName = fmt.Sprintf("%v.address", cmdPrefix)
		}

		addressFlagValue, err := cmd.Flags().GetString(addressFlagName)
		if err != nil {
			return err, false
		}
		m.Address = &addressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterLogTargetLogFormatFlags(depth int, m *models.ClusterLogTarget, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	logFormatFlagName := fmt.Sprintf("%v.log_format", cmdPrefix)
	if cmd.Flags().Changed(logFormatFlagName) {

		var logFormatFlagName string
		if cmdPrefix == "" {
			logFormatFlagName = "log_format"
		} else {
			logFormatFlagName = fmt.Sprintf("%v.log_format", cmdPrefix)
		}

		logFormatFlagValue, err := cmd.Flags().GetString(logFormatFlagName)
		if err != nil {
			return err, false
		}
		m.LogFormat = logFormatFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterLogTargetPortFlags(depth int, m *models.ClusterLogTarget, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portFlagName := fmt.Sprintf("%v.port", cmdPrefix)
	if cmd.Flags().Changed(portFlagName) {

		var portFlagName string
		if cmdPrefix == "" {
			portFlagName = "port"
		} else {
			portFlagName = fmt.Sprintf("%v.port", cmdPrefix)
		}

		portFlagValue, err := cmd.Flags().GetInt64(portFlagName)
		if err != nil {
			return err, false
		}
		m.Port = &portFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterLogTargetProtocolFlags(depth int, m *models.ClusterLogTarget, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	protocolFlagName := fmt.Sprintf("%v.protocol", cmdPrefix)
	if cmd.Flags().Changed(protocolFlagName) {

		var protocolFlagName string
		if cmdPrefix == "" {
			protocolFlagName = "protocol"
		} else {
			protocolFlagName = fmt.Sprintf("%v.protocol", cmdPrefix)
		}

		protocolFlagValue, err := cmd.Flags().GetString(protocolFlagName)
		if err != nil {
			return err, false
		}
		m.Protocol = &protocolFlagValue

		retAdded = true
	}

	return nil, retAdded
}
