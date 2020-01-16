package main
import(
    // "encoding/base64"
	"encoding/hex"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/md5"
	"fmt"
)

func main() {
	md5, sha1, sha256 := multipleReturn("hash this text")
	fmt.Println("Md5:", md5)
	fmt.Println()

	fmt.Println("Sha1:", sha1)
	fmt.Println()

	fmt.Println("Sha256:", sha256)
}

func multipleReturn(text string) (string, string, string) {
	md5 := encryptMd5(text)
	sha1 := encryptSha1(text)
	sha256 := encryptSha256(text)
    return md5, sha1, sha256
}

func encryptMd5(text string) string {
	hasher := md5.New()
    hasher.Write([]byte(text))
    md5 := hex.EncodeToString(hasher.Sum(nil))

    return md5
}

func encryptSha1(text string) string {
    hasher := sha1.New()
    hasher.Write([]byte(text))
    sha1 := hex.EncodeToString(hasher.Sum(nil))
    return sha1
}

func encryptSha256(text string) string {
    hasher := sha256.New()
    hasher.Write([]byte(text))
    sha256 := hex.EncodeToString(hasher.Sum(nil))
    return sha256
}