# pdfiler

`pdfiler` is a Command-Line Interface (CLI) tool written in Go using the Cobra framework. It converts Markdown files into beautifully formatted PDF documents. This tool is lightweight, fast, and easy to integrate into your workflows.

## Features

- Converts Markdown (`.md`) files into PDF documents.
- Supports basic Markdown syntax for headers, lists, links, and more.
- Customizable output file naming and location.
- Lightweight and efficient, built with the Go programming language.

## Installation

### Prerequisites

- Go (1.20 or higher recommended)

### Installation via `go install`

```bash
go install github.com/CreepySunny/pdfiler@latest
```

Ensure `$GOPATH/bin` is in your `PATH` to access `pdfiler` globally.

## Usage

### Basic Syntax

```bash
pdfiler convert [flags] <markdown_file>
```

### Examples

#### Convert a Markdown File to PDF

```bash
pdfiler convert example.md
```

This generates `example.pdf` in the current directory.

#### Specify Output Filename

```bash
pdfiler convert example.md -o custom_name.pdf
```

#### Set Output Directory

```bash
pdfiler convert example.md -d ./output
```

#### Full Help

```bash
pdfiler convert --help
```

### Flags

- `-o, --output` (string): Specify the name of the output PDF file. Defaults to the input filename with `.pdf` extension.
- `-d, --directory` (string): Specify the directory for the output PDF. Defaults to the current directory.
- `-h, --help`: Display help information about the command.

## How It Works

`pdfiler` parses the input Markdown file, applies styling and layout, and uses a PDF generation library to produce the final PDF file. It's ideal for developers and writers who need quick and reliable Markdown-to-PDF conversion.

## Development

### Clone the Repository

```bash
git clone https://github.com/CreepySunny/pdfiler.git
cd pdfiler
```

### Build Locally

```bash
go build -o pdfiler
```

### Run Locally

```bash
./pdfiler convert example.md
```

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes and push to your fork.
4. Submit a pull request.

Please adhere to the [Go Code of Conduct](https://golang.org/conduct).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Author

- **Sidney Thiel** – [GitHub](https://github.com/CreepySunny)
