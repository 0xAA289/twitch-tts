package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gempir/go-twitch-irc/v4"
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
	"github.com/hegedustibor/htgo-tts/voices"
)

func deleteFilesInDirectory(dirPath string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and subdirectories
		if info.IsDir() {
			return nil
		}

		// Delete the file
		err = os.Remove(path)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("TwitchTTS Project by Strange")

	fmt.Print("Enter Twitch Name: ")
	name, _ := reader.ReadString('\n')

	client := twitch.NewAnonymousClient()
	speech := htgotts.Speech{Folder: "audio", Language: voices.English, Handler: &handlers.Native{}}

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.User.Name + " says: " + message.Message)
		speech.Speak(message.User.Name + " says: " + message.Message)
		deleteFilesInDirectory("audio")
	})

	client.Join(name)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
