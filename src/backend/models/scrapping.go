package models

type ScrappingData struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type ScrappingDatas []ScrappingData
