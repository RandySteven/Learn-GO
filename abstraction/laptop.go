package abstraction

type IComputer interface {
	SetGraphicCard(graphicCard string)
	SetRam(ram float32)
	SetMemory(memory float32)
	GetRam() float32
	GetMemory() float32
}

type ISmartPhone interface {
	SetPixel(pixel float32)
	GetPixel() float32
	SetCamera(camera interface{})
	GetCamera() interface{}
}

type ITech interface {
	makeSmartPhone() ISmartPhone
	makeComputer() IComputer
}

func (s *SmartPhone) SetPixel(pixel float32) {
	s.Pixel = pixel
}

func (s *SmartPhone) GetPixel() float32 {
	return s.Pixel
}

func (c *Computer) SetGraphicCard(graphicCard string) {
	c.GraphicCard = graphicCard
}

func (c *Computer) GetGraphicCard() string {
	return c.GraphicCard
}

func (c *Computer) SetMemory(memory float32) {
	c.Memory = memory
}

func (c *Computer) GetMemory() float32 {
	return c.Memory
}

func (c *Computer) SetRam(ram float32) {
	c.Ram = ram
}

func (c *Computer) GetRam() float32 {
	return c.Ram
}

type Asus struct{}

func (a *Asus) makeSmartPhone() ISmartPhone {
	return &AsusSmartPhone{
		&SmartPhone{
			32,
			"Mantap",
		},
	}
}

func (a *Asus) makeComputer() IComputer {
	return &AsusComputer{
		&Computer{
			"NVidia",
			64,
			256,
		},
	}
}

type Computer struct {
	GraphicCard string
	Ram         float32
	Memory      float32
}

type SmartPhone struct {
	Pixel  float32
	Camera interface{}
}

type AsusComputer struct {
	Computer *Computer
}

func (a *AsusComputer) SetGraphicCard(graphicCard string) {
	a.Computer.SetGraphicCard(graphicCard)
}

func (a *AsusComputer) SetRam(ram float32) {
	a.Computer.SetRam(ram)
}

func (a *AsusComputer) SetMemory(memory float32) {
	a.Computer.SetMemory(memory)
}

func (a *AsusComputer) GetRam() float32 {
	return a.Computer.GetRam()
}

func (a *AsusComputer) GetMemory() float32 {
	return a.Computer.GetMemory()
}

type AsusSmartPhone struct {
	SmartPhone *SmartPhone
}

func (a *AsusSmartPhone) SetPixel(pixel float32) {
	a.SmartPhone.SetPixel(pixel)
}

func (a *AsusSmartPhone) GetPixel() float32 {
	return a.SmartPhone.GetPixel()
}

func (a *AsusSmartPhone) SetCamera(camera interface{}) {
	a.SetCamera(camera)
}

func (a *AsusSmartPhone) GetCamera() interface{} {
	return a.GetCamera()
}

type AppleComputer struct {
	Computer
}

type AppleSmartPhone struct {
	SmartPhone
}
