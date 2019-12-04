
package service

// struct - interface structure
type db struct{}

// DB is fake database interface.
type DB interface {
	FetchMessage(lang string) (string, error)
	FetchDefaultMessage() (string, error)
}
// struct - interface structure
type greeter struct {
	database DB
	lang     string
}

// GreeterService is service to greet your friends.
type GreeterService interface {
	Greet() string
	GreetInDefaultMsg() string
}

// Method to implement at database
func (d *db) FetchMessage(lang string) (string, error) {
    // in real life, this code will call an external db
    // but for this sample we will just return the hardcoded example value
	if lang == "en" {
		return "hello", nil
	}
	if lang == "es" {
		return "holla", nil
	}
	return "bzzzz", nil
}

// Method to implement at database
func (d *db) FetchDefaultMessage() (string, error) {
	return "default message", nil
}

// Implements FetchMessage
func (g greeter) Greet() string {
	msg, _ := g.database.FetchMessage(g.lang) // call database to get the message based on the lang
	return "Message is: " + msg
}

//Implements FetchDefaultMessageaultMsg
func (g greeter) GreetInDefaultMsg() string {
	msg, _ := g.database.FetchDefaultMessage() // call database to get the default message
	return "Message is: " + msg
}

// Factory for DB
func NewDB() DB {
	return new(db)
}

// Factory for Greeter
func NewGreeter(db DB, lang string) GreeterService {
	return greeter{db, lang}
}