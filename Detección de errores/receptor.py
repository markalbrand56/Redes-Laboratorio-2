import socket

HOST = "127.0.0.1"
PORT = 8080

def make_crc_table() -> list:
    """
    Crea una tabla de 256 elementos con los valores de CRC-32.
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

def crc32_check(data, checksum):
    """
    Verifica si el checksum recibido es correcto.
    """
    table = make_crc_table()
    crc = 0xFFFFFFFF
    for byte in data:
        crc = table[(crc ^ ord(byte)) & 0xFF] ^ (crc >> 8) 
    crc ^= 0xFFFFFFFF 
    return format(crc, '032b') == checksum

# message = input("Ingrese un mensaje en binario: ")
# data_length = len(message) - 32
# original_message = message[:data_length]
# received_checksum = message[data_length:]

# if crc32_check(original_message, received_checksum):
#     print("No se detectaron errores.")
#     print(f"Mensaje original: {original_message}")
# else:
#     print("Se detectaron errores y el mensaje se descarta.")

if __name__ == "__main__":
    # Se crea un socket para la comunicación. Solo debe recibir un mensaje.

    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        print(HOST, PORT)
        s.bind((HOST, PORT))
        s.listen()
        conn, addr = s.accept()
        with conn:
            print(f"Conexión establecida con {addr}")
            
            data = conn.recv(1024)
            message = data.decode("utf-8")
            data_length = len(message) - 32
            
            original_message = message[:data_length]
            received_checksum = message[data_length:]

            print(f"Mensaje recibido: {message} ({len(message)} bytes)")
            print(f"Mensaje original: {original_message} ({len(original_message)} bytes)")
            print(f"Checksum recibido: {received_checksum} ({len(received_checksum)} bytes)")

            if crc32_check(original_message, received_checksum):
                print("No se detectaron errores.")
                print(f"Mensaje original: {original_message}")
            else:
                print("Se detectaron errores y el mensaje se descarta.")