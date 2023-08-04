package legado

import (
	"SourceConvert/tools"
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func (c SourceConvert) ConvertToZhishu() {
	for _, source := range c.srcBookRule {
		createSearch(source)
	}
}

// ==UserScript==
// @name          示例
// @domain        example.com
// @description   这是一个示例
// @version       1.0.0
// @icon          https://example.com/favicon.ico
// @supportURL    https://github.com/open-book-source/booksource/issues
// @require       https://lf6-cdn-tos.bytecdntp.com/cdn/expire-1-M/crypto-js/4.1.1/crypto-js.min.js
// @function      categories
// @function      search
// @function      detail
// @function      toc
// @function      chapter
// @function      authorization
// @function      profile
// @function      author
// @function      schedules
// ==/UserScript==
func createSearch(source BookSource) {
	builder := bytes.Buffer{}
	builder.WriteString(`async function search(keyword, opaque) {
`)
	if strings.Contains(source.SearchUrl, "{{page}}") {
		builder.WriteString(`    let page = opaque ? opaque.page : 1;
`)
	}
	if strings.HasPrefix(source.SearchUrl, "@js") {
		panic("unsupport type f5umadzvez")
	}
	var urlOption UrlOption
	urlRegex := regexp.MustCompile("(\\S+)\\s*,+\\s*(?is:(\\{.+))")
	matchString := urlRegex.FindStringSubmatch(source.SearchUrl)

	builder.WriteString("    let resp = await fetch(`")
	if len(matchString) == 0 {
		builder.WriteString(formatSearchUrl(source.SearchUrl) + "`)\n")
	} else {
		builder.WriteString(formatSearchUrl(matchString[1]) + "`")
		builder.WriteString(", {\n")
		err := json.Unmarshal([]byte(strings.ReplaceAll(matchString[2], "'", "\"")), &urlOption)
		tools.PanicIfError(err)
		if urlOption.Method != "" {
			builder.WriteString("        method: '" + urlOption.Method + "',\n")
		}
		if len(urlOption.Headers) > 0 {
			builder.WriteString("        headers: {\n")
			for key, value := range urlOption.Headers {
				builder.WriteString("            '" + key + "': '" + value + "',")
			}
			builder.WriteString("    },\n")
		}
		if urlOption.Body != "" {
			builder.WriteString("        data: `" + formatSearchUrl(urlOption.Body) + "`,\n")
		}
		builder.WriteString("    });")
	}
	builder.WriteString(`    if (resp.status !== 200) {
        return {
            code: resp.status,
            message: 'Network error!',
        };
    }`)

	fmt.Println(builder.String())
}

func formatSearchUrl(url string) string {
	searchUrl := strings.ReplaceAll(url, `{{key}}`, `${keyword}`)
	searchUrl = strings.ReplaceAll(url, `{{page}}`, `${page}`)
	return searchUrl
}
