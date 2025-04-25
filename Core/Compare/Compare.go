package Compare

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	Backups "github.com/XanaOG/Cleaner/Core/Backups"
	Client "github.com/XanaOG/Cleaner/Core/Client"
	Copy "github.com/XanaOG/Cleaner/Core/Copy"
)

func FileSizes(directoryPath string) (deleted, renamed int) {
	Config := Client.GetConfig(Client.ConfigFile)

	err := Backups.Clean()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err1 := os.MkdirAll(Config.Cleaner.BackupDirectory, os.ModePerm)
	if err1 != nil {
		fmt.Printf("Error creating backup directory: %v\n", err)
		return
	}
	err2 := os.MkdirAll(Config.Cleaner.SaveDirectory, os.ModePerm)
	if err2 != nil {
		fmt.Printf("Error creating backup directory: %v\n", err)
		return
	}

	files, err := os.ReadDir(directoryPath)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return 0, 0
	}

	if len(files) == 0 {
		fmt.Println("No files found in the save directory.")
		return 0, 0
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		baseName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		if !strings.HasSuffix(file.Name(), ".arktribe") {
			continue
		}

		arkFilePath := filepath.Join(directoryPath, file.Name())
		backupFilePath := filepath.Join(directoryPath, baseName+".tribebak")
		backupDirFilePath := filepath.Join(Config.Cleaner.BackupDirectory, file.Name())

		arkStats, err := file.Info()
		if err != nil {
			fmt.Printf("Error getting file info for %s: %v\n", file.Name(), err)
			continue
		}

		backupStats, err := os.Stat(backupFilePath)
		if err != nil {
			fmt.Printf("Error getting file info for %s: %v\n", backupFilePath, err)
			continue
		}

		if arkStats.Size() != backupStats.Size() {
			err := Copy.File(arkFilePath, backupDirFilePath)
			if err != nil {
				fmt.Printf("Error backing up file %s: %v\n", arkFilePath, err)
				continue
			}

			err = os.Remove(arkFilePath)
			if err != nil {
				fmt.Printf("Error deleting file %s: %v\n", file.Name(), err)
				continue
			}

			err = os.Rename(backupFilePath, arkFilePath)
			if err != nil {
				fmt.Printf("Error renaming file %s to %s: %v\n", backupFilePath, file.Name(), err)
				continue
			}

			deleted++
			renamed++
		}
	}

	fmt.Printf("Completed: %d files deleted, %d files renamed.\n", deleted, renamed)
	return deleted, renamed
}
