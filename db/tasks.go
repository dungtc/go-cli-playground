package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

func Init(dbPath string) (*bolt.DB, error) {
	// Init db
	// It will be created if doesn't exist
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	if err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("task"))
		return err
	}); err != nil {
		return nil, err
	}

	return db, nil
}

type Task struct {
	ID      string
	Message string
}

type TaskRepository struct {
	db *bolt.DB
}

// NewTaskRepository creates task repository instance
func NewTaskRepository(db *bolt.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

// CreateTask creates a new task
func (t *TaskRepository) CreateTask(content string) (task *Task, err error) {
	err = t.db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("task"))

		id, _ := b.NextSequence()
		task = &Task{
			ID:      strconv.Itoa(int(id)),
			Message: content,
		}

		err := b.Put([]byte(task.ID), []byte(task.Message))
		return err
	})
	return task, err
}

// ListTask gets list of tasks
func (t *TaskRepository) ListTask() (tasks []*Task, err error) {
	err = t.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("task"))
		if b == nil {
			return nil
		}
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("%v. %v\n", string(k), string(v))
			tasks = append(tasks, &Task{ID: string(k), Message: string(v)})
		}
		return nil
	})
	return tasks, err
}

// Count gets total key/value pairs
func (t *TaskRepository) Count() (count int, err error) {
	err = t.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("task"))
		if b == nil {
			return nil
		}
		count = b.Stats().KeyN
		return nil
	})
	return count, err
}
