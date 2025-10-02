# Ayah of the Day CLI

A simple command-line application that displays a random Ayah (verse) from the Quran.

## Description

This CLI tool is built with Go and provides a random verse from the Quran to inspire you every time you open your terminal. The data is sourced from the `alquran` directory, which contains various editions of the Quran in JSON format.

## Installation

1.  Clone the repository:
    ```sh
    git clone https://github.com/your-username/ayahoftheday.git
    cd ayahoftheday
    ```

2.  Build the application:
    ```sh
    go build
    ```

## Usage

Run the application from your terminal:

```sh
./ayahoftheday
```

This will display a random Ayah and its corresponding Surah and verse number.

**Example Output:**

```
(Dan Dialah yang menjadikan bintang-bintang bagimu, agar kamu menjadikannya petunjuk dalam kegelapan di darat dan di laut. Sesungguhnya Kami telah menjelaskan tanda-tanda kebesaran (Kami) kepada orang-orang yang mengetahui.)
Al-An'aam : 97
```

## Data Sources

The Quranic data used in this application is from the JSON files located in the `alquran/quran` directory. The default version used is the Indonesian translation.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
