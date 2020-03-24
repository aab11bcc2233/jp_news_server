package model

type TopNews struct {
	TopPriorityNumber     string `json:"top_priority_number"`
	TopDisplayFlag        bool   `json:"top_display_flag"`
	NewsID                string `json:"news_id"`
	NewsPrearrangedTime   Time   `json:"news_prearranged_time"`
	Title                 string `json:"title"`
	TitleWithRuby         string `json:"title_with_ruby"`
	OutlineWithRuby       string `json:"outline_with_ruby"`
	NewsFileVer           bool   `json:"news_file_ver"`
	NewsPublicationStatus bool   `json:"news_publication_status"`
	HasNewsWebImage       bool   `json:"has_news_web_image"`
	HasNewsWebMovie       bool   `json:"has_news_web_movie"`
	HasNewsEasyImage      bool   `json:"has_news_easy_image"`
	HasNewsEasyMovie      bool   `json:"has_news_easy_movie"`
	HasNewsEasyVoice      bool   `json:"has_news_easy_voice"`
	NewsWebImageURI       string `json:"news_web_image_uri"`
	NewsWebMovieURI       string `json:"news_web_movie_uri"`
	NewsEasyImageURI      string `json:"news_easy_image_uri"`
	NewsEasyMovieURI      string `json:"news_easy_movie_uri"`
	NewsEasyVoiceURI      string `json:"news_easy_voice_uri"`
	ImageURL              string `json:"image_url"`
	VoiceURL              string `json:"voice_url"`
}
