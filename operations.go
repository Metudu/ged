package main

import (
	"errors"
	"os"
	"path"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	// desktopFilesLocation is where the .desktop files are located
	desktopFilesLocation = path.Join(os.Getenv("HOME"), ".local/share/applications")
)

// Retrieves the .desktop file names in the $HOME/.local/share/applications directory if this directory exists.
func GetDesktopFiles() ([]string, error) {
	if checkIfDirectoryExists(desktopFilesLocation) {
		entries, err := os.ReadDir(desktopFilesLocation)
		if err != nil {	
			return []string{}, err
		}

		var files []string
		for _, val := range entries {
			// Add only the .desktop files
			if strings.Contains(val.Name(), ".desktop") {
				files = append(files, val.Name())	
			}
		}

		return files, nil
	}

	return []string{}, errors.New("directory does not exist")
}

func SetVisibility(filename string, visible bool) error {
	if filename == "" {
		return errors.New("filename cannot be empty")
	}

	cfg, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, path.Join(desktopFilesLocation, filename))
	if err != nil {
		return err
	}

	if visible {
		if cfg.Section("Desktop Entry").HasKey("NoDisplay") {
			cfg.Section("Desktop Entry").DeleteKey("NoDisplay")
		}
	} else {
		if !cfg.Section("Desktop Entry").HasKey("NoDisplay") {
			if _, err = cfg.Section("Desktop Entry").NewKey("NoDisplay", "true"); err != nil {
				return err
			}
		} else {
			cfg.Section("Desktop Entry").Key("NoDisplay").SetValue("true")
		}
	}
	
	if err = cfg.SaveTo(path.Join(desktopFilesLocation, filename)); err != nil {
		return err
	}

	return nil
}

func GetVisibility(filename string) bool {
	if filename == "" {
		return false
	}

	cfg, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, path.Join(desktopFilesLocation, filename))
	if err != nil {
		return false
	}

	if cfg.Section("Desktop Entry").HasKey("NoDisplay") && cfg.Section("Desktop Entry").Key("NoDisplay").String() == "true" {
		return false
	}

	return true
}

func checkIfDirectoryExists(directory string) bool {
	if directory == "" {
		return false
	}

	_, err := os.Stat(directory)
	return !os.IsNotExist(err)
}