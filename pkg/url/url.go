package url

import "regexp"

const URL_REGEXP = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`

func IsValidURL(url string) bool {
	if len(url) == 0 {
		return false
	}

	matched, err := regexp.Match(URL_REGEXP, []byte(url))
	if err != nil {
		return false
	}
	return matched
}
