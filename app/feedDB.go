package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func feedDBwVideo() (result bool) {
	//`C:\data\*.mp4`
	matches, err := filepath.Glob(`/app/data/*.mp4`)
	if err == nil {
		for _, match := range matches {
			path := filepath.ToSlash(match)
			formatedName := strings.Replace(path, "/app/data/", "", -1)
			// /app /..
			formatedPath := strings.Replace(path, "/app", "..", -1)

			var videoA Video
			videoA.NameVideo = formatedName
			videoA.PathVideo = formatedPath
			videoA.DescVideo = formatedPath
			if !insertVideoIntoDB(videoA) {
				fmt.Println("Something went wrong adding a new video...")
			}
		}
		return true
	}
	fmt.Println("There's no file inside the volume folder.")
	return false
}
