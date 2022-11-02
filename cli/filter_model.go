// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/haproxytech/models"
	"github.com/spf13/cobra"
)

// Schema cli for Filter

// register flags to command
func registerModelFilterFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerFilterCacheName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFilterIndex(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFilterSpoeConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFilterSpoeEngine(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFilterTraceHexdump(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFilterTraceName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFilterTraceRndForwarding(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFilterTraceRndParsing(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFilterType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerFilterCacheName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	cacheNameDescription := ``

	var cacheNameFlagName string
	if cmdPrefix == "" {
		cacheNameFlagName = "cache_name"
	} else {
		cacheNameFlagName = fmt.Sprintf("%v.cache_name", cmdPrefix)
	}

	var cacheNameFlagDefault string

	_ = cmd.PersistentFlags().String(cacheNameFlagName, cacheNameFlagDefault, cacheNameDescription)

	return nil
}

func registerFilterIndex(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	indexDescription := `Required. `

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

func registerFilterSpoeConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	spoeConfigDescription := ``

	var spoeConfigFlagName string
	if cmdPrefix == "" {
		spoeConfigFlagName = "spoe_config"
	} else {
		spoeConfigFlagName = fmt.Sprintf("%v.spoe_config", cmdPrefix)
	}

	var spoeConfigFlagDefault string

	_ = cmd.PersistentFlags().String(spoeConfigFlagName, spoeConfigFlagDefault, spoeConfigDescription)

	return nil
}

func registerFilterSpoeEngine(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	spoeEngineDescription := ``

	var spoeEngineFlagName string
	if cmdPrefix == "" {
		spoeEngineFlagName = "spoe_engine"
	} else {
		spoeEngineFlagName = fmt.Sprintf("%v.spoe_engine", cmdPrefix)
	}

	var spoeEngineFlagDefault string

	_ = cmd.PersistentFlags().String(spoeEngineFlagName, spoeEngineFlagDefault, spoeEngineDescription)

	return nil
}

func registerFilterTraceHexdump(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	traceHexdumpDescription := ``

	var traceHexdumpFlagName string
	if cmdPrefix == "" {
		traceHexdumpFlagName = "trace_hexdump"
	} else {
		traceHexdumpFlagName = fmt.Sprintf("%v.trace_hexdump", cmdPrefix)
	}

	var traceHexdumpFlagDefault bool

	_ = cmd.PersistentFlags().Bool(traceHexdumpFlagName, traceHexdumpFlagDefault, traceHexdumpDescription)

	return nil
}

func registerFilterTraceName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	traceNameDescription := ``

	var traceNameFlagName string
	if cmdPrefix == "" {
		traceNameFlagName = "trace_name"
	} else {
		traceNameFlagName = fmt.Sprintf("%v.trace_name", cmdPrefix)
	}

	var traceNameFlagDefault string

	_ = cmd.PersistentFlags().String(traceNameFlagName, traceNameFlagDefault, traceNameDescription)

	return nil
}

func registerFilterTraceRndForwarding(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	traceRndForwardingDescription := ``

	var traceRndForwardingFlagName string
	if cmdPrefix == "" {
		traceRndForwardingFlagName = "trace_rnd_forwarding"
	} else {
		traceRndForwardingFlagName = fmt.Sprintf("%v.trace_rnd_forwarding", cmdPrefix)
	}

	var traceRndForwardingFlagDefault bool

	_ = cmd.PersistentFlags().Bool(traceRndForwardingFlagName, traceRndForwardingFlagDefault, traceRndForwardingDescription)

	return nil
}

func registerFilterTraceRndParsing(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	traceRndParsingDescription := ``

	var traceRndParsingFlagName string
	if cmdPrefix == "" {
		traceRndParsingFlagName = "trace_rnd_parsing"
	} else {
		traceRndParsingFlagName = fmt.Sprintf("%v.trace_rnd_parsing", cmdPrefix)
	}

	var traceRndParsingFlagDefault bool

	_ = cmd.PersistentFlags().Bool(traceRndParsingFlagName, traceRndParsingFlagDefault, traceRndParsingDescription)

	return nil
}

func registerFilterType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["trace","compression","spoe","cache"]. Required. `

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
			if err := json.Unmarshal([]byte(`["trace","compression","spoe","cache"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelFilterFlags(depth int, m *models.Filter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, cacheNameAdded := retrieveFilterCacheNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || cacheNameAdded

	err, indexAdded := retrieveFilterIndexFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || indexAdded

	err, spoeConfigAdded := retrieveFilterSpoeConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || spoeConfigAdded

	err, spoeEngineAdded := retrieveFilterSpoeEngineFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || spoeEngineAdded

	err, traceHexdumpAdded := retrieveFilterTraceHexdumpFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || traceHexdumpAdded

	err, traceNameAdded := retrieveFilterTraceNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || traceNameAdded

	err, traceRndForwardingAdded := retrieveFilterTraceRndForwardingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || traceRndForwardingAdded

	err, traceRndParsingAdded := retrieveFilterTraceRndParsingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || traceRndParsingAdded

	err, typeAdded := retrieveFilterTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveFilterCacheNameFlags(depth int, m *models.Filter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	cacheNameFlagName := fmt.Sprintf("%v.cache_name", cmdPrefix)
	if cmd.Flags().Changed(cacheNameFlagName) {

		var cacheNameFlagName string
		if cmdPrefix == "" {
			cacheNameFlagName = "cache_name"
		} else {
			cacheNameFlagName = fmt.Sprintf("%v.cache_name", cmdPrefix)
		}

		cacheNameFlagValue, err := cmd.Flags().GetString(cacheNameFlagName)
		if err != nil {
			return err, false
		}
		m.CacheName = cacheNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFilterIndexFlags(depth int, m *models.Filter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	indexFlagName := fmt.Sprintf("%v.index", cmdPrefix)
	if cmd.Flags().Changed(indexFlagName) {

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
		m.Index = &indexFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFilterSpoeConfigFlags(depth int, m *models.Filter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	spoeConfigFlagName := fmt.Sprintf("%v.spoe_config", cmdPrefix)
	if cmd.Flags().Changed(spoeConfigFlagName) {

		var spoeConfigFlagName string
		if cmdPrefix == "" {
			spoeConfigFlagName = "spoe_config"
		} else {
			spoeConfigFlagName = fmt.Sprintf("%v.spoe_config", cmdPrefix)
		}

		spoeConfigFlagValue, err := cmd.Flags().GetString(spoeConfigFlagName)
		if err != nil {
			return err, false
		}
		m.SpoeConfig = spoeConfigFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFilterSpoeEngineFlags(depth int, m *models.Filter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	spoeEngineFlagName := fmt.Sprintf("%v.spoe_engine", cmdPrefix)
	if cmd.Flags().Changed(spoeEngineFlagName) {

		var spoeEngineFlagName string
		if cmdPrefix == "" {
			spoeEngineFlagName = "spoe_engine"
		} else {
			spoeEngineFlagName = fmt.Sprintf("%v.spoe_engine", cmdPrefix)
		}

		spoeEngineFlagValue, err := cmd.Flags().GetString(spoeEngineFlagName)
		if err != nil {
			return err, false
		}
		m.SpoeEngine = spoeEngineFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFilterTraceHexdumpFlags(depth int, m *models.Filter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	traceHexdumpFlagName := fmt.Sprintf("%v.trace_hexdump", cmdPrefix)
	if cmd.Flags().Changed(traceHexdumpFlagName) {

		var traceHexdumpFlagName string
		if cmdPrefix == "" {
			traceHexdumpFlagName = "trace_hexdump"
		} else {
			traceHexdumpFlagName = fmt.Sprintf("%v.trace_hexdump", cmdPrefix)
		}

		traceHexdumpFlagValue, err := cmd.Flags().GetBool(traceHexdumpFlagName)
		if err != nil {
			return err, false
		}
		m.TraceHexdump = traceHexdumpFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFilterTraceNameFlags(depth int, m *models.Filter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	traceNameFlagName := fmt.Sprintf("%v.trace_name", cmdPrefix)
	if cmd.Flags().Changed(traceNameFlagName) {

		var traceNameFlagName string
		if cmdPrefix == "" {
			traceNameFlagName = "trace_name"
		} else {
			traceNameFlagName = fmt.Sprintf("%v.trace_name", cmdPrefix)
		}

		traceNameFlagValue, err := cmd.Flags().GetString(traceNameFlagName)
		if err != nil {
			return err, false
		}
		m.TraceName = traceNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFilterTraceRndForwardingFlags(depth int, m *models.Filter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	traceRndForwardingFlagName := fmt.Sprintf("%v.trace_rnd_forwarding", cmdPrefix)
	if cmd.Flags().Changed(traceRndForwardingFlagName) {

		var traceRndForwardingFlagName string
		if cmdPrefix == "" {
			traceRndForwardingFlagName = "trace_rnd_forwarding"
		} else {
			traceRndForwardingFlagName = fmt.Sprintf("%v.trace_rnd_forwarding", cmdPrefix)
		}

		traceRndForwardingFlagValue, err := cmd.Flags().GetBool(traceRndForwardingFlagName)
		if err != nil {
			return err, false
		}
		m.TraceRndForwarding = traceRndForwardingFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFilterTraceRndParsingFlags(depth int, m *models.Filter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	traceRndParsingFlagName := fmt.Sprintf("%v.trace_rnd_parsing", cmdPrefix)
	if cmd.Flags().Changed(traceRndParsingFlagName) {

		var traceRndParsingFlagName string
		if cmdPrefix == "" {
			traceRndParsingFlagName = "trace_rnd_parsing"
		} else {
			traceRndParsingFlagName = fmt.Sprintf("%v.trace_rnd_parsing", cmdPrefix)
		}

		traceRndParsingFlagValue, err := cmd.Flags().GetBool(traceRndParsingFlagName)
		if err != nil {
			return err, false
		}
		m.TraceRndParsing = traceRndParsingFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFilterTypeFlags(depth int, m *models.Filter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

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
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
