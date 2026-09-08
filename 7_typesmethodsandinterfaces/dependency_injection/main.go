package main

import "fmt"

func LogOutput(message string){
	fmt.Println(message)
}

type SimpleDataStore struct{
	userData map[string]string
}

func (ds SimpleDataStore) UserNameForID(userID string) (string, bool){
	userName, ok := ds.userData[ID]
	return userName
}

func NewSimpleDataSource() SimpleDataStore{
	return SimpleDataStore{
		userData: map[string]string{
			"1" : "Fred",
			"2" : "Mary",
			"3" : "Pat",
		}
	}
}


type DataStore interface{
	UserNameForID(userID string) (string, bool)
}

type Logger interface{
	Log(message string)
}

type LogAdapter func(string) string

func (lg LogAdapter) Log(message string){
	lg(message)
}

type SimpleLogic struct{
	l Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error){
	sl.l.Log("in SayHello for " + userID)
	userName, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return errors.New("unknown user")
	}
	
	return "Hello, " + name, nil
}

func (sl SimpleLogic) sayGoodbye(userID string) (string, error){
	sl.l.Log("in SayGoodbye for UserID " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	
	if !ok {
		return "", errors.New("unknown user")
	}
	
	return "Goodbye, " + name, nil
}

func main(){
	
}