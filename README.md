# github.com/maruware/go-kml

Go KML Encoder / Decoder

## Usage

### Decode

```go
p := filepath.Join("KML_Samples.kml")
f, err := os.Open(p)
k, err := kml.Decode(f)
```

### Encode

```go
// var k KML
buf := &bytes.Buffer{}
Encode(buf, k)
kmlString := buf.String()
```