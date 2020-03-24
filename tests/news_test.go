package tests

import (
	"encoding/json"
	"fmt"
	"jpnews/internal/app/model"
	"sort"
	"testing"
)

//func TestTopNewsUrl(t *testing.T) {
//	fmt.Println(news.TopNews(time.Now().UnixNano() / 1e6))
//}

//func TestParseNews(t *testing.T) {
//	news, newsWithRuby := news.ParseNews(htmlContent)
//
//	fmt.Println("news =============")
//	fmt.Println(news)
//	fmt.Println("newsWithRuby =============")
//	fmt.Println(newsWithRuby)
//
//}

//func TestUnmarshalJSONTopNews(t *testing.T) {
//
//	//s := "\uFEFFGG"
//	//fmt.Println(len(s))
//	//fmt.Println(utf8.RuneCountInString(s))
//	//
//	//b := []byte(s)
//	//fmt.Println(b)
//	//
//	//fmt.Printf("%x\n", []rune(string(b)))
//	//
//	//for i, r := range s {
//	//	fmt.Printf("%d\t%q\t%d\n", i, r, r)
//	//}
//	//
//	//r, size := utf8.DecodeRune(b)
//	//fmt.Println(r)
//	//fmt.Println(size)
//
//	jsonStr := `
//[{"top_priority_number":"1","top_display_flag":true,"news_id":"k10012033071000","news_prearranged_time":"2019-08-13 12:35:00","title":"\u53f0\u98a8\uff11\uff10\u53f7\u304c\u6765\u308b\u3000\u96e8\u3084\u98a8\u306b\u6c17\u3092\u3064\u3051\u3066","title_with_ruby":"<ruby>\u53f0\u98a8<rt>\u305f\u3044\u3075\u3046<\/rt><\/ruby>\uff11\uff10<ruby>\u53f7<rt>\u3054\u3046<\/rt><\/ruby>\u304c<ruby>\u6765<rt>\u304f<\/rt><\/ruby>\u308b\u3000<ruby>\u96e8<rt>\u3042\u3081<\/rt><\/ruby>\u3084<ruby>\u98a8<rt>\u304b\u305c<\/rt><\/ruby>\u306b<ruby>\u6c17<rt>\u304d<\/rt><\/ruby>\u3092\u3064\u3051\u3066","outline_with_ruby":"<p><ruby>\u6c17\u8c61\u5e81<rt>\u304d\u3057\u3087\u3046\u3061\u3087\u3046<\/rt><\/ruby>\u306b\u3088\u308b\u3068\u3001\u3068\u3066\u3082<ruby>\u5927<rt>\u304a\u304a<\/rt><\/ruby>\u304d\u304f\u3066<ruby>\u5f37<rt>\u3064\u3088<\/rt><\/ruby>\u3044<ruby>\u53f0\u98a8<rt>\u305f\u3044\u3075\u3046<\/rt><\/ruby>\uff11\uff10<ruby>\u53f7<rt>\u3054\u3046<\/rt><\/ruby>\u304c<ruby>\uff11\uff14\u65e5<rt>\u3058\u3085\u3046\u3088\u3063\u304b<\/rt><\/ruby>\u304b\u3089\uff11\uff15<ruby>\u65e5<rt>\u306b\u3061<\/rt><\/ruby>\u306b<a href='javascript:void(0)' class='dicWin' id='id-0000'><ruby>\u8fd1\u757f\u5730\u65b9<rt>\u304d\u3093\u304d\u3061\u307b\u3046<\/rt><\/ruby><\/a>\u304b\u3089<ruby>\u897f<rt>\u306b\u3057<\/rt><\/ruby><ruby>\u5074<rt>\u304c\u308f<\/rt><\/ruby>\u306b<ruby>\u6765<rt>\u304d<\/rt><\/ruby>\u305d\u3046\u3067\u3059\u3002","news_file_ver":false,"news_publication_status":true,"has_news_web_image":false,"has_news_web_movie":false,"has_news_easy_image":true,"has_news_easy_movie":false,"has_news_easy_voice":true,"news_web_image_uri":"","news_web_movie_uri":"","news_easy_image_uri":"k10012033071000.jpg","news_easy_movie_uri":"''","news_easy_voice_uri":"k10012033071000.mp4"}]
//`
//
//
//	var news []model.TopNews
//	err := json.Unmarshal([]byte(jsonStr), &news)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	fmt.Println("un json TopNews ===== ")
//	fmt.Printf("%#v\n", news)
//
//}

func TestNewsUrl(t *testing.T) {
	//fmt.Println(newsurl.NewsDetails("K1232131"))

}

func TestNewsListJSON(t *testing.T) {
	var news []map[string][]model.NewsItem
	unJsonErr := json.Unmarshal([]byte(newsList), &news)
	if unJsonErr != nil {
		t.Fatal(unJsonErr)
	}

	//mapData := news[0]
	//fmt.Printf("NewsList first = %#v\n", mapData["2019-08-13"])

	var respData []model.NewsDate

	for _, item := range news {
		var keys []string
		for key := range item {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, v := range keys {
			data := model.NewsDate{
				Date: v,
				Data: item[v],
			}
			respData = append(respData, data)
		}
	}

	bytes, err := json.Marshal(respData)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("NewsList ======")
	fmt.Println(string(bytes))
	//fmt.Printf(" %#v\n", respData)

}

const newsList = `
[{
	"2019-08-13": [{
		"news_priority_number": "1",
		"news_prearranged_time": "2019-08-13 12:35:00",
		"news_id": "k10012033071000",
		"title": "\u53f0\u98a8\uff11\uff10\u53f7\u304c\u6765\u308b\u3000\u96e8\u3084\u98a8\u306b\u6c17\u3092\u3064\u3051\u3066",
		"title_with_ruby": "<ruby>\u53f0\u98a8<rt>\u305f\u3044\u3075\u3046<\/rt><\/ruby>\uff11\uff10<ruby>\u53f7<rt>\u3054\u3046<\/rt><\/ruby>\u304c<ruby>\u6765<rt>\u304f<\/rt><\/ruby>\u308b\u3000<ruby>\u96e8<rt>\u3042\u3081<\/rt><\/ruby>\u3084<ruby>\u98a8<rt>\u304b\u305c<\/rt><\/ruby>\u306b<ruby>\u6c17<rt>\u304d<\/rt><\/ruby>\u3092\u3064\u3051\u3066",
		"news_file_ver": false,
		"news_creation_time": "2019-08-13 16:47:48",
		"news_preview_time": "2019-08-13 16:47:48",
		"news_publication_time": "2019-08-13 13:51:02",
		"news_publication_status": true,
		"has_news_web_image": false,
		"has_news_web_movie": false,
		"has_news_easy_image": true,
		"has_news_easy_movie": false,
		"has_news_easy_voice": true,
		"news_web_image_uri": "",
		"news_web_movie_uri": "",
		"news_easy_image_uri": "k10012033071000.jpg",
		"news_easy_movie_uri": "''",
		"news_easy_voice_uri": "k10012033071000.mp4",
		"news_display_flag": true,
		"news_web_url": "https:\/\/www3.nhk.or.jp\/news\/html\/20190813\/k10012033071000.html"
	}]
}]
`

const htmlContent = `
<!DOCTYPE HTML>
<!--[if lt IE 7]><html class="no-js lt-ie9 lt-ie8 lt-ie7 eq-ie6"><![endif]-->
<!--[if IE 7]><html class="no-js lt-ie9 lt-ie8 eq-ie7"> <![endif]-->
<!--[if IE 8]><html class="no-js lt-ie9 eq-ie8"> <![endif]-->
<!--[if gt IE 8]><!--><html class="no-js"><!--<![endif]-->
<head prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# website: http://ogp.me/ns/website#">
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta http-equiv="X-UA-Compatible" content="requiresActiveX=true">
<meta name="viewport" content="width=device-width,initial-scale=1">
<meta name="description" content="９日、長崎に原爆が落とされてから７４年になりました。原爆で亡くなった人は１８万２６０１人です。この中で、去年からの１年に亡くなった人は３４０...">
<title>NEWS WEB EASY|長崎に原爆が落とされてから７４年　平和を祈る式</title>
<meta property="og:site_name" content="NEWS WEB EASY" />
<meta property="og:title" content="NEWS WEB EASY|長崎に原爆が落とされてから７４年　平和を祈る式" />
<meta property="og:locale" content="ja_JP" />
<meta property="og:description" content="９日、長崎に原爆が落とされてから７４年になりました。原爆で亡くなった人は１８万２６０１人です。この中で、去年からの１年に亡くなった人は３４０..." />
<meta property="og:url" content="https://www3.nhk.or.jp/news/easy/k10012029391000/k10012029391000.html" />
<meta property="og:image" content="https://www3.nhk.or.jp/news/easy/images/og-image.png" />
<meta property="og:type" content="article" />
<link rel="shortcut icon" type="image/x-icon" href="//www.nhk.or.jp/favicon.ico">
<link rel="canonical" href="https://www3.nhk.or.jp/news/easy/k10012029391000/k10012029391000.html" />
<script src="/common/js/nol_tools.js"></script>
<link rel="stylesheet" href="../css/style-v2.css">
<script src="https://www3.nhk.or.jp/common/jquery/jquery-2.2.js"></script>
<script src="https://www3.nhk.or.jp/common/sp/nol_SmartPhone.js" charset="utf-8"></script>
<script src="https://www3.nhk.or.jp/common/sns/nol_share.js" charset="UTF-8"></script>
<script src="https://www3.nhk.or.jp/news/r/js/exist.js" charset="UTF-8"></script>
<script>NHKSNS.initSNSOnLoad();</script>
<script type="text/javascript">var _sf_startpt = (new Date()).getTime();var _cb_sections_preset = 'EASY';</script>
</head>
<body id="news20190809_k10012029391000">
<aside id="nhkheader">
<script>
nol_showCmnHeader({
design: 'gray'
      });
</script>
</aside>
<div id="wrapper">
<div id="content">
<div class="easy-wrapper" id="easy-wrapper">
<div id="mediaFlg"></div>
<header class="easy-header">
<div class="easy-header-title">
<div class="easy-header-title__inner">
<h1 class="easy-header-title__title">
<a href="../">
<img src="../images/logo_pc.png" alt="NEWS WEB EASY" class="easy-logo">
</a>
<img src="../images/subtitle_pc.png" alt="やさしい日本語で書いたニュース" class="easy-subtitle">
</h1>
</div>
</div>
<a href="#disaster-list" class="link-to-disaster js-smooth-scroll show-sp"><ruby>災害<rt>さいがい</rt></ruby>に<ruby>気<rt>き</rt></ruby>をつけて</a>
<div class="header-about-easy js-accordion-wrapper">
<div class="header-about-easy__inner">
<a href="#" class="header-about-easy__toggle js-toggle-accordion">
<!-- <i><img src="../images/icon_about_pc.png" alt="?"></i> -->
NEWS WEB EASYについて
</a>
</div>
<div class="header-about-easy__body about-easy-body js-accordion-body">
<div class="about-easy-body__inner">
<p>
「ＮＥＷＳ ＷＥＢ ＥＡＳＹ」は<ruby>外国人<rt>がいこくじん</rt></ruby>の<ruby>皆<rt>みな</rt></ruby>さんや、<ruby>小学生<rt>しょうがくせい</rt></ruby>・<ruby>中学生<rt>ちゅうがくせい</rt></ruby>の<ruby>皆<rt>みな</rt></ruby>さんのために、わかりやすいことばでニュースを<ruby>伝<rt>つた</rt></ruby>えます。
</p>
</div>
</div>
</div>
</header>
<div class="l-container">
<main class="l-main">
<article class="article-main">
<figure class="article-main__figure" id="js-article-figure">
</figure>
<h1 class="article-main__title">
<ruby>長崎<rt>ながさき</rt></ruby>に<ruby>原爆<rt>げんばく</rt></ruby>が<ruby>落<rt>お</rt></ruby>とされてから７４<ruby>年<rt>ねん</rt></ruby>　<ruby>平和<rt>へいわ</rt></ruby>を<ruby>祈<rt>いの</rt></ruby>る<ruby>式<rt>しき</rt></ruby>
</h1>
<p class="article-main__date" id="js-article-date">[08月09日 16時50分]</p>
<div class="article-main__tools">
<div class="tool-buttons">
<a href="#" class="tool-buttons__item btn toggle-audio js-open-audio">ニュースを<ruby>聞<rt>き</rt></ruby>く</a>
<a href="#" class="tool-buttons__item btn toggle-ruby js-toggle-ruby hide-sp"><ruby>漢字<rt>かんじ</rt></ruby>の<ruby>読<rt>よ</rt></ruby>み<ruby>方<rt>かた</rt></ruby>を<ruby>消<rt>け</rt></ruby>す</a>
</div>
<div class="audio-player" id="js-audio-wrapper">
<a href="#" class="audio-player__close js-close-audio"></a>
<div id="js-audio-inner"></div>
</div>
<div class="tool-buttons show-sp">
<a href="#" class="tool-buttons__item btn toggle-ruby js-toggle-ruby"><ruby>漢字<rt>かんじ</rt></ruby>の<ruby>読<rt>よ</rt></ruby>み<ruby>方<rt>かた</rt></ruby>を<ruby>消<rt>け</rt></ruby>す</a>
</div>
</div>
<div class="article-main__body article-body" id="js-article-body">
<p><ruby>９日<rt>ここのか</rt></ruby>、<span class='colorL'><ruby>長崎<rt>ながさき</rt></ruby></span>に<a href='javascript:void(0)' class='dicWin' id='id-0000'><ruby><span class="under">原爆</span><rt>げんばく</rt></ruby></a>が<ruby>落<rt>お</rt></ruby>とされてから７４<ruby>年<rt>ねん</rt></ruby>になりました。<a href='javascript:void(0)' class='dicWin' id='id-0000'><ruby><span class="under">原爆</span><rt>げんばく</rt></ruby></a>で<ruby>亡<rt>な</rt></ruby>くなった<ruby>人<rt>ひと</rt></ruby>は１８<ruby>万<rt>まん</rt></ruby>２６０１<ruby>人<rt>にん</rt></ruby>です。この<ruby>中<rt>なか</rt></ruby>で、<ruby>去年<rt>きょねん</rt></ruby>からの１<ruby>年<rt>ねん</rt></ruby>に<ruby>亡<rt>な</rt></ruby>くなった<ruby>人<rt>ひと</rt></ruby>は３４０２<ruby>人<rt>にん</rt></ruby>です。</p>
<p><span class='colorL'><ruby>長崎市<rt>ながさきし</rt></ruby></span>が<ruby>行<rt>おこな</rt></ruby>った<a href='javascript:void(0)' class='dicWin' id='id-0001'><ruby><span class="under">平和</span><rt>へいわ</rt></ruby></a>を<ruby>祈<rt>いの</rt></ruby>る<ruby>式<rt>しき</rt></ruby>には、５９００<ruby>人<rt>にん</rt></ruby>が<ruby>出席<rt>しゅっせき</rt></ruby>しました。<a href='javascript:void(0)' class='dicWin' id='id-0000'><ruby><span class="under">原爆</span><rt>げんばく</rt></ruby></a>が<ruby>落<rt>お</rt></ruby>とされた<ruby>午前<rt>ごぜん</rt></ruby>１１<ruby>時<rt>じ</rt></ruby>２<ruby>分<rt>ふん</rt></ruby>になると、みんなで<ruby>静<rt>しず</rt></ruby>かに<ruby>祈<rt>いの</rt></ruby>りました。</p>
<p><span class='colorL'><ruby>長崎市<rt>ながさきし</rt></ruby></span>の<span class='colorN'><ruby>田上<rt>たうえ</rt></ruby></span><a href='javascript:void(0)' class='dicWin' id='id-0002'><ruby><span class="under">市長</span><rt>しちょう</rt></ruby></a>は、<a href='javascript:void(0)' class='dicWin' id='id-0003'><ruby><span class="under">核兵器</span><rt>かくへいき</rt></ruby></a>を<ruby>作<rt>つく</rt></ruby>ったり<ruby>使<rt>つか</rt></ruby>ったりすることを<a href='javascript:void(0)' class='dicWin' id='id-0004'><ruby><span class="under">禁止</span><rt>きんし</rt></ruby></a>すると<a href='javascript:void(0)' class='dicWin' id='id-0005'><ruby><span class="under">国連</span><rt>こくれん</rt></ruby></a>で<ruby>決<rt>き</rt></ruby>めた<a href='javascript:void(0)' class='dicWin' id='id-0006'><ruby><span class="under">条約</span><rt>じょうやく</rt></ruby></a>に<ruby>入<rt>はい</rt></ruby>るように、<span class='colorL'><ruby>日本<rt>にっぽん</rt></ruby></span><a href='javascript:void(0)' class='dicWin' id='id-0007'><ruby><span class="under">政府</span><rt>せいふ</rt></ruby></a>に<ruby>言<rt>い</rt></ruby>いました。</p>
<p><span class='colorN'><ruby>安倍<rt>あべ</rt></ruby></span><a href='javascript:void(0)' class='dicWin' id='id-0008'><ruby><span class="under">総理大臣</span><rt>そうりだいじん</rt></ruby></a>は「<a href='javascript:void(0)' class='dicWin' id='id-0003'><ruby><span class="under">核兵器</span><rt>かくへいき</rt></ruby></a>をなくすために、<a href='javascript:void(0)' class='dicWin' id='id-0003'><ruby><span class="under">核兵器</span><rt>かくへいき</rt></ruby></a>を<ruby>持<rt>も</rt></ruby>っている<ruby>国<rt>くに</rt></ruby>と<ruby>持<rt>も</rt></ruby>っていない<ruby>国<rt>くに</rt></ruby>の<ruby>両方<rt>りょうほう</rt></ruby>と<a href='javascript:void(0)' class='dicWin' id='id-0009'><ruby><span class="under">話</span><rt>はな</rt></ruby><span class="under">し</span><ruby><span class="under">合</span><rt>あ</rt></ruby><span class="under">い</span></a>をします」と<ruby>言<rt>い</rt></ruby>いましたが、<a href='javascript:void(0)' class='dicWin' id='id-0005'><ruby><span class="under">国連</span><rt>こくれん</rt></ruby></a>の<a href='javascript:void(0)' class='dicWin' id='id-0006'><ruby><span class="under">条約</span><rt>じょうやく</rt></ruby></a>については<ruby>何<rt>なに</rt></ruby>も<ruby>話<rt>はな</rt></ruby>しませんでした。</p>
<p><span class='colorL'><ruby>長崎市<rt>ながさきし</rt></ruby></span>の７５<ruby>歳<rt>さい</rt></ruby>の<ruby>女性<rt>じょせい</rt></ruby>は「<ruby>私<rt>わたし</rt></ruby>は１<ruby>歳<rt>さい</rt></ruby><ruby>半<rt>はん</rt></ruby>だったので、<ruby>何<rt>なに</rt></ruby>も<ruby>覚<rt>おぼ</rt></ruby>えていません。<ruby>母<rt>はは</rt></ruby>は<a href='javascript:void(0)' class='dicWin' id='id-0000'><ruby><span class="under">原爆</span><rt>げんばく</rt></ruby></a>について<ruby>何<rt>なに</rt></ruby>も<ruby>話<rt>はな</rt></ruby>さないで<ruby>亡<rt>な</rt></ruby>くなりました。<ruby>父<rt>ちち</rt></ruby>は<a href='javascript:void(0)' class='dicWin' id='id-0000'><ruby><span class="under">原爆</span><rt>げんばく</rt></ruby></a>で<ruby>亡<rt>な</rt></ruby>くなったので<ruby>顔<rt>かお</rt></ruby>を<ruby>見<rt>み</rt></ruby>ることもできませんでした。<ruby>戦争<rt>せんそう</rt></ruby>のない<a href='javascript:void(0)' class='dicWin' id='id-0001'><ruby><span class="under">平和</span><rt>へいわ</rt></ruby></a>な<ruby>世界<rt>せかい</rt></ruby>になってほしいです」と<ruby>話<rt>はな</rt></ruby>していました。</p>
<p></p>
<p></p>
</div>
<div class="article-main__colors">
<div class="about-word-color">
<ul class="about-word-color__list word-color-list">
<li class="word-color-list__item is-person">
…<ruby>人<rt>ひと</rt></ruby>の<ruby>名前<rt>なまえ</rt></ruby>
</li>
<li class="word-color-list__item is-place">
…<ruby>国<rt>くに</rt></ruby>や<ruby>県<rt>けん</rt></ruby>、<ruby>町<rt>まち</rt></ruby>、<ruby>場所<rt>ばしょ</rt></ruby>などの<ruby>名前<rt>なまえ</rt></ruby>
</li>
<li class="word-color-list__item is-group">
…<ruby>会社<rt>かいしゃ</rt></ruby>やグループなどの<ruby>名前<rt>なまえ</rt></ruby>
</li>
</ul>
<div class="about-word-color__toggle word-color-toggle">
<a href="#" class="word-color-toggle__switch btn" id="js-toggle-color">ことばの<ruby>色<rt>いろ</rt></ruby>を<ruby>消<rt>け</rt></ruby>す</a>
</div>
</div>
<p class="about-word-color__caution hide-sp">
※<span class="dicWin"><ruby>下<rt>した</rt></ruby>に<ruby>線<rt>せん</rt></ruby>があることば</span>は<ruby>辞書<rt>じしょ</rt></ruby>の<ruby>説明<rt>せつめい</rt></ruby>を<ruby>見<rt>み</rt></ruby>ることができます。
</p>
<p class="about-word-color__caution show-sp">
※<span class="dicWin"><ruby>下<rt>した</rt></ruby>に<ruby>線<rt>せん</rt></ruby>があることば</span>を<ruby>押<rt>お</rt></ruby>すと、<ruby>辞書<rt>じしょ</rt></ruby>の<ruby>説明<rt>せつめい</rt></ruby>を<ruby>見<rt>み</rt></ruby>ることができます。
</p>
</div>
<nav class="article-main__links">
<div class="share">
<div class="nhk-snsbtn" data-nhksns-disable="google" data-nhksns-description=" "></div>
</div>
<div class="link-to-normal" id="js-regular-news-wrapper">
<a href="https://www3.nhk.or.jp/news/html/20190809/k10012029391000.html" class="btn" target="_blank" id="js-regular-news"><ruby>普通<rt>ふつう</rt></ruby>のニュースを<ruby>読<rt>よ</rt></ruby>む</a>
</div>
</nav>
</article>
<section class="enquete">
<h2 class="enquete__title">アンケート</h2>
<iframe style="border: none;width:100%;" scrolling="no" src="https://cgi2.nhk.or.jp/news/easy/easy_enq/bin/form/enqform.html?id=k10012029391000&amp;title=長崎に原爆が落とされてから７４年　平和を祈る式" id="js-enquete-frame"></iframe>
</section>
</main>
<aside class="l-sidebar sidebar">
<section class="sidebar-section side-latest">
<h1 class="sidebar-section__title"><ruby>新<rt>あたら</rt></ruby>しいニュース</h1>
<div class="side-latest__list">
<div class="side-news-list" id="js-latest-news">
</div>
</div>
</section>
<section class="sidebar-section archives">
<h1 class="sidebar-section__title">１か<ruby>月<rt>げつ</rt></ruby>のニュース</h1>
<div class="archives-pager">
<a href="#" class="archives-pager__arrow is-prev js-pager-nav"><ruby>次<rt>つぎ</rt></ruby>の<ruby>日<rt>ひ</rt></ruby></a>
<p class="archives-pager__date" id="js-pager-date"></p>
<a href="#" class="archives-pager__arrow is-next js-pager-nav"><ruby>前<rt>まえ</rt></ruby>の<ruby>日<rt>ひ</rt></ruby></a>
</div>
<div class="archives-list" id="js-archives-list">
</div>
</section>
<section class="sidebar-section side-disaster" id="disaster-list">
<h1 class="sidebar-section__title"><ruby>災害<rt>さいがい</rt></ruby>に<ruby>気<rt>き</rt></ruby>をつけて</h1>
<ul class="disaster-list js-disaster-list">
</ul>
</section>
<section class="sidebar-section side-about-easy show-sp js-accordion-wrapper">
<a href="#" class="side-about-easy__toggle js-toggle-accordion">NEWS WEB EASYについて</a>
<div class="side-about-easy__body js-accordion-body">
<p>
「ＮＥＷＳ ＷＥＢ ＥＡＳＹ」は<ruby>外国人<rt>がいこくじん</rt></ruby>の<ruby>皆<rt>みな</rt></ruby>さんや、<ruby>小学生<rt>しょうがくせい</rt></ruby>・<ruby>中学生<rt>ちゅうがくせい</rt></ruby>の<ruby>皆<rt>みな</rt></ruby>さんのために、わかりやすいことばでニュースを<ruby>伝<rt>つた</rt></ruby>えるウェブサイトです。
</p>
</div>
</section>
<section class="sidebar-section about-dictionary js-accordion-wrapper">
<a href="#" class="about-dictionary__toggle js-toggle-accordion"><ruby>辞書<rt>じしょ</rt></ruby>について</a>
<div class="about-dictionary__body js-accordion-body">
<p>
「NEWS WEB EASY」では、<ruby>三省堂<rt>さんせいどう</rt></ruby>の『<ruby>例解小学国語辞典第<rt>れいかいしょうがくこくごじてんだい</rt></ruby>5<ruby>版<rt>はん</rt></ruby>』を<ruby>使<rt>つか</rt></ruby>っています。<br>
<ruby>辞書<rt>じしょ</rt></ruby>の<ruby>著作権<rt>ちょさくけん</rt></ruby>は、<ruby>辞書<rt>じしょ</rt></ruby>を<ruby>作<rt>つく</rt></ruby>った<ruby>株式会社<rt>かぶしきがいしゃ</rt></ruby>　<ruby>三省堂<rt>さんせいどう</rt></ruby>にあります。
</p>
</div>
</section>
<nav class="sidebar-section link-banners">
<a href="https://www3.nhk.or.jp/news/" class="link-banners__item is-newsweb">
<img src="../images/banner_newsweb_pc.png" alt="NHK NEWS WEB" class="js-responsive-image">
</a>
</nav>
</aside>
<div class="dictionary-box" id="js-dictionary-box">
<div id="dicContent">
</div>
</div>
</div>
</div>
<aside id="nhkfooter" style="max-width: 100%; width: 100%; padding: 0px; margin: 0px;">
<script>
nol_showCmnFooter({
mode: 'simple'
          });
</script>
</aside>
</div>
<!--/ #content -->
</div>
<script src="../js/js.cookie.js"></script>
<script src="../js/common.js"></script>
<script src="../js/article.js"></script>
<script src="/news/parts16/js/chartbeat_config.js" charset="UTF-8"></script>
</body>
</html>
`
