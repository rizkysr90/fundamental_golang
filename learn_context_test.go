package learn_context_test

import (
	"context"
	"log"
	"testing"
)

type contextKey string

func (c contextKey) String() string {
	return "context_key_" + string(c)
}

var learn_key contextKey = "learning"
var learn_key1 contextKey = "learning1"

func GetValContext(c context.Context, key contextKey) (string, bool) {
	res, isExist := c.Value(key).(string)
	return res, isExist
}
func TestContextWithValue(t *testing.T) {
	// create context
	// context is immutabilty
	parentContext := context.Background()

	contextA := context.WithValue(parentContext, learn_key, "xyz")
	contextB := context.WithValue(parentContext, learn_key1, "yaza")
	log.Println(parentContext)
	log.Println(contextA)
	log.Println(contextB)
	// it will be nil value because context is look up from parent context values not from child values
	log.Println(GetValContext(parentContext, learn_key))
	// there will be a value because context is look up from parent context or current context
	log.Println(GetValContext(contextA, learn_key))
	// it will be nil value because context is look up from parent context values not from child values or sibling values
	// and key value is not same as created contextB value
	log.Println(GetValContext(contextB, learn_key))
}
