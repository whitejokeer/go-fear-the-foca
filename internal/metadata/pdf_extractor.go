package metadata

// PDFExtractor implementa la interfaz Extractor para archivos PDF.
type PDFExtractor struct{}

// NewPDFExtractor crea una nueva instancia de PDFExtractor.
// Retorna:
//   - Una instancia de PDFExtractor.
func NewPDFExtractor() *PDFExtractor {
	return &PDFExtractor{}
}

// Extract extrae metadatos de los archivos PDF proporcionados.
// Parámetros:
//   - files: Lista de rutas de archivos PDF locales.
//
// Retorna:
//   - Una lista de objetos Metadata.
//   - Un error si ocurre algún problema durante la extracción.
func (p *PDFExtractor) Extract(files []string) ([]Metadata, error) {
	var metadataList []Metadata

	return metadataList, nil
}
