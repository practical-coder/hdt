// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/spf13/cobra"
)

// Schema cli for Cookie

// register flags to command
func registerModelCookieFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCookieDomains(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCookieDynamic(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCookieHttponly(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCookieIndirect(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCookieMaxidle(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCookieMaxlife(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCookieName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCookieNocache(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCookiePostonly(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCookiePreserve(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCookieSecure(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCookieType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCookieDomains(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Domains []*Domain array type is not supported by go-swagger cli yet

	return nil
}

func registerCookieDynamic(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	dynamicDescription := ``

	var dynamicFlagName string
	if cmdPrefix == "" {
		dynamicFlagName = "dynamic"
	} else {
		dynamicFlagName = fmt.Sprintf("%v.dynamic", cmdPrefix)
	}

	var dynamicFlagDefault bool

	_ = cmd.PersistentFlags().Bool(dynamicFlagName, dynamicFlagDefault, dynamicDescription)

	return nil
}

func registerCookieHttponly(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	httponlyDescription := ``

	var httponlyFlagName string
	if cmdPrefix == "" {
		httponlyFlagName = "httponly"
	} else {
		httponlyFlagName = fmt.Sprintf("%v.httponly", cmdPrefix)
	}

	var httponlyFlagDefault bool

	_ = cmd.PersistentFlags().Bool(httponlyFlagName, httponlyFlagDefault, httponlyDescription)

	return nil
}

func registerCookieIndirect(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	indirectDescription := ``

	var indirectFlagName string
	if cmdPrefix == "" {
		indirectFlagName = "indirect"
	} else {
		indirectFlagName = fmt.Sprintf("%v.indirect", cmdPrefix)
	}

	var indirectFlagDefault bool

	_ = cmd.PersistentFlags().Bool(indirectFlagName, indirectFlagDefault, indirectDescription)

	return nil
}

func registerCookieMaxidle(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxidleDescription := ``

	var maxidleFlagName string
	if cmdPrefix == "" {
		maxidleFlagName = "maxidle"
	} else {
		maxidleFlagName = fmt.Sprintf("%v.maxidle", cmdPrefix)
	}

	var maxidleFlagDefault int64

	_ = cmd.PersistentFlags().Int64(maxidleFlagName, maxidleFlagDefault, maxidleDescription)

	return nil
}

func registerCookieMaxlife(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxlifeDescription := ``

	var maxlifeFlagName string
	if cmdPrefix == "" {
		maxlifeFlagName = "maxlife"
	} else {
		maxlifeFlagName = fmt.Sprintf("%v.maxlife", cmdPrefix)
	}

	var maxlifeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(maxlifeFlagName, maxlifeFlagDefault, maxlifeDescription)

	return nil
}

func registerCookieName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerCookieNocache(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nocacheDescription := ``

	var nocacheFlagName string
	if cmdPrefix == "" {
		nocacheFlagName = "nocache"
	} else {
		nocacheFlagName = fmt.Sprintf("%v.nocache", cmdPrefix)
	}

	var nocacheFlagDefault bool

	_ = cmd.PersistentFlags().Bool(nocacheFlagName, nocacheFlagDefault, nocacheDescription)

	return nil
}

func registerCookiePostonly(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	postonlyDescription := ``

	var postonlyFlagName string
	if cmdPrefix == "" {
		postonlyFlagName = "postonly"
	} else {
		postonlyFlagName = fmt.Sprintf("%v.postonly", cmdPrefix)
	}

	var postonlyFlagDefault bool

	_ = cmd.PersistentFlags().Bool(postonlyFlagName, postonlyFlagDefault, postonlyDescription)

	return nil
}

func registerCookiePreserve(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	preserveDescription := ``

	var preserveFlagName string
	if cmdPrefix == "" {
		preserveFlagName = "preserve"
	} else {
		preserveFlagName = fmt.Sprintf("%v.preserve", cmdPrefix)
	}

	var preserveFlagDefault bool

	_ = cmd.PersistentFlags().Bool(preserveFlagName, preserveFlagDefault, preserveDescription)

	return nil
}

func registerCookieSecure(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	secureDescription := ``

	var secureFlagName string
	if cmdPrefix == "" {
		secureFlagName = "secure"
	} else {
		secureFlagName = fmt.Sprintf("%v.secure", cmdPrefix)
	}

	var secureFlagDefault bool

	_ = cmd.PersistentFlags().Bool(secureFlagName, secureFlagDefault, secureDescription)

	return nil
}

func registerCookieType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["rewrite","insert","prefix"]. `

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
			if err := json.Unmarshal([]byte(`["rewrite","insert","prefix"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCookieFlags(depth int, m *models.Cookie, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, domainsAdded := retrieveCookieDomainsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || domainsAdded

	err, dynamicAdded := retrieveCookieDynamicFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dynamicAdded

	err, httponlyAdded := retrieveCookieHttponlyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || httponlyAdded

	err, indirectAdded := retrieveCookieIndirectFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || indirectAdded

	err, maxidleAdded := retrieveCookieMaxidleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxidleAdded

	err, maxlifeAdded := retrieveCookieMaxlifeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxlifeAdded

	err, nameAdded := retrieveCookieNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, nocacheAdded := retrieveCookieNocacheFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nocacheAdded

	err, postonlyAdded := retrieveCookiePostonlyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || postonlyAdded

	err, preserveAdded := retrieveCookiePreserveFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || preserveAdded

	err, secureAdded := retrieveCookieSecureFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || secureAdded

	err, typeAdded := retrieveCookieTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveCookieDomainsFlags(depth int, m *models.Cookie, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	domainsFlagName := fmt.Sprintf("%v.Domains", cmdPrefix)
	if cmd.Flags().Changed(domainsFlagName) {
		// warning: Domains array type []*Domain is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveCookieDynamicFlags(depth int, m *models.Cookie, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dynamicFlagName := fmt.Sprintf("%v.dynamic", cmdPrefix)
	if cmd.Flags().Changed(dynamicFlagName) {

		var dynamicFlagName string
		if cmdPrefix == "" {
			dynamicFlagName = "dynamic"
		} else {
			dynamicFlagName = fmt.Sprintf("%v.dynamic", cmdPrefix)
		}

		dynamicFlagValue, err := cmd.Flags().GetBool(dynamicFlagName)
		if err != nil {
			return err, false
		}
		m.Dynamic = dynamicFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCookieHttponlyFlags(depth int, m *models.Cookie, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	httponlyFlagName := fmt.Sprintf("%v.httponly", cmdPrefix)
	if cmd.Flags().Changed(httponlyFlagName) {

		var httponlyFlagName string
		if cmdPrefix == "" {
			httponlyFlagName = "httponly"
		} else {
			httponlyFlagName = fmt.Sprintf("%v.httponly", cmdPrefix)
		}

		httponlyFlagValue, err := cmd.Flags().GetBool(httponlyFlagName)
		if err != nil {
			return err, false
		}
		m.Httponly = httponlyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCookieIndirectFlags(depth int, m *models.Cookie, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	indirectFlagName := fmt.Sprintf("%v.indirect", cmdPrefix)
	if cmd.Flags().Changed(indirectFlagName) {

		var indirectFlagName string
		if cmdPrefix == "" {
			indirectFlagName = "indirect"
		} else {
			indirectFlagName = fmt.Sprintf("%v.indirect", cmdPrefix)
		}

		indirectFlagValue, err := cmd.Flags().GetBool(indirectFlagName)
		if err != nil {
			return err, false
		}
		m.Indirect = indirectFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCookieMaxidleFlags(depth int, m *models.Cookie, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxidleFlagName := fmt.Sprintf("%v.maxidle", cmdPrefix)
	if cmd.Flags().Changed(maxidleFlagName) {

		var maxidleFlagName string
		if cmdPrefix == "" {
			maxidleFlagName = "maxidle"
		} else {
			maxidleFlagName = fmt.Sprintf("%v.maxidle", cmdPrefix)
		}

		maxidleFlagValue, err := cmd.Flags().GetInt64(maxidleFlagName)
		if err != nil {
			return err, false
		}
		m.Maxidle = maxidleFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCookieMaxlifeFlags(depth int, m *models.Cookie, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxlifeFlagName := fmt.Sprintf("%v.maxlife", cmdPrefix)
	if cmd.Flags().Changed(maxlifeFlagName) {

		var maxlifeFlagName string
		if cmdPrefix == "" {
			maxlifeFlagName = "maxlife"
		} else {
			maxlifeFlagName = fmt.Sprintf("%v.maxlife", cmdPrefix)
		}

		maxlifeFlagValue, err := cmd.Flags().GetInt64(maxlifeFlagName)
		if err != nil {
			return err, false
		}
		m.Maxlife = maxlifeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCookieNameFlags(depth int, m *models.Cookie, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCookieNocacheFlags(depth int, m *models.Cookie, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nocacheFlagName := fmt.Sprintf("%v.nocache", cmdPrefix)
	if cmd.Flags().Changed(nocacheFlagName) {

		var nocacheFlagName string
		if cmdPrefix == "" {
			nocacheFlagName = "nocache"
		} else {
			nocacheFlagName = fmt.Sprintf("%v.nocache", cmdPrefix)
		}

		nocacheFlagValue, err := cmd.Flags().GetBool(nocacheFlagName)
		if err != nil {
			return err, false
		}
		m.Nocache = nocacheFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCookiePostonlyFlags(depth int, m *models.Cookie, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	postonlyFlagName := fmt.Sprintf("%v.postonly", cmdPrefix)
	if cmd.Flags().Changed(postonlyFlagName) {

		var postonlyFlagName string
		if cmdPrefix == "" {
			postonlyFlagName = "postonly"
		} else {
			postonlyFlagName = fmt.Sprintf("%v.postonly", cmdPrefix)
		}

		postonlyFlagValue, err := cmd.Flags().GetBool(postonlyFlagName)
		if err != nil {
			return err, false
		}
		m.Postonly = postonlyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCookiePreserveFlags(depth int, m *models.Cookie, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	preserveFlagName := fmt.Sprintf("%v.preserve", cmdPrefix)
	if cmd.Flags().Changed(preserveFlagName) {

		var preserveFlagName string
		if cmdPrefix == "" {
			preserveFlagName = "preserve"
		} else {
			preserveFlagName = fmt.Sprintf("%v.preserve", cmdPrefix)
		}

		preserveFlagValue, err := cmd.Flags().GetBool(preserveFlagName)
		if err != nil {
			return err, false
		}
		m.Preserve = preserveFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCookieSecureFlags(depth int, m *models.Cookie, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	secureFlagName := fmt.Sprintf("%v.secure", cmdPrefix)
	if cmd.Flags().Changed(secureFlagName) {

		var secureFlagName string
		if cmdPrefix == "" {
			secureFlagName = "secure"
		} else {
			secureFlagName = fmt.Sprintf("%v.secure", cmdPrefix)
		}

		secureFlagValue, err := cmd.Flags().GetBool(secureFlagName)
		if err != nil {
			return err, false
		}
		m.Secure = secureFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCookieTypeFlags(depth int, m *models.Cookie, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for Domain

// register flags to command
func registerModelDomainFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDomainValue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDomainValue(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	valueDescription := ``

	var valueFlagName string
	if cmdPrefix == "" {
		valueFlagName = "value"
	} else {
		valueFlagName = fmt.Sprintf("%v.value", cmdPrefix)
	}

	var valueFlagDefault string

	_ = cmd.PersistentFlags().String(valueFlagName, valueFlagDefault, valueDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDomainFlags(depth int, m *models.Domain, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, valueAdded := retrieveDomainValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || valueAdded

	return nil, retAdded
}

func retrieveDomainValueFlags(depth int, m *models.Domain, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	valueFlagName := fmt.Sprintf("%v.value", cmdPrefix)
	if cmd.Flags().Changed(valueFlagName) {

		var valueFlagName string
		if cmdPrefix == "" {
			valueFlagName = "value"
		} else {
			valueFlagName = fmt.Sprintf("%v.value", cmdPrefix)
		}

		valueFlagValue, err := cmd.Flags().GetString(valueFlagName)
		if err != nil {
			return err, false
		}
		m.Value = valueFlagValue

		retAdded = true
	}

	return nil, retAdded
}
