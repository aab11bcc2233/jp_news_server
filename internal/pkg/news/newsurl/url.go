package newsurl

import (
	"strings"
)

const NewsWebUrl = "https://www3.nhk.or.jp/news/html"

//var v = &url.URL{
//	Scheme: "https",
//	Host: "www3.nhk.or.jp",
//	Path: "news/easy",
//}

const baseUrl = "https://www3.nhk.or.jp"
const urlPath = "/news/easy"
const NewsEasyUrl = baseUrl + urlPath

const topNewsList = NewsEasyUrl + "/top-list.json"
const newsList = NewsEasyUrl + "/news-list.json"

const VoiceBaseUrl = "https://nhks-vh.akamaihd.net"
const VoiceEasyPath = "i/news/easy/"
const VoiceEasyUrl = VoiceBaseUrl + "/" + VoiceEasyPath

//https://nhks-vh.akamaihd.net/i/news/easy/k10012041781000.mp4/master.m3u8

func TopNews() string {
	return topNewsList
}

func News() string {
	return newsList
}

func NewsDetails(newsId string) string {
	return strings.Join(
		[]string{
			baseUrl,
			urlPath,
			"/",
			newsId,
			"/",
			newsId,
			".html",
		},
		"",
	)
}

func NewsEasyImage(newsId string, newsEasyImageUri string) string {
	return strings.Join(
		[]string{
			NewsEasyUrl,
			"/",
			newsId,
			"/",
			newsEasyImageUri,
		},
		"",
	)
}
