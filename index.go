package main

import (
	"Crawl_Web/SaveIntoMongo"
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	trTags = regexp.MustCompile(`<tr[^>]*>(?:.|\n)*<\/tr>`)
	tagA   = regexp.MustCompile(`<\s*a[^>]*>(.*?)<\s*/\s*a>`)
)

func getStringURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	StringURL := string(body)
	return StringURL
}

func getTagTrElement(url string) string {
	div := trTags.FindAllStringSubmatch(url, -1)
	return div[0][0]
}

func getTagAElement(fatherElement string) string {
	trTag := getTagTrElement(fatherElement)
	tags := tagA.FindAllStringSubmatch(trTag, -1)
	return tags[0][0]
}

func getLinkDomain(URL string) string {
	StringURL := getStringURL(URL)
	tags := getTagAElement(StringURL)
	links := strings.Split(tags, "\"")
	IndexIncludeLinkOnHref := 1
	return links[IndexIncludeLinkOnHref]
}

func downloadFromUrl(url string) {
	tokens := strings.Split(url, "/")
	file := tokens[len(tokens)-1]
	fileName := file + ".zip"
	fmt.Println("Downloading", url, "to", fileName)

	//Check file existence with io.IsExit
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating: ", err)
		return
	}
	defer output.Close()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading: ", err)
		return
	}
	defer resp.Body.Close()

	n, err := io.Copy(output, resp.Body)
	if err != nil {
		fmt.Println("Error while downloading: ", err)
		return
	}
	fmt.Println(n, "Downloaded")

	unZip(fileName)
}

func unZip(name string) {
	fileName := fmt.Sprintf(name)
	reader, _ := zip.OpenReader(fileName)
	defer reader.Close()

	for _, file := range reader.File {
		in, _ := file.Open()
		defer in.Close()

		year, month, day := time.Now().Date()
		yearString := strconv.Itoa(year)
		dayString := strconv.Itoa(day)

		folderName := path.Join(yearString, month.String(), dayString, file.Name)

		dir := path.Dir(folderName)
		os.MkdirAll(dir, 0777)
		out, err := os.Create(folderName)
		if err != nil {
			fmt.Println("Error while UnZip")
		}
		defer out.Close()

		n, err := io.Copy(out, in)
		if err != nil {
			fmt.Println("Error while UnZip: ", err)
		}
		fmt.Println(n, "Successful Unzip")
	}
}

func addViperConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	addViperConfig("config")

	URL := viper.GetString("domain.URL")
	link := getLinkDomain(URL)
	fmt.Println("Link: ", link)
	resp, err := http.Get(link)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status: ", resp.Status)
	if resp.StatusCode != 200 {
		return
	}

	downloadFromUrl(link)
	domainName := viper.GetString("domain.fileTXT")

	year, month, day := time.Now().Date()
	pathFile := strconv.Itoa(year) + "/" + month.String() + "/" + strconv.Itoa(day) + "/" + domainName
	SaveIntoMongo.Saving(pathFile, year, month.String(), day)

}
