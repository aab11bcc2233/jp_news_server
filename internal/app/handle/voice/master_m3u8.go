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

const PathMasterM3u8 string = router.PathPrefix + "/media/voice/master/"

func HandleMasterM3u8(writer http.ResponseWriter, request *http.Request) {

	voiceUrl := LocalVoiceUrlToNewsVoiceUrl(request.RequestURI)
	if len(voiceUrl) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, `{"error": "音频地址错误"}`)
		return
	}

	masterM3u8RespContent, err := myrequest.RequestGet(voiceUrl)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "internal server error"}`)
		return
	}

	masterM3u8 := string(masterM3u8RespContent)

	urlIndex := strings.Index(masterM3u8, "http")
	if urlIndex == -1 {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "internal server error"}`)
		return
	}

	indexMasterM3u8Url := masterM3u8[urlIndex:]
	indexMasterM3u8Url = strings.TrimFunc(indexMasterM3u8Url, func(r rune) bool {
		return r == '\n' || r == ' '
	})

	respContent, err := myrequest.RequestGet(indexMasterM3u8Url)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "internal server error"}`)
		return
	}

	indexMasterM3u8 := string(respContent)

	indexMasterM3u8 = NewsEasyPlayList2LocalPlayList(indexMasterM3u8)

	//fmt.Fprintf(writer, indexMasterM3u8)
	writer.Write([]byte(indexMasterM3u8))
}

func LocalVoiceUrlToNewsVoiceUrl(voiceUrl string) string {
	if strings.Contains(voiceUrl, PathMasterM3u8) {
		return strings.Replace(voiceUrl, PathMasterM3u8, newsurl.VoiceBaseUrl+"/", 1)
	} else {
		return ""
	}
}
func LocalVoiceURL(voiceName string) string {
	return config.GetConfig().MediaDomain + PathMasterM3u8 + newsurl.VoiceEasyPath + voiceName + "/master.m3u8"
}
