package voice

import (
	"fmt"
	"jpnews/internal/app/router"
	"jpnews/internal/pkg/config"
	"jpnews/internal/pkg/news/newsurl"
	"jpnews/internal/pkg/util/myrequest"
	"net/http"
	"strings"
)

const PathMasterM3u8PlayList string = router.PathPrefix + "/media/voice/playlist/"

func HandleMasterM3u8PlayList(writer http.ResponseWriter, request *http.Request) {
	voiceUrl := LocalPlayListUrl2NewsEasyPlayListUrl(request.RequestURI)
	if len(voiceUrl) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, `{"error": "音频地址错误"}`)
		return
	}

	respContent, err := myrequest.RequestGet(voiceUrl)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "internal server error"}`)
		return
	}

	//fmt.Fprintf(writer, voiceUrl)
	writer.Write(respContent)
}

func LocalPlayListUrl2NewsEasyPlayListUrl(url string) string {
	if strings.Contains(url, PathMasterM3u8PlayList) {
		return strings.Replace(url, PathMasterM3u8PlayList, newsurl.VoiceBaseUrl + "/", 1)
	} else {
		return ""
	}
}

func NewsEasyPlayList2LocalPlayList(newsPlayList string) string {
	return strings.Replace(newsPlayList, newsurl.VoiceBaseUrl+"/", config.GetConfig().MediaDomain+PathMasterM3u8PlayList, -1)
}
