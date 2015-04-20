package textproc

import (
	"strings"
	"testing"
)

func TestSentenceWithNames(t *testing.T) {
	paragraph := `We know that the periods in Mr. Smith and Johann S. Bach
do not mark sentence boundaries.  And sometimes sentences
can start with non-capitalized words.  i is a good variable
name.`

	expected := []string{`We know that the periods in Mr. Smith and Johann S. Bach
do not mark sentence boundaries.`,
		`And sometimes sentences
can start with non-capitalized words.`,
		`i is a good variable name.`}

	split := sentencesV1(paragraph)

	if len(expected) != len(split) {
		t.Errorf("Expected to be split into %d sentences, but got %d.", len(expected), len(split))
	}

	if strings.Join(expected, "\n") != strings.Join(split, "\n") {
		t.Errorf("Paragraph incorrectly segmented: \n%s", strings.Join(split, "\n--------\n"))
	}
}
