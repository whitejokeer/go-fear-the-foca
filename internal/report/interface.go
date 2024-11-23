package report

import "github.com/whitejokeer/go-fear-the-foca/internal/analyzer"

// Reporter define la interfaz para generar informes.
type Reporter interface {
    // Generate crea un informe basado en los resultados del análisis.
    // Parámetros:
    //   - results: Lista de objetos AnalysisResult.
    //   - outputDir: Directorio donde se guardará el informe.
    // Retorna:
    //   - Un error si ocurre algún problema durante la generación del informe.
    Generate(results []analyzer.AnalysisResult, outputDir string) error
}
