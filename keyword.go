package gotools
import (
    "crypto/rand"
    "math/big"
)

func SenhasComplexas(size int) string {
    const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
    var password strings.Builder
    for i := 0; i < size; i++ {
        index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
        password.WriteByte(symbols[index.Int64()])
    }
    return password.String()
}
