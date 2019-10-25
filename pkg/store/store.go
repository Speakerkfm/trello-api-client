package store

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"github.com/go-redis/redis"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/mkideal/pkg/netutil/protocol"
	"gopkg.in/boj/redistore.v1"
	"io"
	"time"
)

var Storage *storage

type storage struct {
	gorm *gorm.DB
	SessionStore sessions.Store
	redisClient *redis.Client
	secret string
}

func Config(secretKey, redisHost, redisPwd string, redisSize int, db *gorm.DB) {
	var rst *redistore.RediStore
	var err error

	redisOpt := redis.Options{Addr: ":6379"}
	redisClient := redis.NewClient(&redisOpt)

	rst, err = redistore.NewRediStore(redisSize, protocol.TCP, redisHost, redisPwd, []byte(secretKey))
	if err != nil {
		panic(err)
	}

	rst.SetSerializer(&hashSerializer{
		jsonSerializer: &redistore.JSONSerializer{},
		secretKey: secretKey,
	})

	Storage = &storage{
		gorm: db,
		SessionStore: rst,
		redisClient: redisClient,
		secret: secretKey,
	}
}

func (s *storage)Set(key string, value string, expiration time.Duration){
	value = string(encrypt([]byte(value), s.secret))

	s.redisClient.Set(key, value, expiration)
}

func (s *storage)Get(key string) string{
	value := s.redisClient.Get(key).Val()

	return string(decrypt([]byte(value), s.secret))
}

type hashSerializer struct {
	jsonSerializer *redistore.JSONSerializer
	secretKey string
}

func (hs *hashSerializer) Serialize(ss *sessions.Session) ([]byte, error) {
	bits, err := hs.jsonSerializer.Serialize(ss)
	if err != nil {
		return nil, err
	}

	return encrypt(bits, hs.secretKey), nil
}

func (hs *hashSerializer) Deserialize(d []byte, ss *sessions.Session) error {
	bits := decrypt(d, hs.secretKey)

	return hs.jsonSerializer.Deserialize(bits, ss)
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

