package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// HTTPDownloader implementa la interfaz Downloader utilizando HTTP.
type HTTPDownloader struct{}

// NewHTTPDownloader crea una nueva instancia de HTTPDownloader.
// Retorna una instancia de HTTPDownloader.
func NewHTTPDownloader() *HTTPDownloader {
	return &HTTPDownloader{}
}

// Download descarga archivos desde las URLs proporcionadas.
// Parámetros:
//   - urls: Lista de URLs de los archivos a descargar.
//
// Retorna:
//   - Una lista de rutas de archivos locales descargados.
//   - Un error si ocurre algún problema durante la descarga.
func (d *HTTPDownloader) Download(urls []string) ([]string, error) {
	var files []string

	// Crear la carpeta "temp" si no existe
	err := os.MkdirAll("temp", os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("error al crear la carpeta temp: %v", err)
	}

	// Descargar cada archivo
	for _, url := range urls {
		fileName := filepath.Join("temp", filepath.Base(url))
		outFile, err := os.Create(fileName)
		if err != nil {
			return nil, fmt.Errorf("error al crear el archivo %s: %v", fileName, err)
		}
		defer outFile.Close()

		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error al descargar %s: %v", url, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error en la respuesta HTTP al descargar %s: %s", url, resp.Status)
		}

		_, err = io.Copy(outFile, resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error al guardar el archivo %s: %v", fileName, err)
		}

		files = append(files, fileName)
	}
	return files, nil
}
