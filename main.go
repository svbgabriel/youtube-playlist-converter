package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"youtube-playlist-converter/config"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Playlist ID not informed!")
	}
	id := args[0]

	configuration := config.ReadConfig()

	playlist, err := readPlaylist(id, configuration.Youtube.Key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Looking for songs in playlist \"%s\"\n", playlist.Items[0].Snippet.Title)

	titleList, err := readPlaylistItems(id, configuration.Youtube.Key)
	if err != nil {
		log.Fatal(err)
	}

	token := initiateAuth(configuration)

	currentUser, err := getCurrentUser(token)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Using the user %s\n", currentUser)

	title := playlist.Items[0].Snippet.Title + " by " + playlist.Items[0].Snippet.ChannelTitle
	description := playlist.Items[0].Snippet.Description

	in := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter a name for the playlist in Spotify [%s]:", title)
	newTitle, _ := in.ReadString('\n')
	newTitle = strings.TrimSpace(newTitle)
	if len(newTitle) > 0 {
		title = newTitle
	}
	fmt.Printf("Enter a description for the playlist [%s]:", description)
	newDescription, _ := in.ReadString('\n')
	newDescription = strings.TrimSpace(newDescription)
	if len(newDescription) > 0 {
		description = newDescription
	}

	playlistID, err := createPlaylist(currentUser, title, description, token)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The playlist %s was created", title)

	for _, item := range titleList {
		tracks, _ := searchTrack(item, token)
		for _, t := range tracks {
			_ = addItemPlaylist(playlistID, t.ID, token)
		}
	}
}
