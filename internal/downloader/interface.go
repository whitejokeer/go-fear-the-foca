package downloader

// Downloader define la interfaz para descargar archivos desde URLs.
type Downloader interface {
	// Download descarga archivos desde las URLs proporcionadas.
	// Parámetros:
	//   - urls: Lista de URLs de los archivos a descargar.
	// Retorna:
	//   - Una lista de rutas de archivos locales descargados.
	//   - Un error si ocurre algún problema durante la descarga.
	Download(urls []string) ([]string, error)
}
