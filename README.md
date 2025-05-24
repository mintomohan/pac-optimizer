# pac-optimizer

A simple command-line utility written in Go to optimize PAC (Proxy Auto-Config) files by removing unnecessary comments and blank lines. Useful for reducing file size, improving readability, and eliminating noise in PAC scripts.

## Features

- Removes single-line `//` comments
- Eliminates blank or whitespace-only lines
- Detects and preserves original line endings (Unix, Windows, or old Mac)
- Outputs a cleaned and optimized PAC file
- Written in idiomatic and efficient Go


## Installation

### Clone & Build

```bash
git clone https://github.com/mintomohan/pac-optimizer.git
cd pac-optimizer
go build pac-optimizer.go
```

### Or install directly

```bash
go install github.com/mintomohan/pac-optimizer@latest
```

## Usage
```bash
./pac-optimizer [options] <input-file> <output-file>

Options:
  --version     Show version information
  --help        Show this help message

Arguments:
  <input-file>  Path to the PAC file to optimize
  <output-file> Path where the optimized PAC file will be saved
```

Example
```bash
./pac-optimizer proxy.pac proxy.optimized.pac
```

## Limitations
Only supports single-line (```//```) comments.

Does not support block (```/* ... */```) comment removal or parse full JavaScript syntax (like string literals with embedded //).

Not suitable for PAC files with dynamic, complex JS constructs.


## License
MIT License. See LICENSE for details.


## Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you’d like to improve.

