package main

type PlaylistItemListResponse struct {
	Kind          string          `json:"kind"`
	Etag          string          `json:"etag"`
	NextPageToken string          `json:"nextPageToken"`
	PrevPageToken string          `json:"prevPageToken"`
	PageInfo      PageInfo        `json:"pageInfo"`
	Items         []PlaylistItems `json:"items"`
}

type PageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
}

type PlaylistItems struct {
	Kind    string  `json:"kind"`
	Etag    string  `json:"etag"`
	Id      string  `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type Snippet struct {
	PublishedAt            string         `json:"publishedAt"`
	ChannelId              string         `json:"channelId"`
	Title                  string         `json:"title"`
	Description            string         `json:"description"`
	ChannelTitle           string         `json:"channelTitle"`
	VideoOwnerChannelTitle string         `json:"videoOwnerChannelTitle"`
	VideoOwnerChannelId    string         `json:"videoOwnerChannelId"`
	PlaylistId             string         `json:"playlistId"`
	Position               uint64         `json:"position"`
	ResourceId             ResourceId     `json:"resourceId"`
	ContentDetails         ContentDetails `json:"contentDetails"`
	Status                 Status         `json:"status"`
}

type ResourceId struct {
	Kind    string `json:"kind"`
	VideoId string `json:"videoId"`
}

type ContentDetails struct {
	VideoId          string `json:"videoId"`
	StartAt          string `json:"startAt"`
	EndAt            string `json:"endAt"`
	Note             string `json:"note"`
	VideoPublishedAt string `json:"videoPublishedAt"`
}

type Status struct {
	PrivacyStatus string `json:"privacyStatus"`
}

type PlaylistListResponse struct {
	Kind     string     `json:"kind"`
	Etag     string     `json:"etag"`
	PageInfo PageInfo   `json:"pageInfo"`
	Items    []Playlist `json:"items"`
}

type Playlist struct {
	Kind    string          `json:"kind"`
	Etag    string          `json:"etag"`
	Id      string          `json:"id"`
	Snippet SnippetPlaylist `json:"snippet"`
}

type SnippetPlaylist struct {
	PublishedAt  string    `json:"publishedAt"`
	ChannelId    string    `json:"channelId"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	ChannelTitle string    `json:"channelTitle"`
	Localized    Localized `json:"localized"`
}

type Localized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
