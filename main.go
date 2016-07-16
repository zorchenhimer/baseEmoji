package baseEmoji

import (
    "bytes"
    "encoding/base64"
    "fmt"
    "io/ioutil"
    "unicode/utf8"
)

// Emoji is a single character "rune" encoded as a byte slice
type Emoji string
type EmojiString []Emoji

type Encoding struct {
    encode      EmojiString
    encodeMap   map[byte]Emoji
    decodeMap   map[Emoji]byte
    padChar     Emoji
}

var DefaultEncoding *Encoding

// Build the default encoding
func init() {
    encode := EmojiString{
        u1F601, u1F602, u1F603, u1F604, u1F605, u1F606, u1F609, u1F60A, u1F60B, u1F60C,
        u1F60D, u1F60F, u1F612, u1F613, u1F614, u1F616, u1F618, u1F61A, u1F61C, u1F61D,
        u1F61E, u1F620, u1F621, u1F622, u1F623, u1F624, u1F625, u1F628, u1F629, u1F62A,
        u1F62B, u1F62D, u1F630, u1F631, u1F632, u1F633, u1F635, u1F637, u1F638, u1F639,
        u1F63A, u1F63B, u1F63C, u1F63D, u1F63E, u1F63F, u2702, u2705, u2708, u2709,
        u270A, u270B, u270C, u270F, u2712, u2714, u2716, u2728, u2733, u2734,
        u2744, u2747, u274C, u274E, u2753,
    }

    var err error
    DefaultEncoding, err = NewEncode(encode)
    if err != nil {
        panic(fmt.Sprintf("Default encoding failed to init(): %s", err))
    }
}

// NewEncode returns a new Encoding given the EmojiString provided.  An error is
// returned if there are not 65 characters in the EmojiString.  The last
// character is the padding character.
func NewEncode(emString EmojiString) (*Encoding, error) {
    if len(emString) != 65 {
        return nil, fmt.Errorf("EmojiString is not 65 characters long! (the 65th is the padding character)")
    }
    
    enc := &Encoding{
        encode:     emString,
        encodeMap:  make(map[byte]Emoji),
        decodeMap:  make(map[Emoji]byte),
        padChar:    u2753,
    }

    for idx, em := range enc.encode {
        enc.decodeMap[em] = encodeStd[idx]
        enc.encodeMap[encodeStd[idx]] = em
    }
    return enc, nil
}

// Encode takes a byte slice and encodes it into an EmojiString encoded in 
// UTF-8 bytes.  This does not add the UTF-8 Byte Order Mark (BOM).
func (enc *Encoding) Encode(input []byte) (*EmojiString, error) {
    b64 := base64.StdEncoding.EncodeToString(input)
    es := EmojiString{}

    for i := 0; i < len(b64); i++ {
        em := enc.encodeMap[b64[i]]
        es = append(es, em)
    }

    return &es, nil
}

// Decode takes the input byte slice of UTF-8 encoded emoji and decodes it into
// it's original binary blob as a byte slice.
func (enc *Encoding) Decode(input []byte) ([]byte, error) {
    b64 := []byte{}
    inputEmoji := []Emoji{}

    // Remove the UTF8 BOM if it exists
    if len(input) >= len(Utf8Header) && bytes.Equal(input[0:len(Utf8Header)-1], Utf8Header) {
        input = input[len(Utf8Header):]
    }

    for len(input) > 0 {
        r, size := utf8.DecodeRune(input)
        inputEmoji = append(inputEmoji, Emoji(r))

        input = input[size:]
    }

    for idx, em := range inputEmoji {
        r, _ := utf8.DecodeRune([]byte(em))
        char, ok := enc.decodeMap[em]
        if !ok {
            return nil, fmt.Errorf("Invalid character at offset %d: %U: %c", idx + 1, r, r)
        }
        b64 = append(b64, char)
    }

    return base64.StdEncoding.DecodeString(string(b64))
}

// WriteFile writes the emoji string to the given filename.  This prepends the
// UTF-8 BOM automatically.
func (es *EmojiString) WriteFile(filename string) error {
    return ioutil.WriteFile(filename, append(Utf8Header, es.ToBytes()...), 0655)
}

// ToByte returns the EmojiString as a byte slice.  This does not prepend the UTF-8 BOM.
func (es *EmojiString) ToBytes() []byte {
    b := []byte{}

    for _, e := range *es {
        b = append(b, []byte(e)...)
    }

    return b
}

func (es EmojiString) String() string {
    s := ""
    for _, e := range es {
        s += string(e)
    }
    return s
}
