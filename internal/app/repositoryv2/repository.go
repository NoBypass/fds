package repositoryv2

type Repository interface {
	insert(ID string, content any) (any, error)
	delete(ID string) (any, error)
	relate() (any, error)
}

func Insert[Record any](r Repository, ID string, content *Record) error {
	// return r.insert()
	return nil
}

func Relate[In, Out, Edge any](r Repository, inID, outID string, options ...RelateOptions[Edge]) error {
	// return r.relate()
	return nil
}

func Delete[Record any](r Repository, ID string) error {
	// return r.delete()
	return nil
}

func Select[Record any](r Repository, options ...SelectOptions) (*Record, error) {
	// return r.select()
	return nil, nil
}

func Update[Record any](r Repository, ID string, record *Record, options ...UpdateOptions[Record]) error {
	// return r.update()
	return nil
}
