// internal/analyzer/openai_analyzer.go

package analyzer

import (
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
	"github.com/whitejokeer/go-fear-the-foca/internal/metadata"
)

// OpenAIAnalyzer implementa la interfaz Analyzer utilizando la API de OpenAI.
type OpenAIAnalyzer struct {
	Client *openai.Client
	Prompt string
}

// NewOpenAIAnalyzer crea una nueva instancia de OpenAIAnalyzer.
// Parámetros:
//   - apiKey: Clave de API de OpenAI.
//   - prompt: Prompt personalizado para el análisis.
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
func (o *OpenAIAnalyzer) Analyze(metadataList []metadata.Metadata) (*AnalysisResult, error) {
	ctx := context.Background()

	// Convertir todos los metadatos a una sola cadena
	allMetadataStr := formatAllMetadata(metadataList)

	// Crear el prompt personalizado utilizando el que ya tienes
	prompt := fmt.Sprintf(prompt, allMetadataStr)

	// Crear la solicitud a la API de OpenAI
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Eres un experto en ciberseguridad que analiza metadatos de documentos para identificar posibles riesgos.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}

	// Llamar a la API de OpenAI
	resp, err := o.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error al llamar a la API de OpenAI: %v", err)
	}

	// Extraer la respuesta del asistente
	summary := resp.Choices[0].Message.Content

	// Guardar el resumen en un archivo
	file, err := os.Create("internal/report/summary.md")
	if err != nil {
		return nil, fmt.Errorf("error al crear el archivo: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(summary)
	if err != nil {
		return nil, fmt.Errorf("error al escribir en el archivo: %v", err)
	}

	results := AnalysisResult{
		Summary: summary,
	}

	return &results, nil
}

// formatAllMetadata convierte la lista de metadatos en una cadena legible.
// Parámetros:
//   - metadataList: Lista de objetos Metadata.
//
// Retorna:
//   - Una cadena formateada de todos los metadatos.
func formatAllMetadata(metadataList []metadata.Metadata) string {
	var builder strings.Builder
	for _, md := range metadataList {
		builder.WriteString(fmt.Sprintf("Archivo: %s\n", md.FileName))
		for key, value := range md.Properties {
			builder.WriteString(fmt.Sprintf("%s: %v\n", key, value))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
