package goExamples

func GetStringOccurrences(searchString []string) (map[string]int, error) {

	duplicationFrequency := make(map[string]int)

	for _, item := range searchString {
		// check if item is in the duplication list already

		_, exists := duplicationFrequency[item]

		if exists {
			duplicationFrequency[item] += 1
		} else {
			duplicationFrequency[item] = 1
		}
	}
	return duplicationFrequency, nil
}
