package omnivoice_test

import (
	"os"
	"testing"

	"github.com/plexusone/omnivoice-openai/omnivoice"
	"github.com/plexusone/omnivoice-core/stt"
	"github.com/plexusone/omnivoice-core/stt/providertest"
)

func TestSTTConformance(t *testing.T) {
	apiKey := os.Getenv("OPENAI_API_KEY")

	provider := omnivoice.NewSTTProvider(apiKey)

	providertest.RunAll(t, providertest.Config{
		Provider:        provider,
		SkipIntegration: apiKey == "",
		TestAudioConfig: stt.TranscriptionConfig{
			Encoding:   "mp3",
			SampleRate: 16000,
			Channels:   1,
		},
	})
}
