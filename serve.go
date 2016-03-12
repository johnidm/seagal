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

    html, err := ReadURL(url)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    og, err := GetOG(html)

    if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
      }

    text := fmt.Sprintf("<!channel> *Content Marketing is amazing* :smile:, *compartilhem*\n<%s|%s>", og.URL, og.Title)

    facebook := Attachments{
      "#36a64f",
      "Share on Facebook",
      fmt.Sprintf("https://www.facebook.com/sharer/sharer.php?u=%s", og.URL),
    }

    twitter := Attachments{
      "#1a53ff",
      "Share on Twitter",

      fmt.Sprintf("https://twitter.com/home?status=%s", og.URL),
    }

    linkedin := Attachments{
        "#ffcc00",
        "Share on Linkedin",

        fmt.Sprintf("https://www.linkedin.com/shareArticle?mini=true&url=%s", og.URL),
    }

    googleplus := Attachments{
      "#ff5050",
      "Share on G+",
      fmt.Sprintf("https://plus.google.com/share?url=%s", og.URL),
    }

    m := Response{
        text,
        true,
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

func FacebookShareCount(URL string) int {
  return 0
}

func TwitterShareCount(URL string) int {
  return 0
}

func LinkedinShareCount(URL string) int {
  return 0
}

func GooglePlusShareCount(URL string) int {
  return 0
}
