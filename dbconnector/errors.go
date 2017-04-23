package dbconnector

type connectionError struct{}

func (err *connectionError) Error() string {
	return "There was a problem while connecting to the database"
}

type connectionDriverError struct{}

func (err *connectionDriverError) Error() string {
	return "There was a problem while validating the database driver"
}
