# github.com/maruware/go-kml

Go KML Encoder / Decoder

## Usage

### Decode

```go
p := filepath.Join("KML_Samples.kml")
f, err := os.Open(p)
kml, err := Decode(f)
```

### Encode

```go
buf := &bytes.Buffer{}
Encode(buf, kml)
kmlString := buf.String()
```