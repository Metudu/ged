package db

import (
	"gorm.io/gorm"

	f "github.com/Metudu/ged/internal/file"
)

var (
	db  *gorm.DB
	err error
)

func Initialize() {
	// First removes the database completely, then recreates it.
	removeDB()
	connect()

	files := f.GetDesktopFiles()
	for _, file := range files {
		f := f.GetFileProperties(file)
		insert(f)
	}
}

func Refresh() {
	// Only deletes the records, then adds them again.
	connect()
	deleteRecords()

	files := f.GetDesktopFiles()
	for _, file := range files {
		f := f.GetFileProperties(file)
		insert(f)
	}
}

func List() {
	connect()
	list()
}

func GetNames() []string {
	return getFileNames()
}

func Show(name string) f.File {
	file := getRecordByName(name)
	update(file, true)

	return file
}

func Hide(name string) f.File {
	file := getRecordByName(name)
	update(file, true)

	return file
}
