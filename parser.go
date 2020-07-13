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
	"strings"

	"github.com/arnumina/failure"
	"github.com/arnumina/options"
)

func parseCfgString(cs string) (string, options.Options, error) {
	if cs == "" {
		return "", nil,
			failure.New(nil).
				Msg("the configuration string is empty") ///////////////////////////////////////////////////////////////
	}

	opts := options.New()
	ls := strings.Split(cs, ":")

	if len(ls) != 1 {
		if len(ls) != 2 {
			return "", nil,
				failure.New(nil).
					Set("string", cs).
					Msg("this configuration string is not valid") //////////////////////////////////////////////////////
		}

		for _, opt := range strings.Split(ls[1], ",") {
			kv := strings.Split(opt, "=")
			if len(kv) != 2 {
				return "", nil,
					failure.New(nil).
						Set("string", cs).
						Set("option", opt).
						Msg("this option of this configuration string is not valid") ///////////////////////////////////
			}

			opts[kv[0]] = kv[1]
		}
	}

	return ls[0], opts, nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
