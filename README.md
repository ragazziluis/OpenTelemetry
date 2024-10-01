# 🙋‍ OpenTelemetry - 1 [Luis Miranda]

## :mag: Introdução

Este repositório contém o código do tutorial "OpenTelemetry com Go", focado na aplicação de rastreamento distribuído e observabilidade em aplicações Golang. O tutorial ensina como configurar o OpenTelemetry, criar spans de rastreamento e analisar a execução do código, com exemplos documentados.

## :dart: Objetivo

O objetivo deste repositório é demonstrar a execução dos exemplos fornecidos no tutorial de OpenTelemetry, documentar os resultados com prints, e ampliar os comentários no código para explicar as técnicas e conceitos de instrumentação e observabilidade usando OpenTelemetry.

## :jigsaw: Estrutura do Repositório

- `src/`: Contém o código-fonte dos exemplos do tutorial.
- `assets/`: Contém prints das execuções dos exemplos.
- `docs/`: Contém documentação relacionada ao projeto e as etapas do tutorial.

## :gear: Passos Realizados

### 1. Criação do Repositório no GitHub

Este repositório foi criado para armazenar o código e a documentação do tutorial de OpenTelemetry com Go. O tutorial fonte pode ser encontrado [aqui](https://signoz.io/opentelemetry/go/).

### 2. Execução dos Exemplos e Documentação

Todos os exemplos de rastreamento com OpenTelemetry foram executados e documentados. Os prints das execuções estão disponíveis na pasta `assets/`.

### Exemplos Executados

- **Exemplo 1: Instrumentação Básica com Spans**
    - Configuração do rastreamento e geração de spans.
    - Documentação com o print do rastreamento em JSON no terminal.
- **Exemplo 2: Definição de Atributos de Serviço**
    - Modificação do código para definir o nome do serviço rastreado.

### 3. Ampliação dos Comentários do Código

Os comentários no código foram ampliados para explicar as técnicas e conceitos de OpenTelemetry e rastreamento distribuído. Veja abaixo um exemplo de como os comentários foram aplicados no código:

```go
package main

import (
    "context"
    "fmt"
    "time"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/trace"
)

// main inicia o processo de rastreamento com OpenTelemetry.
// Utilizamos spans para rastrear as operações do sistema.
func main() {
    tracer := otel.Tracer("example-tracer") // Definindo um tracer para rastreamento
    ctx, span := tracer.Start(context.Background(), "main") // Iniciando um span principal
    defer span.End() // Finalizando o span no final da execução

    work(ctx) // Chama uma função simulando uma tarefa
    fmt.Println("Processo completado com sucesso!")
}

// work simula uma operação que será rastreada com um novo span.
func work(ctx context.Context) {
    tracer := otel.Tracer("example-tracer")
    _, span := tracer.Start(ctx, "work") // Inicia um span para a função de trabalho
    defer span.End() // Finaliza o span quando o trabalho termina

    time.Sleep(1 * time.Second) // Simula uma operação que leva 1 segundo
}

```

## :memo: Conclusão

Este repositório fornece um guia prático para aplicar rastreamento distribuído e observabilidade em Go usando OpenTelemetry. A documentação detalha a execução dos exemplos, com prints e comentários ampliados para facilitar a compreensão dos conceitos envolvidos, como a criação de spans, a configuração de serviços, e a importância da instrumentação para monitoramento e depuração.
