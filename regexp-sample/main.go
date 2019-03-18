package main

import (
	"log"

	"github.com/sugamasao/go-misc/regexp-sample/parse"
)

func main() {
	str := `
	<html>
		<meta charset="EUC_JP" />
	</html>
	`
	log.Printf("[%s]", parse.DetectMetaTagCharset(str))
}
