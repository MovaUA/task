package store

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/movaua/task/pkg/model"
	bolt "go.etcd.io/bbolt"
)

var (
	// bucket is tasks backet name
	bucket = []byte("tasks")
)

// Store interacts with persistent storage
type Store interface {
	Close() error
	Read(id uint64) (*model.Task, error)
	Update(task *model.Task) error
	List() ([]*model.Task, error)
}

// New creates Store
func New(filename string) (Store, error) {
	db, err := bolt.Open(filename, 600, &bolt.Options{
		Timeout: time.Second,
	})
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket(bucket)
		if err != nil {
			return fmt.Errorf("create bucket: %w", err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &storage{db: db}, nil
}

type storage struct {
	db *bolt.DB
}

func (s *storage) Close() error {
	return s.db.Close()
}

func (s *storage) Read(id uint64) (*model.Task, error) {
	var task model.Task

	if err := s.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket(bucket)

		c := b.Cursor()

		id := itob(id)
		k, v := c.Seek(id)
		if bytes.Compare(k, id) != 0 {
			return errors.New("task is not found")
		}

		return proto.Unmarshal(v, &task)
	}); err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *storage) Update(task *model.Task) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		// Retrieve the tasks bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket(bucket)

		// Generate ID for the task.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id, _ := b.NextSequence()
		task.Id = id

		buf, err := proto.Marshal(task)
		if err != nil {
			return err
		}

		// Persist bytes to tasks bucket.
		return b.Put(itob(task.Id), buf)
	})
}

func (s *storage) List() ([]*model.Task, error) {
	var tasks []*model.Task

	if err := s.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket(bucket)

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task model.Task
			if err := proto.Unmarshal(v, &task); err != nil {
				return err
			}
			tasks = append(tasks, &task)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return tasks, nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
