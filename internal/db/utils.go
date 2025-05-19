package db

import (
	"log"
	"os"
	"path"

	"github.com/Metudu/ged/internal/file"
	"github.com/olekukonko/tablewriter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connect() {
	homeDir, _ := os.UserHomeDir()
	db, err = gorm.Open(
		sqlite.Open(
			path.Join(
				homeDir, os.Getenv("GED_DIR"), os.Getenv("GED_DB_FILENAME"),
			),
		),
		&gorm.Config{},
	)

	db.AutoMigrate(&file.File{})

	if err != nil {
		log.Fatal("Error connecting the database: ", err)
	}
}

func insert(file file.File) {
	db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&file).Error
	})
}

func removeDB() {
	homeDir, _ := os.UserHomeDir()
	_ = os.Remove(path.Join(homeDir, os.Getenv("GED_DIR"), os.Getenv("GED_DB_FILENAME")))
}

func deleteRecords() {
	db.Transaction(func(tx *gorm.DB) error {
		return tx.Exec("DELETE FROM files").Error
	})
}

func list() {
	var files []file.File
	var data []struct {
		Name       string
		Visibility bool
	}
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Select("launcher_name", "visibility").Find(&files).Error; err != nil {
			return err
		}
		return nil
	})

	for _, file := range files {
		data = append(data, struct {
			Name       string
			Visibility bool
		}{
			Name:       file.LauncherName,
			Visibility: file.Visibility,
		})
	}

	table := tablewriter.NewTable(os.Stdout)
	table.Header("Name", "Visibility")
	table.Bulk(data)
	table.Render()
}

func getFileNames() []string {
	var files []file.File
	var data []string
	if err := db.Transaction(func(tx *gorm.DB) error {
		return tx.Select("launcher_name").Find(&files).Error
	}); err != nil {
		panic(err)
	}

	for _, file := range files {
		data = append(data, file.LauncherName)
	}

	return data
}

func getRecordByName(name string) file.File {
	var f file.File
	if err := db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("launcher_name = ?", name).First(&f).Error
	}); err != nil {
		panic(err)
	}

	return f
}

func update(f file.File, visibility bool) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := f.ChangeVisibility(visibility); err != nil {
			return err
		}
		if err := tx.Model(&f).Update("visibility", visibility).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		panic(err)
	}
}
