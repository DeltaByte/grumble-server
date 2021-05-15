package channel

import bolt "go.etcd.io/bbolt"

func GetAll(db *bolt.DB) ([]Channel, error) {
	var channels []Channel

	err := db.View(func(tx *bolt.Tx) (error) {
		dbb := tx.Bucket([]byte(DBBucket))

		// iterate over all channels, decode, and add to result
		dbb.ForEach(func(k, v []byte) error {
			decoded, err := Decode(v)
			if (err != nil) { return err }
			channels = append(channels, decoded)
			return nil
		})

		return nil
	})

	return channels, err
}