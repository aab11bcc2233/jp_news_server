package list_news

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
	"sort"
)

const Path string = router.PathPrefix + "/list"

func Handle(writer http.ResponseWriter, request *http.Request) {
	content, err := myrequest.RequestGet(newsurl.News())
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "internal server error"}`)
		return
	}

	/*
		[
			{
				"2019-08-15": [
					{
					}
				]
			}
		]
		为了把上面的数据，变成下面的
		[
			{
				"date": "2019-08-15"
				"data": [
					{
					}
				]
			}
		]
	*/

	var news []map[string][]model.NewsItem
	unJsonErr := json.Unmarshal(content, &news)
	if unJsonErr != nil {
		log.Fatalf("JSON marshaling failed: %s", unJsonErr)
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "internal server error"}`)
		return
	}

	var respData []model.NewsDate

	for _, item := range news {
		len := len(item)

		// 因为 map 遍历顺序是随机的，所以这里先把 map 的 key 存入数组
		var keys = make([]string, len)
		keysIndex := 0
		for key := range item {
			keys[keysIndex] = key
			keysIndex++
		}

		// 对数组，也就是 map 的 key 排序
		sort.Strings(keys)

		// 逆序遍历 key
		for k_i := len - 1; k_i >= 0; k_i-- {
			date := keys[k_i]
			list := item[date]

			// 将 nhk 的 ImageUrl 转变成 自己的 URL
			for i := range list {
				data := &list[i]
				if data.HasNewsEasyImage {
					data.NewsEasyImageURI = image.EasyImageUrlToLocalImageUrl(data.NewsID, data.NewsEasyImageURI)
					data.ImageURL = data.NewsEasyImageURI
				} else if data.HasNewsWebImage {
					data.NewsWebImageURI = image.WebImageUrlToLocalImageUrl(data.NewsWebImageURI)
					data.ImageURL = data.NewsWebImageURI
				}

				if data.HasNewsEasyVoice {
					data.VoiceURL = voice.LocalVoiceURL(data.NewsEasyVoiceURI)
				}
			}

			// 重新封装数据
			data := model.NewsDate{
				Date: date,
				Data: list,
			}
			respData = append(respData, data)
		}
	}

	bytes, err := json.Marshal(respData)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "internal server error"}`)
		return
	}

	fmt.Fprintf(writer, string(bytes))
}
