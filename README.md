# System Utils
This terminal-based program allows you to automatically close applications or shut down your system after a specified period. Originally created to provide a sleep timer for the Spotify desktop app, which lacks this feature, it can be used for other general task management purposes as well. Designed to be run from the command line, this utility offers a convenient way to automate system tasks directly from your terminal.

## Features
- Set a timer to close any running application.
- Schedule system shutdowns.
- Supports multiple applications.
- Simple command-line interface for easy usage.


## Installation 
<img src="docs/readme-assets/windows_logo.svg" width="100" height="75"> <br /> Latest Release 
[Download](https://github.com/nekorionebula/system-utils/releases/latest/download/system-utils.exe) <br />

## Usage
Run the program from the command line and specify the application you want to close, the duration in minutes, or schedule a system shutdown
### Kill App
- Schedule the application to close after a specified duration (in seconds).
- Provide the name of the application without the `.exe` extension.

### Shutdown
Equivalent to:
```bash
shutdown -s -t <time>
