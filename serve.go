package main

import (
    "fmt"
  "strings"
    "net/http"
    "io/ioutil"

  "github.com/zenazn/goji"
        // "github.com/zenazn/goji/web"
  "github.com/dyatlov/go-opengraph/opengraph"
)


func GetRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Content Marketing is Amazing")
}

func PostRoot(w http.ResponseWriter, r *http.Request) {
    response := `
      {
            "text": "*Content Marketing is amazing* :smile: <http://contentools.com.br/marketing-de-conteudo/como-comecar-uma-estrategia-eficaz-de-marketing-de-conteudo-para-cursos-online/|Como começar uma estratégia eficaz de marketing de conteúdo para cursos online>",
            "unfurl_links": false,
            "response_type": "in_channel",
            "parse":"full",
            "mrkdwn": true,
            "attachments": [
                {
                    "color": "#36a64f",
                    "title": "Share on Facebook",
                    "title_link": "https://api.slack.com/"
                },
                {
                    "color": "#1a53ff",
                    "title": "Share on Twitter",
                    "title_link": "https://api.slack.com/"
                },
                {
                    "color": "#ffcc00",
                    "title": "Share on Linkedin",
                    "title_link": "https://api.slack.com/"
                },
                {
                    "color": "#ff5050",
                    "title": "Share on G+",
                    "title_link": "https://api.slack.com/"
                }
            ]
        }
    `
    fmt.Fprintf(w, response)
}

func main() {
      goji.Get("/", GetRoot)
      goji.Post("/", PostRoot)
      goji.Serve()


    // content, err := ReadURL("http://www.johnidouglas.com.br/django-admin-using-list_filter-and-search_fields/")
    // if err != nil {
    //   fmt.Println(err)
    //   return
    // }
    // og, err := GetOG(content)

    // fmt.Printf("Title: %s\n", og.Title)
    // fmt.Printf("URL: %s\n", og.URL)

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
