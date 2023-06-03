package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"fyne.io/fyne/v2/widget"
	"github.com/gempir/go-twitch-irc/v4"
	htgotts "github.com/hegedustibor/htgo-tts"
)

type TWITCH_WHITELIST struct {
	Enabled bool
	List    []string
}

type TWITCH_RUNTIME_STRUCT struct {
	client         *twitch.Client
	nameWidget     *widget.Entry
	languageWidget *widget.Select
	speechManager  htgotts.Speech
}

type TWITCH_TTS_STRUCT struct {
	Channel   string
	Language  string
	Whitelist TWITCH_WHITELIST
}

var TTS_DATA = TWITCH_TTS_STRUCT{
	Channel:  "none",
	Language: "English",
	Whitelist: TWITCH_WHITELIST{
		Enabled: false,
		List:    []string{},
	},
}

var RUNTIME = TWITCH_RUNTIME_STRUCT{
	client:         nil,
	nameWidget:     nil,
	languageWidget: nil,
}

func SaveConfig() {
	b, err := json.MarshalIndent(TTS_DATA, "", " ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = ioutil.WriteFile("config.json", b, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func LoadConfig() {
	f, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal(b, &TTS_DATA)
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	RUNTIME.nameWidget.SetText(TTS_DATA.Channel)
	RUNTIME.languageWidget.SetSelected(TTS_DATA.Language)
}
