package db

type Database struct {
	Val string
}

func New() Database {
	return Database{}
}

func (db *Database) FindMembers() {
	println(db.Val)
}
