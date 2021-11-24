# binarycursor

This is a small library that implements a stateful cursor which can read binary data of various types. Handy for parsing binary formats.

## usage

```go
reader := bytes.NewReader(data)
c := NewBinaryCursor(reader)

// basic types
_, err := c.ReadUint8()
_, err := c.ReadUint16()
_, err := c.ReadUint32()
_, err := c.ReadUint64()


// LEB128
_, err := c.ReadUleb128()
_, err := c.ReadSleb128()
```

### Suggestions or Feedback

We'd love to hear from you! Get in touch via [twitter](https://twitter.com/chimehq), an issue, or a pull request.

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project you agree to abide by its terms.
