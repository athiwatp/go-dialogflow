# Dialogflow Tools

Some util functions that will make your life easier.



## Dependencies

```go
import (
	_struct "github.com/golang/protobuf/ptypes/struct"
	dialogflow "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)
```



## Functions



### Create values

```go
NewTextValue (value) *_struct.Value
```

| Argument | Type   | Description      |
| -------- | ------ | :--------------- |
| value    | string | given text value |

------

```go
NewBoolValue (value) *_struct.Value
```

| Argument | Type | Description      |
| -------- | ---- | :--------------- |
| value    | bool | given bool value |

------

```go
NewNumberValue (value) *_struct.Value
```

| Argument | Type    | Description        |
| -------- | ------- | :----------------- |
| value    | float64 | given number value |



### Create parameters

```go
NewParameters (parameters) *_struct.Struct
```

| Argument   | Type                      | Description                  |
| ---------- | ------------------------- | :--------------------------- |
| parameters | map[string]*_struct.Value | struct containing parameters |



### Find parameter

```go
FindParameter (name, parameters) *_struct.Value
```

| Argument   | Type            | Description                  |
| ---------- | --------------- | :--------------------------- |
| name       | string          | context name                 |
| parameters | *_struct.Struct | struct containing parameters |



### Create context

```go
NewContext (session, name, lifespan, parameters) *dialogflow.Context
```

| Argument   | Type            | Description                  |
| ---------- | --------------- | :--------------------------- |
| session    | string          | session id                   |
| name       | string          | context name                 |
| lifespan   | int32           | context lifetime             |
| parameters | *_struct.Struct | struct containing parameters |



### Find context

```go
FindContext (name, contexts) *dialogflow.Context
```

| Argument | Type                  | Description        |
| -------- | --------------------- | :----------------- |
| name     | string                | context name       |
| contexts | []*dialogflow.Context | context collection |



### Create text fulfillment message

```go
NewTextMessage (messages) *dialogflow.Intent_Message_Text_
```

| Argument | Type     | Description            |
| -------- | -------- | :--------------------- |
| messages | []string | collection of messages |



### Create card button

```go
NewCardButton (text, postback) *dialogflow.Intent_Message_Card_Button
```

| Argument | Type   | Description                                      |
| -------- | ------ | :----------------------------------------------- |
| text     | string | Button text                                      |
| postback | string | A text sent back to Dialogflow or a URL to open. |



### Create card fulfillment message

```go
NewCardMessage (title, subtitle, imageURI, buttons) *dialogflow.Intent_Message_Card_
```

| Argument | Type                                     | Description    |
| -------- | ---------------------------------------- | :------------- |
| title    | string                                   | card title     |
| subtitle | string                                   | card subtitle  |
| imageURI | string                                   | card image URI |
| buttons  | []*dialogflow.Intent_Message_Card_Button | card buttons   |





