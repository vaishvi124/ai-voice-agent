package stt

// DetectLanguage inspects string runes to identify Indic scripts or default to English.
func DetectLanguage(text string) string {
	for _, r := range text {
		if r >= '\u0A80' && r <= '\u0AFF' {
			return "gu-IN"
		}
		if r >= '\u0900' && r <= '\u097F' {
			return "hi-IN"
		}
	}
	return "en-IN"
}
