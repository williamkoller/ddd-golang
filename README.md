[![Continuous Integration](https://github.com/williamkoller/ddd-golang/actions/workflows/continuous-integration.yml/badge.svg)](https://github.com/williamkoller/ddd-golang/actions/workflows/continuous-integration.yml)

# 🏗️ Domain-Driven Design com GoLang

Este projeto é uma implementação de um sistema simples utilizando os princípios de Domain-Driven Design (DDD) em GoLang. Ele simula a criação e manipulação de ordens de compra (Orders), aplicando conceitos de Entities, Value Objects, Repositories, e Use Cases.

# 🛠️ Estrutura do Projeto

```bash
ddd-golang/
├── cmd/
│   └── main.go                # Ponto de entrada da aplicação
├── domain/
│   ├── entities/
│   │   └── order.go           # Entidade Order
│   ├── value_objects/
│   │   └── price.go           # Value Object Price
│   └── repositories/
│       └── order_repository.go # Interface para o repositório de Orders
├── infrastructure/
│   └── persistence/
│       └── order_repository.go # Implementação do repositório
├── usecases/
│   └── create_order.go        # Caso de uso para criar uma Order
└── go.mod                     # Configuração do módulo Go
```

# ✨ Funcionalidades 
* Criação de Ordens: Simula a criação de uma ordem com ID único e preço total.

* Manipulação de Preços: Validação e cálculos utilizando o Value Object Price.

* Repositórios: Implementa o padrão de repositório para armazenar as ordens.

* Arquitetura Limpa: Divisão clara entre camadas (Domínio, Casos de Uso, Infraestrutura).


# 🚀 Como Executar
Pré-requisitos
* GoLang instalado na versão mais recente.

* Ambiente configurado para desenvolvimento em Go.

# Passos
1. Clone o repositório:

```bash
git clone https://github.com/williamkoller/ddd-golang.git
cd ddd-golang
```

2. Instale as dependências:

```bash
go mod tidy
```

3. Execute o projeto:

```bash

go run cmd/main.go
```

# 🧩 Componentes Principais
1. Entidades
   Representam os objetos principais do domínio, como Order:

```go
type Order struct {
	ID         string
	TotalPrice value_objects.Price
	Products   []string
}
```

2. Value Objects
   Representam conceitos do domínio com validações e comportamentos específicos, como Price:

```go
type Price struct {
	Amount float64
}

func NewPrice(amount float64) (Price, error) {
	if amount < 0 {
		return Price{}, fmt.Errorf("price cannot be negative")
	}
	return Price{Amount: amount}, nil
}
```

3. Repositórios
   Interfaces que definem como interagir com a infraestrutura:

```go
type OrderRepository interface {
	Save(order entities.Order) error
	FindById(id string) (entities.Order, error)
}
```

# 🔍 Exemplo de Uso
Criação de uma ordem utilizando o caso de uso CreateOrder:

```go
package main

import (
	"ddd-golang/domain/value_objects"
	"ddd-golang/infrastructure/persistence"
	"ddd-golang/usecases"
	"fmt"
)

func main() {
	repo := persistence.NewInMemoryOrderRepository()
	useCase := usecases.NewCreateOrderUseCase(repo)

	price, _ := value_objects.NewPrice(100.0)
	order, _ := useCase.Execute("1234", price)

	fmt.Printf("Ordem criada com ID: %s\n", order.ID)
}
```

# Testes

Para executar os testes, utilize o comando:

```bash
go test ./...

ok      ddd-golang/cmd  0.002s
ok      ddd-golang/domain/entities      0.002s
ok      ddd-golang/domain/repositories  0.002s
ok      ddd-golang/domain/value_objects 0.002s
ok      ddd-golang/usecases     0.002s

```

