package app

import (
	"jpnews/internal/app/handle/details_news"
	"jpnews/internal/app/handle/image"
	"jpnews/internal/app/handle/list_news"
	"jpnews/internal/app/handle/top_news"
	"jpnews/internal/app/handle/voice"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc(top_news.Path, top_news.Handle)
	http.HandleFunc(list_news.Path, list_news.Handle)
	http.HandleFunc(details_news.Path, details_news.Handle)
	http.HandleFunc(image.Path, image.Handle)
	http.HandleFunc(voice.PathMasterM3u8, voice.HandleMasterM3u8)
	http.HandleFunc(voice.PathMasterM3u8PlayList, voice.HandleMasterM3u8PlayList)
	log.Fatal(http.ListenAndServe("localhost:8088", nil))
}
