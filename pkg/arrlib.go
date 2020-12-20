package pkg

func CopyArrString (arr []string) []string {
	var copied  []string
	for _, e := range arr {
		copied = append(copied, e)
	}
	return copied
}

func MergeMaps(original, mapToMerge map[string]bool) map[string]bool {
	for k, v := range mapToMerge {
		original[k] = v
	}
	return original
}