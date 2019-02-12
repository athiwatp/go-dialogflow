package tools

import (
	"testing"

	_struct "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

const session = "dfv2ef2eccrfbr3bfv1e11ed1evef1f1"

var contexts = []*dialogflow.Context{}
var parameters *_struct.Struct

func TestNewParameters(testing *testing.T) {
	parameters = NewParameters(map[string]*_struct.Value{
		"name": NewTextValue("test"),
	})

	if parameters == nil || parameters.Fields == nil {
		testing.Fatal("invalid parameters created")
	}
}

func TestNewContext(testing *testing.T) {
	ctx := NewContext(session, "testing", 5, parameters)
	if ctx == nil {
		testing.Fatal("invalid context created")
	}

	contexts = append(contexts, ctx)
}

func TestFindContext(testing *testing.T) {
	ctx := FindContext("testing", contexts)
	if ctx == nil {
		testing.Fatal("invalid context found")
	}
}

func TestFindParameter(testing *testing.T) {
	ctx := FindContext("testing", contexts)
	param := FindParameter("name", ctx.Parameters)
	if param == nil {
		testing.Fatal("invalid param found")
	}
}

func TestNewTextMessage(testing *testing.T) {
	msg := NewTextMessage([]string{"Hey"})
	if msg == nil {
		testing.Fatal("invalid text message")
	}
}

func TestNewCardButton(testing *testing.T) {
	cbtn := NewCardButton("Test", "Subtest")
	if cbtn == nil {
		testing.Fatal("invalid card button")
	}
}

func TestNewCardMessage(testing *testing.T) {
	cbtn := NewCardButton("Test", "Subtest")
	msg := NewCardMessage([]*dialogflow.Intent_Message_Card_Button{cbtn}, "Test", "Subtest", "")
	if msg == nil {
		testing.Fatal("invalid card message")
	}
}

func TestNewTextValue(testing *testing.T) {
	txtv := NewTextValue("test")
	if txtv == nil || txtv.GetStringValue() != "test" {
		testing.Fatal("invalid text value")
	}
}

func TestNewNumberValue(testing *testing.T) {
	txtv := NewNumberValue(500)
	if txtv == nil || txtv.GetNumberValue() != 500 {
		testing.Fatal("invalid number value")
	}
}

func TestNewBoolValue(testing *testing.T) {
	txtv := NewBoolValue(true)
	if txtv == nil || txtv.GetBoolValue() != true {
		testing.Fatal("invalid bool value")
	}
}

func TestNewTextMessageCollection(testing *testing.T) {
	msgc := NewTextMessageCollection([]string{"test"})
	if msgc == nil {
		testing.Fatal("invalid text message collection")
	}
}
