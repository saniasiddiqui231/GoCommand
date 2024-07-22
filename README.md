# GoShell

GoShell is a simple command-line shell written in Go. It allows users to execute system commands interactively. This shell reads user input, executes commands, and displays output and errors.

## Features

- **Command Execution**: Execute system commands and view their output directly from the shell.
- **Interactive Prompt**: A simple `>` prompt for entering commands.
- **Cross-Platform**: Should work on both Unix-like systems and Windows.

## Installation

To build and run GoCommand, follow these steps:

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/yourusername/goshell.git
    ```

2. **Navigate to the Directory**:
    ```bash
    cd ShellGo
    ```

3. **Build the Application**:
    ```bash
    go build -o ShellGo
    ```

4. **Run the Shell**:
    ```bash
    ./ShellGo
    ```

   On Windows, use:
    ```bash
    ShellGo.exe
    ```

## Usage

Once running, the shell will display a `>` prompt. You can enter system commands at the prompt, and the output will be displayed in the terminal. For example:

```bash
* ls
