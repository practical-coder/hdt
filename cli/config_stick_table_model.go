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

// Schema cli for ConfigStickTable

// register flags to command
func registerModelConfigStickTableFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerConfigStickTableExpire(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerConfigStickTableKeylen(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerConfigStickTableNopurge(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerConfigStickTablePeers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerConfigStickTableSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerConfigStickTableStore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerConfigStickTableType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerConfigStickTableExpire(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	expireDescription := ``

	var expireFlagName string
	if cmdPrefix == "" {
		expireFlagName = "expire"
	} else {
		expireFlagName = fmt.Sprintf("%v.expire", cmdPrefix)
	}

	var expireFlagDefault int64

	_ = cmd.PersistentFlags().Int64(expireFlagName, expireFlagDefault, expireDescription)

	return nil
}

func registerConfigStickTableKeylen(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	keylenDescription := ``

	var keylenFlagName string
	if cmdPrefix == "" {
		keylenFlagName = "keylen"
	} else {
		keylenFlagName = fmt.Sprintf("%v.keylen", cmdPrefix)
	}

	var keylenFlagDefault int64

	_ = cmd.PersistentFlags().Int64(keylenFlagName, keylenFlagDefault, keylenDescription)

	return nil
}

func registerConfigStickTableNopurge(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nopurgeDescription := ``

	var nopurgeFlagName string
	if cmdPrefix == "" {
		nopurgeFlagName = "nopurge"
	} else {
		nopurgeFlagName = fmt.Sprintf("%v.nopurge", cmdPrefix)
	}

	var nopurgeFlagDefault bool

	_ = cmd.PersistentFlags().Bool(nopurgeFlagName, nopurgeFlagDefault, nopurgeDescription)

	return nil
}

func registerConfigStickTablePeers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	peersDescription := ``

	var peersFlagName string
	if cmdPrefix == "" {
		peersFlagName = "peers"
	} else {
		peersFlagName = fmt.Sprintf("%v.peers", cmdPrefix)
	}

	var peersFlagDefault string

	_ = cmd.PersistentFlags().String(peersFlagName, peersFlagDefault, peersDescription)

	return nil
}

func registerConfigStickTableSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sizeDescription := ``

	var sizeFlagName string
	if cmdPrefix == "" {
		sizeFlagName = "size"
	} else {
		sizeFlagName = fmt.Sprintf("%v.size", cmdPrefix)
	}

	var sizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sizeFlagName, sizeFlagDefault, sizeDescription)

	return nil
}

func registerConfigStickTableStore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	storeDescription := ``

	var storeFlagName string
	if cmdPrefix == "" {
		storeFlagName = "store"
	} else {
		storeFlagName = fmt.Sprintf("%v.store", cmdPrefix)
	}

	var storeFlagDefault string

	_ = cmd.PersistentFlags().String(storeFlagName, storeFlagDefault, storeDescription)

	return nil
}

func registerConfigStickTableType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["ip","ipv6","integer","string","binary"]. `

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
			if err := json.Unmarshal([]byte(`["ip","ipv6","integer","string","binary"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelConfigStickTableFlags(depth int, m *models.ConfigStickTable, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, expireAdded := retrieveConfigStickTableExpireFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || expireAdded

	err, keylenAdded := retrieveConfigStickTableKeylenFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || keylenAdded

	err, nopurgeAdded := retrieveConfigStickTableNopurgeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nopurgeAdded

	err, peersAdded := retrieveConfigStickTablePeersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || peersAdded

	err, sizeAdded := retrieveConfigStickTableSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeAdded

	err, storeAdded := retrieveConfigStickTableStoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || storeAdded

	err, typeAdded := retrieveConfigStickTableTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveConfigStickTableExpireFlags(depth int, m *models.ConfigStickTable, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	expireFlagName := fmt.Sprintf("%v.expire", cmdPrefix)
	if cmd.Flags().Changed(expireFlagName) {

		var expireFlagName string
		if cmdPrefix == "" {
			expireFlagName = "expire"
		} else {
			expireFlagName = fmt.Sprintf("%v.expire", cmdPrefix)
		}

		expireFlagValue, err := cmd.Flags().GetInt64(expireFlagName)
		if err != nil {
			return err, false
		}
		m.Expire = &expireFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveConfigStickTableKeylenFlags(depth int, m *models.ConfigStickTable, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	keylenFlagName := fmt.Sprintf("%v.keylen", cmdPrefix)
	if cmd.Flags().Changed(keylenFlagName) {

		var keylenFlagName string
		if cmdPrefix == "" {
			keylenFlagName = "keylen"
		} else {
			keylenFlagName = fmt.Sprintf("%v.keylen", cmdPrefix)
		}

		keylenFlagValue, err := cmd.Flags().GetInt64(keylenFlagName)
		if err != nil {
			return err, false
		}
		m.Keylen = &keylenFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveConfigStickTableNopurgeFlags(depth int, m *models.ConfigStickTable, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nopurgeFlagName := fmt.Sprintf("%v.nopurge", cmdPrefix)
	if cmd.Flags().Changed(nopurgeFlagName) {

		var nopurgeFlagName string
		if cmdPrefix == "" {
			nopurgeFlagName = "nopurge"
		} else {
			nopurgeFlagName = fmt.Sprintf("%v.nopurge", cmdPrefix)
		}

		nopurgeFlagValue, err := cmd.Flags().GetBool(nopurgeFlagName)
		if err != nil {
			return err, false
		}
		m.Nopurge = nopurgeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveConfigStickTablePeersFlags(depth int, m *models.ConfigStickTable, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	peersFlagName := fmt.Sprintf("%v.peers", cmdPrefix)
	if cmd.Flags().Changed(peersFlagName) {

		var peersFlagName string
		if cmdPrefix == "" {
			peersFlagName = "peers"
		} else {
			peersFlagName = fmt.Sprintf("%v.peers", cmdPrefix)
		}

		peersFlagValue, err := cmd.Flags().GetString(peersFlagName)
		if err != nil {
			return err, false
		}
		m.Peers = peersFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveConfigStickTableSizeFlags(depth int, m *models.ConfigStickTable, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sizeFlagName := fmt.Sprintf("%v.size", cmdPrefix)
	if cmd.Flags().Changed(sizeFlagName) {

		var sizeFlagName string
		if cmdPrefix == "" {
			sizeFlagName = "size"
		} else {
			sizeFlagName = fmt.Sprintf("%v.size", cmdPrefix)
		}

		sizeFlagValue, err := cmd.Flags().GetInt64(sizeFlagName)
		if err != nil {
			return err, false
		}
		m.Size = &sizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveConfigStickTableStoreFlags(depth int, m *models.ConfigStickTable, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	storeFlagName := fmt.Sprintf("%v.store", cmdPrefix)
	if cmd.Flags().Changed(storeFlagName) {

		var storeFlagName string
		if cmdPrefix == "" {
			storeFlagName = "store"
		} else {
			storeFlagName = fmt.Sprintf("%v.store", cmdPrefix)
		}

		storeFlagValue, err := cmd.Flags().GetString(storeFlagName)
		if err != nil {
			return err, false
		}
		m.Store = storeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveConfigStickTableTypeFlags(depth int, m *models.ConfigStickTable, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
