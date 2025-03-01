package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Limpa o ambiente antes de cada teste
	os.Clearenv()

	// Caso 1: Configuração válida com todas as chaves
	t.Run("ValidConfig", func(t *testing.T) {
		// Define variáveis de ambiente
		os.Setenv("YOUTUBE_API_KEY", "youtube-test-key")
		os.Setenv("OPENAI_API_KEY", "openai-test-key")
		os.Setenv("TRANSCRIPTION_API_KEY", "transcription-test-key")

		// Chama a função
		err := LoadConfig()
		if err != nil {
			t.Errorf("esperava nil, mas obteve erro: %v", err)
		}

		// Verifica se as chaves foram carregadas corretamente
		if GlobalConfig.YouTubeAPIKey != "youtube-test-key" {
			t.Errorf("esperava YouTubeAPIKey 'youtube-test-key', mas obteve '%s'", GlobalConfig.YouTubeAPIKey)
		}
		if GlobalConfig.OpenAIAPIKey != "openai-test-key" {
			t.Errorf("esperava OpenAIAPIKey 'openai-test-key', mas obteve '%s'", GlobalConfig.OpenAIAPIKey)
		}
		if GlobalConfig.TranscriptionAPIKey != "transcription-test-key" {
			t.Errorf("esperava TranscriptionAPIKey 'transcription-test-key', mas obteve '%s'", GlobalConfig.TranscriptionAPIKey)
		}
	})

	// Caso 2: Falta a chave obrigatória da OpenAI
	t.Run("MissingOpenAIKey", func(t *testing.T) {
		// Limpa o ambiente novamente
		os.Clearenv()

		// Define apenas algumas chaves, mas não a OpenAI
		os.Setenv("YOUTUBE_API_KEY", "youtube-test-key")
		os.Setenv("TRANSCRIPTION_API_KEY", "transcription-test-key")

		// Chama a função
		err := LoadConfig()
		if err == nil {
			t.Error("esperava erro por falta de OPENAI_API_KEY, mas obteve nil")
		}
	})

	// Caso 3: Sem .env, mas com variáveis de ambiente
	t.Run("NoDotEnvButEnvVars", func(t *testing.T) {
		// Limpa o ambiente
		os.Clearenv()

		// Define variáveis diretamente
		os.Setenv("OPENAI_API_KEY", "openai-test-key")

		// Chama a função
		err := LoadConfig()
		if err != nil {
			t.Errorf("esperava nil com variáveis definidas, mas obteve erro: %v", err)
		}
		if GlobalConfig.OpenAIAPIKey != "openai-test-key" {
			t.Errorf("esperava OpenAIAPIKey 'openai-test-key', mas obteve '%s'", GlobalConfig.OpenAIAPIKey)
		}
	})
}
