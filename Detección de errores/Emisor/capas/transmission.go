package capas

import (
	"bufio"
	"fmt"
	"net"
)

type Server struct {
	host string
	port int
}

func NewServer(host string, port int) *Server {
	return &Server{host, port}
}

// Capa de transmisión: Enviar información
// Servidor TCP para conectarse a un socket y enviar un mensaje
func (s *Server) Send(message string) {
	fmt.Printf("Enviando mensaje: %s\n", message)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		fmt.Println("Error al conectar con el servidor:", err)
		return
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	_, err = writer.WriteString(message)
	if err != nil {
		fmt.Println("Error al enviar mensaje:", err)
		return
	}
	writer.Flush()
	fmt.Println("Mensaje enviado con éxito")
}
