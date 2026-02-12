# Blu

<p align="center">
  <strong>The next-generation terminal-based AI assistant for software engineers.</strong>
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#configuration">Configuration</a> •
  <a href="#usage">Usage</a>
</p>

---

## What is Blu?

**Blu** is a high-performance, Go-powered terminal assistant designed to bridge the gap between your local environment and powerful AI models. It provides a fluid, interactive TUI that allows you to chat with AI, let it analyze your codebase, and execute complex developer workflows without ever leaving your terminal.

> [!IMPORTANT]
> **Built with on top of [OpenCode](https://github.com/opencode-ai/opencode).**  
> Blu is a specialized distribution and evolution of the original OpenCode repository, optimized for a premium user experience and unique developer features.

---

## The Blu Experience

### Intelligence, Unbound
Supported by every world-class AI provider. Whether you prefer **Claude 3.7**, **Gemini 2.5 Pro**, or **GPT-4o**, Blu speaks them all. Mix and match models for different tasks—titles, summaries, or deep-dives.

### Tool-Integrated Workflow
Blu isn't just a chatbot; it's an operator.
- **File System Mastery**: Glob, grep, read, and write with surgical precision.
- **System Execution**: Run bash commands directly from the chat.
- **Smart Patching**: Apply complex diffs and refactors automatically.

### Premium Terminal UI
Built with the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework, Blu offers a state-of-the-art interface:
- **Session Focus**: Switch between multiple context-heavy sessions instantly (`Ctrl+S`).
- **Interactive Sidebar**: Visualize file changes and diffs in real-time.
- **Vim-like Command Pallet**: Execute custom logic at the speed of thought.

---

## Feature Showcase

| Pillar | Capability |
| :--- | :--- |
| **Deep Integration** | **LSP Support & Diagnostics**: Real-time error checking and code intelligence via Language Server Protocol. |
| | **MCP Readiness**: Extend Blu with any Model Context Protocol server for infinite tool potential. |
| **Workflow Power** | **Named Custom Commands**: Create reusable prompt templates with dynamic placeholders (`Ctrl+K`). |
| | **Self-Hosting**: Full support for local models via OpenAI-compatible endpoints. |
| **Smart Management** | **Auto-Compaction**: Intelligent context window management that summarizes long sessions so you never lose track. |
| | **SQLite Persistence**: Your entire history is stored locally and securely. |

---

## Installation

### Quick Install
```bash
curl -fsSL https://raw.githubusercontent.com/Get-Blu/blu-code/main/install | bash
```

### From Source
```bash
git clone https://github.com/Get-Blu/blu-code.git
cd blu-code
go build -o blu
```

---

## Configuration

Blu keeps its soul in `.blu.json`. You can place this in your home directory or locally in a project.

```json
{
  "agents": {
    "coder": {
      "model": "claude-3.7-sonnet",
      "maxTokens": 8192
    }
  },
  "autoCompact": true
}
```

### Environment Variables
Configure your keys instantly:
- `ANTHROPIC_API_KEY`
- `OPENAI_API_KEY`
- `GEMINI_API_KEY`
- `BLU_DEBUG=true`

---

## Acknowledgments & Credits

Blu is stands on the shoulders of giants:
- **Base Engine**: Created by [opencode-ai/opencode](https://github.com/opencode-ai/opencode).
- **Core Design**: Inspired by the amazing work of [@adamdottv](https://github.com/adamdottv).
- **LSP Foundation**: Powered by logic from [@isaacphi's](https://github.com/isaacphi) `mcp-language-server`.

---

## License

Distributed under the **MIT License**. See `LICENSE` for more information.

<p align="center">
  Made by Garv Agnihotri
</p>
