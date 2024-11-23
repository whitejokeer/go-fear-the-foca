package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/whitejokeer/go-fear-the-foca/internal/app"
	"github.com/whitejokeer/go-fear-the-foca/internal/config"

	"github.com/joho/godotenv"
)

// main es el punto de entrada de la aplicación.
// Carga las variables de entorno, procesa los argumentos de línea de comandos
// y ejecuta la aplicación principal.
func main() {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró un archivo .env, se usarán las variables de entorno del sistema")
	}

	// Cargar configuración
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error al cargar la configuración: %v", err)
	}

	// Procesar argumentos de línea de comandos
	domain := flag.String("domain", "", "Dominio de la empresa a analizar (ejemplo.com)")
	output := flag.String("output", "reports", "Directorio donde se guardarán los informes")
	timeout := flag.Int("timeout", 60, "Tiempo máximo de ejecución en segundos")
	flag.Parse()

	if *domain == "" {
		fmt.Println("Uso: go-fear-the-foca -domain ejemplo.com")
		os.Exit(1)
	}

	cfg.Domain = *domain
	cfg.OutputDir = *output
	cfg.Timeout = *timeout

	// Ejecutar la aplicación
	application := app.NewApp(cfg)
	if err := application.Run(); err != nil {
		log.Fatalf("Error al ejecutar la aplicación: %v", err)
	}
}
