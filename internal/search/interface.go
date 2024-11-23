package search

// Searcher define la interfaz para buscar archivos en línea.
type Searcher interface {
    // Search busca archivos relacionados con el dominio proporcionado.
    // Parámetros:
    //   - domain: Dominio de la empresa a analizar.
    // Retorna:
    //   - Una lista de URLs de archivos encontrados.
    //   - Un error si ocurre algún problema durante la búsqueda.
    Search(domain string) ([]string, error)
}
