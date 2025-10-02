package main

import (
	"encoding/json"
	"fmt"
)

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

// func loadFile(file string) (*QuranResponse, error) {
// 	data, err := os.ReadFile(file)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read file: %w", err)
// 	}
//
// 	var quranResponse QuranResponse
// 	err = json.Unmarshal(data, &quranResponse)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal file: %w", err)
// 	}
// 	return &quranResponse, nil
// }

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
