package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Link      string    `xml:"link"`
	Enclosure Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	Url string `xml:"url,attr"`
}

func main() {

	file := ""
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	v := new(Rss)
	xml.Unmarshal(data, v)

	for i := 0; i < len(v.Channel.Items); i++ {

		item := v.Channel.Items[i]
		pid := strings.Replace(item.Link, "http://www.bbc.co.uk/programmes/", "", -1)

		fmt.Printf("curl -L -s %s > %s.mp3\n", item.Enclosure.Url, pid)

	}

}
