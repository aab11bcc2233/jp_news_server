package model

type NewsDate struct {
	Date string     `json:"date"`
	Data []NewsItem `json:"data"`
}

type NewsItem struct {
	NewsPriorityNumber    string `json:"news_priority_number"`
	NewsPrearrangedTime   Time   `json:"news_prearranged_time"`
	NewsID                string `json:"news_id"`
	Title                 string `json:"title"`
	TitleWithRuby         string `json:"title_with_ruby"`
	NewsFileVer           bool   `json:"news_file_ver"`
	NewsCreationTime      Time   `json:"news_creation_time"`
	NewsPreviewTime       Time   `json:"news_preview_time"`
	NewsPublicationTime   Time   `json:"news_publication_time"`
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
	NewsDisplayFlag       bool   `json:"news_display_flag"`
	NewsWebURL            string `json:"news_web_url"`
	ImageURL              string `json:"image_url"`
	VoiceURL              string `json:"voice_url"`
}

type News struct {
	NewsID              		string `json:"news_id"`
	Content             		string `json:"content"`
	ContentHTML         		string `json:"content_html"`
	ContentHTMLStyle        	string `json:"content_html_style"`
	ContentHTMLWithRuby         string `json:"content_html_with_ruby"`
	ContentHTMLWithRubyStyle	string `json:"content_html_with_ruby_style"`
}
