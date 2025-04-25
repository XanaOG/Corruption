package Declared

type Config struct {
	Cleaner struct {
		SaveDirectory   string `json:"savedirectory"`
		BackupDirectory string `json:"backupdirectory"`

		CleanOnRun bool `json:"cleanonrun"`
	}
}
