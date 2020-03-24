package image

import (
	"fmt"
	"jpnews/internal/app/router"
	"jpnews/internal/pkg/config"
	"jpnews/internal/pkg/news/newsurl"
	"jpnews/internal/pkg/util/myrequest"
	"net/http"
	"strings"
)

const Path string = router.PathPrefix + "/media/image/"

func Handle(writer http.ResponseWriter, request *http.Request) {
	imageUrl := LocalImageUrlToNewsImageUrl(request.RequestURI)
	if len(imageUrl) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, `{"error": "图片地址错误"}`)
		return
	}

	respContent, err := myrequest.RequestGet(imageUrl)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "internal server error"}`)
		return
	}

	writer.Write(respContent)
}

const newsEasyLocalImagePath = Path + "newseasy"

func NewsEasyLocalImageUrl() string {
	return config.GetConfig().MediaDomain + newsEasyLocalImagePath
}

const newsWebLocalImagePath = Path + "newsweb"

func NewsWebLocalImageUrl() string {
	return config.GetConfig().MediaDomain + newsWebLocalImagePath
}

func EasyImageUrlToLocalImageUrl(newsId string, newsEasyImageUri string) string {
	imageUrl := newsurl.NewsEasyImage(newsId, newsEasyImageUri)
	return strings.Replace(imageUrl, newsurl.NewsEasyUrl, NewsEasyLocalImageUrl(), 1)
}

func WebImageUrlToLocalImageUrl(newsWebImageUri string) string {
	return strings.Replace(newsWebImageUri, newsurl.NewsWebUrl, NewsWebLocalImageUrl(), 1)
}

func LocalImageUrlToNewsImageUrl(imageUrl string) string {
	if strings.Contains(imageUrl, newsEasyLocalImagePath) {
		return strings.Replace(imageUrl, newsEasyLocalImagePath, newsurl.NewsEasyUrl, 1)
	} else if strings.Contains(imageUrl, newsWebLocalImagePath) {
		return strings.Replace(imageUrl, newsWebLocalImagePath, newsurl.NewsWebUrl, 1)
	} else {
		return ""
	}
}
