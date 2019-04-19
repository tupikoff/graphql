package graphql

import (
	"encoding/json"
	"github.com/matryer/is"
	"strings"
	"testing"
)

func TestFillMultipartRequestSpec(t *testing.T) {
	is := is.New(t)

	req := NewRequest("{query}")
	f := strings.NewReader(`This is a file`)
	req.File("file1", "filename1.txt", f)
	req.File("file2", "filename2.txt", f)

	mprs := req.fillMultipartRequestSpecQuery()

	operations, e := json.Marshal(mprs.Operations)
	is.NoErr(e)
	is.Equal(`{"query":"{query}","variables":{"files":[null,null]}}`, string(operations))

	maps, e := json.Marshal(mprs.Map)
	is.NoErr(e)
	is.Equal(`{"file1":["variables.files.0"],"file2":["variables.files.1"]}`, string(maps))
}

func TestFillMultipartRequestSpecNoFiles(t *testing.T) {
	is := is.New(t)

	req := NewRequest("{query}")

	mprs := req.fillMultipartRequestSpecQuery()

	operations, e := json.Marshal(mprs.Operations)
	is.NoErr(e)
	is.Equal(`{"query":"{query}","variables":{}}`, string(operations))

	maps, e := json.Marshal(mprs.Map)
	is.NoErr(e)
	is.Equal(`{}`, string(maps))
}
