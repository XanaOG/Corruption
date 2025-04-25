package main

import (
	Client "github.com/XanaOG/Cleaner/Core/Client"
	Compare "github.com/XanaOG/Cleaner/Core/Compare"
)

func main() {
	Config := Client.GetConfig(Client.ConfigFile)
	Compare.FileSizes(Config.Cleaner.SaveDirectory)
}
