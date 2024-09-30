package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace" // Alias para evitar conflito com a variável trace
)

// initTracer configura o rastreamento para exportar para o terminal.
func initTracer() func() {
	// Exporta os dados de rastreamento para o terminal.
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalf("failed to initialize stdouttrace exporter %v", err)
	}

	tp := sdktrace.NewTracerProvider( // Usando alias para evitar conflito com a variável trace
		sdktrace.WithBatcher(exp),
	)

	otel.SetTracerProvider(tp)

	// Função para encerrar o TracerProvider de forma segura.
	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatalf("failed to shutdown TracerProvider %v", err)
		}
	}
}

func main() {
	cleanup := initTracer()
	defer cleanup()

	tracer := otel.Tracer("example-tracer")
	ctx, span := tracer.Start(context.Background(), "main")
	defer span.End()

	// Simulação de uma função que executa algum processo
	work(ctx)

	fmt.Println("Processo completado com sucesso!")
}

func work(ctx context.Context) {
	tracer := otel.Tracer("example-tracer")
	_, span := tracer.Start(ctx, "work")
	defer span.End()

	// Simula trabalho com espera
	time.Sleep(1 * time.Second)
}
