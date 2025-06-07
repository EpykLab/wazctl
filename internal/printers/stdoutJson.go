package printers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func PrintJsonFormatted(data []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "	")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(out.String())
}
