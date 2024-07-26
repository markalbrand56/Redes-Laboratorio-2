import socket

HOST = "127.0.0.1"
PORT = 8080

def make_crc_table() -> list[int]:
    """
    Crea una tabla de 256 elementos con los valores de CRC-32.
    :return: Lista con los valores de CRC-32.
    """
    poly = 0xEDB88320
    table = []
    for i in range(256):
        crc = i
        for j in range(8):
            if crc & 1:
                crc = (crc >> 1) ^ poly
            else:
                crc >>= 1
        table.append(crc)
    return table

def crc32_check(data: str, checksum: str) -> bool:
    """
    Capa de enlace: Verificar integridad
    Verifica si el checksum recibido es correcto.
    :param data: Mensaje original.
    :param checksum: Checksum recibido.
    :return: True si el checksum es correcto, False en caso contrario.
    """
    table = make_crc_table()
    crc = 0xFFFFFFFF
    for byte in data:
        crc = table[(crc ^ ord(byte)) & 0xFF] ^ (crc >> 8) 
    crc ^= 0xFFFFFFFF 
    return format(crc, '032b') == checksum


def decode_message(message: str) -> str:
    """
    Capa de presentación: Decodificación
    Convierte una cadena de 0s y 1s a una cadena de caracteres.
    :param message: Mensaje en binario.
    :return: Mensaje decodificado.
    """
    return "".join([chr(int(message[i:i+8], 2)) for i in range(0, len(message), 8)])


if __name__ == "__main__":
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.bind((HOST, PORT))
        s.listen()
        conn, addr = s.accept()
        
        with conn:
            print(f"Conexión establecida con {addr}")
            
            data = conn.recv(1024)  # Capa de transmisión: Recibir información
            message = data.decode("utf-8")
            data_length = len(message) - 32
            
            original_message = message[:data_length]
            received_checksum = message[data_length:]

            print(f"Mensaje original: {original_message} ({len(original_message)} bytes)")
            print(f"Checksum recibido: {received_checksum} ({len(received_checksum)} bytes)")

            # Capa de enlace: Verificar integridad
            if crc32_check(original_message, received_checksum):
                print("No se detectaron errores.")

                # Capa de presentación: Decodificación
                decoded_message = decode_message(original_message)

                # Capa de aplicación: Mostrar mensaje
                print(f"Mensaje decodificado: \n{decoded_message}\n")
            else:
                print("Se detectaron errores y el mensaje se descarta.")