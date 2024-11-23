package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

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
	query := fmt.Sprintf("site:%s filetype:pdf", domain)
	searchURL := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?q=%s&key=%s&cx=%s",
		url.QueryEscape(query), g.APIKey, g.EngineID)

	fmt.Println("searchURL: ", searchURL)

	resp, err := http.Get(searchURL)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la respuesta HTTP: %s", resp.Status)
	}

	var result struct {
		Items []struct {
			Link string `json:"link"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error al parsear la respuesta JSON: %v", err)
	}

	var pdfURLs []string
	for _, item := range result.Items {
		pdfURLs = append(pdfURLs, item.Link)
	}

	return pdfURLs, nil
}
