package parse_test

import (
	"io/ioutil"
	"testing"

	"github.com/sugamasao/go-misc/regexp-sample/parse"
)

func TestDetectMetaTagCharset(t *testing.T) {
	cases := []struct {
		path   string
		encode string
	}{
		{path: "./../testdata/qiita.html", encode: "UTF-8"},
		{path: "./../testdata/github.html", encode: "UTF-8"},
		{path: "./../testdata/takebouki_01.html", encode: "EUC_JP"},
		{path: "./../testdata/takebouki_02.html", encode: "UTF-8"},
	}

	for _, tt := range cases {
		file, err := ioutil.ReadFile(tt.path)
		if err != nil {
			t.Errorf("file open failed. %v", err)
		}
		ret := parse.DetectMetaTagCharset(string(file))
		if tt.encode != ret {
			t.Errorf("got %v, wont %s(%s)", ret, tt.path, tt.encode)
		}
	}
}
