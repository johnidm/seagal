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
}

func main() {

      goji.Get("/", GetRoot)
      goji.Post("/", PostRoot)
      goji.Serve()



}


func GetRoot(w http.ResponseWriter, r *http.Request) {

      // m := Response{"Welcome to the SandovalEffect API, build v0.0.001.992, 6/22/2015 0340 UTC."}
      // b, err := json.Marshal(m)

      // if err != nil {
      //     http.Error(w, err.Error(), http.StatusInternalServerError)
      //     return
      // }

      // w.Header().Set("Content-Type", "application/json")
      // w.Write(b)
}

func PostRoot(c web.C, w http.ResponseWriter, r *http.Request) {

    url := r.FormValue("text")

    html, _ := ReadURL(url)

    og, _ := GetOG(html)

    text := fmt.Sprintf("*Content Marketing is amazing* :smile: <%s|%s>",og.URL, og.Title)

    m := Response{
        text,
        false,
        "in_channel",
        "full",
        true,
      }

    b, err := json.Marshal(m)

      if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
      }

    // response := `
    //   {

    //         "attachments": [
    //             {
    //                 "color": "#36a64f",
    //                 "title": "Share on Facebook",
    //                 "title_link": "https://api.slack.com/"
    //             },
    //             {
    //                 "color": "#1a53ff",
    //                 "title": "Share on Twitter",
    //                 "title_link": "https://api.slack.com/"
    //             },
    //             {
    //                 "color": "#ffcc00",
    //                 "title": "Share on Linkedin",
    //                 "title_link": "https://api.slack.com/"
    //             },
    //             {
    //                 "color": "#ff5050",
    //                 "title": "Share on G+",
    //                 "title_link": "https://api.slack.com/"
    //             }
    //         ]
    //     }
    // `
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
