package main

import (
    "fmt"
    "os"

    "../"
)

func main() {
    text := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam
nec hendrerit lectus. Curabitur justo est, auctor ac tempus id, laoreet
sed eros. Pellentesque vehicula mollis ipsum, vel euismod nibh commodo vel.
Sed sit amet dui id augue iaculis molestie ut a justo. Vivamus in dignissim
risus. Etiam tincidunt, quam egestas sagittis ullamcorper, orci orci congue
sapien, in pharetra quam quam non tellus. Donec luctus facilisis magna a
efficitur. Vestibulum iaculis tempor sem sagittis molestie. Mauris fermentum
ut purus vel vulputate. Donec congue tempus felis, sed placerat nulla
viverra a. Sed et ultrices velit.

Nunc quam quam, egestas vitae euismod et, porttitor non urna. Pellentesque
ligula odio, gravida sed ligula et, suscipit placerat urna. Pellentesque
habitant morbi tristique senectus et netus et malesuada fames ac turpis
egestas. Sed ut massa velit. Aliquam imperdiet, ipsum a accumsan ornare,
lacus risus vehicula ipsum, accumsan pharetra quam nibh et ligula. Mauris id
mattis libero. Phasellus sed elit dui. Etiam accumsan justo velit, sed
feugiat sapien lacinia eu. Vestibulum facilisis ullamcorper est, eu
scelerisque lorem accumsan a. Integer lorem dolor, euismod ut lobortis et,
feugiat et enim. Sed condimentum lectus ac magna eleifend, non volutpat dui
faucibus. Cras tristique nisl mi, nec vestibulum lectus tincidunt ac. Aenean
lectus tortor, placerat ac diam nec, cursus dignissim libero. Nam ex sem,
dignissim vel lobortis id, luctus ut ante.

Praesent sit amet justo cursus, varius ante non, rhoncus ex. Fusce lacinia
odio sit amet enim ultrices eleifend. Morbi egestas, sapien in feugiat
consectetur, purus lacus bibendum libero, non commodo purus erat eu elit.
Sed vitae blandit odio. Phasellus iaculis, mi a placerat accumsan, tellus
est sodales nisl, quis varius neque ante ut eros. Vestibulum ultrices, dui
porta rutrum dictum, lectus lectus malesuada nibh, sed fermentum quam nibh
nec est. Cum sociis natoque penatibus et magnis dis parturient montes,
nascetur ridiculus mus. Proin non dolor odio. Pellentesque mollis in eros
a semper.`

    file, err := os.Create("outfile.txt")
    if err != nil {
        fmt.Printf("Unable to open file for writing: %s\n", err)
        return
    }
    defer file.Close()
    fmt.Fprintf(file, "%s", baseEmoji.Utf8Header)

    encoded, err := baseEmoji.DefaultEncoding.Encode([]byte(text))
    if err != nil {
        fmt.Printf("Error encoding: %s\n", err)
        return
    }
    fmt.Fprintf(file, "%s\n\n", encoded)
    
    decoded, err := baseEmoji.DefaultEncoding.Decode(encoded.ToBytes())
    if err != nil {
        fmt.Printf("Error decoding: %s\n", err)
        return
    }
    fmt.Fprintf(file, "%s\n", decoded)
}
