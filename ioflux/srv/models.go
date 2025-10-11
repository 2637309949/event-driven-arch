package main

import "time"

type File struct {
	FileId    int64
	SavePath  string `json:"save_path"`
	NewName   string `json:"new_name"`
	OrigName  string `json:"orig_name"`
	Ext       string `json:"ext"`
	MimeType  string `json:"mime_type"`
	CreatedAt time.Time
}
