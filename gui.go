package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func gui() {
	a := app.New()
	w := a.NewWindow("Twitch Chat TTS")

	w.SetFixedSize(true)
	w.SetFullScreen(false)
	w.Resize(fyne.NewSize(488, 259))

	twitchAccName := widget.NewEntry()
	twitchAccName.PlaceHolder = "yourname"
	RUNTIME.nameWidget = twitchAccName

	langs := []string{
		"English", "EnglishUK", "EnglishAU", "Japanese", "German", "Spanish", "Russian", "Arabic", "bengali", "Czech", "Danish", "Dutch", "Finnish", "Greek", "Hindi", "Hungarian", "Indonesian", "Khmer", "Latin", "Italian", "Norwegian", "Polish", "Slovak", "Swedish", "Thai", "Turkish", "Ukrainian", "Vietnamese", "Afrikaans", "Bulgarian", "Catalan", "Welsh", "Estonian", "French", "Gujarati", "Icelandic", "Javanese", "Kannada", "Korean", "Latvian", "Malayalam", "Marathi", "Malay", "Nepali", "Portuguese", "Romanian", "Sinhala", "Serbian", "Sundanese", "Tamil", "Telugu", "Tagalog", "Urdu", "Chinese", "Swahili", "Albanian", "Burmese", "Macedonian", "Armenian", "Croatian", "Esperanto", "Bosnian"}

	languageSelection := widget.NewSelect(langs, nil)
	RUNTIME.languageWidget = languageSelection

	submitButton := widget.NewButton("Connect", func() {
		RUNTIME.client.Depart(TTS_DATA.Channel)
		TTS_DATA.Channel = twitchAccName.Text
		TTS_DATA.Language = languageSelection.Selected
		RUNTIME.client.Join(TTS_DATA.Channel)

		twitchAccName.Disable()
		SaveConfig()
	})

	whitelistList := widget.NewList(
		func() int { return len(TTS_DATA.Whitelist.List) },
		func() fyne.CanvasObject { return widget.NewLabel("Item") },
		func(i widget.ListItemID, o fyne.CanvasObject) { o.(*widget.Label).SetText(TTS_DATA.Whitelist.List[i]) },
	)

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
				widget.NewCheck("Enabled", nil),
			),
		)),
		container.NewTabItem("Config", container.New(layout.NewHBoxLayout(),
			widget.NewButton("Load Config", LoadConfig),
			widget.NewButton("Save Config", SaveConfig),
		)),
	))
	w.SetContent(main)
	w.CenterOnScreen()

	LoadConfig()
	w.ShowAndRun()
}
