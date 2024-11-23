package analyzer

import (
	openai "github.com/sashabaranov/go-openai"
	"github.com/whitejokeer/go-fear-the-foca/internal/metadata"
)

// OpenAIAnalyzer implementa la interfaz Analyzer utilizando la API de OpenAI.
type OpenAIAnalyzer struct {
	Client *openai.Client
}

// NewOpenAIAnalyzer crea una nueva instancia de OpenAIAnalyzer.
// Parámetros:
//   - apiKey: Clave de API de OpenAI.
//
// Retorna:
//   - Una instancia de OpenAIAnalyzer.
func NewOpenAIAnalyzer(apiKey string) *OpenAIAnalyzer {
	client := openai.NewClient(apiKey)
	return &OpenAIAnalyzer{
		Client: client,
	}
}

// Analyze procesa los metadatos utilizando la API de OpenAI.
// Parámetros:
//   - metadataList: Lista de objetos Metadata a analizar.
//
// Retorna:
//   - Una lista de objetos AnalysisResult.
//   - Un error si ocurre algún problema durante el análisis.
func (o *OpenAIAnalyzer) Analyze(metadataList []metadata.Metadata) ([]AnalysisResult, error) {
	var results []AnalysisResult

	return results, nil
}
