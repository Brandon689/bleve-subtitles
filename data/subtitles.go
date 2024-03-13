package data

import (
	"fmt"
	"github.com/Brandon689/bleve-subtitles/types"
	"github.com/asticode/go-astisub"
	"os"
	"path/filepath"
	"strings"
)

func GetSubtitles(dir string) []types.Subtitle {
	var subs []types.Subtitle
	err := filepath.Walk(dir, func(fp string, fi os.FileInfo, err error) error {
		_, currentSubs := visitFile(fp, fi, err)
		subs = append(subs, currentSubs...)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", dir, err)
	}
	return subs
}

func visitFile(fp string, fi os.FileInfo, err error) (error, []types.Subtitle) {
	var Subs []types.Subtitle
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil, Subs // ignore error
	}
	if fi.IsDir() {
		return nil, Subs // ignore directories
	}
	s1, _ := astisub.OpenFile(fp)

	for i := range s1.Items {
		var allLines []string

		for j := range s1.Items[i].Lines {
			allLines = append(allLines, s1.Items[i].Lines[j].String())
		}
		fileName := strings.TrimSuffix(filepath.Base(fp), filepath.Ext(fp))
		if len(allLines) != 0 {
			r := types.Subtitle{
				Lines2:  allLines[0],
				Lines:   allLines,
				Episode: fileName,
				StartAt: s1.Items[i].StartAt,
				EndAt:   s1.Items[i].EndAt,
			}
			//fmt.Println(s1.Items[i].Region)
			//fmt.Println(s1.Items[i].Lines[0].VoiceName)
			Subs = append(Subs, r)
		}
		//} else {
		//	fmt.Println(s1.Items[i].StartAt)
		//	}
	}
	return nil, Subs
}
