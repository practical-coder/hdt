// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/practical-coder/hdc/models"
	"github.com/spf13/cobra"
)

// Schema cli for StatsOptions

// register flags to command
func registerModelStatsOptionsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerStatsOptionsStatsAdmin(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsAdminCond(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsAdminCondTest(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsAuths(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsEnable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsHideVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsHTTPRequests(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsMaxconn(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsRealm(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsRealmRealm(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsRefreshDelay(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsShowDesc(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsShowLegends(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsShowModules(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsShowNodeName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatsOptionsStatsURIPrefix(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerStatsOptionsStatsAdmin(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statsAdminDescription := ``

	var statsAdminFlagName string
	if cmdPrefix == "" {
		statsAdminFlagName = "stats_admin"
	} else {
		statsAdminFlagName = fmt.Sprintf("%v.stats_admin", cmdPrefix)
	}

	var statsAdminFlagDefault bool

	_ = cmd.PersistentFlags().Bool(statsAdminFlagName, statsAdminFlagDefault, statsAdminDescription)

	return nil
}

func registerStatsOptionsStatsAdminCond(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statsAdminCondDescription := `Enum: ["if","unless"]. `

	var statsAdminCondFlagName string
	if cmdPrefix == "" {
		statsAdminCondFlagName = "stats_admin_cond"
	} else {
		statsAdminCondFlagName = fmt.Sprintf("%v.stats_admin_cond", cmdPrefix)
	}

	var statsAdminCondFlagDefault string

	_ = cmd.PersistentFlags().String(statsAdminCondFlagName, statsAdminCondFlagDefault, statsAdminCondDescription)

	if err := cmd.RegisterFlagCompletionFunc(statsAdminCondFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerStatsOptionsStatsAdminCondTest(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statsAdminCondTestDescription := ``

	var statsAdminCondTestFlagName string
	if cmdPrefix == "" {
		statsAdminCondTestFlagName = "stats_admin_cond_test"
	} else {
		statsAdminCondTestFlagName = fmt.Sprintf("%v.stats_admin_cond_test", cmdPrefix)
	}

	var statsAdminCondTestFlagDefault string

	_ = cmd.PersistentFlags().String(statsAdminCondTestFlagName, statsAdminCondTestFlagDefault, statsAdminCondTestDescription)

	return nil
}

func registerStatsOptionsStatsAuths(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: stats_auths []*StatsAuth array type is not supported by go-swagger cli yet

	return nil
}

func registerStatsOptionsStatsEnable(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statsEnableDescription := ``

	var statsEnableFlagName string
	if cmdPrefix == "" {
		statsEnableFlagName = "stats_enable"
	} else {
		statsEnableFlagName = fmt.Sprintf("%v.stats_enable", cmdPrefix)
	}

	var statsEnableFlagDefault bool

	_ = cmd.PersistentFlags().Bool(statsEnableFlagName, statsEnableFlagDefault, statsEnableDescription)

	return nil
}

func registerStatsOptionsStatsHideVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statsHideVersionDescription := ``

	var statsHideVersionFlagName string
	if cmdPrefix == "" {
		statsHideVersionFlagName = "stats_hide_version"
	} else {
		statsHideVersionFlagName = fmt.Sprintf("%v.stats_hide_version", cmdPrefix)
	}

	var statsHideVersionFlagDefault bool

	_ = cmd.PersistentFlags().Bool(statsHideVersionFlagName, statsHideVersionFlagDefault, statsHideVersionDescription)

	return nil
}

func registerStatsOptionsStatsHTTPRequests(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: stats_http_requests []*StatsHTTPRequest array type is not supported by go-swagger cli yet

	return nil
}

func registerStatsOptionsStatsMaxconn(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statsMaxconnDescription := ``

	var statsMaxconnFlagName string
	if cmdPrefix == "" {
		statsMaxconnFlagName = "stats_maxconn"
	} else {
		statsMaxconnFlagName = fmt.Sprintf("%v.stats_maxconn", cmdPrefix)
	}

	var statsMaxconnFlagDefault int64

	_ = cmd.PersistentFlags().Int64(statsMaxconnFlagName, statsMaxconnFlagDefault, statsMaxconnDescription)

	return nil
}

func registerStatsOptionsStatsRealm(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statsRealmDescription := ``

	var statsRealmFlagName string
	if cmdPrefix == "" {
		statsRealmFlagName = "stats_realm"
	} else {
		statsRealmFlagName = fmt.Sprintf("%v.stats_realm", cmdPrefix)
	}

	var statsRealmFlagDefault bool

	_ = cmd.PersistentFlags().Bool(statsRealmFlagName, statsRealmFlagDefault, statsRealmDescription)

	return nil
}

func registerStatsOptionsStatsRealmRealm(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statsRealmRealmDescription := ``

	var statsRealmRealmFlagName string
	if cmdPrefix == "" {
		statsRealmRealmFlagName = "stats_realm_realm"
	} else {
		statsRealmRealmFlagName = fmt.Sprintf("%v.stats_realm_realm", cmdPrefix)
	}

	var statsRealmRealmFlagDefault string

	_ = cmd.PersistentFlags().String(statsRealmRealmFlagName, statsRealmRealmFlagDefault, statsRealmRealmDescription)

	return nil
}

func registerStatsOptionsStatsRefreshDelay(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statsRefreshDelayDescription := ``

	var statsRefreshDelayFlagName string
	if cmdPrefix == "" {
		statsRefreshDelayFlagName = "stats_refresh_delay"
	} else {
		statsRefreshDelayFlagName = fmt.Sprintf("%v.stats_refresh_delay", cmdPrefix)
	}

	var statsRefreshDelayFlagDefault int64

	_ = cmd.PersistentFlags().Int64(statsRefreshDelayFlagName, statsRefreshDelayFlagDefault, statsRefreshDelayDescription)

	return nil
}

func registerStatsOptionsStatsShowDesc(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statsShowDescDescription := ``

	var statsShowDescFlagName string
	if cmdPrefix == "" {
		statsShowDescFlagName = "stats_show_desc"
	} else {
		statsShowDescFlagName = fmt.Sprintf("%v.stats_show_desc", cmdPrefix)
	}

	var statsShowDescFlagDefault string

	_ = cmd.PersistentFlags().String(statsShowDescFlagName, statsShowDescFlagDefault, statsShowDescDescription)

	return nil
}

func registerStatsOptionsStatsShowLegends(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statsShowLegendsDescription := ``

	var statsShowLegendsFlagName string
	if cmdPrefix == "" {
		statsShowLegendsFlagName = "stats_show_legends"
	} else {
		statsShowLegendsFlagName = fmt.Sprintf("%v.stats_show_legends", cmdPrefix)
	}

	var statsShowLegendsFlagDefault bool

	_ = cmd.PersistentFlags().Bool(statsShowLegendsFlagName, statsShowLegendsFlagDefault, statsShowLegendsDescription)

	return nil
}

func registerStatsOptionsStatsShowModules(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statsShowModulesDescription := ``

	var statsShowModulesFlagName string
	if cmdPrefix == "" {
		statsShowModulesFlagName = "stats_show_modules"
	} else {
		statsShowModulesFlagName = fmt.Sprintf("%v.stats_show_modules", cmdPrefix)
	}

	var statsShowModulesFlagDefault bool

	_ = cmd.PersistentFlags().Bool(statsShowModulesFlagName, statsShowModulesFlagDefault, statsShowModulesDescription)

	return nil
}

func registerStatsOptionsStatsShowNodeName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statsShowNodeNameDescription := ``

	var statsShowNodeNameFlagName string
	if cmdPrefix == "" {
		statsShowNodeNameFlagName = "stats_show_node_name"
	} else {
		statsShowNodeNameFlagName = fmt.Sprintf("%v.stats_show_node_name", cmdPrefix)
	}

	var statsShowNodeNameFlagDefault string

	_ = cmd.PersistentFlags().String(statsShowNodeNameFlagName, statsShowNodeNameFlagDefault, statsShowNodeNameDescription)

	return nil
}

func registerStatsOptionsStatsURIPrefix(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statsUriPrefixDescription := ``

	var statsUriPrefixFlagName string
	if cmdPrefix == "" {
		statsUriPrefixFlagName = "stats_uri_prefix"
	} else {
		statsUriPrefixFlagName = fmt.Sprintf("%v.stats_uri_prefix", cmdPrefix)
	}

	var statsUriPrefixFlagDefault string

	_ = cmd.PersistentFlags().String(statsUriPrefixFlagName, statsUriPrefixFlagDefault, statsUriPrefixDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelStatsOptionsFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, statsAdminAdded := retrieveStatsOptionsStatsAdminFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsAdminAdded

	err, statsAdminCondAdded := retrieveStatsOptionsStatsAdminCondFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsAdminCondAdded

	err, statsAdminCondTestAdded := retrieveStatsOptionsStatsAdminCondTestFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsAdminCondTestAdded

	err, statsAuthsAdded := retrieveStatsOptionsStatsAuthsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsAuthsAdded

	err, statsEnableAdded := retrieveStatsOptionsStatsEnableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsEnableAdded

	err, statsHideVersionAdded := retrieveStatsOptionsStatsHideVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsHideVersionAdded

	err, statsHttpRequestsAdded := retrieveStatsOptionsStatsHTTPRequestsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsHttpRequestsAdded

	err, statsMaxconnAdded := retrieveStatsOptionsStatsMaxconnFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsMaxconnAdded

	err, statsRealmAdded := retrieveStatsOptionsStatsRealmFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsRealmAdded

	err, statsRealmRealmAdded := retrieveStatsOptionsStatsRealmRealmFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsRealmRealmAdded

	err, statsRefreshDelayAdded := retrieveStatsOptionsStatsRefreshDelayFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsRefreshDelayAdded

	err, statsShowDescAdded := retrieveStatsOptionsStatsShowDescFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsShowDescAdded

	err, statsShowLegendsAdded := retrieveStatsOptionsStatsShowLegendsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsShowLegendsAdded

	err, statsShowModulesAdded := retrieveStatsOptionsStatsShowModulesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsShowModulesAdded

	err, statsShowNodeNameAdded := retrieveStatsOptionsStatsShowNodeNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsShowNodeNameAdded

	err, statsUriPrefixAdded := retrieveStatsOptionsStatsURIPrefixFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statsUriPrefixAdded

	return nil, retAdded
}

func retrieveStatsOptionsStatsAdminFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsAdminFlagName := fmt.Sprintf("%v.stats_admin", cmdPrefix)
	if cmd.Flags().Changed(statsAdminFlagName) {

		var statsAdminFlagName string
		if cmdPrefix == "" {
			statsAdminFlagName = "stats_admin"
		} else {
			statsAdminFlagName = fmt.Sprintf("%v.stats_admin", cmdPrefix)
		}

		statsAdminFlagValue, err := cmd.Flags().GetBool(statsAdminFlagName)
		if err != nil {
			return err, false
		}
		m.StatsAdmin = statsAdminFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsAdminCondFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsAdminCondFlagName := fmt.Sprintf("%v.stats_admin_cond", cmdPrefix)
	if cmd.Flags().Changed(statsAdminCondFlagName) {

		var statsAdminCondFlagName string
		if cmdPrefix == "" {
			statsAdminCondFlagName = "stats_admin_cond"
		} else {
			statsAdminCondFlagName = fmt.Sprintf("%v.stats_admin_cond", cmdPrefix)
		}

		statsAdminCondFlagValue, err := cmd.Flags().GetString(statsAdminCondFlagName)
		if err != nil {
			return err, false
		}
		m.StatsAdminCond = statsAdminCondFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsAdminCondTestFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsAdminCondTestFlagName := fmt.Sprintf("%v.stats_admin_cond_test", cmdPrefix)
	if cmd.Flags().Changed(statsAdminCondTestFlagName) {

		var statsAdminCondTestFlagName string
		if cmdPrefix == "" {
			statsAdminCondTestFlagName = "stats_admin_cond_test"
		} else {
			statsAdminCondTestFlagName = fmt.Sprintf("%v.stats_admin_cond_test", cmdPrefix)
		}

		statsAdminCondTestFlagValue, err := cmd.Flags().GetString(statsAdminCondTestFlagName)
		if err != nil {
			return err, false
		}
		m.StatsAdminCondTest = statsAdminCondTestFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsAuthsFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsAuthsFlagName := fmt.Sprintf("%v.stats_auths", cmdPrefix)
	if cmd.Flags().Changed(statsAuthsFlagName) {
		// warning: stats_auths array type []*StatsAuth is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsEnableFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsEnableFlagName := fmt.Sprintf("%v.stats_enable", cmdPrefix)
	if cmd.Flags().Changed(statsEnableFlagName) {

		var statsEnableFlagName string
		if cmdPrefix == "" {
			statsEnableFlagName = "stats_enable"
		} else {
			statsEnableFlagName = fmt.Sprintf("%v.stats_enable", cmdPrefix)
		}

		statsEnableFlagValue, err := cmd.Flags().GetBool(statsEnableFlagName)
		if err != nil {
			return err, false
		}
		m.StatsEnable = statsEnableFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsHideVersionFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsHideVersionFlagName := fmt.Sprintf("%v.stats_hide_version", cmdPrefix)
	if cmd.Flags().Changed(statsHideVersionFlagName) {

		var statsHideVersionFlagName string
		if cmdPrefix == "" {
			statsHideVersionFlagName = "stats_hide_version"
		} else {
			statsHideVersionFlagName = fmt.Sprintf("%v.stats_hide_version", cmdPrefix)
		}

		statsHideVersionFlagValue, err := cmd.Flags().GetBool(statsHideVersionFlagName)
		if err != nil {
			return err, false
		}
		m.StatsHideVersion = statsHideVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsHTTPRequestsFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsHttpRequestsFlagName := fmt.Sprintf("%v.stats_http_requests", cmdPrefix)
	if cmd.Flags().Changed(statsHttpRequestsFlagName) {
		// warning: stats_http_requests array type []*StatsHTTPRequest is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsMaxconnFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsMaxconnFlagName := fmt.Sprintf("%v.stats_maxconn", cmdPrefix)
	if cmd.Flags().Changed(statsMaxconnFlagName) {

		var statsMaxconnFlagName string
		if cmdPrefix == "" {
			statsMaxconnFlagName = "stats_maxconn"
		} else {
			statsMaxconnFlagName = fmt.Sprintf("%v.stats_maxconn", cmdPrefix)
		}

		statsMaxconnFlagValue, err := cmd.Flags().GetInt64(statsMaxconnFlagName)
		if err != nil {
			return err, false
		}
		m.StatsMaxconn = statsMaxconnFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsRealmFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsRealmFlagName := fmt.Sprintf("%v.stats_realm", cmdPrefix)
	if cmd.Flags().Changed(statsRealmFlagName) {

		var statsRealmFlagName string
		if cmdPrefix == "" {
			statsRealmFlagName = "stats_realm"
		} else {
			statsRealmFlagName = fmt.Sprintf("%v.stats_realm", cmdPrefix)
		}

		statsRealmFlagValue, err := cmd.Flags().GetBool(statsRealmFlagName)
		if err != nil {
			return err, false
		}
		m.StatsRealm = statsRealmFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsRealmRealmFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsRealmRealmFlagName := fmt.Sprintf("%v.stats_realm_realm", cmdPrefix)
	if cmd.Flags().Changed(statsRealmRealmFlagName) {

		var statsRealmRealmFlagName string
		if cmdPrefix == "" {
			statsRealmRealmFlagName = "stats_realm_realm"
		} else {
			statsRealmRealmFlagName = fmt.Sprintf("%v.stats_realm_realm", cmdPrefix)
		}

		statsRealmRealmFlagValue, err := cmd.Flags().GetString(statsRealmRealmFlagName)
		if err != nil {
			return err, false
		}
		m.StatsRealmRealm = &statsRealmRealmFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsRefreshDelayFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsRefreshDelayFlagName := fmt.Sprintf("%v.stats_refresh_delay", cmdPrefix)
	if cmd.Flags().Changed(statsRefreshDelayFlagName) {

		var statsRefreshDelayFlagName string
		if cmdPrefix == "" {
			statsRefreshDelayFlagName = "stats_refresh_delay"
		} else {
			statsRefreshDelayFlagName = fmt.Sprintf("%v.stats_refresh_delay", cmdPrefix)
		}

		statsRefreshDelayFlagValue, err := cmd.Flags().GetInt64(statsRefreshDelayFlagName)
		if err != nil {
			return err, false
		}
		m.StatsRefreshDelay = &statsRefreshDelayFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsShowDescFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsShowDescFlagName := fmt.Sprintf("%v.stats_show_desc", cmdPrefix)
	if cmd.Flags().Changed(statsShowDescFlagName) {

		var statsShowDescFlagName string
		if cmdPrefix == "" {
			statsShowDescFlagName = "stats_show_desc"
		} else {
			statsShowDescFlagName = fmt.Sprintf("%v.stats_show_desc", cmdPrefix)
		}

		statsShowDescFlagValue, err := cmd.Flags().GetString(statsShowDescFlagName)
		if err != nil {
			return err, false
		}
		m.StatsShowDesc = &statsShowDescFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsShowLegendsFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsShowLegendsFlagName := fmt.Sprintf("%v.stats_show_legends", cmdPrefix)
	if cmd.Flags().Changed(statsShowLegendsFlagName) {

		var statsShowLegendsFlagName string
		if cmdPrefix == "" {
			statsShowLegendsFlagName = "stats_show_legends"
		} else {
			statsShowLegendsFlagName = fmt.Sprintf("%v.stats_show_legends", cmdPrefix)
		}

		statsShowLegendsFlagValue, err := cmd.Flags().GetBool(statsShowLegendsFlagName)
		if err != nil {
			return err, false
		}
		m.StatsShowLegends = statsShowLegendsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsShowModulesFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsShowModulesFlagName := fmt.Sprintf("%v.stats_show_modules", cmdPrefix)
	if cmd.Flags().Changed(statsShowModulesFlagName) {

		var statsShowModulesFlagName string
		if cmdPrefix == "" {
			statsShowModulesFlagName = "stats_show_modules"
		} else {
			statsShowModulesFlagName = fmt.Sprintf("%v.stats_show_modules", cmdPrefix)
		}

		statsShowModulesFlagValue, err := cmd.Flags().GetBool(statsShowModulesFlagName)
		if err != nil {
			return err, false
		}
		m.StatsShowModules = statsShowModulesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsShowNodeNameFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsShowNodeNameFlagName := fmt.Sprintf("%v.stats_show_node_name", cmdPrefix)
	if cmd.Flags().Changed(statsShowNodeNameFlagName) {

		var statsShowNodeNameFlagName string
		if cmdPrefix == "" {
			statsShowNodeNameFlagName = "stats_show_node_name"
		} else {
			statsShowNodeNameFlagName = fmt.Sprintf("%v.stats_show_node_name", cmdPrefix)
		}

		statsShowNodeNameFlagValue, err := cmd.Flags().GetString(statsShowNodeNameFlagName)
		if err != nil {
			return err, false
		}
		m.StatsShowNodeName = &statsShowNodeNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatsOptionsStatsURIPrefixFlags(depth int, m *models.StatsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statsUriPrefixFlagName := fmt.Sprintf("%v.stats_uri_prefix", cmdPrefix)
	if cmd.Flags().Changed(statsUriPrefixFlagName) {

		var statsUriPrefixFlagName string
		if cmdPrefix == "" {
			statsUriPrefixFlagName = "stats_uri_prefix"
		} else {
			statsUriPrefixFlagName = fmt.Sprintf("%v.stats_uri_prefix", cmdPrefix)
		}

		statsUriPrefixFlagValue, err := cmd.Flags().GetString(statsUriPrefixFlagName)
		if err != nil {
			return err, false
		}
		m.StatsURIPrefix = statsUriPrefixFlagValue

		retAdded = true
	}

	return nil, retAdded
}