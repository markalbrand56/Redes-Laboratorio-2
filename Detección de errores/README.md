# Detección de errores

## Algoritmo CRC-32

### Emisor (Go)

El emisor para el algoritmo CRC-32 se encarga de calcular el CRC-32 de una cadena binaria, y retornar la cadena con el CRC-32 calculado.

Para ejecutar el emisor, se debe ejecutar el siguiente comando, dentro de la carpeta `Detección de Errores\Emisor`:

```bash
go run .
```

### Receptor (Python)

El receptor para el algoritmo CRC-32 se encarga de verificar si la cadena recibida contiene errores, y en caso de que los contenga, corregirlos.

Para ejecutar el receptor, se debe ejecutar el siguiente comando:

```bash
python3 receptor.py
```
