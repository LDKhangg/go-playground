package functions

import (
	"io"
)

func ReadAndClose(rc io.ReadCloser) ([]byte, error) {
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	return data, nil
}
