package utils

type ArticleInfo struct {
	Article       string
	ParentArticle string
}

type Node struct {
	Title string
	URL   string
	Path  []string
}

type Result struct {
	Results           []string `json:"results"`
	ArticlesChecked   int      `json:"articlesChecked"`
	ArticlesTraversed int      `json:"articlesTraversed"`
	NumberPath        int      `json:"numberPath"`
	ElapsedTime       float64  `json:"elapsedTime"`
}
