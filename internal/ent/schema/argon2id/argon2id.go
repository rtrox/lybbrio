package argon2id

import (
	"bytes"
	"crypto/rand"
	"database/sql/driver"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// OWASP Recommended Values:
// https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id
// m=12288 (12 MiB), t=3, p=1
// m=9216  (9 MiB),  t=4, p=1
// m=7168  (7 MiB),  t=5, p=1
type Config struct {
	Iterations  uint32 `koanf:"iterations" validate:"min=3"`
	Memory      uint32 `koanf:"memory" validate:"min=7168"`
	Parallelism uint8  `koanf:"parallelism" validate:"min=1"`
	KeyLen      uint32 `koanf:"keyLength" validate:"min=32"`
	SaltLen     uint32 `koanf:"saltLength" validate:"min=16"`
}

type Argon2IDHash struct {
	Version     uint32 // "V"
	Memory      uint32 // "M"
	Iterations  uint32 // "T"
	Parallelism uint8  // "P"
	Salt        []byte
	Hash        []byte
}

type WrongPasswordError struct{}

func (e WrongPasswordError) Error() string {
	return "wrong password"
}

type InvalidFormatError struct{}

func (e InvalidFormatError) Error() string {
	return "invalid format"
}

func generateSalt(length uint32) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func NewArgon2idHashFromPassword(password []byte, conf Config) (*Argon2IDHash, error) {
	salt, err := generateSalt(conf.SaltLen)
	if err != nil {
		return nil, err
	}
	hash := argon2.IDKey(password, salt, conf.Iterations, conf.Memory, conf.Parallelism, conf.KeyLen)
	ret := &Argon2IDHash{
		Version:     argon2.Version,
		Memory:      conf.Memory,
		Iterations:  conf.Iterations,
		Parallelism: conf.Parallelism,
		Salt:        salt,
		Hash:        hash,
	}
	return ret, nil
}

func (a *Argon2IDHash) Verify(password []byte) error {
	hashed := argon2.IDKey(password, a.Salt, a.Iterations, a.Memory, a.Parallelism, uint32(len(a.Hash)))
	if !bytes.Equal(hashed, a.Hash) {
		return WrongPasswordError{}
	}
	return nil
}

func (a Argon2IDHash) Value() (driver.Value, error) {
	return a.String(), nil
}

func (a Argon2IDHash) String() string {
	b64Salt := base64.RawStdEncoding.EncodeToString(a.Salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(a.Hash)
	return fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		a.Memory,
		a.Iterations,
		a.Parallelism,
		b64Salt,
		b64Hash,
	)
}

func (a *Argon2IDHash) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	switch src := src.(type) {
	case []byte:
		return a.UnmarshalText(string(src))
	case string:
		return a.UnmarshalText(src)
	case Argon2IDHash:
		*a = src
	default:
		return fmt.Errorf("argon2id: unexpected type %T", src)
	}
	return nil
}

func (a *Argon2IDHash) UnmarshalText(encodedHash string) error {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return InvalidFormatError{}
	}
	if vals[1] != "argon2id" {
		return InvalidFormatError{}
	}
	var err error
	_, err = fmt.Sscanf(vals[2], "v=%d", &a.Version)
	if err != nil {
		return InvalidFormatError{}
	}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &a.Memory, &a.Iterations, &a.Parallelism)
	if err != nil {
		return InvalidFormatError{}
	}
	a.Salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return InvalidFormatError{}
	}

	a.Hash, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return InvalidFormatError{}
	}
	return nil
}
