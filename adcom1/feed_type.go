package adcom1

// FeedType represents types of feeds for audio.
type FeedType int8

// Types of feeds for audio content.
const (
	FeedMusicService   FeedType = 1 // AOD - Music Streaming Service
	FeedRadioBroadcast FeedType = 2 // LIVE - FM/AM broadcast: Live content broadcast over the air but also available via online streaming.
	FeedPodcast        FeedType = 3 // AOD - Podcast: Original, pre-recorded content distributed as episodes in a series.
	FeedCatchUpRadio   FeedType = 4 // AOD - Catch-up Radio: recorded segment of a radio show that was originally broadcast live.
	FeedWebRadio       FeedType = 5 // LIVE - Web Radio: Live content only available via online streaming, not as AM/FM broadcast.
	FeedVideoGame      FeedType = 6 // MISC - Video Game: Background audio in video games.
	FeedTextToSpeech   FeedType = 7 // MISC - Text To Speech: Audio books, website plugin that can read article.
)
