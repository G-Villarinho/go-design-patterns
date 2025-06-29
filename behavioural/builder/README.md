# Padrão Builder - Construção de Computadores

O **Builder** é um padrão de projeto criacional que permite construir objetos complexos passo a passo. Este exemplo demonstra o padrão usando a construção de computadores, um cenário muito comum no dia a dia.

## Problema Real

Imagine uma loja de computadores que precisa montar diferentes tipos de máquinas:

- **Gaming**: Processadores potentes, muita RAM, placas de vídeo dedicadas
- **Escritório**: Configurações básicas, gráficos integrados, preço acessível
- **Desenvolvimento**: Muita RAM, armazenamento rápido, monitores grandes
- **Econômico**: Componentes básicos, preço baixo

Criar construtores separados para cada tipo seria ineficiente e difícil de manter.

## Solução com Builder

O padrão Builder permite construir computadores passo a passo, com flexibilidade total para personalizar cada componente.

## Estrutura do Exemplo

- **Computer**: O objeto complexo (computador) com processador, memória, armazenamento, gráficos e monitor
- **ComputerBuilder**: Interface que define métodos para configurar cada componente
- **GamingComputerBuilder**: Builder especializado para computadores gaming
- **OfficeComputerBuilder**: Builder especializado para computadores de escritório
- **ComputerDirector**: Coordena a construção de configurações padrão

## Vantagens Demonstradas

1. **Flexibilidade**: Mesmo builder pode criar configurações completamente diferentes
2. **Reutilização**: Director oferece configurações padrão prontas para uso
3. **Controle**: Construção passo a passo permite personalização total
4. **Legibilidade**: Código fica mais claro e fácil de entender

## Exemplos de Uso

### Construção usando Director (configurações padrão)

```go
// Computador Gaming padrão
gamingBuilder := NewGamingComputerBuilder()
director := NewComputerDirector(gamingBuilder)
gamingComputer := director.BuildGamingComputer()

// Computador de Escritório padrão
officeBuilder := NewOfficeComputerBuilder()
director = NewComputerDirector(officeBuilder)
officeComputer := director.BuildOfficeComputer()
```

### Construção manual personalizada

```go
customComputer := NewGamingComputerBuilder().
    SetProcessor("AMD Ryzen 9 7900X").
    SetMemory("64GB DDR5 6000MHz").
    SetStorage("2TB NVMe SSD").
    SetGraphics("RTX 4090").
    SetMonitor("32\" 4K 144Hz").
    Build()
```

## Cenários Reais

- **Loja de computadores**: Montagem personalizada para clientes
- **Empresa de TI**: Configuração de máquinas para diferentes departamentos
- **Desenvolvedor**: Setup de ambiente de desenvolvimento
- **Gamer**: Montagem de PC para jogos específicos

## Executando o Exemplo

```bash
cd creational/builder/example
go run main.go
```

## Saída Esperada

O exemplo mostrará diferentes configurações de computadores:

- Computador Gaming com componentes de alta performance
- Computador de Escritório com configuração básica
- Computador Econômico com preço acessível
- Configurações personalizadas para diferentes necessidades

## Quando Usar o Builder

- Objetos complexos com muitos parâmetros opcionais
- Diferentes representações do mesmo objeto
- Controle total sobre o processo de criação
- Código de construção reutilizável
