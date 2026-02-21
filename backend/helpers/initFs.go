package helpers

import "os"

func InitStorage(dir string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return nil
}
