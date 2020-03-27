package main

import (
	"fmt"
	"path/filepath"
)

func feedDBwVideo() (result bool) {
	matches, err := filepath.Glob(`/app/data/*.mp4`)
	if err == nil {
		for _, match := range matches {
			var videoA Video
			videoA.NameVideo = filepath.ToSlash(match)
			videoA.PathVideo = filepath.ToSlash(match)
			videoA.DescVideo = filepath.ToSlash(match)
			if insertVideoIntoDB(videoA) {
				fmt.Println("Success adding a new video...!")
			} else {
				fmt.Println("Something went wrong adding a new video...")
			}
		}
		return true
	}
	fmt.Println("There's no file inside the volume folder.")
	return false
}
