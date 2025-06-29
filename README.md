# Design Patterns em Go

Este repositório contém exemplos práticos de padrões de design implementados em Go. Cada padrão demonstra soluções para problemas comuns de desenvolvimento de software, com exemplos reais e aplicáveis.

## 🚀 Como Usar

### Pré-requisitos

- Go 1.24.2 ou superior
- Make (opcional, mas recomendado)

### Comandos Disponíveis

O projeto inclui um Makefile para facilitar a execução dos exemplos:

```bash
# Ver todos os comandos disponíveis
make help

# Executar um exemplo específico
make builder
make strategy

# Executar todos os exemplos
make all

# Listar padrões disponíveis
make list

# Limpar arquivos compilados
make clean
```

### Execução Manual

Se preferir não usar o Makefile, você pode executar os exemplos diretamente:

```bash
# Builder Pattern
cd creational/builder/example
go run main.go

# Strategy Pattern
cd behavioural/strategy
go run main.go
```

## 📚 Padrões de Design Implementados

### 1. Builder Pattern (Padrão Construtor)

**Localização:** `creational/builder/`
**📖 [Documentação Detalhada](creational/builder/README.md)**

**Descrição:** O Builder é um padrão criacional que permite construir objetos complexos passo a passo. Este exemplo demonstra a construção de computadores com diferentes configurações.

**Problema Resolvido:**

- Construção de objetos complexos com muitos parâmetros
- Flexibilidade para criar diferentes variações do mesmo objeto
- Eliminação de construtores com muitos parâmetros

**Cenário Real:** Loja de computadores que precisa montar diferentes tipos de máquinas (Gaming, Escritório, Desenvolvimento, Econômico).

**Principais Benefícios:**

- ✅ Flexibilidade total na construção
- ✅ Reutilização de código
- ✅ Controle passo a passo
- ✅ Código mais legível

**Como Executar:**

```bash
make builder
```

### 2. Strategy Pattern (Padrão Estratégia)

**Localização:** `behavioural/strategy/`
**📖 [Documentação Detalhada](behavioural/strategy/README.md)**

**Descrição:** O Strategy é um padrão comportamental que permite definir uma família de algoritmos, encapsulá-los e torná-los intercambiáveis. Este exemplo implementa um serviço de notificações com diferentes canais.

**Problema Resolvido:**

- Eliminação de múltiplos `if/else` statements
- Extensibilidade para novos tipos de notificação
- Manutenibilidade do código

**Cenário Real:** Sistema de notificações que precisa enviar mensagens através de diferentes canais (Discord, Email, Instagram, Twitter, WhatsApp).

**Principais Benefícios:**

- ✅ Eliminação de condicionais complexos
- ✅ Fácil adição de novos canais
- ✅ Código mais limpo e testável
- ✅ Segue o princípio Open/Closed

**Como Executar:**

```bash
make strategy
```

## 🏗️ Estrutura do Projeto

```
design-patterns/
├── creational/            # Padrões criacionais
│   └── builder/          # Padrão Builder
│       ├── builder.go    # Implementação do padrão
│       ├── example/      # Exemplo prático
│       └── README.md     # 📖 [Documentação](creational/builder/README.md)
├── behavioural/           # Padrões comportamentais
│   └── strategy/         # Padrão Strategy
│       ├── main.go       # Exemplo de uso
│       ├── service/      # Implementação do serviço
│       └── README.md     # 📖 [Documentação](behavioural/strategy/README.md)
├── structural/           # Padrões estruturais (futuro)
├── Makefile             # Comandos para execução
├── go.mod               # Dependências do Go
└── README.md            # Este arquivo
```

## 🎯 Objetivos do Projeto

- **Aprendizado:** Demonstrar padrões de design de forma prática
- **Referência:** Servir como guia para implementações em Go
- **Exemplos Reais:** Usar cenários do mundo real
- **Documentação:** Explicar quando e como usar cada padrão

## 🔧 Desenvolvimento

### Adicionando Novos Padrões

Para adicionar um novo padrão de design:

1. Crie uma nova pasta na categoria apropriada
2. Implemente o padrão com exemplos práticos
3. Adicione documentação detalhada
4. Atualize o Makefile com o novo comando
5. Atualize este README

### Estrutura Recomendada para Novos Padrões

```
padrao-nome/
├── main.go          # Exemplo de uso
├── implementation/  # Implementação do padrão
├── example/         # Exemplos práticos
└── README.md        # Documentação
```

## 📖 Recursos Adicionais

- [Go by Example](https://gobyexample.com/)
- [Design Patterns - Gang of Four](https://en.wikipedia.org/wiki/Design_Patterns)
- [Go Design Patterns](https://github.com/tmrts/go-patterns)

## 🤝 Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para:

- Adicionar novos padrões de design
- Melhorar a documentação
- Corrigir bugs
- Sugerir melhorias

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

---

**Nota:** Este projeto está em desenvolvimento ativo. Novos padrões serão adicionados conforme necessário.
