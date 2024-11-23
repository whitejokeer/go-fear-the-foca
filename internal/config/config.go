package config

import (
    "os"
    "strconv"
)

// Config contiene la configuración de la aplicación.
type Config struct {
    OpenAIAPIKey       string
    SearchEngineAPIKey string
    SearchEngineID     string
    Domain             string
    OutputDir          string
    Timeout            int
}

// LoadConfig carga la configuración desde las variables de entorno.
// Retorna un puntero a Config y un error si ocurre algún problema.
func LoadConfig() (*Config, error) {
    timeout, err := strconv.Atoi(getEnv("TIMEOUT", "60"))
    if err != nil {
        timeout = 60
    }

    cfg := &Config{
        OpenAIAPIKey:       os.Getenv("OPENAI_API_KEY"),
        SearchEngineAPIKey: os.Getenv("SEARCH_ENGINE_API_KEY"),
        SearchEngineID:     os.Getenv("SEARCH_ENGINE_ID"),
        Timeout:            timeout,
    }

    return cfg, nil
}

// getEnv obtiene una variable de entorno o devuelve un valor por defecto si no está definida.
// Parámetros:
//   - key: Nombre de la variable de entorno.
//   - defaultValue: Valor por defecto si la variable no está definida.
// Retorna el valor de la variable de entorno o el valor por defecto.
func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}
