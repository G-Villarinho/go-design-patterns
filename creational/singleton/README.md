# Padrão Singleton - Database Connection Manager

Este é um exemplo prático do padrão Singleton implementando um **Database Connection Manager** (Gerenciador de Conexão com Banco de Dados), que é um dos casos de uso mais comuns e reais do Singleton.

## 🎯 Problema Resolvido

Em aplicações que precisam se conectar a bancos de dados, é comum que múltiplos componentes (serviços, controllers, etc.) precisem acessar o banco simultaneamente. Sem o padrão Singleton, cada componente criaria sua própria conexão, resultando em:

- **Múltiplas conexões desnecessárias** ao banco
- **Consumo excessivo de recursos** (memória, conexões TCP)
- **Dificuldade de gerenciamento** do estado da conexão
- **Possíveis problemas de performance** e escalabilidade

## 🏗️ Solução com Singleton

O `DatabaseManager` implementa o padrão Singleton para garantir:

- ✅ **Uma única instância** em toda a aplicação
- ✅ **Thread safety** usando `sync.Once`
- ✅ **Controle centralizado** do estado da conexão
- ✅ **Economia de recursos** (apenas uma conexão ativa)
- ✅ **Fácil gerenciamento** do ciclo de vida da conexão

## 📁 Estrutura dos Arquivos

```
singleton/
├── database_manager.go  # Implementação do Singleton
├── example.go          # Demonstrações de uso
└── README.md           # Este arquivo
```

## 🚀 Como Executar

```bash
# Navegue para o diretório singleton
cd creational/singleton

# Execute o exemplo
go run database_manager.go example.go
```

## 🔍 Características Demonstradas

### 1. **Garantia de Instância Única**

- Múltiplas chamadas para `GetInstance()` retornam a mesma instância
- Endereços de memória idênticos confirmam o comportamento

### 2. **Thread Safety**

- Uso de `sync.Once` garante inicialização thread-safe
- Múltiplas goroutines podem acessar simultaneamente
- Mutex protege operações críticas (connect, disconnect, queries)

### 3. **Estado Compartilhado**

- Todos os serviços compartilham o mesmo estado de conexão
- Contador de conexões é mantido globalmente
- Última conexão é registrada para todos

### 4. **Cenário Real**

- Simula múltiplos serviços (UserService, OrderService, etc.)
- Cada serviço executa queries específicas
- Demonstra concorrência real em ambiente de produção

## 💡 Casos de Uso Reais

O padrão Singleton é ideal para:

- **Database Connection Pools**
- **Configuration Managers**
- **Logger Systems**
- **Cache Managers**
- **Service Locators**
- **Resource Managers**

## ⚠️ Considerações Importantes

### Vantagens:

- ✅ Controle de recursos centralizado
- ✅ Economia de memória
- ✅ Facilita o gerenciamento de estado global
- ✅ Thread safety garantida

### Desvantagens:

- ❌ Pode dificultar testes unitários
- ❌ Pode criar acoplamento forte
- ❌ Pode mascarar problemas de design
- ❌ Difícil de estender ou modificar

## 🎯 Quando Usar

Use o Singleton quando:

- Você precisa garantir que uma classe tenha apenas uma instância
- Essa instância deve ser acessível globalmente
- O controle de acesso à instância é importante
- Você quer economizar recursos do sistema

## 🔧 Exemplo de Saída

```
🚀 Iniciando demonstração do padrão Singleton - Database Manager
============================================================

📋 Demonstração 1: Verificando se múltiplas instâncias são a mesma
✅ DatabaseManager Singleton criado pela primeira vez
db1: 0xc000010230
db2: 0xc000010230
db3: 0xc000010230
db1 == db2: true
db2 == db3: true

📋 Demonstração 2: Uso concorrente com múltiplas goroutines
🔧 UserService iniciando operações...
🔌 Conectando ao banco de dados...
✅ Conectado com sucesso! (Conexão #1)
🔍 Executando query: SELECT * FROM UserService_users
✅ UserService: Resultado da query: 'SELECT * FROM UserService_users' - 32 registros encontrados
...
```

Este exemplo demonstra como o padrão Singleton resolve problemas reais de gerenciamento de recursos em aplicações de produção!
