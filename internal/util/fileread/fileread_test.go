package fileread

import (
	"strings"
	"testing"
)

func TestJSON(t *testing.T) {
	// WHEN unable to open the JSON file (since it doesn't exist)
	data, missingFileErr := JSON[any]("internal/util/test/unknown_file.json")
	// THEN data should be nil (or other default value) and err not nil
	if data != nil && missingFileErr == nil {
		t.Error("JSON() unexpectedly succeeded")
	}

	// WHEN JSON file exists but underlying JSON is malformed
	_, badJsonErr := JSON[any]("internal/util/test/bad.json")
	if badJsonErr == nil { // THEN should receive an error back
		t.Error("JSON() unexpectedly succeeded in unmarshaling JSON")
	}
	// WHEN a GraphQL query placed in a JSON file is input
	_, graphqlErr := JSON[any]("internal/util/test/graphql_query.json")
	if graphqlErr == nil { // THEN should get an unexpected key char error
		t.Error("JSON() unexpectedly failed with graphql formatted JSON")
	}
	// WHEN any other file type is input
	_, fileTypeErr := JSON[any]("internal/util/test/json.go")
	// THEN should get an incorrect file type error
	if fileTypeErr == nil || !strings.HasPrefix(fileTypeErr.Error(), "Incorrect File Type") {
		t.Error("JSON() unexpectedly read a non-JSON file")
	}

	// map[string][]map = parent jsonObj with key to an array of jsonObjs
	// WHEN JSON file exists with valid JSON, matching the generic type
	goodFile := "internal/util/test/good.json"
	jsonMap, err := JSON[map[string][]map[string]any](goodFile)
	if err != nil { // THEN no err returned
		t.Error("JSON() unexpected failed to return a map")
	}
	jsonArr := jsonMap["objs"] // AND the returned type should be usable
	if jsonArr == nil || jsonArr[0] == nil || jsonArr[0]["foo"] != "bar" {
		t.Errorf("Json Map contains different values than expected = %v", jsonMap)
	}
}

func TestText(t *testing.T) {
	// WHEN unable to open file (since it doesn't exist)
	data, missingFileErr := Text("internal/util/test/unknown_file.json")
	// THEN data should be nil (or other default value) and err not nil
	if data != "" && missingFileErr == nil {
		t.Error("Text() unexpectedly succeeded")
	}

	// WHEN file exists but underlying JSON is malformed
	badJsonStr, badJsonErr := Text("internal/util/test/bad.json")
	if badJsonStr == "" && badJsonErr != nil { // THEN should STILL get text back, no err
		t.Error("Text() unexpected failed to read malformed JSON into string")
	}

	// WHEN a GraphQL query is packed into a json file
	graphqlText, graphqlErr := Text("internal/util/test/graphql_query.json")
	// THEN should get a perfectly valid query string back w/out error
	if graphqlText == "" && graphqlErr != nil {
		t.Error("Text() unexpectedly failed to read basic GraphQL Query")
	}
}
