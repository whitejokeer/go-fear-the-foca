package report

import (
	"github.com/whitejokeer/go-fear-the-foca/internal/analyzer"
)

// HTMLReporter implementa la interfaz Reporter para generar informes en HTML.
type HTMLReporter struct{}

// NewHTMLReporter crea una nueva instancia de HTMLReporter.
// Retorna:
//   - Una instancia de HTMLReporter.
func NewHTMLReporter() *HTMLReporter {
	return &HTMLReporter{}
}

// Generate crea un informe en HTML basado en los resultados del análisis.
// Parámetros:
//   - results: Lista de objetos AnalysisResult.
//   - outputDir: Directorio donde se guardará el informe.
//
// Retorna:
//   - Un error si ocurre algún problema durante la generación del informe.
func (h *HTMLReporter) Generate(results []analyzer.AnalysisResult, outputDir string) error {
	return nil
}
