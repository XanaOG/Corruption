package Backups

import (
	"os"
	"path/filepath"

	"github.com/XanaOG/Cleaner/Core/Client"
)

func Clean() error {
	Config := Client.GetConfig(Client.ConfigFile)
	dir, err := os.Open(Config.Cleaner.BackupDirectory)
	if err != nil {
		return err
	}
	defer dir.Close()

	files, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, d := range files {
		err = os.RemoveAll(filepath.Join(Config.Cleaner.BackupDirectory, d))
		if err != nil {
			return err
		}
	}

	return nil
}
