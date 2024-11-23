package app

import (
    "log"

    "github.com/whitejokeer/go-fear-the-foca/internal/analyzer"
    "github.com/whitejokeer/go-fear-the-foca/internal/config"
    "github.com/whitejokeer/go-fear-the-foca/internal/downloader"
    "github.com/whitejokeer/go-fear-the-foca/internal/metadata"
    "github.com/whitejokeer/go-fear-the-foca/internal/report"
    "github.com/whitejokeer/go-fear-the-foca/internal/search"
)

// App es la estructura principal de la aplicación.
// Contiene la configuración y las implementaciones de las interfaces.
type App struct {
    Config     *config.Config
    Searcher   search.Searcher
    Downloader downloader.Downloader
    Extractor  metadata.Extractor
    Analyzer   analyzer.Analyzer
    Reporter   report.Reporter
}

// NewApp crea una nueva instancia de App con las implementaciones predeterminadas.
// Parámetros:
//   - cfg: Configuración de la aplicación.
// Retorna una instancia de App.
func NewApp(cfg *config.Config) *App {
    return &App{
        Config:     cfg,
        Searcher:   search.NewGoogleSearcher(cfg.SearchEngineAPIKey, cfg.SearchEngineID),
        Downloader: downloader.NewHTTPDownloader(),
        Extractor:  metadata.NewPDFExtractor(),
        Analyzer:   analyzer.NewOpenAIAnalyzer(cfg.OpenAIAPIKey),
        Reporter:   report.NewHTMLReporter(),
    }
}

// Run ejecuta el flujo principal de la aplicación.
// Retorna un error si ocurre algún problema durante la ejecución.
func (a *App) Run() error {
    // Paso 1: Buscar archivos
    urls, err := a.Searcher.Search(a.Config.Domain)
    if err != nil {
        return err
    }
    log.Printf("Se encontraron %d archivos", len(urls))

    // Paso 2: Descargar archivos
    files, err := a.Downloader.Download(urls)
    if err != nil {
        return err
    }
    log.Printf("Se descargaron %d archivos", len(files))

    // Paso 3: Extraer metadatos
    metadataList, err := a.Extractor.Extract(files)
    if err != nil {
        return err
    }
    log.Printf("Se extrajeron metadatos de %d archivos", len(metadataList))

    // Paso 4: Analizar metadatos
    analysis, err := a.Analyzer.Analyze(metadataList)
    if err != nil {
        return err
    }
    log.Println("Análisis completado")

    // Paso 5: Generar reporte
    err = a.Reporter.Generate(analysis, a.Config.OutputDir)
    if err != nil {
        return err
    }
    log.Println("Reporte generado exitosamente")

    return nil
}
