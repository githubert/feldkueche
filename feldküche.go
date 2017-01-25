package main

import (
	"encoding/xml"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"strings"
	"path/filepath"
	img "image"
)

type Enclosure struct {
	Url string `xml:"url,attr"`
}

func (e *Enclosure) BaseName() string {
	index := strings.LastIndex(e.Url, "/")

	if index == -1 {
		return ""
	}

	return e.Url[index+1:]
}

func (e *Enclosure) ImageSize() (img.Point) {
	filename := "./enclosures" + e.BaseName()

	f, err := os.Open(filename)

	if err != nil {
		return img.Point{}
	}

	i, _, err := img.Decode(f)

	if err != nil {
		return img.Point{}
	}

	return i.Bounds().Max
}

type Item struct {
	XMLName     xml.Name  `xml:"item"`
	Enclosure   Enclosure `xml:"enclosure"`
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	PubDate     string    `xml:"pubDate"`
	Link        string    `xml:"link"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

var rss Rss

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./html/feldküche.html")
	t.Execute(w, nil)
}

func posts(w http.ResponseWriter, r *http.Request) {
	offset, _ := strconv.Atoi(r.URL.Path[len("/posts/"):])
	t, _ := template.ParseFiles("./html/feldküche_posts.html")
	p := rss.Channel.Items[offset*10 : offset*10+10]
	t.Execute(w, p)
}

func image(w http.ResponseWriter, r *http.Request) {
	// TODO: Is this secure?
	filename := r.URL.Path[len("/image/"):]
	basename := filepath.Base(filepath.Clean(filename));

	if basename == "." {
		return
	}

	http.ServeFile(w, r, "./enclosures/" + basename)
}

func main() {
	file, _ := os.Open("backup.rss")

	decoder := xml.NewDecoder(file)

	decoder.Decode(&rss)

	http.HandleFunc("/posts/", posts)
	http.HandleFunc("/image/", image)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
