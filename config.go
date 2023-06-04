package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveConfig() {
	TTS_DATA.Whitelist.List = unique(TTS_DATA.Whitelist.List)

	b, err := json.MarshalIndent(TTS_DATA, "", " ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = os.WriteFile("config.json", b, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	whitelist = stringArrayToKV(TTS_DATA.Whitelist.List)
}

func LoadConfig() {
	b, err := os.ReadFile("config.json")
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
	RUNTIME.whitelistSwitch.SetChecked(TTS_DATA.Whitelist.Enabled)
}
