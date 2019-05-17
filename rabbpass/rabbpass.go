//tool for generation password hash for rabbitmq

package main

import (
    "crypto/rand"
    "crypto/sha256"
    "encoding/base64"
    "flag"
    "fmt"
    mRand "math/rand"
    "time"
)

var src = mRand.NewSource(time.Now().UnixNano())

func main() {

    input := flag.String("password", "", "The password to be encoded. One will be generated if not supplied")

    flag.Parse()

    salt := [4]byte{}
    _, err := rand.Read(salt[:])
    if err != nil {
        panic(err)
    }

    pass := *input
    if len(pass) == 0 {
        pass = randomString(32)
    }

    saltedP := append(salt[:], []byte(pass)...)

    hash := sha256.New()

    _, err = hash.Write(saltedP)

    if err != nil {
        panic(err)
    }

    hashPass := hash.Sum(nil)

    saltedP = append(salt[:], hashPass...)

    b64 := base64.StdEncoding.EncodeToString(saltedP)

    fmt.Printf("Password: %s\n", string(pass))
    fmt.Printf("Hash: %s\n", b64)
}

const (
    letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randomString(size int) string {
    b := make([]byte, size)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax letters!
    for i, cache, remain := size-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return string(b)

}
