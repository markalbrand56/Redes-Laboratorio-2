# Laboratorio 2: Esquemas de detección y corrección

- Mark Albrand - 21004
- Jimena Hernández - 21199

Este proyecto incluye implementaciones de dos algoritmos fundamentales en la teoría de la comunicación y la informática: el algoritmo de Hamming para la corrección de errores y el algoritmo CRC-32 para la detección de errores. Ambos algoritmos son esenciales para garantizar la integridad de los datos en la transmisión y el almacenamiento de información.

## Estructura del Proyecto

El proyecto está dividido en dos subcarpetas principales, cada una dedicada a uno de los algoritmos:

- `Corrección de Errores`: Contiene la implementación del algoritmo de Hamming, con scripts separados para el emisor y el receptor.
- `Detección de Errores`: Contiene la implementación del algoritmo CRC-32, también con scripts separados para el emisor y el receptor, además de un programa adicional para la generación de colisiones de hash.

## Uso

Para utilizar las implementaciones de los algoritmos, se debe navegar a la subcarpeta correspondiente y seguir las instrucciones específicas en el de cada subcarpeta.

### Corrección de Errores con Hamming

- **Emisor (Python):** Calcula los bits de paridad para una cadena binaria.
- **Receptor (Go):** Verifica y corrige errores en la cadena recibida.

### Detección de Errores con CRC-32

- **Emisor (Go):** Calcula el CRC-32 de una cadena binaria.
- **Receptor (Python):** Verifica la cadena recibida para detectar errores.
- **Hash Collider (Go):** Encuentra colisiones de hash para el algoritmo CRC-32.

## Requisitos

Para ejecutar los scripts de este proyecto, necesitará tener instalados Python 3 y Go en su sistema. Las instrucciones específicas de ejecución se encuentran en los `README` de las subcarpetas correspondientes.

## Contribuciones

Este proyecto es abierto a contribuciones. Si desea mejorar las implementaciones existentes o agregar nuevas funcionalidades, siéntase libre de hacer un fork del repositorio y enviar sus pull requests.

## Licencia

Este proyecto está licenciado bajo [incluir tipo de licencia aquí], lo que permite su uso, modificación y distribución bajo ciertas condiciones.