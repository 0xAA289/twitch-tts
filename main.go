package main

import (
	"os"
	"path/filepath"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
	"github.com/hegedustibor/htgo-tts/voices"
)

func initLangmap() {
	langMap = map[string]string{
		"English":                  "en",
		"English (Australia)":      "en-AU",
		"English (United Kingdom)": "en-GB",
		"Japanese":                 "ja",
		"German":                   "de",
		"Spanish":                  "es",
		"Russian":                  "ru",
		"Arabic":                   "ar",
		"Bengali":                  "bn",
		"Czech":                    "cs",
		"Danish":                   "da",
		"Dutch":                    "nl",
		"Finnish":                  "fi",
		"Greek":                    "el",
		"Hindi":                    "hi",
		"Hungarian":                "hu",
		"Indonesian":               "id",
		"Khmer":                    "km",
		"Latin":                    "la",
		"Italian":                  "it",
		"Norwegian":                "no",
		"Polish":                   "pl",
		"Slovak":                   "sk",
		"Swedish":                  "sv",
		"Thai":                     "th",
		"Turkish":                  "tr",
		"Ukrainian":                "uk",
		"Vietnamese":               "vi",
		"Afrikaans":                "af",
		"Bulgarian":                "bg",
		"Catalan":                  "ca",
		"Welsh":                    "cy",
		"Estonian":                 "et",
		"French":                   "fr",
		"Gujarati":                 "gu",
		"Icelandic":                "is",
		"Javanese":                 "jv",
		"Kannada":                  "kn",
		"Korean":                   "ko",
		"Latvian":                  "lv",
		"Malayalam":                "ml",
		"Marathi":                  "mr",
		"Malay":                    "ms",
		"Nepali":                   "ne",
		"Portuguese":               "pt",
		"Romanian":                 "ro",
		"Sinhala":                  "si",
		"Serbian":                  "sr",
		"Sundanese":                "su",
		"Tamil":                    "ta",
		"Telugu":                   "te",
		"Tagalog":                  "tl",
		"Urdu":                     "ur",
		"Chinese":                  "zh",
		"Swahili":                  "sw",
		"Albanian":                 "sq",
		"Burmese":                  "my",
		"Macedonian":               "mk",
		"Armenian":                 "hy",
		"Croatian":                 "hr",
		"Esperanto":                "eo",
		"Bosnian":                  "bs",
	}
}

func arrRemove(array []string, index int) []string {
	// Check if the index is within the bounds of the array.
	if index < 0 || index >= len(array) {
		return array
	}

	// Remove the element at the specified index.
	array = append(array[:index], array[index+1:]...)

	// Return the new array.
	return array
}

func unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

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

var whitelist map[string]bool

func stringArrayToKV(array []string) map[string]bool {
	// Create a new map to store the key-value pairs.
	keyValueArray := make(map[string]bool)

	// Iterate over the array and add each string as a key in the map.
	for _, value := range array {
		keyValueArray[value] = true
	}

	// Return the map.
	return keyValueArray
}

func main() {
	initLangmap()

	RUNTIME.speechManager = htgotts.Speech{Folder: "audio", Language: voices.English, Handler: &handlers.Native{}}

	gui()
}
