package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")

var db *bolt.DB

type Task struct {
	Key int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	} 

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})

}

func CreateTask(task string)(int, error){
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		// Access database bucket
		b := tx.Bucket(taskBucket)

		// Get next id
		id64, _ := b.NextSequence()

		// Store id
		id = int(id64)

		// Convert id to binary
		key := itob(int(id))

		// Update database bucket
		return b.Put(key, []byte(task))
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

func GetTasks()([]Task, error){
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error{
		b := tx.Bucket(taskBucket)

		// Get a cursor pointing to keys
		c := b.Cursor()

		// Read and store all key value pairs
		for k, v := c.First(); k != nil; k, v = c.Next(){
			tasks = append(tasks, Task{
				Key: btoi(k),
				Value: string(v),
			})
		}

		return nil
	})

	// Handle error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte{
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}