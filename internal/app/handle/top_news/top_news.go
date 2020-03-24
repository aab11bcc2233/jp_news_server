package top_news

import (
	"encoding/json"
	"fmt"
	"jpnews/internal/app/handle/image"
	"jpnews/internal/app/handle/voice"
	"jpnews/internal/app/model"
	"jpnews/internal/app/router"
	"jpnews/internal/pkg/news/newsurl"
	"jpnews/internal/pkg/util/myrequest"
	"log"
	"net/http"
)

const Path string = router.PathPrefix + "/top"

func Handle(writer http.ResponseWriter, request *http.Request) {
	content, err := myrequest.RequestGet(newsurl.TopNews())
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "internal server error"}`)
		return
	}

	var topNews []model.TopNews
	err = json.Unmarshal(content, &topNews)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "internal server error"}`)
		return
	}

	for i := range topNews {
		v := &topNews[i]
		if v.HasNewsEasyImage {
			v.NewsEasyImageURI = image.EasyImageUrlToLocalImageUrl(v.NewsID, v.NewsEasyImageURI)
			v.ImageURL = v.NewsEasyImageURI
		} else if v.HasNewsWebImage {
			v.NewsWebImageURI = image.WebImageUrlToLocalImageUrl(v.NewsWebImageURI)
			v.ImageURL = v.NewsWebImageURI
		}

		if v.HasNewsEasyVoice {
			v.VoiceURL = voice.LocalVoiceURL(v.NewsEasyVoiceURI)
		}
	}

	jsonBytes, err := json.Marshal(topNews)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "internal server error"}`)
		return
	}

	fmt.Fprintf(writer, string(jsonBytes))

}
