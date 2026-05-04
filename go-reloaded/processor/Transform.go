package processor

func Transform(text string) string {
	text = ApplyHex(text)
	text = ApplyBin(text)
	text = Capitalize(text)
	text = ApplyCases(text)
	text = ApplyArticle(text)
	text = ApplyQuotes(text)
	text = ApplyPunctuation(text)
	
	return text
}
