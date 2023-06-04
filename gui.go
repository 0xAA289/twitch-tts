package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/gempir/go-twitch-irc/v4"
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
)

var langMap map[string]string
var connId int

func gui() {
	a := app.New()
	w := a.NewWindow("Twitch Chat TTS")

	w.SetFixedSize(true)
	w.SetFullScreen(false)
	w.Resize(fyne.NewSize(488, 248))

	twitchAccName := widget.NewEntry()
	twitchAccName.PlaceHolder = "yourname"
	RUNTIME.nameWidget = twitchAccName

	langs := []string{
		"English", "Japanese", "German", "Spanish", "Russian", "Arabic", "bengali", "Czech", "Danish", "Dutch", "Finnish", "Greek", "Hindi", "Hungarian", "Indonesian", "Khmer", "Latin", "Italian", "Norwegian", "Polish", "Slovak", "Swedish", "Thai", "Turkish", "Ukrainian", "Vietnamese", "Afrikaans", "Bulgarian", "Catalan", "Welsh", "Estonian", "French", "Gujarati", "Icelandic", "Javanese", "Kannada", "Korean", "Latvian", "Malayalam", "Marathi", "Malay", "Nepali", "Portuguese", "Romanian", "Sinhala", "Serbian", "Sundanese", "Tamil", "Telugu", "Tagalog", "Urdu", "Chinese", "Swahili", "Albanian", "Burmese", "Macedonian", "Armenian", "Croatian", "Esperanto", "Bosnian"}

	languageSelection := widget.NewSelect(langs, func(s string) {
		TTS_DATA.Language = s

		var selectedLanguage string = "English"
		if len(langMap[TTS_DATA.Language]) != 0 {
			selectedLanguage = langMap[TTS_DATA.Language]
		}

		RUNTIME.speechManager = htgotts.Speech{Folder: "audio", Language: selectedLanguage, Handler: &handlers.Native{}}
	})

	RUNTIME.languageWidget = languageSelection

	submitButton := widget.NewButton("Connect", nil)

	submitButton.OnTapped = func() {
		connId++
		if submitButton.Text == "Connect" {
			if RUNTIME.client != nil {
				RUNTIME.client.Disconnect()
			}
			RUNTIME.client = twitch.NewAnonymousClient()
			TTS_DATA.Channel = twitchAccName.Text
			TTS_DATA.Language = languageSelection.Selected
			RUNTIME.client.Join(TTS_DATA.Channel)

			go RUNTIME.client.Connect()

			var statConn int = connId
			RUNTIME.client.OnPrivateMessage(func(message twitch.PrivateMessage) {
				if TTS_DATA.Whitelist.Enabled && !whitelist[message.User.Name] {
					return
				}

				if statConn != connId {
					return
				}

				fmt.Println(message.User.Name + " says: " + message.Message)
				RUNTIME.speechManager.Speak(message.User.Name + " says: " + message.Message)
				deleteFilesInDirectory("audio")
			})

			SaveConfig()
			submitButton.SetText("Disconnect")
			submitButton.Importance = 4
		} else if submitButton.Text == "Disconnect" {
			RUNTIME.client.Depart(twitchAccName.Text)
			submitButton.SetText("Connect")
			submitButton.Importance = 0
		}

		submitButton.Refresh()
	}

	whitelistList := widget.NewList(
		func() int { return len(TTS_DATA.Whitelist.List) },
		func() fyne.CanvasObject { return widget.NewLabel("Item") },
		func(i widget.ListItemID, o fyne.CanvasObject) { o.(*widget.Label).SetText(TTS_DATA.Whitelist.List[i]) },
	)

	whitelistEnabled := widget.NewCheck("Enable Whitelist", nil)
	whitelistEnabled.OnChanged = func(b bool) {
		TTS_DATA.Whitelist.Enabled = b
		SaveConfig()
	}

	RUNTIME.whitelistSwitch = whitelistEnabled

	var selected int = -1
	whitelistList.OnSelected = func(id widget.ListItemID) {
		selected = id
	}

	removeSelectedButton := widget.NewButton("Remove Selected", func() {
		if selected == -1 {
			return
		}
		TTS_DATA.Whitelist.List = arrRemove(TTS_DATA.Whitelist.List, selected)
		whitelistList.Refresh()
		SaveConfig()
	})

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Username")

	main := container.NewBorder(nil, nil, nil, nil, container.NewAppTabs(
		container.NewTabItem("Setup", container.New(layout.NewVBoxLayout(),
			widget.NewLabel("Enter Channel ID"),
			twitchAccName,
			widget.NewLabel("TTS Language (Voice/Accent)"),
			languageSelection,
			submitButton,
		)),
		container.NewTabItem("Whitelist", container.New(layout.NewGridLayout(2),
			whitelistList,
			container.New(layout.NewVBoxLayout(),
				whitelistEnabled,
				nameEntry,
				widget.NewButton("Add", func() {
					for _, user := range TTS_DATA.Whitelist.List {
						if user == nameEntry.Text {
							return
						}
					}

					TTS_DATA.Whitelist.List = append(TTS_DATA.Whitelist.List, nameEntry.Text)

					whitelistList.Refresh()
					SaveConfig()
				}),
				removeSelectedButton,
			),
		)),
		// container.NewTabItem("Config", container.New(layout.NewHBoxLayout(),
		// 	widget.NewButton("Load Config", LoadConfig),
		// 	widget.NewButton("Save Config", SaveConfig),
		// )),
	))
	w.SetContent(main)
	w.CenterOnScreen()

	LoadConfig()
	w.ShowAndRun()
}
