package builder

import "fmt"

type Computer struct {
	processor string
	memory    string
	storage   string
	graphics  string
	monitor   string
}

type ComputerBuilder interface {
	SetProcessor(processor string) ComputerBuilder
	SetMemory(memory string) ComputerBuilder
	SetStorage(storage string) ComputerBuilder
	SetGraphics(graphics string) ComputerBuilder
	SetMonitor(monitor string) ComputerBuilder
	Build() *Computer
}

type GamingComputerBuilder struct {
	computer *Computer
}

func NewGamingComputerBuilder() *GamingComputerBuilder {
	return &GamingComputerBuilder{
		computer: &Computer{},
	}
}

func (b *GamingComputerBuilder) SetProcessor(processor string) ComputerBuilder {
	b.computer.processor = processor
	return b
}

func (b *GamingComputerBuilder) SetMemory(memory string) ComputerBuilder {
	b.computer.memory = memory
	return b
}

func (b *GamingComputerBuilder) SetStorage(storage string) ComputerBuilder {
	b.computer.storage = storage
	return b
}

func (b *GamingComputerBuilder) SetGraphics(graphics string) ComputerBuilder {
	b.computer.graphics = graphics
	return b
}

func (b *GamingComputerBuilder) SetMonitor(monitor string) ComputerBuilder {
	b.computer.monitor = monitor
	return b
}

func (b *GamingComputerBuilder) Build() *Computer {
	return b.computer
}

type OfficeComputerBuilder struct {
	computer *Computer
}

func NewOfficeComputerBuilder() *OfficeComputerBuilder {
	return &OfficeComputerBuilder{
		computer: &Computer{},
	}
}

func (b *OfficeComputerBuilder) SetProcessor(processor string) ComputerBuilder {
	b.computer.processor = processor
	return b
}

func (b *OfficeComputerBuilder) SetMemory(memory string) ComputerBuilder {
	b.computer.memory = memory
	return b
}

func (b *OfficeComputerBuilder) SetStorage(storage string) ComputerBuilder {
	b.computer.storage = storage
	return b
}

func (b *OfficeComputerBuilder) SetGraphics(graphics string) ComputerBuilder {
	b.computer.graphics = "Gráficos integrados"
	return b
}

func (b *OfficeComputerBuilder) SetMonitor(monitor string) ComputerBuilder {
	b.computer.monitor = monitor
	return b
}

func (b *OfficeComputerBuilder) Build() *Computer {
	return b.computer
}

type ComputerDirector struct {
	builder ComputerBuilder
}

func NewComputerDirector(builder ComputerBuilder) *ComputerDirector {
	return &ComputerDirector{
		builder: builder,
	}
}

func (d *ComputerDirector) BuildGamingComputer() *Computer {
	return d.builder.
		SetProcessor("Intel Core i7-12700K").
		SetMemory("32GB DDR4 3600MHz").
		SetStorage("1TB NVMe SSD").
		SetGraphics("RTX 4070 Ti").
		SetMonitor("27\" 1440p 165Hz").
		Build()
}

func (d *ComputerDirector) BuildOfficeComputer() *Computer {
	return d.builder.
		SetProcessor("Intel Core i5-12400").
		SetMemory("16GB DDR4 3200MHz").
		SetStorage("512GB SATA SSD").
		SetGraphics("Gráficos integrados").
		SetMonitor("24\" 1080p 60Hz").
		Build()
}

func (d *ComputerDirector) BuildBudgetComputer() *Computer {
	return d.builder.
		SetProcessor("Intel Core i3-12100").
		SetMemory("8GB DDR4 2666MHz").
		SetStorage("256GB SATA SSD").
		SetGraphics("Gráficos integrados").
		SetMonitor("21.5\" 1080p 60Hz").
		Build()
}

func (c *Computer) String() string {
	return fmt.Sprintf("Computador:\n"+
		"  Processador: %s\n"+
		"  Memória: %s\n"+
		"  Armazenamento: %s\n"+
		"  Gráficos: %s\n"+
		"  Monitor: %s",
		c.processor, c.memory, c.storage, c.graphics, c.monitor)
}
