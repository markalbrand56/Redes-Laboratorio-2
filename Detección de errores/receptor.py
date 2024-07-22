def make_crc_table():
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
    table = make_crc_table()
    crc = 0xFFFFFFFF
    for byte in data:
        crc = table[(crc ^ ord(byte)) & 0xFF] ^ (crc >> 8)
    crc ^= 0xFFFFFFFF
    return format(crc, '032b') == checksum

message = input("Ingrese un mensaje en binario: ")
data_length = len(message) - 32
original_message = message[:data_length]
received_checksum = message[data_length:]

if crc32_check(original_message, received_checksum):
    print("No se detectaron errores.")
    print(f"Mensaje original: {original_message}")
else:
    print("Se detectaron errores y el mensaje se descarta.")
