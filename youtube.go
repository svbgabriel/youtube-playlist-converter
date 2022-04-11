package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func readPlaylistItems(id string, key string) ([]string, error) {
	url := "https://www.googleapis.com/youtube/v3/playlistItems?key=%s&part=snippet&playlistId=%s&maxResults=25"

	titleList := make([]string, 0)
	nextPageToken := ""

	for {
		resp, err := http.Get(fmt.Sprintf(url, key, id))
		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		playlistItems := PlaylistItemListResponse{}
		if err := json.Unmarshal(body, &playlistItems); err != nil {
			return nil, err
		}

		for _, item := range playlistItems.Items {
			title := cleanTitle(item.Snippet.Title)
			titleList = append(titleList, title)
		}

		nextPageToken = playlistItems.NextPageToken

		if nextPageToken == "" {
			break
		}

		url = fmt.Sprintf("%s&pageToken=%s", url, nextPageToken)
	}

	return titleList, nil
}

func readPlaylist(id string, key string) (PlaylistListResponse, error) {
	url := "https://www.googleapis.com/youtube/v3/playlists?key=%s&part=snippet&id=%s"

	url = fmt.Sprintf(url, key, id)

	resp, err := http.Get(url)
	if err != nil {
		return PlaylistListResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PlaylistListResponse{}, err
	}

	playlist := PlaylistListResponse{}
	if err := json.Unmarshal(body, &playlist); err != nil {
		return PlaylistListResponse{}, err
	}

	return playlist, nil
}

func cleanTitle(title string) string {
	regexSquareBrackets, _ := regexp.Compile("(\\[).*?(])")
	regexFeat, _ := regexp.Compile("\\((?i)feat.*?\\)")

	altered := regexSquareBrackets.ReplaceAllString(title, "")
	altered = regexFeat.ReplaceAllString(altered, "")
	altered = strings.Trim(altered, " - ")

	return altered
}
