package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config armazena as configurações carregadas
type Config struct {
	YouTubeAPIKey       string
	OpenAIAPIKey        string
	TranscriptionAPIKey string
}

// GlobalConfig é a variável global para acessar as configurações
var GlobalConfig Config

// LoadConfig carrega as configurações do arquivo .env
func LoadConfig() error {
	// Carrega o arquivo .env (se existir)
	err := godotenv.Load()
	if err != nil {
		// Ignora erro se .env não existir (pode usar variáveis de ambiente diretamente)
	}

	// Carrega as chaves de API das variáveis de ambiente
	GlobalConfig = Config{
		YouTubeAPIKey:       os.Getenv("YOUTUBE_API_KEY"),
		OpenAIAPIKey:        os.Getenv("OPENAI_API_KEY"),
		TranscriptionAPIKey: os.Getenv("TRANSCRIPTION_API_KEY"),
	}

	// Validação básica: verifica se as chaves principais estão presentes
	if GlobalConfig.OpenAIAPIKey == "" {
		return os.ErrInvalid
	}

	return nil
}
