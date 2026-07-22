package adcom1

// PlaybackMethod represents media playback methods.
type PlaybackMethod int8

// Media playback methods.
const (
	PlaybackPageLoadSoundOn  PlaybackMethod = 1  // Initiates without any specific user interaction with Sound On
	PlaybackPageLoadSoundOff PlaybackMethod = 2  // Initiates without any specific user interaction Sound Off by Default
	PlaybackClickSoundOn     PlaybackMethod = 3  // Initiates on Click with Sound On
	PlaybackMouseOverSoundOn PlaybackMethod = 4  // Initiates on Mouse-Over or cursor over with Sound On
	PlaybackViewportSoundOn  PlaybackMethod = 5  // Initiates on Entering Viewport or scrolling into view with Sound On
	PlaybackViewportSoundOff PlaybackMethod = 6  // Initiates on Entering Viewport or scrolling into view with Sound Off by Default
	PlaybackContinuous       PlaybackMethod = 7  // Continuous Playback - Media playback is set to play additional media automatically without user interaction. The media player will keep playing additional media (playlist or generated) for the user until the user actively stops this from happening.
	PlaybackPauseSoundOn     PlaybackMethod = 8  // Initiated by user pausing content, Sound on
	PlaybackPauseSoundOff    PlaybackMethod = 9  // Initiated by user pausing content, Sound off
	PlaybackIdlingsoundOn    PlaybackMethod = 10 // Initiated by idling (e.g. screensaver), Sound on
	PlaybackIdlingSoundOff   PlaybackMethod = 11 // Initiated by idling (e.g. screensaver), Sound off
)
