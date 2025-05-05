package prot

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type UniqueSkipper struct {
	StorePath string

	log   *LogBehavior
	store SkipStore
}

func (s *UniqueSkipper) Inject(log *LogBehavior) {
	s.log = log
}

func (s *UniqueSkipper) StartUp() error {
	if s.StorePath == "" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		s.StorePath = filepath.Join(homedir, "tmp", "speakitskip.json")
	}
	if err := s.read(); err != nil {
		return err
	}
	s.flush()
	s.write()

	return nil
}

func (s *UniqueSkipper) flush() {
	now := int(time.Now().Unix())

	for key, expire := range s.store.Items {
		if expire > now {
			delete(s.store.Items, key)
		}
	}
}

func (s *UniqueSkipper) read() error {
	s.store = SkipStore{
		Items: make(map[string]int),
	}
	fbytes, err := os.ReadFile(s.StorePath)
	if err != nil {
		// file not found
		return nil
	}
	return json.Unmarshal(fbytes, &s.store)
}

func (s *UniqueSkipper) write() error {
	f, err := os.Create(s.StorePath)
	if err != nil {
		return err
	}
	fbytes, err := json.Marshal(s.store)
	if err != nil {
		return err
	}
	_, err = f.Write(fbytes)
	return err
}

func (s *UniqueSkipper) ShouldSkip(record Record) bool {
	hash := sha256.Sum256([]byte(record.Text))
	uniqueKey := hex.EncodeToString(hash[:])

	// already exists
	if _, ok := s.store.Items[uniqueKey]; ok {
		s.log.Info("skip")
		return true
	}
	s.store.Items[uniqueKey] = int(time.Now().Unix())

	if err := s.write(); err != nil {
		s.log.Err(err)
	}
	return false
}

func (s *UniqueSkipper) Close() error {
	return nil
}

type SkipStore struct {
	Items map[string]int `json:"skipped"`
}
type SkipStoreItem struct {
	Key    string `json:"key"`
	Expire int    `json:"expire"` // unix time
}
