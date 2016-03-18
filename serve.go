package main

import (
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
    "github.com/dyatlov/go-opengraph/opengraph"
    "github.com/parnurzeal/gorequest"
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

type AttachmentsMetrics struct {
    Color string              `json:"color" bson:"color"`
    Title string              `json:"title" bson:"title"`
}

type ResponseMetrics struct {
    Text string             `json:"text" bson:"text"`
    ResponseType string     `json:"response_type" bson:"response_type"`
    Markdown bool           `json:"mrkdwn" bson:"mrkdwn"`
    AttachmentsMetrics []AttachmentsMetrics  `json:"attachments" bson:"attachments"`
}

type FacebookShare struct {
  ID     string `json:"id"`
  Shares int    `json:"shares"`
}

type LinkedinShare struct {
  Count       int    `json:"count"`
  FCnt        string `json:"fCnt"`
  FCntPlusOne string `json:"fCntPlusOne"`
  URL         string `json:"url"`
}

func main() {
      goji.Post("/share", PostShare)
      goji.Post("/metric", PostMetrics)
      goji.Serve()
}


func PostMetrics(c web.C, w http.ResponseWriter, r *http.Request) {

    user := r.FormValue("user_name")
    user_id := r.FormValue("user_id")

    url := r.FormValue("text")

    facebook := AttachmentsMetrics{
      "#36a64f",
      fmt.Sprintf("Total Shares on Facebook: %d", FacebookShareCount(url)),
    }

    // twitter := AttachmentsMetrics{
    //   "#1a53ff",
    //   "Total Shares on Twitter: 1",
    // }

    linkedin := AttachmentsMetrics{
        "#ffcc00",
        fmt.Sprintf("Total Shares on Linkedin: %d", LinkedinShareCount(url)),
    }

    // googleplus := AttachmentsMetrics{
    //   "#ff5050",
    //   "Total Shares on G+",
    // }

    text := fmt.Sprintf("Hi <@%s|%s>, we have been analyzing the URL you gave us, and this is the result of our analysis:", user_id, user)

    m := ResponseMetrics{
        text,
        "in_channel",
        true,
        []AttachmentsMetrics{
          facebook,
          // twitter,
          linkedin,
          // googleplus,
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
func PostShare(c web.C, w http.ResponseWriter, r *http.Request) {

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

    if len(og.URL) == 0 {
      og.URL = url
    }

    text := fmt.Sprintf("<!channel> *Content Marketing is amazing* :smile: <%s|%s>",og.URL, og.Title)

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

  var fs FacebookShare;

  url := fmt.Sprintf("http://graph.facebook.com/?id=%s", URL)

  _, body, _ := gorequest.New().Get(url).End()

  json.Unmarshal([]byte(body), &fs)

  return fs.Shares
}

func TwitterShareCount(URL string) int {
  return 0
}

func LinkedinShareCount(URL string) int {

  var ls LinkedinShare

  url := fmt.Sprintf("https://www.linkedin.com/countserv/count/share?url=%s&format=json", URL)

  _, body, _ := gorequest.New().Get(url).End()

  fmt.Println(body)
  json.Unmarshal([]byte(body), &ls)

  return ls.Count
}

func GooglePlusShareCount(URL string) int {
  return 0
}
