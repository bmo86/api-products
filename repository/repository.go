package repository


type Repo interface{
	Close() error
}

var implementation Repo

//set repo
func SetRepo(repo Repo)  {
	implementation = repo
}

func Close() error {
	return implementation.Close()
}