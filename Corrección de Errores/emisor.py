import time
def hamming(message, m, n):
    """
    Función hamming(m, n) que recibe un mensaje de m bits y añade n bits de paridad
    :param message: Mensaje binario
    :param m: Número de bits del mensaje
    :param n: Número de bits de paridad
    :return: Mensaje con bits de paridad
    """
    
    print("\n - Mensaje original: ", message)
    print("\n -Número de bits en la cadena: ", m)
    print("\n -Número de bits de paridad: ", n)

    list_bits = list(message) # Convertimos el mensaje en una lista de bits
    list_bits = list_bits[::-1 ] 

    for i in range(n):  # Iteramos sobre los bits de paridad
        pos = 2**i
        list_bits.insert(pos - 1, 'p')

    print("\n Mensaje con bits de paridad: ", ''.join(list_bits[::-1]))

    for i in range(n):  # Iteramos sobre los bits de paridad
        pos = 2**i
        count = 0

        for j in range(1, len(list_bits) + 1):
            if j & pos == pos and list_bits[j - 1] == '1':
                # Si el bit j tiene el bit i-ésimo activado (pos) y es 1
                count += 1

        # Asignación del bit de paridad
        if count % 2 == 0:
            list_bits[pos - 1] = '0'
        else:
            list_bits[pos - 1] = '1'

    return ''.join(list_bits[::-1])


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

    start_time = time.time()  
    # Calculamos los bits de paridad
    hamming = hamming(message, m, r)
    end_time = time.time()  

    print("\nMensaje con bits de paridad: ", hamming)
    print("Tiempo de ejecución: {:.6f} segundos".format(end_time - start_time)) 