package utils

const WIKIPEDIA_URL_EN = "https://en.wikipedia.org"

const URL_SCRAPPING_WIKIPEDIA = WIKIPEDIA_URL_EN + "/wiki/"

const DEFAULT_ROOT_VAL = "root"

var noToVisit = map[string]bool{
	"/wiki/Special:Random": true,
}

const NumOfNodeWORKERS = 32
const MAX_HOS_CONNECTION = 100
const TIMEOUT = 10
const MAX_MEMO_SIZE = 5000000
const SIZE_OF_CONTAINER = 500
