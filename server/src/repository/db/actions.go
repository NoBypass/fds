package db

// Find will find a node in the database by the values
// of the struct passed in. It will return the node
// found in the database, or an error if one occurred.
func (db *DB[T]) Find(entity *T) (*T, error) {
	return db.common(entity, "MATCH")
}

// Create will create a node in the database by the struct
// passed in. It will return the node created in the database,
// or an error if one occurred.
func (db *DB[T]) Create(entity *T) (*T, error) {
	return db.common(entity, "CREATE")
}
