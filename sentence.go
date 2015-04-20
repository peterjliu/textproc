package textproc

import "strings"

var punctuation map[string]bool = map[string]bool{
	".": true,
	"?": true,
	"!": true,
}
var markers map[string]string = map[string]string{
	"\"": "\"",
	"(":  ")",
	"[":  "]",
	"{":  "}",
}

func sentencesV1(input string) (sentences []string) {
	var current string
	currentMarker := ""
	r := strings.NewReplacer("\n", "", "\r", "")

	for _, c := range input {
		current += string(c)

		_, isOpener := markers[string(c)]
		if isOpener && len(currentMarker) == 0 {
			currentMarker = string(c)
		} else if len(currentMarker) > 0 {
			closer := markers[currentMarker]
			if closer == string(c) {
				currentMarker = ""
			}
		}

		_, ok := punctuation[string(c)]
		if ok {
			// don't split if we're inside markers:
			if len(currentMarker) > 0 {
				continue
			}
			current = r.Replace(current)
			sentences = append(sentences, current)
			current = ""
		}
	}
	if len(current) > 0 {
		current = r.Replace(current)
		sentences = append(sentences, current)
	}
	return sentences
}
