
package llm

type SplitResponse struct {
	Gujarati string
	English  string
	Hindi    string
}

func SplitLanguages(text string) (*SplitResponse, error) {
	return &SplitResponse{
		Gujarati: text,
		English:  "",
		Hindi:    "",
	}, nil
}
