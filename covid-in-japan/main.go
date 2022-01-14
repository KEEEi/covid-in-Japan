package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type CovidData struct {
	Date  string
	Japan string
	Tokyo string
	Aichi string
	Osaka string
}

func main() {
	fileUrl := "https://covid19.mhlw.go.jp/public/opendata/newly_confirmed_cases_daily.csv"

	if err := DownloadFile("covidData.csv", fileUrl); err != nil {
		panic(err)
	}
	file, err := os.Open("covidData.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var info []CovidData

	for i := 0; ; i++ {
		var line, err = reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if i == 0 {
			continue
		}

		var covidData CovidData
		for i, v := range line {
			switch i {
			case 0:
				covidData.Date = v
			case 1:
				covidData.Japan = v
			case 14:
				covidData.Tokyo = v
			case 24:
				covidData.Aichi = v
			case 28:
				covidData.Osaka = v
			}
		}
		info = append(info, covidData)
	}
	var latestInfo = (info[len(info)-1])
	fmt.Printf("    日付    全国 東京 愛知 大阪\n")
	fmt.Println(latestInfo)
}

func DownloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
