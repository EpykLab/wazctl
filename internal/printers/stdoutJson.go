package printers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	// "github.com/EpykLab/wazctl/internal/bolterr"
)

func PrintJsonFormattedOrError(data []byte, wazctlErr error) {

	if wazctlErr != nil {
		// TODO: Logic for parsing error better, for now defaulting to standard
		log.Println(wazctlErr)
	} else {
		var out bytes.Buffer
		err := json.Indent(&out, data, "", "	")
		if err != nil {
			log.Println(err)
		}

		fmt.Println(out.String())
	}
}
