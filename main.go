package main

import (
	"embed"
	"fmt"
	"math/rand"
)

//go:embed alquran/quran/*
var embededFiles embed.FS

func main() {
	quranEn, err := loadEmbedFile("alquran/quran/en.asad.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	quranId, err := loadEmbedFile("alquran/quran/id.indonesian.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// get random surahs
	surahs := quranEn.Data.Surahs
	randomSurahNumber := rand.Intn(len(surahs) - 1)
	randomSurah := surahs[randomSurahNumber]

	// get random ayahs
	ayahs := randomSurah.Ayahs
	randomAyahNumber := rand.Intn(len(ayahs) - 1)
	randomAyah := ayahs[randomAyahNumber]

	randomAyahId := quranId.Data.Surahs[randomSurahNumber].Ayahs[randomAyahNumber]

	surahName := randomSurah.EnglishName
	ayahNumber := randomAyah.NumberInSurah

	fmt.Println(randomAyah.Text)
	fmt.Println("---")
	fmt.Println(randomAyahId.Text)
	fmt.Printf("%s : %d\n", surahName, ayahNumber)
}
