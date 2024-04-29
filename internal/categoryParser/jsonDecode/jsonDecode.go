package jsonDecode

import (
	"encoding/json"
	"io"
	"os"
)

// TODO: вынести потом в отдельный файл потому что где-то еще юзается эта хуйня
const (
	math  = "math"
	rus   = "rus"
	soc   = "soc"
	bio   = "bio"
	chem  = "chem"
	info  = "info"
	hist  = "hist"
	lit   = "lit"
	phys  = "phys"
	mathb = "mathb"
)

type Category struct {
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type Topic struct {
	TopicId    string     `json:"topic_id"`
	Categories []Category `json:"categories"`
	TopicName  string     `json:"topic_name"`
}

type SubjInfo struct {
	subjStr string
	topics  []Topic
}

func New(subjName string, path string) *SubjInfo {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var result []Topic

	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	jsonErr := json.Unmarshal(data, &result)
	if jsonErr != nil {
		panic(jsonErr)
	}

	return &SubjInfo{
		subjStr: subjName,
		topics:  result,
	}

}
