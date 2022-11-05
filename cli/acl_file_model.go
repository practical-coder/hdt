// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/spf13/cobra"
)

// Schema cli for ACLFile

// register flags to command
func registerModelACLFileFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerACLFileDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerACLFileID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerACLFileStorageName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerACLFileDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerACLFileID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerACLFileStorageName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	storageNameDescription := ``

	var storageNameFlagName string
	if cmdPrefix == "" {
		storageNameFlagName = "storage_name"
	} else {
		storageNameFlagName = fmt.Sprintf("%v.storage_name", cmdPrefix)
	}

	var storageNameFlagDefault string

	_ = cmd.PersistentFlags().String(storageNameFlagName, storageNameFlagDefault, storageNameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelACLFileFlags(depth int, m *models.ACLFile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrieveACLFileDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, idAdded := retrieveACLFileIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, storageNameAdded := retrieveACLFileStorageNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || storageNameAdded

	return nil, retAdded
}

func retrieveACLFileDescriptionFlags(depth int, m *models.ACLFile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveACLFileIDFlags(depth int, m *models.ACLFile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveACLFileStorageNameFlags(depth int, m *models.ACLFile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	storageNameFlagName := fmt.Sprintf("%v.storage_name", cmdPrefix)
	if cmd.Flags().Changed(storageNameFlagName) {

		var storageNameFlagName string
		if cmdPrefix == "" {
			storageNameFlagName = "storage_name"
		} else {
			storageNameFlagName = fmt.Sprintf("%v.storage_name", cmdPrefix)
		}

		storageNameFlagValue, err := cmd.Flags().GetString(storageNameFlagName)
		if err != nil {
			return err, false
		}
		m.StorageName = storageNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}
