package k6sendudp

import (
    "fmt"
    "net"
    "go.k6.io/k6/js/modules"
)

func init() {
    modules.Register("k6/x/k6sendudp", new(UdpModule))
}

type UdpModule struct {
    conn *net.UDPConn
    addr *net.UDPAddr
}

// Init метод для создания UDP-соединения
func (u *UdpModule) Init(host string, port int) error {
    // Создаем UDP адрес на основе переданных хоста и порта
    addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", host, port))
    if err != nil {
        return fmt.Errorf("Ошибка при разрешении UDP адреса: %v", err)
    }

    // Создаем UDP соединение
    conn, err := net.DialUDP("udp", nil, addr)
    if err != nil {
        return fmt.Errorf("Ошибка при установке соединения: %v", err)
    }

    u.conn = conn
    u.addr = addr

    fmt.Printf("Успешное подключение к %s:%d\n", host, port)
    return nil
}

// SendMessage метод для отправки строки и числа через UDP
//func (u *UdpModule) SendMessage(message string, number int) error {
func (u *UdpModule) SendMessage(message string) error {
    if u.conn == nil {
        return fmt.Errorf("UDP соединение не установлено. Вызовите Init() для настройки соединения.")
    }

    // Формируем сообщение
    //fullMessage := fmt.Sprintf("%s %d", message, number)
	fullMessage := fmt.Sprintf("%s", message)

    // Отправляем данные
    _, err := u.conn.Write([]byte(fullMessage))
    if err != nil {
        return fmt.Errorf("Ошибка при отправке данных: %v", err)
    }

    fmt.Printf("Сообщение отправлено: %s\n", fullMessage)
    return nil
}