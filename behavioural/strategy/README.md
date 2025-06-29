# Padrão Strategy - Notification Service

## Descrição

Este exemplo demonstra a implementação do padrão Strategy para um serviço de notificações. O objetivo é eliminar a necessidade de múltiplos `if/else` statements, tornando o código mais limpo, extensível e seguindo o princípio Open/Closed (aberto para extensão, fechado para modificação).

## Estrutura do Projeto

```
behavioural/strategy/
├── service/
│   ├── interfaces/
│   │   └── notification_strategy.go    # Interface principal
│   ├── notifications/
│   │   ├── discord.go                  # Estratégia para Discord
│   │   ├── email.go                    # Estratégia para Email
│   │   ├── instagram.go                # Estratégia para Instagram
│   │   ├── twitter.go                  # Estratégia para Twitter
│   │   ├── whatsapp.go                 # Estratégia para WhatsApp
│   │   └── unsupported.go              # Estratégia para canais não suportados
│   ├── notification.go                 # Serviço principal
│   ├── notification_strategy.go        # Alias para a interface
│   └── request.go                      # Estruturas de dados
├── main.go                             # Exemplo de uso
└── README.md                           # Esta documentação
```

## Antes vs Depois

### Antes (com if/else)

```go
func (s *NotificationService) Notify(channel string, destination string, message string) {
    if channel == "discord" {
        // lógica para Discord
    } else if channel == "instagram" {
        // lógica para Instagram
    } else if channel == "twitter" {
        // lógica para Twitter
    } else if channel == "whatsapp" {
        // lógica para WhatsApp
    } else if channel == "email" {
        // lógica para Email
    } else {
        // erro para canal não suportado
    }
}
```

### Depois (com Strategy Pattern)

```go
func (s *NotificationService) Notify(channel string, destination string, message string) {
    channel = strings.ToLower(channel)

    strategy, exists := s.strategies[channel]
    if !exists {
        unsupportedStrategy := notifications.NewUnsupportedNotificationStrategy()
        unsupportedStrategy.SendNotification(channel, destination, message)
        return
    }

    strategy.SendNotification(channel, destination, message)
}
```

## Benefícios

1. **Eliminação de if/else**: O código principal não precisa mais de múltiplos condicionais
2. **Extensibilidade**: Para adicionar um novo canal, basta criar uma nova estratégia
3. **Manutenibilidade**: Cada estratégia é isolada e pode ser modificada independentemente
4. **Testabilidade**: Cada estratégia pode ser testada isoladamente
5. **Princípio Open/Closed**: O código está aberto para extensão (novas estratégias) mas fechado para modificação

## Como Usar

```go
// Criar o serviço
notificationService := service.NewNotificationService()

// Enviar notificações
notificationService.Notify("discord", "usuario@exemplo.com", "Mensagem de teste")
notificationService.Notify("email", "usuario@exemplo.com", "Mensagem de teste")
notificationService.Notify("telegram", "usuario@exemplo.com", "Mensagem de teste") // Canal não suportado
```

## Como Adicionar um Novo Canal

1. Criar um novo arquivo em `service/notifications/` (ex: `telegram.go`)
2. Implementar a interface `NotificationStrategy`
3. Adicionar a nova estratégia ao map no construtor `NewNotificationService()`

```go
// Em service/notifications/telegram.go
type TelegramNotificationStrategy struct {}

func NewTelegramNotificationStrategy() interfaces.NotificationStrategy {
    return &TelegramNotificationStrategy{}
}

func (s *TelegramNotificationStrategy) SendNotification(channel string, destination string, message string) {
    // Implementação específica para Telegram
}

// Em service/notification.go, adicionar ao map:
"telegram": notifications.NewTelegramNotificationStrategy(),
```

## Executando o Exemplo

```bash
cd behavioural/strategy
go run main.go
```

O programa demonstrará o envio de notificações para diferentes canais, incluindo um canal não suportado.
