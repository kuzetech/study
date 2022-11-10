package badger

import (
	"errors"
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

func openDB() {
	db, openErr := badger.Open(badger.DefaultOptions("./badger"))
	if openErr != nil {
		log.Fatal(openErr)
	}

	writeErr := db.Update(func(txn *badger.Txn) error {
		err1 := txn.Set([]byte("answer-1"), []byte("42"))
		if err1 != nil {
			return err1
		}
		err2 := txn.Set([]byte("answer-2"), []byte("1111"))
		if err2 != nil {
			return err2
		}
		return nil
	})

	if writeErr != nil {
		log.Fatal(writeErr)
	}

	readErr := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("answer-1"))
		if err != nil {
			return err
		}

		itemErr := item.Value(func(val []byte) error {
			// Accessing val here is valid.
			fmt.Printf("The answer-1 is: %s\n", val)
			return nil
		})

		if itemErr != nil {
			return itemErr
		}

		return nil
	})

	if readErr != nil {
		log.Fatal(readErr)
	}

	scanErr := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("answer-")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	if scanErr != nil {
		log.Fatal(scanErr)
	}

	defer db.Close()
}

func printAllKeys() {
	db, openErr := badger.Open(badger.DefaultOptions("/Users/huangsw/code/funny/turbine/sources/ingest-client-file/badger"))

	if openErr != nil {
		log.Fatal(openErr)
	}

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func keyNotExist() {
	db, _ := badger.Open(badger.DefaultOptions("./badger"))

	getErr := db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte("aaaaa"))
		if err != nil {
			return err
		}
		return nil
	})

	if getErr == nil {
		log.Println("key 存在")
		return
	}

	if errors.As(getErr, &badger.ErrKeyNotFound) {
		log.Println("key 不存在")
		return
	}

	log.Panic("出现意外的错误")
}

func keyExist() {
	db, _ := badger.Open(badger.DefaultOptions("./badger"))

	db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("answer-1"), []byte("42"))
		if err != nil {
			return err
		}
		return nil
	})

	getErr := db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte("answer-1"))
		if err != nil {
			return err
		}
		return nil
	})

	if getErr == nil {
		log.Println("key 存在")
		return
	}

	if errors.As(getErr, &badger.ErrKeyNotFound) {
		log.Println("key 不存在")
		return
	}

	log.Panic("出现意外的错误")
}
