package dialogflow

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

const (
	mockPath = "../mock/v2"
)

func TestSigninEvent(t *testing.T) {
	var req *Request
	filePath := mockPath + "/SignInEvent.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Error(err)
	}
	if err = json.Unmarshal(data, &req); err != nil {
		t.Error(err)
	}

	want := map[string]interface{}{
		"intent":      "actions.intent.SIGN_IN",
		"accessToken": "Your OAuth service access token",
	}

	got := map[string]interface{}{
		"intent":      req.OriginalDetectIntentRequest.Payload.Inputs[0].Intent,
		"accessToken": req.OriginalDetectIntentRequest.Payload.User.AccessToken,
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestWelcomeEvent(t *testing.T) {
	var req *Request
	filePath := mockPath + "/WelcomeEvent.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Error(err)
	}
	if err = json.Unmarshal(data, &req); err != nil {
		t.Error(err)
	}

	want := map[string]interface{}{
		"action": "input.welcome",
		"intent": "actions.intent.MAIN",
		"permissions": []string{
			"UPDATE",
		},
	}

	got := map[string]interface{}{
		"action":      req.QueryResult.Action,
		"intent":      req.OriginalDetectIntentRequest.Payload.Inputs[0].Intent,
		"permissions": req.OriginalDetectIntentRequest.Payload.User.Permissions,
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
func TestPermissionEvent(t *testing.T) {
	var req *Request
	filePath := mockPath + "/PermissionEvent.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Error(err)
	}
	if err = json.Unmarshal(data, &req); err != nil {
		t.Error(err)
	}

	want := map[string]interface{}{
		"textValue": "true",
		"name":      "PERMISSION",
		"boolValue": true,
		"intent":    "actions.intent.PERMISSION",
		"lastSeen":  "2018-02-26T01:38:19Z",
		"permissions": []string{
			"NAME",
			"DEVICE_PRECISE_LOCATION",
		},
		"displayName": "Matt Carroll",
		"givenName":   "Matt",
		"familyName":  "Carroll",
		"locale":      "en-US",
		"userId":      "ABwppHEvwoXs18xBNzumk18p5h02bhRDp_riW0kTZKYdxB6-LfP3BJRjgPjHf1xqy1lxqS2uL8Z36gT6JLXSrSCZ",
		"latitude":    37.4219806,
		"longitude":   -122.0841979,
	}

	input := req.OriginalDetectIntentRequest.Payload.Inputs[0]
	arg := input.Arguments[0]
	user := req.OriginalDetectIntentRequest.Payload.User
	profile := user.Profile
	location := req.OriginalDetectIntentRequest.Payload.Device.Location

	got := map[string]interface{}{
		"textValue":   arg.TextValue,
		"name":        arg.Name,
		"boolValue":   arg.BoolValue,
		"intent":      input.Intent,
		"lastSeen":    user.LastSeen,
		"permissions": user.Permissions,
		"displayName": profile.DisplayName,
		"givenName":   profile.GivenName,
		"familyName":  profile.FamilyName,
		"locale":      user.Locale,
		"userId":      user.UserID,
		"latitude":    location.Coordinates.Latitude,
		"longitude":   location.Coordinates.Longitude,
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}

}

func TestNoInputEvent(t *testing.T) {
	var req *Request
	filePath := mockPath + "/NoInputEvent.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Error(err)
	}
	if err = json.Unmarshal(data, &req); err != nil {
		t.Error(err)
	}

	want := "actions.intent.NO_INPUT"
	got := req.OriginalDetectIntentRequest.Payload.Inputs[0].Intent
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestListSelectionEvent(t *testing.T) {
	var req *Request
	filePath := mockPath + "/ListSelectionEvent.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Error(err)
	}
	if err = json.Unmarshal(data, &req); err != nil {
		t.Error(err)
	}

	want := "Key of selected item"
	got := req.OriginalDetectIntentRequest.Payload.Inputs[0].Arguments[0].TextValue
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}

	want = "OPTION"
	got = req.OriginalDetectIntentRequest.Payload.Inputs[0].Arguments[0].Name
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestEndConversation(t *testing.T) {
	var req *Request
	filePath := mockPath + "/EndConversationEvent.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Error(err)
	}
	if err = json.Unmarshal(data, &req); err != nil {
		t.Error(err)
	}

	want := "actions.intent.CANCEL"
	got := req.OriginalDetectIntentRequest.Payload.Inputs[0].Intent
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestConfirmatonEvent(t *testing.T) {
	var req *Request
	filePath := mockPath + "/ConfirmationEvent.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Error(err)
	}
	if err = json.Unmarshal(data, &req); err != nil {
		t.Error(err)
	}

	input := req.OriginalDetectIntentRequest.Payload.Inputs[0]
	arg := input.Arguments[0]
	want := "CONFIRMATION"
	if arg.Name != want {
		t.Errorf("want: %v, got: %v", want, arg.Name)
	}

	if !arg.BoolValue {
		t.Error("boolValue should be true")
	}
}

func TestDatetimeEvent(t *testing.T) {
	var req *Request
	filePath := mockPath + "/DatetimeEvent.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Error(err)
	}
	if err = json.Unmarshal(data, &req); err != nil {
		t.Error(err)
	}

	input := req.OriginalDetectIntentRequest.Payload.Inputs[0]
	arg := input.Arguments[0]
	dt := arg.DatetimeValue

	want := map[string]int{
		"month": 4,
		"year":  2018,
		"day":   19,
		"hours": 15,
	}

	got := map[string]int{
		"month": dt.Date.Month,
		"year":  dt.Date.Year,
		"day":   dt.Date.Day,
		"hours": dt.Time.Hours,
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestBase(t *testing.T) {

	var req *Request

	filePath := mockPath + "/request.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(data, &req)
	if err != nil {
		t.Error(err)
	}

	want := map[string]interface{}{
		"responseId": "68efa569-4ba1-4b7f-9b1b-ac2865deb539",
		"queryText":  "query from the user",
		"action":     "action.name.of.matched.dialogflow.intent",
		"outputContexts": []string{
			"projects/integrationfulfillmenttest/agent/sessions/1522951193000/contexts/actions_capability_screen_output",
			"projects/integrationfulfillmenttest/agent/sessions/1522951193000/contexts/actions_capability_audio_output",
			"projects/integrationfulfillmenttest/agent/sessions/1522951193000/contexts/google_assistant_input_type_keyboard",
			"projects/integrationfulfillmenttest/agent/sessions/1522951193000/contexts/actions_capability_media_response_audio",
			"projects/integrationfulfillmenttest/agent/sessions/1522951193000/contexts/actions_capability_web_browser",
		},
		"intentname":                "projects/integrationfulfillmenttest/agent/intents/1f4e5bd9-a670-4161-a22e-2c97b077f29f",
		"intentDisplayName":         "Name of Dialogflow Intent",
		"intentDetectionConfidence": 1,
		"languageCode":              "en-us",
		"source":                    "google",
		"version":                   "2",
		"isInSandbox":               true,
		"capabilities": []string{
			"actions.capability.SCREEN_OUTPUT",
			"actions.capability.AUDIO_OUTPUT",
			"actions.capability.WEB_BROWSER",
			"actions.capability.MEDIA_RESPONSE_AUDIO",
		},
		"query":             "query from the user",
		"inputType":         "KEYBOARD",
		"rawText":           "query from the user",
		"textValue":         "query from the user",
		"argumentName":      "text",
		"intent":            "actions.intent.TEXT",
		"lastSeen":          "2017-10-06T01:06:56Z",
		"locale":            "en-US",
		"userId":            "AI_yXq-AtrRh3mJX5D-G0MsVhqun",
		"conversationId":    "1522951193000",
		"conversationType":  "ACTIVE",
		"conversationToken": "[]",
		"surfaceCapabilities": []string{
			"actions.capability.SCREEN_OUTPUT",
			"actions.capability.AUDIO_OUTPUT",
		},
		"session": "projects/integrationfulfillmenttest/agent/sessions/1522951193000",
	}

	outputContexts := []string{}
	for _, k := range req.QueryResult.OutputContexts {
		outputContexts = append(outputContexts, k.Name)
	}

	payload := req.OriginalDetectIntentRequest.Payload
	capabilities := []string{}
	for _, k := range payload.Surface.Capabilities {
		capabilities = append(capabilities, k.Name)
	}

	input := payload.Inputs[0]
	arg := input.Arguments[0]
	user := payload.User
	convo := payload.Conversation

	availableCapabilities := []string{}
	for _, k := range payload.AvailableSurfaces[0].Capabilities {
		availableCapabilities = append(availableCapabilities, k.Name)
	}

	got := map[string]interface{}{
		"responseId":                req.ResponseID,
		"queryText":                 req.QueryResult.QueryText,
		"action":                    req.QueryResult.Action,
		"outputContexts":            outputContexts,
		"intentname":                req.QueryResult.Intent.Name,
		"intentDisplayName":         req.QueryResult.Intent.DisplayName,
		"intentDetectionConfidence": req.QueryResult.IntentDetectionConfidence,
		"languageCode":              req.QueryResult.LanguageCode,
		"source":                    req.OriginalDetectIntentRequest.Source,
		"version":                   req.OriginalDetectIntentRequest.Version,
		"isInSandbox":               payload.IsInSandbox,
		"capabilities":              capabilities,
		"query":                     input.RawInputs[0].Query,
		"inputType":                 input.RawInputs[0].InputType,
		"rawText":                   arg.RawText,
		"textValue":                 arg.TextValue,
		"argumentName":              arg.Name,
		"intent":                    input.Intent,
		"lastSeen":                  user.LastSeen,
		"locale":                    user.Locale,
		"userId":                    user.UserID,
		"conversationId":            convo.ConversationID,
		"conversationType":          convo.Type,
		"conversationToken":         convo.ConversationToken,
		"surfaceCapabilities":       availableCapabilities,
		"session":                   req.Session,
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
