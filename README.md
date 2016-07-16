# baseEmoji

Encode things in Emoji because reasons.

You probably shouldn't use this in any kind of serious production environment.

## Usage

[![img](https://godoc.org/github.com/zorchenhimer/baseEmoji?status.svg)](https://godoc.org/github.com/zorchenhimer/baseEmoji)

```
    text := "Homura did nothing wrong."
    enc := baseEmoji.DefaultEncoding
    
    // Encode the text string
    encoded, err := enc.Encode([]byte(text))
    if err != nil {
        fmt.Printf("Error encoding: %s\n", err)
        return
    }
    fmt.Printf(file, "%s\n", encoded)
    
    // Write the encoded string to a text file
    encoding.WriteFile("output.txt")
    
    // Decode the dext string
    decoded, err := enc.Decode(encoded.ToBytes())
    if err != nil {
        fmt.Printf("Error decoding: %s\n", err)
        return
    }
    fmt.Printf(file, "%s\n", decoded)
```

## License

See LICENSE.md
