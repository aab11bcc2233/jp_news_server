package details_news

import (
	"encoding/json"
	"fmt"
	"jpnews/internal/app/router"
	"jpnews/internal/pkg/news"
	"jpnews/internal/pkg/news/newsurl"
	"jpnews/internal/pkg/util/myrequest"
	"net/http"
	"strings"
)

const Path string = router.PathPrefix + "/details"

func Handle(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id := strings.Trim(strings.Join(request.Form["id"], ""), " ")
	if len(id) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, `{"error": "缺少必要参数"}`)
		return
	}

	content, err := myrequest.RequestGet(newsurl.NewsDetails(id))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "internal server error"}`)
		return
	}

	parseNews, err := news.ParseNews(string(content))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "internal server error"}`)
		return
	}

	parseNews.NewsID = id

	respBytes, err := json.Marshal(parseNews)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "internal server error"}`)
		return
	}

	fmt.Fprintf(writer, string(respBytes))
}
