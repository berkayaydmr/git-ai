package errors

import "fmt"

var (
	ErrRepositoryNotFound       = fmt.Errorf("selected github repository not found")
	ErrRepositoryExists         = fmt.Errorf("repository already exists")
	ErrRepositoryUrlNotProvided = fmt.Errorf("repository url not provided")
	ErrUnexpectedStatusCode     = fmt.Errorf("unexpected status code")
	ErrBranchNotFound           = fmt.Errorf("branch not found")
	ErrRepositoryDirNotProvided = fmt.Errorf("repository directory not provided")

	ErrNoneOfApiKeysFound = fmt.Errorf("none of the api keys found")
	ErrApiKeyNotFound     = fmt.Errorf("api key not found")

	ErrInvalidSelection = fmt.Errorf("invalid selection")

	ErrNameNotProvided  = fmt.Errorf("name not provided")
	ErrKeyNotProvided   = fmt.Errorf("key not provided")
	ErrApiKeyNameExists = fmt.Errorf("api key name already exists")

	ErrBranchesShouldBeProvided = fmt.Errorf("branches should be provided")
	ErrFirstBranchNotProvided   = fmt.Errorf("first branch not provided")
	ErrSecondBranchNotProvided  = fmt.Errorf("second branch not provided")
	ErrFirstBranchDoesNotExist  = fmt.Errorf("first branch does not exist on the repository")
	ErrSecondBranchDoesNotExist = fmt.Errorf("second branch does not exist on the repository")
)
