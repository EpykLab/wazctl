package files

import (
	"bytes"
	"os"
)

func FileCreateWithSpecifiedNameAndContent(name string, content bytes.Buffer) error {

	return os.WriteFile(
		name,
		content.Bytes(),
		0644)
}
