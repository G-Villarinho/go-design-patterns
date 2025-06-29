package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// DatabaseManager representa o Singleton para gerenciar conexões com banco de dados
type DatabaseManager struct {
	connectionString string
	isConnected      bool
	lastConnection   time.Time
	connectionCount  int
}

var (
	instance *DatabaseManager
	once     sync.Once
	mu       sync.Mutex
)

// GetInstance retorna a única instância do DatabaseManager
func GetInstance() *DatabaseManager {
	once.Do(func() {
		instance = &DatabaseManager{
			connectionString: "postgresql://localhost:5432/myapp",
			isConnected:      false,
			lastConnection:   time.Time{},
			connectionCount:  0,
		}
		fmt.Println("✅ DatabaseManager Singleton criado pela primeira vez")
	})
	return instance
}

// Connect estabelece uma conexão com o banco de dados
func (dm *DatabaseManager) Connect() error {
	mu.Lock()
	defer mu.Unlock()

	if dm.isConnected {
		fmt.Println("ℹ️  Conexão já está ativa")
		return nil
	}

	// Simula o processo de conexão
	fmt.Println("🔌 Conectando ao banco de dados...")
	time.Sleep(1 * time.Second) // Simula delay de conexão

	dm.isConnected = true
	dm.lastConnection = time.Now()
	dm.connectionCount++

	fmt.Printf("✅ Conectado com sucesso! (Conexão #%d)\n", dm.connectionCount)
	return nil
}

// Disconnect fecha a conexão com o banco de dados
func (dm *DatabaseManager) Disconnect() error {
	mu.Lock()
	defer mu.Unlock()

	if !dm.isConnected {
		fmt.Println("ℹ️  Nenhuma conexão ativa para desconectar")
		return nil
	}

	fmt.Println("🔌 Desconectando do banco de dados...")
	time.Sleep(500 * time.Millisecond) // Simula delay de desconexão

	dm.isConnected = false
	fmt.Println("✅ Desconectado com sucesso!")
	return nil
}

// ExecuteQuery executa uma query no banco de dados
func (dm *DatabaseManager) ExecuteQuery(query string) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	if !dm.isConnected {
		return "", fmt.Errorf("❌ Não há conexão ativa com o banco de dados")
	}

	fmt.Printf("🔍 Executando query: %s\n", query)
	time.Sleep(200 * time.Millisecond) // Simula tempo de execução

	result := fmt.Sprintf("Resultado da query: '%s' - %d registros encontrados", query, len(query))
	return result, nil
}

// GetStatus retorna o status atual da conexão
func (dm *DatabaseManager) GetStatus() string {
	mu.Lock()
	defer mu.Unlock()

	status := "Desconectado"
	if dm.isConnected {
		status = "Conectado"
	}

	return fmt.Sprintf("Status: %s | Conexões realizadas: %d | Última conexão: %s",
		status, dm.connectionCount, dm.lastConnection.Format("15:04:05"))
}

// GetConnectionString retorna a string de conexão
func (dm *DatabaseManager) GetConnectionString() string {
	return dm.connectionString
}

// Example demonstra o uso do padrão Singleton
func Example() {
	fmt.Println("🚀 Iniciando demonstração do padrão Singleton - Database Manager")
	fmt.Println(strings.Repeat("=", 60))

	// Demonstração 1: Múltiplas instâncias apontam para o mesmo objeto
	fmt.Println("\n📋 Demonstração 1: Verificando se múltiplas instâncias são a mesma")

	db1 := GetInstance()
	db2 := GetInstance()
	db3 := GetInstance()

	fmt.Printf("db1: %p\n", db1)
	fmt.Printf("db2: %p\n", db2)
	fmt.Printf("db3: %p\n", db3)
	fmt.Printf("db1 == db2: %t\n", db1 == db2)
	fmt.Printf("db2 == db3: %t\n", db2 == db3)

	// Demonstração 2: Uso concorrente do Singleton
	fmt.Println("\n📋 Demonstração 2: Uso concorrente com múltiplas goroutines")

	var wg sync.WaitGroup

	// Simula múltiplos serviços tentando usar o banco simultaneamente
	services := []string{"UserService", "OrderService", "PaymentService", "NotificationService"}

	for _, service := range services {
		wg.Add(1)
		go func(serviceName string) {
			defer wg.Done()
			useDatabase(serviceName)
		}(service)
	}

	wg.Wait()

	// Demonstração 3: Verificação final do estado
	fmt.Println("\n📋 Demonstração 3: Estado final do Singleton")
	fmt.Println(db1.GetStatus())
}

// useDatabase simula o uso do banco de dados por um serviço
func useDatabase(serviceName string) {
	db := GetInstance()

	fmt.Printf("\n🔧 %s iniciando operações...\n", serviceName)

	// Conecta ao banco
	err := db.Connect()
	if err != nil {
		fmt.Printf("❌ %s: Erro ao conectar: %v\n", serviceName, err)
		return
	}

	// Executa algumas queries
	queries := []string{
		fmt.Sprintf("SELECT * FROM %s_users", serviceName),
		fmt.Sprintf("SELECT * FROM %s_config", serviceName),
	}

	for _, query := range queries {
		result, err := db.ExecuteQuery(query)
		if err != nil {
			fmt.Printf("❌ %s: Erro na query: %v\n", serviceName, err)
		} else {
			fmt.Printf("✅ %s: %s\n", serviceName, result)
		}
		time.Sleep(100 * time.Millisecond) // Simula processamento
	}

	fmt.Printf("🏁 %s finalizou operações\n", serviceName)
}

// ThreadSafeExample demonstra a thread-safety do Singleton
func ThreadSafeExample() {
	fmt.Println("\n🛡️  Demonstração de Thread Safety")
	fmt.Println(strings.Repeat("=", 40))

	var wg sync.WaitGroup
	instances := make([]*DatabaseManager, 10)

	// Cria múltiplas goroutines para obter instâncias simultaneamente
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			instances[index] = GetInstance()
			fmt.Printf("Goroutine %d obteve instância: %p\n", index, instances[index])
		}(i)
	}

	wg.Wait()

	// Verifica se todas as instâncias são iguais
	firstInstance := instances[0]
	allSame := true

	for i, instance := range instances {
		if instance != firstInstance {
			allSame = false
			fmt.Printf("❌ Instância %d é diferente!\n", i)
		}
	}

	if allSame {
		fmt.Println("✅ Todas as instâncias são iguais - Thread safety confirmada!")
	} else {
		fmt.Println("❌ Instâncias diferentes encontradas!")
	}
}

func main() {
	fmt.Println("🎯 Padrão Singleton - Exemplo do Mundo Real")
	fmt.Println("Database Connection Manager")
	fmt.Println()

	// Executa a demonstração principal
	Example()

	// Aguarda um pouco antes da próxima demonstração
	time.Sleep(2 * time.Second)

	// Executa a demonstração de thread safety
	ThreadSafeExample()

	fmt.Println("\n🎉 Demonstração concluída!")
	fmt.Println("\n💡 Pontos importantes do Singleton demonstrados:")
	fmt.Println("   • Garantia de uma única instância")
	fmt.Println("   • Thread safety com sync.Once")
	fmt.Println("   • Estado compartilhado entre todas as chamadas")
	fmt.Println("   • Controle centralizado de recursos")
	fmt.Println("   • Economia de memória e recursos")
}
