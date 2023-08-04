package legado

type BookSource struct {
	BookSourceComment string `json:"bookSourceComment,omitempty"`
	BookSourceGroup   string `json:"bookSourceGroup"`
	BookSourceName    string `json:"bookSourceName"`
	BookSourceType    int    `json:"bookSourceType"`
	BookSourceUrl     string `json:"bookSourceUrl"`
	BookUrlPattern    string `json:"bookUrlPattern,omitempty"`
	CustomOrder       int    `json:"customOrder"`
	Enabled           bool   `json:"enabled"`
	EnabledCookieJar  bool   `json:"enabledCookieJar"`
	EnabledExplore    bool   `json:"enabledExplore"`
	ExploreUrl        string `json:"exploreUrl,omitempty"`
	Header            string `json:"header,omitempty"`
	LastUpdateTime    int64  `json:"lastUpdateTime"`
	RespondTime       int    `json:"respondTime"`
	RuleBookInfo      struct {
		Author      string `json:"author,omitempty"`
		CoverUrl    string `json:"coverUrl,omitempty"`
		Init        string `json:"init,omitempty"`
		Intro       string `json:"intro,omitempty"`
		Kind        string `json:"kind,omitempty"`
		LastChapter string `json:"lastChapter,omitempty"`
		Name        string `json:"name,omitempty"`
		TocUrl      string `json:"tocUrl,omitempty"`
		WordCount   string `json:"wordCount,omitempty"`
		CanReName   string `json:"canReName,omitempty"`
	} `json:"ruleBookInfo"`
	RuleContent struct {
		Content        string `json:"content,omitempty"`
		NextContentUrl string `json:"nextContentUrl,omitempty"`
		ReplaceRegex   string `json:"replaceRegex,omitempty"`
		SourceRegex    string `json:"sourceRegex,omitempty"`
		ImageStyle     string `json:"imageStyle,omitempty"`
		WebJs          string `json:"webJs,omitempty"`
	} `json:"ruleContent"`
	RuleExplore struct {
		Author      string `json:"author,omitempty"`
		BookList    string `json:"bookList,omitempty"`
		BookUrl     string `json:"bookUrl,omitempty"`
		CoverUrl    string `json:"coverUrl,omitempty"`
		Intro       string `json:"intro,omitempty"`
		Kind        string `json:"kind,omitempty"`
		LastChapter string `json:"lastChapter,omitempty"`
		Name        string `json:"name,omitempty"`
		WordCount   string `json:"wordCount,omitempty"`
	} `json:"ruleExplore"`
	RuleReview struct {
	} `json:"ruleReview"`
	RuleSearch struct {
		Author       string `json:"author"`
		BookList     string `json:"bookList"`
		BookUrl      string `json:"bookUrl"`
		CheckKeyWord string `json:"checkKeyWord,omitempty"`
		CoverUrl     string `json:"coverUrl,omitempty"`
		Intro        string `json:"intro,omitempty"`
		Kind         string `json:"kind,omitempty"`
		LastChapter  string `json:"lastChapter,omitempty"`
		Name         string `json:"name"`
		WordCount    string `json:"wordCount,omitempty"`
	} `json:"ruleSearch"`
	RuleToc struct {
		ChapterList string `json:"chapterList"`
		ChapterName string `json:"chapterName"`
		ChapterUrl  string `json:"chapterUrl"`
		UpdateTime  string `json:"updateTime,omitempty"`
		IsVip       string `json:"isVip,omitempty"`
		IsVolume    string `json:"isVolume,omitempty"`
		NextTocUrl  string `json:"nextTocUrl,omitempty"`
		PreUpdateJs string `json:"preUpdateJs,omitempty"`
		IsPay       string `json:"isPay,omitempty"`
	} `json:"ruleToc"`
	SearchUrl       string `json:"searchUrl"`
	Weight          int    `json:"weight"`
	LoginUrl        string `json:"loginUrl,omitempty"`
	VariableComment string `json:"variableComment,omitempty"`
}

type UrlOption struct {
	Method  string            `json:"method,omitempty"`
	Charset string            `json:"charset,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    string            `json:"body,omitempty"`
}
