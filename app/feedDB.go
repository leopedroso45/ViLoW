package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func feedDBwVideo() (result bool) {
	matches, err := filepath.Glob(`/app/data/*.mp4`)
	if err == nil {
		for _, match := range matches {
			path := filepath.ToSlash(match)
			formatedName := strings.Replace(path, "/app/data/", "", -1)
			staticPath := "http://192.168.43.216:8000"
			fmt.Println(staticPath)
			// /app /..
			//http://localhost:8000/static/data/FINNEAS%20-%20I%20Lost%20A%20Friend%20(Official%20Video).mp4
			formatedPath := strings.Replace(path, "/app", staticPath, -1)

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
