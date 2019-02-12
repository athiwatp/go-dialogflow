package tools

import (
	"net/http"
	"strings"

	"github.com/golang/protobuf/jsonpb"

	_struct "github.com/golang/protobuf/ptypes/struct"
	dialogflow "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// NewTextMessageCollection Creates new message collection
func NewTextMessageCollection(messages []string) []*dialogflow.Intent_Message {
	msgs := []*dialogflow.Intent_Message{}

	for i := 0; i < len(messages); i++ {
		msgs = append(msgs, &dialogflow.Intent_Message{
			Message: NewTextMessage([]string{
				messages[i],
			}),
		})
	}

	return msgs
}

// FindContext Finds context on given collection
func FindContext(name string, contexts []*dialogflow.Context) *dialogflow.Context {
	var result *dialogflow.Context
	for _, context := range contexts {
		r := strings.Split(context.Name, "/")
		foundName := r[len(r)-1]
		if foundName == name {
			result = context
			break
		}
	}
	return result
}

// FindParameter Finds parameter by name on given parameter struct
func FindParameter(name string, parameters *_struct.Struct) *_struct.Value {
	var result *_struct.Value
	pfields := parameters.GetFields()
	if val, ok := pfields[name]; ok {
		result = val
	}
	return result
}

// NewTextMessage Creates new simple text message
func NewTextMessage(messages []string) *dialogflow.Intent_Message_Text_ {
	return &dialogflow.Intent_Message_Text_{
		Text: &dialogflow.Intent_Message_Text{
			Text: messages,
		},
	}
}

// NewCardButton Creates new card button
func NewCardButton(text string, postback string) *dialogflow.Intent_Message_Card_Button {
	return &dialogflow.Intent_Message_Card_Button{Text: text, Postback: postback}
}

// NewCardMessage Creates card message
func NewCardMessage(buttons []*dialogflow.Intent_Message_Card_Button, title, subtitle, imageURI string) *dialogflow.Intent_Message_Card_ {
	return &dialogflow.Intent_Message_Card_{
		Card: &dialogflow.Intent_Message_Card{
			Title:    title,
			Subtitle: subtitle,
			ImageUri: imageURI,
			Buttons:  buttons,
		},
	}
}

// NewParameters Creates parameters collection
func NewParameters(keyv map[string]*_struct.Value) *_struct.Struct {
	parameters := &_struct.Struct{
		Fields: keyv,
	}
	return parameters
}

// NewTextValue Creates new text value
func NewTextValue(value string) *_struct.Value {
	val := &_struct.Value{
		Kind: &_struct.Value_StringValue{
			StringValue: value,
		},
	}
	return val
}

// NewBoolValue Creates new bool value
func NewBoolValue(value bool) *_struct.Value {
	val := &_struct.Value{
		Kind: &_struct.Value_BoolValue{
			BoolValue: value,
		},
	}
	return val
}

// NewNumberValue Creates new number value
func NewNumberValue(value float64) *_struct.Value {
	val := &_struct.Value{
		Kind: &_struct.Value_NumberValue{
			NumberValue: value,
		},
	}
	return val
}

// NewContext Creates new context
func NewContext(session string, name string, lifespan int32, parameters *_struct.Struct) *dialogflow.Context {
	path := "/contexts/"
	fullname := session + path + name
	context := &dialogflow.Context{
		Name:          fullname,
		LifespanCount: lifespan,
		Parameters:    parameters,
	}
	return context
}

// RespondJSON Sends json response back to dialogflow
func RespondJSON(w http.ResponseWriter, r *dialogflow.WebhookResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	Respond(w, r)
}

// Respond Sends response back to dialogflow
func Respond(w http.ResponseWriter, response *dialogflow.WebhookResponse) {
	jmarshaller := jsonpb.Marshaler{}
	jmarshaller.Marshal(w, response)
}
