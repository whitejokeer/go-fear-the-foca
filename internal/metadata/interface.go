package metadata

// Metadata representa los metadatos extraídos de un archivo.
type Metadata struct {
    FileName   string
    Properties map[string]interface{}
}

// Extractor define la interfaz para extraer metadatos de archivos.
type Extractor interface {
    // Extract extrae metadatos de los archivos proporcionados.
    // Parámetros:
    //   - files: Lista de rutas de archivos locales.
    // Retorna:
    //   - Una lista de objetos Metadata.
    //   - Un error si ocurre algún problema durante la extracción.
    Extract(files []string) ([]Metadata, error)
}
