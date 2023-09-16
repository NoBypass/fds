package db

func (db *DB[T]) Find(entity *T) (*T, error) {
	return db.common(entity, "MATCH")
}

func (db *DB[T]) Create(entity *T) (*T, error) {
	return db.common(entity, "CREATE")
}
