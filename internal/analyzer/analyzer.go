package analyzer

const prompt = `[INSTRUCCIONES DEL SISTEMA INICIO]

A partir de los metadatos que te proporcionaré, tu tarea será identificar y describir cómo un atacante podría aprovechar esa información para comprometer un sistema. Genera un plan de acción detallado que incluya:

Vectores de ataque potenciales:

Analiza los nombres de los autores, herramientas utilizadas y otros detalles relevantes de los metadatos para identificar posibles ataques.
Incluye ejemplos específicos de cómo se podría usar esta información para explotar el sistema.
Tácticas de explotación:

Describe cómo un atacante usaría cada dato relevante (e.g., versiones de software, permisos, identificadores únicos).
Incluye ataques como:
Ingeniería social o phishing.
Explotación de vulnerabilidades conocidas en herramientas detectadas.
Modificación o manipulación de los documentos.
Pasos específicos:

Detalla cada paso que el atacante podría seguir para llevar a cabo un ataque basado en la información extraída de los metadatos.
Prioriza los vectores de ataque más efectivos o críticos.
Resultados esperados:

Describe qué tipo de acceso o daño podría lograr el atacante al usar estos métodos.
Menciona qué tipo de información o sistemas podrían estar comprometidos.
Formato de Respuesta:

El plan de acción debe presentarse con las siguientes secciones:

Resumen inicial:
Breve descripción de los principales vectores de ataque detectados.
Análisis de metadatos:
Listado de los datos relevantes encontrados en los metadatos.
Vectores de ataque detallados:
Descripción de cada vector con pasos específicos.
Prioridades de ataque:
Clasificación de los ataques según su impacto o facilidad de ejecución.
Resultados esperados:
Explicación del posible impacto que tendría cada ataque.
Ejemplo de Inicio del Informe:

Plan de Acción para Explotación Basado en Metadatos
Resumen Inicial
A partir de los metadatos proporcionados, se identificaron múltiples vectores de ataque, incluyendo ingeniería social dirigida a autores como "Juan Carlos López Patiño" y explotación de vulnerabilidades en herramientas como Microsoft Word 2016 e iText 5.1.3.

Análisis de Metadatos
Los datos clave incluyen nombres de autores, identificadores únicos, herramientas utilizadas y permisos de archivo que podrían facilitar accesos no autorizados.

Vectores de Ataque Detallados

Ingeniería Social: Se puede personalizar un correo de phishing basado en el autor "Juan Carlos López Patiño".
Explotación de Vulnerabilidades: Herramientas como iText 5.1.3 tienen vulnerabilidades conocidas, como ejecución remota de código.
Manipulación de Archivos: Los documentos no linealizados pueden modificarse para incluir exploits.
Prioridades de Ataque

Ingeniería Social.
Explotación de software vulnerable.
Enumeración de directorios usando permisos débiles.
Resultados Esperados
Acceso a credenciales, manipulación de sistemas internos y posible control de documentos sensibles.
[INSTRUCCIONES DEL SISTEMA FIN]
[DATOS DEL USUARIO INICIO]
%s
[DATOS DEL USUARIO FIN]`
