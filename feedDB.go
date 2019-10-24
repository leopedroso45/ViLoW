package main

import (
	"fmt"
	"path/filepath"
)

func feedDBwVideo() {
	matches, _ := filepath.Glob(`D:\Videos\Test\*.mp4`)
	for _, match := range matches {
		var videoA Video
		videoA.NameVideo = filepath.ToSlash(match)
		videoA.PathVideo = filepath.ToSlash(match)
		videoA.DescVideo = filepath.ToSlash(match)
		if insertVideoIntoDB(videoA) {
			fmt.Println("Success adding a new video...")
		} else {
			fmt.Println("Something went wrong adding a new video...")
		}
	}
}
