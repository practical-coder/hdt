// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/haproxytech/client-native/v4/models"
	"github.com/spf13/cobra"
)

// Schema cli for SslCertificate

// register flags to command
func registerModelSslCertificateFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSslCertificateDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSslCertificateFile(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSslCertificateStorageName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSslCertificateDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerSslCertificateFile(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	fileDescription := ``

	var fileFlagName string
	if cmdPrefix == "" {
		fileFlagName = "file"
	} else {
		fileFlagName = fmt.Sprintf("%v.file", cmdPrefix)
	}

	var fileFlagDefault string

	_ = cmd.PersistentFlags().String(fileFlagName, fileFlagDefault, fileDescription)

	return nil
}

func registerSslCertificateStorageName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelSslCertificateFlags(depth int, m *models.SslCertificate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrieveSslCertificateDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, fileAdded := retrieveSslCertificateFileFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fileAdded

	err, storageNameAdded := retrieveSslCertificateStorageNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || storageNameAdded

	return nil, retAdded
}

func retrieveSslCertificateDescriptionFlags(depth int, m *models.SslCertificate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveSslCertificateFileFlags(depth int, m *models.SslCertificate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fileFlagName := fmt.Sprintf("%v.file", cmdPrefix)
	if cmd.Flags().Changed(fileFlagName) {

		var fileFlagName string
		if cmdPrefix == "" {
			fileFlagName = "file"
		} else {
			fileFlagName = fmt.Sprintf("%v.file", cmdPrefix)
		}

		fileFlagValue, err := cmd.Flags().GetString(fileFlagName)
		if err != nil {
			return err, false
		}
		m.File = fileFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSslCertificateStorageNameFlags(depth int, m *models.SslCertificate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
