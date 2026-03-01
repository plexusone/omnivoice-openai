# omnivoice-openai

[![Go Reference](https://pkg.go.dev/badge/github.com/plexusone/omnivoice-openai.svg)](https://pkg.go.dev/github.com/plexusone/omnivoice-openai)
[![CI](https://github.com/plexusone/omnivoice-openai/actions/workflows/ci.yaml/badge.svg)](https://github.com/plexusone/omnivoice-openai/actions/workflows/ci.yaml)
[![Lint](https://github.com/plexusone/omnivoice-openai/actions/workflows/lint.yaml/badge.svg)](https://github.com/plexusone/omnivoice-openai/actions/workflows/lint.yaml)

OpenAI audio provider for the [OmniVoice](https://github.com/plexusone/omnivoice-core) voice pipeline framework.

## Features

- **STT (Speech-to-Text)**: Whisper transcription with word and segment timestamps
- **TTS (Text-to-Speech)**: OpenAI audio synthesis with multiple voices
- **OmniVoice Integration**: Implements `stt.Provider` and `tts.Provider` interfaces

## Installation

```bash
go get github.com/plexusone/omnivoice-openai
```

## Usage

### Direct Client Usage

```go
package main

import (
    "context"
    "log"

    "github.com/plexusone/omnivoice-openai"
)

func main() {
    // Create client from environment variable
    client, err := openai.NewClientFromEnv()
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Transcribe audio
    resp, err := client.Transcribe(ctx, openai.TranscriptionRequest{
        Audio:    audioData,
        Filename: "audio.mp3",
    })
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Transcription: %s", resp.Text)

    // Synthesize speech
    ttsResp, err := client.Synthesize(ctx, openai.TTSRequest{
        Input: "Hello, world!",
        Voice: openai.VoiceAlloy,
    })
    if err != nil {
        log.Fatal(err)
    }
    // ttsResp.Audio contains the MP3 audio data
}
```

### OmniVoice Provider Usage

```go
package main

import (
    "context"

    "github.com/plexusone/omnivoice-core/stt"
    "github.com/plexusone/omnivoice-core/tts"
    openaistt "github.com/plexusone/omnivoice-openai/omnivoice/stt"
    openaitts "github.com/plexusone/omnivoice-openai/omnivoice/tts"
)

func main() {
    ctx := context.Background()

    // Create STT provider
    sttProvider := openaistt.NewProvider()
    transcription, err := sttProvider.Transcribe(ctx, audioData)

    // Create TTS provider
    ttsProvider := openaitts.NewProvider()
    audio, err := ttsProvider.Synthesize(ctx, "Hello, world!")
}
```

## Configuration

Set the `OPENAI_API_KEY` environment variable or pass the API key directly:

```go
client := openai.NewClient("your-api-key")
```

## Available Voices

| Voice | Description |
|-------|-------------|
| alloy | Neutral, balanced |
| ash | Warm, engaging |
| ballad | Melodic, expressive |
| coral | Clear, articulate |
| echo | Smooth, natural |
| fable | Storytelling, dramatic |
| nova | Bright, energetic |
| onyx | Deep, resonant |
| sage | Calm, wise |
| shimmer | Light, cheerful |
| verse | Poetic, rhythmic |
| marin | Friendly, approachable |
| cedar | Grounded, trustworthy |

## License

MIT License - see [LICENSE](LICENSE) for details.
