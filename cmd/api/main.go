package main

import (
	"internal/interfaces/router"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
    Env         string
    Addr        string
    DatabaseURL string
}

func loadConfig() Config {
    cfg := Config{
        Env:         getEnv("APP_ENV", "dev"),
        Addr:        getEnv("ADDR", ":8080"),
        DatabaseURL: mustEnv("DATABASE_URL"),
    }
    return cfg
}

func getEnv(k, def string) string { if v := os.Getenv(k); v != "" { return v }; return def }
func mustEnv(k string) string     { v := os.Getenv(k); if v == "" { log.Fatalf("%s is required", k) }; return v }

func main() {
  cfg := loadConfig()

  r := router.NewRouter()
  server := &http.Server{
    Addr: cfg.Addr,
    Handler: r,
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 15 * time.Second,
    IdleTimeout: 60 * time.Second,
  }
  
  go func() {
    log.Printf("Server starting on addr: %s", server.Addr)
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      log.Fatal("Server error", err)
    }
  }()
}

