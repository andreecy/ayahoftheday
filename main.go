package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

//go:embed alquran/quran/id.indonesian.json
var embededFiles embed.FS

type QuranResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   struct {
		Surahs []struct {
			Number                 int    `json:"number"`
			Name                   string `json:"name"`
			EnglishName            string `json:"englishName"`
			EnglishNameTranslation string `json:"englishNameTranslation"`
			RevelationType         string `json:"revelationType"`
			Ayahs                  []struct {
				Number        int    `json:"number"`
				Text          string `json:"text"`
				NumberInSurah int    `json:"numberInSurah"`
				Juz           int    `json:"juz"`
				Manzil        int    `json:"manzil"`
				Page          int    `json:"page"`
				Ruku          int    `json:"ruku"`
				HizbQuarter   int    `json:"hizbQuarter"`
			} `json:"ayahs"`
		} `json:"surahs"`
	} `json:"data"`
}

func loadFile(file string) (*QuranResponse, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var quranResponse QuranResponse
	err = json.Unmarshal(data, &quranResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal file: %w", err)
	}
	return &quranResponse, nil
}

func loadEmbedFile(file string) (*QuranResponse, error) {
	data, err := embededFiles.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var quranResponse QuranResponse
	err = json.Unmarshal(data, &quranResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal file: %w", err)
	}
	return &quranResponse, nil

}

func main() {
	quranResponse, err := loadEmbedFile("alquran/quran/id.indonesian.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// get random surahs
	surahs := quranResponse.Data.Surahs
	randomSurah := surahs[rand.Intn(len(surahs)-1)]

	// get random ayahs
	ayahs := randomSurah.Ayahs
	randomAyah := ayahs[rand.Intn(len(ayahs)-1)]

	surahName := randomSurah.EnglishName
	ayahNumber := randomAyah.NumberInSurah

	fmt.Println(randomAyah.Text)
	fmt.Printf("%s : %d\n", surahName, ayahNumber)
}
