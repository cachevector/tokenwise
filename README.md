<div align="center">
  <img src="./assets/icons/icon.svg" alt="TokenWise Logo" width="120" />
  <h1>TokenWise</h1> 
  
  A GUI tool for bidirectional conversion between JSON and TOON to optimize LLM token usage.
</div>

---

## Installation

### Prerequisites
- Go 1.20+

### Build from source
```bash
git clone https://github.com/cachevector/tokenwise
cd tokenwise
go build -o bin/tokenwise ./cmd/tokenwise
```

### Run without building
```bash
go run cmd/tokenwise/*.go
```

## Usage
- Launch the application.
- Paste JSON or TOON data into the input pane.
- Select input and output formats.
- Click Convert.
- View converted output and token counts.

## Dependencies

- [**Fyne**](https://fyne.io/) – GUI toolkit for Go.

- [**toon-go**](https://github.com/toon-format/toon-go) – TOON format library.

## License
This project is licensed under the [**MIT License**](./LICENSE).
