def hamming(message, m, n):
    print("Mensaje original: ", message)
    print("Número de bits de datos: ", m)
    print("Número de bits de paridad: ", n)

    list_bits = list(message)
    list_bits = list_bits[::-1]

    for i in range(n):  # Iteramos sobre los bits de paridad
        pos = 2**i
        
        list_bits.insert(pos - 1, 'p')

    print("Mensaje con bits de paridad: ", ''.join(list_bits[::-1]))

    for i in range(n):  # Iteramos sobre los bits de paridad
        pos = 2**i
        count = 0

        for j in range(1, len(list_bits) + 1):
            if j & pos == pos and list_bits[j - 1] == '1':
                count += 1

        if count % 2 == 0:
            list_bits[pos - 1] = '0'
        else:
            list_bits[pos - 1] = '1'

    print("Bits de paridad: ", ''.join(list_bits[::-1]))


if __name__ == "__main__":
    message = input("Introduce el mensaje: ")

    # El mensaje solo debe incluir 1 y 0
    if not all([c in "01" for c in message]):
        print("El mensaje debe ser una cadena binaria")
        exit()

    # Calculamos el número de bits de paridad necesarios
    m = len(message)
    r = 0
    while 2**r < m + r + 1:
        r += 1

    # Calculamos los bits de paridad
    hamming = hamming(message, m, r)