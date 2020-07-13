/*
#######
##                      ____
##       _______  ___  / _(_)__ _
##      / __/ _ \/ _ \/ _/ / _ `/
##      \__/\___/_//_/_//_/\_, /
##                        /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/arnumina/failure"
	"github.com/arnumina/options"
	"gopkg.in/yaml.v3"
)

func readFile(opts options.Options) ([]byte, error) {
	filename, err := opts.String("file")
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func loadJSONFile(opts options.Options) (interface{}, error) {
	content, err := readFile(opts)
	if err != nil {
		return nil, err
	}

	var value interface{}

	if err := json.Unmarshal(content, &value); err != nil {
		return nil, err
	}

	return value, nil
}

func loadYAMLFile(opts options.Options) (interface{}, error) {
	content, err := readFile(opts)
	if err != nil {
		return nil, err
	}

	var value interface{}

	if err := yaml.Unmarshal(content, &value); err != nil {
		return nil, err
	}

	return value, nil
}

// Load AFAIRE.
func Load(cfgString string) (interface{}, error) {
	t, opts, err := parseCfgString(cfgString)
	if err != nil {
		return nil, err
	}

	switch t {
	case "empty":
		return nil, nil
	case "json":
		return loadJSONFile(opts)
	case "yaml":
		return loadYAMLFile(opts)
	default:
		return nil,
			failure.New(nil).
				Set("type", t).
				Msg("there is no configuration loader for this type") //////////////////////////////////////////////////
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
