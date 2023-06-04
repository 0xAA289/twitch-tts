package main

import (
	"fyne.io/fyne/v2/widget"
	"github.com/gempir/go-twitch-irc/v4"
	htgotts "github.com/hegedustibor/htgo-tts"
)

type TWITCH_WHITELIST struct {
	Enabled bool
	List    []string
}

type TWITCH_RUNTIME_STRUCT struct {
	client          *twitch.Client
	nameWidget      *widget.Entry
	languageWidget  *widget.Select
	speechManager   htgotts.Speech
	whitelistSwitch *widget.Check
}

type TWITCH_TTS_STRUCT struct {
	Channel   string
	Language  string
	Whitelist TWITCH_WHITELIST
}
