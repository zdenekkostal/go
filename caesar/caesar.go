package caesar

import (
        "flag"
	"fmt"
	"os"
        "bufio"
)

type RuneMapper func(rune) rune
type CipherGetter func() int32

func Encode(text string, getCipherShift CipherGetter) string {
    return mapRunes(text, func (r rune) rune {
        return 'A' + (r - 'A' + 26 + getCipherShift()) % 26
    })
}

func Decode(text string, getCipherShift CipherGetter) string {
    return mapRunes(text, func (r rune) rune {
        return 'A' + (r - 'A' + 26 - getCipherShift()) % 26
    })
}

func GetCipherGetter(cipher string) CipherGetter {
    var cipherIndex = 0

    return func() int32 {
        var shiftBy = int32(cipher[cipherIndex % len(cipher)] - 'A')
        cipherIndex++
        return shiftBy
    }
}

func mapRunes(text string, runeMapper RuneMapper) string {
    var decodedText = make([]rune, len(text))

    for index, r := range text {
        if (r >= 'A' && r <= 'Z') {
            decodedText[index] = runeMapper(r)
        } else {
            decodedText[index] = r
        }
    }

    return string(decodedText)
}

// Usage:
// echo "String to encode" | go run caesar.go -mode=encode -cipher=CIPHERSTRING
// < message_to_encrypt.txt | go run caesar.go -mode=encode -cipher=CIPHERSTRING > encrypted_message.txt
func main() {
    cipher := flag.String("cipher", "", "Cipher")
    mode := flag.String("mode", "encode", "encode/decode")

    // TODO: is there some lighter alternative to Cobra?
    flag.Parse()

    if (len(*cipher) == 0) {
        fmt.Fprintln(os.Stderr, "Plz provide a cipher")
        os.Exit(1)
    }

    getCipherShift := GetCipherGetter(*cipher)
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        if (*mode == "encode") {
            fmt.Println(Encode(scanner.Text(), getCipherShift))
        } else {
            fmt.Println(Decode(scanner.Text(), getCipherShift))
        }
    }
}
