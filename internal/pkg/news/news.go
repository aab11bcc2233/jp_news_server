package news

import (
	"github.com/PuerkitoBio/goquery"
	"jpnews/internal/app/model"
	"strings"
)

//var colors = map[string]string{
//	".colorN": "#35a16b", // 人名的颜色
//	".colorL": "#ff7f00", // 国家、县、城市、场所的名字的颜色
//	".colorC": "#0041cc", // 公司、集团的名字的颜色
//}



func ParseNews(html string) (model.News, error) {
	dom, e := goquery.NewDocumentFromReader(strings.NewReader(html))
	if e != nil {
		return model.News{}, nil
	}

	newsText, newsHTML, newsHTMLWithRuby, e := parseNewsDetail(dom)
	if e != nil {
		return model.News{}, nil
	}

	titleHTML, titleHTMLWithRuby, e := parseNewsTitle(dom)
	if e != nil {
		return model.News{}, nil
	}

	return model.News{
		Content:             		newsText,
		ContentHTML:         		packData(titleHTML, newsHTML),
		ContentHTMLStyle:        	cssStyle,
		ContentHTMLWithRuby: 		packData(titleHTMLWithRuby, newsHTMLWithRuby),
		ContentHTMLWithRubyStyle: 	cssStyleWithRuby,
	}, nil
}

func parseNewsTitle(dom *goquery.Document) (titleHTML string, titleHTMLWithRuby string, err error) {
	htmlTitle := dom.Find(".article-main__title")

	htmlWithRuby, e := htmlTitle.Html()
	if e != nil {
		return "", "", nil
	}

	htmlTitle.Each(func(_ int, selection *goquery.Selection) {
		removeRtTag(selection)
		replaceRubyTag(selection)
	})

	html, e := htmlTitle.Html()
	if e != nil {
		return "", "", nil
	}

	html = strings.TrimFunc(html, trimFunc)
	htmlWithRuby = strings.TrimFunc(htmlWithRuby, trimFunc)

	addH1 := func(h string) string {
		return `<h1 class="article-main__title">` + h + `</h1>`
	}

	return addH1(html), addH1(htmlWithRuby), nil
}

func parseNewsDetail(dom *goquery.Document) (newsText string, newsHTML string, newsHTMLWithRuby string, err error) {
	htmlArticle := dom.Find("#js-article-body")

	htmlArticle.Each(func(_ int, selection *goquery.Selection) {
		replaceATag(selection)
	})

	htmlWithRuby, e := htmlArticle.Html()
	if e != nil {
		return "", "", "", nil
	}

	htmlArticle.Each(func(_ int, selection *goquery.Selection) {
		removeRtTag(selection)
		replaceRubyTag(selection)
	})

	html, e := htmlArticle.Html()
	if e != nil {
		return "", "", "", nil
	}

	text := strings.TrimFunc(htmlArticle.Text(), trimFunc)
	html = strings.TrimFunc(html, trimFunc)
	htmlWithRuby = strings.TrimFunc(htmlWithRuby, trimFunc)
	return text, html, htmlWithRuby, nil
}

func trimFunc(r rune) bool {
	return r == rune('\n') || r == rune(' ')
}

func packData(titleHtml string, newsHtml string) string {
	return "<body>" + titleHtml + newsHtml +"</body>"
}


func replaceATag(selection *goquery.Selection) {
	selection.Find("a").Each(func(_ int, selection *goquery.Selection) {
		aTagChild, e := selection.Html()
		if e == nil {
			selection.ReplaceWithHtml(aTagChild)
		}
	})
}

func removeRtTag(selection *goquery.Selection) {
	selection.Find("rt").Each(func(_ int, selection *goquery.Selection) {
		selection.Remove()
	})
}

func replaceRubyTag(selection *goquery.Selection) {
	selection.Find("ruby").Each(func(_ int, selection *goquery.Selection) {
		html, e := selection.Html()
		if e == nil {
			selection.ReplaceWithHtml(html)
		}
	})
}

const cssStyle = `
	<style>
		body {
			padding: 2.0rem; 
			font-family: "Helvetica Neue", Helvetica, Arial, "PingFang SC", "Hiragino Sans GB", "Heiti SC", "Microsoft YaHei", "WenQuanYi Micro Hei", sans-serif;
                word-wrap: break-word;
		}
		h1 {
			font-size: 2.0rem;
			line-height: 4.0rem;
		}
		p {
			font-size: 2.0rem;
			line-height: 4.0rem;
			text-indent: 4.0rem;
		}
		rt {
			font-size: 1.8rem;
		}
		.under {
			padding-bottom: 4px;
			border-bottom: 2px solid #000;
		}
		.colorN {
			color: #35a16b;
		}
		.colorL {
			color: #ff7f00;
		}
		.colorC {
			color: #0041cc;
		}
	</style>
`

const cssStyleWithRuby = `
	<style>
		body {
			padding: 2.0rem; 
			font-family: "Helvetica Neue", Helvetica, Arial, "PingFang SC", "Hiragino Sans GB", "Heiti SC", "Microsoft YaHei", "WenQuanYi Micro Hei", sans-serif;
			word-wrap: break-word;
		}
		h1 {
			font-size: 2.0rem;
			line-height: 5.8rem;
		}
		p {
			font-size: 2.0rem;
			line-height: 5.8rem;
			text-indent: 4.0rem;
		}
		rt {
			font-size: 1.8rem;
		}
		.under {
			padding-bottom: 4px;
			border-bottom: 2px solid #000;
		}
		.colorN {
			color: #35a16b;
		}
		.colorL {
			color: #ff7f00;
		}
		.colorC {
			color: #0041cc;
		}
	</style>
`
