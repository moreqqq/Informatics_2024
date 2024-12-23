package laba8

func SearchingString(filename, valueForSearch string) (int, bool) {
	if valueForSearch == "" {
		return 0, false
	}
	values := ReadDataFromFile(filename)

	for _, value := range values {
		if value == valueForSearch {
			return 1, true
		}
	}
	return 0, false
}
