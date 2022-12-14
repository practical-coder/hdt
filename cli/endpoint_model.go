// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/haproxytech/models"
	"github.com/spf13/cobra"
)

// Schema cli for Endpoint

// register flags to command
func registerModelEndpointFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEndpointDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointTitle(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEndpointDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := `Endpoint description`

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

func registerEndpointTitle(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	titleDescription := `Endpoint title`

	var titleFlagName string
	if cmdPrefix == "" {
		titleFlagName = "title"
	} else {
		titleFlagName = fmt.Sprintf("%v.title", cmdPrefix)
	}

	var titleFlagDefault string

	_ = cmd.PersistentFlags().String(titleFlagName, titleFlagDefault, titleDescription)

	return nil
}

func registerEndpointURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	urlDescription := `Path to the endpoint`

	var urlFlagName string
	if cmdPrefix == "" {
		urlFlagName = "url"
	} else {
		urlFlagName = fmt.Sprintf("%v.url", cmdPrefix)
	}

	var urlFlagDefault string

	_ = cmd.PersistentFlags().String(urlFlagName, urlFlagDefault, urlDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEndpointFlags(depth int, m *models.Endpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrieveEndpointDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, titleAdded := retrieveEndpointTitleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || titleAdded

	err, urlAdded := retrieveEndpointURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || urlAdded

	return nil, retAdded
}

func retrieveEndpointDescriptionFlags(depth int, m *models.Endpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveEndpointTitleFlags(depth int, m *models.Endpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	titleFlagName := fmt.Sprintf("%v.title", cmdPrefix)
	if cmd.Flags().Changed(titleFlagName) {

		var titleFlagName string
		if cmdPrefix == "" {
			titleFlagName = "title"
		} else {
			titleFlagName = fmt.Sprintf("%v.title", cmdPrefix)
		}

		titleFlagValue, err := cmd.Flags().GetString(titleFlagName)
		if err != nil {
			return err, false
		}
		m.Title = titleFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointURLFlags(depth int, m *models.Endpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	urlFlagName := fmt.Sprintf("%v.url", cmdPrefix)
	if cmd.Flags().Changed(urlFlagName) {

		var urlFlagName string
		if cmdPrefix == "" {
			urlFlagName = "url"
		} else {
			urlFlagName = fmt.Sprintf("%v.url", cmdPrefix)
		}

		urlFlagValue, err := cmd.Flags().GetString(urlFlagName)
		if err != nil {
			return err, false
		}
		m.URL = urlFlagValue

		retAdded = true
	}

	return nil, retAdded
}
