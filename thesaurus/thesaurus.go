package thesaurus

// Thesaurus returns a list of synonyms for term
type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}
