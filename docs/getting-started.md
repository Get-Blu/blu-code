# Getting Started with Blu

This guide will help you install, configure, and start using Blu for the first time.

## Prerequisites

- Go 1.21 or higher (if building from source)
- Node.js (for npm installation)
- An API key from a supported provider (Anthropic, OpenAI, or Google Gemini)

## Installation

### Via npm

This is the recommended method for most users.

```bash
npm install -g @get-blu/blu-code
```

### Via Raw Script

For environments without npm:

```bash
curl -fsSL https://raw.githubusercontent.com/Get-Blu/blu-code/main/install | bash
```

## First Run

1. Open your terminal in a project directory.
2. Set your API key as an environment variable:
   ```bash
   export ANTHROPIC_API_KEY="your-api-key"
   ```
3. Launch the application:
   ```bash
   blu
   ```

## Initial Setup

When you first launch Blu, it will look for a `.blu.json` configuration file in your home directory or the current working directory. If none is found, it will use default settings. It is recommended to create a configuration file to specify your preferred models and settings.
