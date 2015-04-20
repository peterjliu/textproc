package textproc

// Get sentences stripped of white-space and from text.
func GetSentences(text string) []string {
	candidates := sentencesV1(text)
	//result := make([]string, 0, cap(candidates))
	return candidates
}
