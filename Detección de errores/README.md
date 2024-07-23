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

### Adicional

#### Hash Collider

Para responder a la segunda pregunta del documento, se adjuntó un programa en Go que, a fuerza bruta, encuentra una colisión de hash para el algoritmo CRC-32.
Esto significa que dado un mensaje original, se puede encontrar un mensaje diferente que tiene el mismo checksum CRC-32.

El archivo original es del usuario [fyxme](https://github.com/fyxme/crc-32-hash-collider/blob/master/crc-32-collide.go) y fue modificado para probar con los ejemplos de este laboratorio.

Para ejecutar el programa, se debe ejecutar el siguiente comando, dentro de la carpeta `Detección de Errores\`:

```bash
go run hash_collider.go
```
