package main

import (
	"fmt"

	"github.com/g-villarinho/design-patterns/behavioural/builder"
)

func main() {
	fmt.Println("=== Exemplo do Padrão Builder - Construção de Computadores ===")

	fmt.Println("1. Construção usando Director:")

	gamingBuilder := builder.NewGamingComputerBuilder()
	gamingDirector := builder.NewComputerDirector(gamingBuilder)
	gamingComputer := gamingDirector.BuildGamingComputer()

	fmt.Println("Computador Gaming:")
	fmt.Println(gamingComputer)
	fmt.Println()

	officeBuilder := builder.NewOfficeComputerBuilder()
	officeDirector := builder.NewComputerDirector(officeBuilder)
	officeComputer := officeDirector.BuildOfficeComputer()

	fmt.Println("Computador de Escritório:")
	fmt.Println(officeComputer)
	fmt.Println()

	budgetComputer := officeDirector.BuildBudgetComputer()

	fmt.Println("Computador Econômico:")
	fmt.Println(budgetComputer)
	fmt.Println()

	fmt.Println("2. Construção manual personalizada:")

	customGamingBuilder := builder.NewGamingComputerBuilder()
	customGamingComputer := customGamingBuilder.
		SetProcessor("AMD Ryzen 9 7900X").
		SetMemory("64GB DDR5 6000MHz").
		SetStorage("2TB NVMe SSD").
		SetGraphics("RTX 4090").
		SetMonitor("32\" 4K 144Hz").
		Build()

	fmt.Println("Computador Gaming Personalizado:")
	fmt.Println(customGamingComputer)
	fmt.Println()

	fmt.Println("3. Construção para desenvolvimento:")

	devBuilder := builder.NewGamingComputerBuilder()
	devComputer := devBuilder.
		SetProcessor("Intel Core i9-13900K").
		SetMemory("128GB DDR5 5600MHz").
		SetStorage("4TB NVMe SSD").
		SetGraphics("RTX 4080").
		SetMonitor("34\" Ultrawide 1440p 144Hz").
		Build()

	fmt.Println("Computador para Desenvolvimento:")
	fmt.Println(devComputer)
	fmt.Println()

	fmt.Println("4. Demonstração da flexibilidade:")

	flexibleBuilder := builder.NewGamingComputerBuilder()

	basicComputer := flexibleBuilder.
		SetProcessor("Intel Core i5-13600K").
		SetMemory("16GB DDR4 3200MHz").
		SetStorage("500GB NVMe SSD").
		SetGraphics("RTX 3060").
		SetMonitor("24\" 1080p 144Hz").
		Build()

	fmt.Println("Configuração Básica:")
	fmt.Println(basicComputer)
	fmt.Println()

	flexibleBuilder = builder.NewGamingComputerBuilder()

	premiumComputer := flexibleBuilder.
		SetProcessor("AMD Ryzen 9 7950X").
		SetMemory("128GB DDR5 6400MHz").
		SetStorage("8TB NVMe SSD").
		SetGraphics("RTX 4090").
		SetMonitor("49\" Super Ultrawide 5120x1440 240Hz").
		Build()

	fmt.Println("Configuração Premium:")
	fmt.Println(premiumComputer)
}
