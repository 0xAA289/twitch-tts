package main

var TTS_DATA = TWITCH_TTS_STRUCT{
	Channel:  "none",
	Language: "English",
	Whitelist: TWITCH_WHITELIST{
		Enabled: false,
		List:    []string{},
	},
}

var RUNTIME = TWITCH_RUNTIME_STRUCT{}
