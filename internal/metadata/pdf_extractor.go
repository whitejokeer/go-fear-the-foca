package metadata

import (
	"encoding/json"
	"os/exec"
)

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

	for _, file := range files {
		// Construir el comando exiftool
		cmd := exec.Command("exiftool", "-All", "-H", "-j", file)

		// Ejecutar el comando y capturar la salida
		output, err := cmd.Output()
		if err != nil {
			return nil, err
		}

		// Parsear la salida JSON
		var metadataArray []map[string]interface{}
		err = json.Unmarshal(output, &metadataArray)
		if err != nil {
			return nil, err
		}

		// Procesar los metadatos
		if len(metadataArray) > 0 {
			metadataMap := metadataArray[0]

			// Extraer el nombre del archivo
			fileName, _ := metadataMap["SourceFile"].(string)

			// Remover SourceFile del mapa de propiedades
			delete(metadataMap, "SourceFile")

			// Crear una instancia de Metadata
			metadata := Metadata{
				FileName:   fileName,
				Properties: metadataMap,
			}

			// Agregar a la lista de metadatos
			metadataList = append(metadataList, metadata)
		}
	}

	return metadataList, nil
}
