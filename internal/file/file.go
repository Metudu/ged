package file

import (
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Location     string `gorm:"unique;not null"`
	LauncherName string `gorm:"not null"`
	Visibility   bool   `gorm:"not null"`
}

func GetDesktopFiles() []string {
	var f []string
	// Loop through all the dirs with desktop files, use colon to separate the paths
	dirs := strings.Split(os.Getenv("XDG_DATA_DIRS"), ":")
	for _, dir := range dirs {
		dir = path.Join(dir, "applications")
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		// Loop through the files
		for _, file := range files {
			// Check if the file is a .desktop file
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".desktop") {
				f = append(f, path.Join(dir, file.Name()))
			}
		}
	}

	return f
}

func GetFileProperties(location string) File {
	var file File
	file.Location = location
	cfg, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, location)
	if err != nil {
		log.Fatal(err)
	}

	file.LauncherName = cfg.Section("Desktop Entry").Key("Name").Value()
	if cfg.Section("Desktop Entry").HasKey("NoDisplay") && cfg.Section("Desktop Entry").Key("NoDisplay").Value() == "false" {
		file.Visibility = true
	} else {
		file.Visibility = false
	}

	return file
}

func (f File) ChangeVisibility(visibility bool) error {
	if f.Visibility == visibility {
		log.Printf("Visibility is already set to %v!\n", f.Visibility)
		return nil
	}

	cfg, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, f.Location)
	if err != nil {
		return err
	}

	section := cfg.Section("Desktop Entry")

	switch visibility {
	case true && section.HasKey("NoDisplay"):
		section.Key("NoDisplay").SetValue("false")

	case false && section.HasKey("NoDisplay"):
		section.Key("NoDisplay").SetValue("true")
	case false && !section.HasKey("NoDisplay"):
		if _, err := section.NewKey("NoDisplay", "true"); err != nil {
			return err
		}
	default:
	}

	cfg.SaveTo(f.Location)
	return nil
}
