package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/zmb3/spotify/v2"
	spotifyAuth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"youtube-playlist-converter/config"
)

const redirectURI = "http://localhost:8888/callback"

var (
	auth  *spotifyAuth.Authenticator
	ch    = make(chan *oauth2.Token)
	state string
)

func initiateAuth(configApp config.Configurations) *oauth2.Token {
	// first start an HTTP server
	http.HandleFunc("/callback", completeAuth)
	go func() {
		err := http.ListenAndServe(":8888", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	auth = spotifyAuth.New(
		spotifyAuth.WithRedirectURL(redirectURI),
		spotifyAuth.WithScopes(spotifyAuth.ScopeUserReadPrivate, spotifyAuth.ScopePlaylistModifyPrivate),
		spotifyAuth.WithClientID(configApp.Spotify.ClientId),
		spotifyAuth.WithClientSecret(configApp.Spotify.ClientSecret))

	state = uuid.NewString()

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	token := <-ch

	fmt.Println("Logged!")

	return token
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	token, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	ch <- token

	_, _ = fmt.Fprintf(w, "Login Completed! You can close this tab")
}

func searchTrack(title string, token *oauth2.Token) ([]spotify.FullTrack, error) {
	ctx := context.Background()

	httpClient := spotifyAuth.New(
		spotifyAuth.WithRedirectURL(redirectURI),
		spotifyAuth.WithScopes(spotifyAuth.ScopeUserReadPrivate)).Client(ctx, token)

	client := spotify.New(httpClient)

	results, err := client.Search(ctx, title, spotify.SearchTypeTrack, spotify.Limit(20))
	if err != nil {
		return nil, err
	}

	return results.Tracks.Tracks, nil
}

func getCurrentUser(token *oauth2.Token) (string, error) {
	ctx := context.Background()

	httpClient := spotifyAuth.New(
		spotifyAuth.WithRedirectURL(redirectURI),
		spotifyAuth.WithScopes(spotifyAuth.ScopeUserReadPrivate)).Client(ctx, token)

	client := spotify.New(httpClient)

	result, err := client.CurrentUser(ctx)
	if err != nil {
		return "", err
	}

	return result.ID, nil
}

func createPlaylist(userID string, playlistName string, description string, token *oauth2.Token) (spotify.ID, error) {
	ctx := context.Background()

	httpClient := spotifyAuth.New(
		spotifyAuth.WithRedirectURL(redirectURI),
		spotifyAuth.WithScopes(spotifyAuth.ScopePlaylistModifyPrivate)).Client(ctx, token)

	client := spotify.New(httpClient)

	result, err := client.CreatePlaylistForUser(ctx, userID, playlistName, description, false, false)
	if err != nil {
		return "", err
	}

	return result.ID, nil
}

func addItemPlaylist(playlistID spotify.ID, trackID spotify.ID, token *oauth2.Token) error {
	ctx := context.Background()

	httpClient := spotifyAuth.New(
		spotifyAuth.WithRedirectURL(redirectURI),
		spotifyAuth.WithScopes(spotifyAuth.ScopePlaylistModifyPrivate)).Client(ctx, token)

	client := spotify.New(httpClient)

	_, err := client.AddTracksToPlaylist(ctx, playlistID, trackID)

	return err
}
