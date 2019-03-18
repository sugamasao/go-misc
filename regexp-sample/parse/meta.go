package parse

import (
	"regexp"
	"strings"
)

// DetectMetaTagCharset meta タグから文字コードを抽出する
func DetectMetaTagCharset(file string) string {
	reg := regexp.MustCompile(`(?i)<meta.+charset="?([\w\-]+?)["\W\/]*?>`)
	ret := reg.FindAllStringSubmatch(file, -1)

	return strings.ToUpper(ret[0][1])
}
