
import socket
import sys
import random

def capa_aplicacion():
    """
    Función que simula la capa de aplicación
    """
    print("\n\n--- Capa de Aplicación ---")
    mensaje = input("Introduzca el mensaje: ")
    return mensaje

def capa_presentacion(mensaje):
    """
    Función que simula la capa de presentación
    """
    print("\n--- Capa de Presentación ---")
    #convertir a ASCII
    mensaje = ''.join(format(ord(i), '08b') for i in mensaje)
    print("Mensaje en binario: ", mensaje)
    return mensaje

def enlace_calcular_paridad(message, m, n):
    list_bits = list(message) # Convertimos el mensaje en una lista de bits
    list_bits = list_bits[::-1 ] 

    for i in range(n):  # Iteramos sobre los bits de paridad
        pos = 2**i
        list_bits.insert(pos - 1, 'p')

    print("\nMensaje con bits de paridad: ", ''.join(list_bits[::-1]))

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

def capa_enlace(mensaje):
    """
    Función que simula la capa de enlace
    """
    print("\n--- Capa de Enlace ---")
    # Calculamos el número de bits de paridad necesarios
    m = len(mensaje)
    r = 0
    while 2**r < m + r + 1:
        r += 1
    # Calculamos los bits de paridad
    hamming = enlace_calcular_paridad(mensaje, m, r)
    return hamming

def ruido_aplicar(mensaje, probabilidad_error):
    """
    Función que simula la capa de ruido
    """
    print("\n--- Ruido ---")
    print ("\nMensaje sin errores: ", mensaje)
    mensaje_con_errores = ''
    for bit in mensaje:
        random_num = random.random()
        if random_num < probabilidad_error:
            print(f"- Bit {bit} con error con random {random_num}")
            mensaje_con_errores += '1' if bit == '0' else '0'
        else:
            mensaje_con_errores += bit
    print("Mensaje con errores: ", mensaje_con_errores)
    return mensaje_con_errores


def aplicacion_mostrar_mensaje(mensaje):
    print(f"Mensaje final recibido: {mensaje}")

# Capa de Transmisión
def iniciar_cliente(host, port, mensaje):
    print("\n--- Capa de Transmisión ---")
    mensaje = mensaje + "\n"
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.connect((host, port))
        s.sendall(mensaje.encode('utf-8'))
        print("Mensaje enviado: ", mensaje)

def iniciar_servidor(host, port):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.bind((host, port))
        s.listen()
        print("Servidor escuchando en {}:{}".format(host, port))
        conn, addr = s.accept()
        with conn:
            print('Conectado por', addr)
            while True:
                data = conn.recv(1024)
                if not data:
                    break
                mensaje_recibido = data.decode('utf-8')
                aplicacion_mostrar_mensaje(mensaje_recibido)
                conn.sendall(data)

if __name__ == "__main__":
    while True:
        enlace = capa_enlace(capa_presentacion(capa_aplicacion()))
        mensaje_con_ruido = ruido_aplicar(enlace, 0.001)
        print("Mensaje con ruido: ", mensaje_con_ruido)
        iniciar_cliente('127.0.0.1', 65432, mensaje_con_ruido)