# Pseudolang

A command-line interpreter for the pseudolang language with integrated LLM provider management. Pseudolang allows you to configure and switch between multiple LLM providers (OpenAI, Anthropic, Groq, etc.) and models seamlessly.

## Usage

### Configuring Providers

Before using LLM features, configure API tokens for your providers:

```bash
ps provider anthropic sk-ant-...
ps provider openai sk-proj-...
ps provider groq gsk-...
```

This stores the API tokens without changing your active model.

### Switching Models

Switch between models using the model command. The provider is automatically detected from the model name:

```bash
ps model claude-3-5-sonnet-20240620
ps model gpt-4o
ps model llama-3.1-70b-versatile
```

You can also set a model and its provider token in one command:

```bash
ps model claude-3-5-sonnet --token sk-ant-...
```

### Supported Providers

- **OpenAI**: gpt-4, gpt-4o, gpt-3.5-turbo, o1, o3
- **Anthropic**: claude-3-5-sonnet, claude-3-opus, claude-3-haiku
- **Groq**: llama-3, mixtral, gemma
- **Ollama**: Local models (llama, mistral, codellama, phi, qwen)
- **Mistral**: mistral-large, mistral-medium
- **OpenRouter**: Models via OpenRouter API
- **Azure OpenAI**: Azure-hosted OpenAI models

### Running Code

```bash
ps run <file>     # Run a pseudolang file
ps exec <code>    # Execute pseudolang code directly
```

Note: Interpreter functionality is currently in development.

## Development

- `mise run build`: Build the project (outputs to `out/ps`)
- `mise run test`: Run tests
- `mise run check`: Check code with linters and formatters
- `mise run tidy`: Clean up Go module dependencies

Configuration is stored at `~/.config/pseudolang/config.json`.
