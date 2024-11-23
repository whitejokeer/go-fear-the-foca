package search

// GoogleSearcher implementa la interfaz Searcher utilizando la API de Google Custom Search.
type GoogleSearcher struct {
	APIKey   string
	EngineID string
}

// NewGoogleSearcher crea una nueva instancia de GoogleSearcher.
// Parámetros:
//   - apiKey: Clave de API de Google Custom Search.
//   - engineID: ID del motor de búsqueda personalizado.
//
// Retorna una instancia de GoogleSearcher.
func NewGoogleSearcher(apiKey, engineID string) *GoogleSearcher {
	return &GoogleSearcher{
		APIKey:   apiKey,
		EngineID: engineID,
	}
}

// Search busca archivos PDF relacionados con el dominio utilizando Google Custom Search.
// Parámetros:
//   - domain: Dominio de la empresa a analizar.
//
// Retorna:
//   - Una lista de URLs de archivos PDF encontrados.
//   - Un error si ocurre algún problema durante la búsqueda.
func (g *GoogleSearcher) Search(domain string) ([]string, error) {
	var pdfURLs []string

	return pdfURLs, nil
}
