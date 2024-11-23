package analyzer

import "github.com/whitejokeer/go-fear-the-foca/internal/metadata"

// AnalysisResult representa el resultado del análisis de metadatos.
type AnalysisResult struct {
	Summary string
}

// Analyzer define la interfaz para analizar metadatos.
type Analyzer interface {
	// Analyze procesa los metadatos proporcionados y retorna los resultados del análisis.
	// Parámetros:
	//   - metadataList: Lista de objetos Metadata a analizar.
	// Retorna:
	//   - Una lista de objetos AnalysisResult.
	//   - Un error si ocurre algún problema durante el análisis.
	Analyze(metadataList []metadata.Metadata) (*AnalysisResult, error)
}
