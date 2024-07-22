# Corrección de errores

## Algoritmo de Hamming (n, m)

### Emisor (Python)

El emisor para el algoritmo de Hamming (n, m) se encarga de calcular los bits de paridad necesarios para una cadena binaria de cualquier tamaño, y
retornar la cadena con los bits de paridad calculados.

Para ejecutar el emisor, se debe ejecutar el siguiente comando:

```bash
python3 emisor.py
```

### Receptor (Go)

El receptor para el algoritmo de Hamming (n, m) se encarga de verificar si la cadena recibida contiene errores, y en caso de que los contenga,
corregirlos.

Para ejecutar el receptor, se debe ejecutar el siguiente comando, dentro de la carpeta `Corrección de Errores\Receptor`:

```bash
go run .
```

Además, el receptor contiene con un archivo de pruebas unitarias, el cual se puede ejecutar con el siguiente comando, dentro de la carpeta `Corrección de Errores\Receptor`:

```bash
go test
```
