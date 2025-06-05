package files

import (
	"os"
)

func ReadFileFromSpecifiedPath(path string) ([]byte, error) {
	return os.ReadFile(path)
}
