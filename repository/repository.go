package repository


type Repo interface{
	Close() error
}

var implementation Repo

//set repo
func SetRepo(repo Repo)  {
	implementation = repo
}

//close connection db
func Close() error {
	return implementation.Close()
}