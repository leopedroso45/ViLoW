package main

import (
	"log"
	"path/filepath"
	"strings"
)

/*feedDBwVideo Feeds the database with volume videos
and returns a true if no error occurs.
*/
func feedDBwVideo() (result bool) {
	matches, err := filepath.Glob(`/app/data/*.mp4`)
	if err == nil {
		for _, match := range matches {
			path := filepath.ToSlash(match)

			staticPath := "http://192.168.43.216:8000"
			formatedName := strings.Replace(path, "/app/data/", "", -1)
			formatedPath := strings.Replace(path, "/app", staticPath, -1)

			//
			formatedName = strings.Replace(formatedName, "'", "", -1)
			formatedPath = strings.Replace(formatedPath, "'", "", -1)

			videoA := VideoConstructor(formatedName, formatedName, formatedPath)

			if !insertVideoIntoDB(videoA) {
				log.Println("Something went wrong adding a new video...")
			}
		}
		return true
	}
	log.Println("There's no file inside the volume folder.")
	return false
}
