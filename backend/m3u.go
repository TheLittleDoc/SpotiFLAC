package backend

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	// import metadata.go
)

func AppendTrackToM3U(playlistFile string, audioFile string) interface{} {
	// check if playlistFile exists
	playlistName := ""
	playlistName = strings.TrimSuffix(playlistFile, ".m3u")
	playlistName = strings.TrimSuffix(playlistName, ".M3U")
	playlistName = strings.Split(playlistName, "/")[len(strings.Split(playlistName, "/"))-1]
	if _, err := os.Stat(playlistFile); os.IsNotExist(err) {
		// create the file
		file, err := os.Create(playlistFile)

		if err != nil {
			return err
		}

		file.WriteString("#EXTM3U\n")
		file.WriteString("#PLAYLIST: " + playlistName + "\n")
		file.Close()
	}
	// append audioFile to playlistFile
	file, err := os.OpenFile(playlistFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	trackMetadata, err := ReadMetadata(audioFile)
	if err != nil {
		return err
	}
	_, err = file.WriteString("#EXTINF:" + strconv.Itoa(trackMetadata.SongLength) + "," + trackMetadata.Title + " - " + trackMetadata.Artist + "\n")
	if err != nil {
		return nil
	}
	// strip upper directories
	relativePath := audioFile
	if strings.Contains(audioFile, "/") {
		fmt.Printf("Playlist Name: %s\n", playlistName)
		parts := strings.Split(audioFile, playlistName+"/")
		for _, part := range parts {
			fmt.Printf("Part: %s\n", part)
		}
		relativePath = parts[1]
	}
	_, err = file.WriteString(relativePath + "\n")
	if err != nil {
		return nil
	}
	file.Close()

	return nil
}
