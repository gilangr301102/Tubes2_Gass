package utils

type ScrappingData struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type ScrappingDatas [][]ScrappingData

type JsonError struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type ArticleInfo1 struct {
	Article       string
	ParentArticle string
}

type ArticleNode struct {
	Name  string
	Level int
}

type ArticleInfo2 struct {
	ArticleNameWithParent ArticleInfo1
	Level                 int
}

type Node struct {
	Title string
	URL   string
	Path  []string
}
