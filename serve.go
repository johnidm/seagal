package main

import (
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
    "github.com/dyatlov/go-opengraph/opengraph"
    "encoding/json"
)

type Body struct {
    Text string             `json:"text" bson:"text"`
}

type Response struct {
    Text string             `json:"text" bson:"text"`
    UnfurlLinks bool        `json:"unfurl_links" bson:"unfurl_links"`
    ResponseType string     `json:"response_type" bson:"response_type"`
    Parse string            `json:"parse" bson:"parse"`
    Markdown bool           `json:"mrkdwn" bson:"mrkdwn"`

    Attachments []Attachments  `json:"attachments" bson:"attachments"`
}

type Attachments struct {
    Color string              `json:"color" bson:"color"`
    Title string              `json:"title" bson:"title"`
    TitleLink string          `json:"title_link" bson:"title_link"`
}

func main() {

      goji.Get("/", GetRoot)
      goji.Post("/", PostRoot)
      goji.Serve()



}


func GetRoot(w http.ResponseWriter, r *http.Request) {
}

func PostRoot(c web.C, w http.ResponseWriter, r *http.Request) {

    url := r.FormValue("text")

    html, _ := ReadURL(url)

    og, _ := GetOG(html)

    text := fmt.Sprintf("*Content Marketing is amazing* :smile: <%s|%s>",og.URL, og.Title)

    facebook := Attachments{
      "#36a64f",
      "Share on Facebook",
      "https://www.facebook.com/",
    }

    twitter := Attachments{
      "#1a53ff",
      "Share on Twitter",
      "https://twitter.com/",
    }

    linkedin := Attachments{
        "#ffcc00",
        "Share on Linkedin",
        "https://api.slack.com/",
    }

    googleplus := Attachments{
      "#ff5050",
      "Share on G+",
      "https://api.slack.com/",

    }

    m := Response{
        text,
        false,
        "in_channel",
        "full",
        true,
        []Attachments{
          facebook,
          twitter,
          linkedin,
          googleplus,
        },
      }

    b, err := json.Marshal(m)

      if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
      }

    w.Header().Set("Content-Type", "application/json")
    w.Write(b)
}



func GetOG(HTML string) (*opengraph.OpenGraph, error) {

   og := opengraph.NewOpenGraph()
   err := og.ProcessHTML(strings.NewReader(HTML))

  if err != nil {
    fmt.Println(err)
    return og, err
  }

  return  og, err
}

func ReadURL(URL string) (string, error) {

  response, err := http.Get(URL)

  if err != nil {
      return "", err
  } else {
    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)

    if err != nil {
      return "", err
    }

    return string(contents), nil
  }


}
