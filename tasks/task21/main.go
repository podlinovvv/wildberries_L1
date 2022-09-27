package main

import "fmt"

func main() {
	//клиент, от которого скрываются детали реализации
	client := &Client{}


	//тип, у которого можно вызвать метод без адаптера
	hdmi := &HDMICable{}
	client.connect(hdmi)

	//тип, которому нужен адаптер
	oldCable := &DVIcable{}
	//создаём адаптер
	DVItoHDMIconverter := &DviAdapter{
		dviCable: oldCable,
	}
	//вызываем метод через адаптер
	client.connect(DVItoHDMIconverter)
}

type Client struct {
}

func (c *Client) connect(cable Cable) {
	fmt.Println("Client inserts cable into")
	cable.InsertIntoHDMI()
}


type Cable interface {
	InsertIntoHDMI()
}



type HDMICable struct {
}

func (h *HDMICable) InsertIntoHDMI() {
	fmt.Println("HDMI connector is plugged in")
}




type DVIcable struct{}

func (d *DVIcable) insertIntoUSBPort() {
	fmt.Println("DVI connector is plugged in")
}

type DviAdapter struct {
	dviCable *DVIcable
}

func (d *DviAdapter) InsertIntoHDMI() {
	fmt.Println("Adapter converts DVI into HDMI")
	d.dviCable.insertIntoUSBPort()
}