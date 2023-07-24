package runner

import (
	"github.com/khulnasoft-labs/gologger"
	updateutils "github.com/khulnasoft-labs/utils/update"
)

const banner = `
       __        ________        __       
  ___ / /  __ __/ _/ _/ /__  ___/ /__ ___
 (_-</ _ \/ // / _/ _/ / -_)/ _  / _ \(_-<
/___/_//_/\_,_/_//_//_/\__/ \_,_/_//_/___/
`

// version is the current version of shuffledns
const version = `v1.0.9`

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tkhulnasoft-labs.io\n\n")
}

// GetUpdateCallback returns a callback function that updates shuffledns
func GetUpdateCallback() func() {
	return func() {
		showBanner()
		updateutils.GetUpdateToolCallback("shuffledns", version)()
	}
}
