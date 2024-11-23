package downloader

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

	return files, nil
}
