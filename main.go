package main

import (
	"log"
	"os"
	"youtube-playlist-converter/config"
)

func main() {
	args := os.Args[1:]
	id := args[0]

	configuration := config.ReadConfig()

	playlist, err := readPlaylist(id, configuration.Youtube.Key)
	if err != nil {
		log.Fatal(err)
	}

	titleList, err := readPlaylistItems(id, configuration.Youtube.Key)
	if err != nil {
		log.Fatal(err)
	}

	token := initiateAuth(configuration)

	currentUser, err := getCurrentUser(token)
	if err != nil {
		log.Fatal(err)
	}

	title := playlist.Items[0].Snippet.Title + " by " + playlist.Items[0].Snippet.ChannelTitle
	description := playlist.Items[0].Snippet.Description

	playlistID, err := createPlaylist(currentUser, title, description, token)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range titleList {
		tracks, _ := searchTrack(item, token)
		for _, t := range tracks {
			_ = addItemPlaylist(playlistID, t.ID, token)
		}
	}
}
