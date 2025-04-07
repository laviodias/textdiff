# textdiff

**textdiff** is a lightweight and expressive Go library for comparing two texts and generating inline diffs in a Git-style format. It highlights added and removed words, making it ideal for applications like text editors, document review tools, and change tracking systems.

---

## ✨ Example

```go
package main

import (
    "fmt"
    "textdiff/diff"
)

func main() {
    a := "my name is John and I like apples"
    b := "my name is Josh and I love oranges"

    result := diff.CompareTexts(a, b)
    fmt.Println(result)
}
```

### 🖨️ Output

```
my name is [-John-] {+Josh+} and I [-like apples-] {+love oranges+}
```

---

## 📦 Installation

Add the module to your Go project using:

```bash
go get github.com/laviodias/textdiff/diff
```

---

## 🧪 Tests

Run tests with:

```bash
go test ./...
```

---

## 📜 License

MIT License. See the [LICENSE](LICENSE) file for details.

---

## 🤝 Contributing

Contributions are welcome! Please open issues or pull requests for suggestions or improvements.

---

## 🌐 Package Documentation

Documentation available at:

[https://pkg.go.dev/github.com/laviodias/textdiff](https://pkg.go.dev/github.com/laviodias/textdiff)
---

