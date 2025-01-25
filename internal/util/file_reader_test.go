package util

import (
	"strings"
	"testing"
)

func TestReadJson(t *testing.T) {
	// WHEN unable to open the JSON file (since it doesn't exist)
	data, missingFileErr := ReadJSON[any]("internal/util/test/unknown_file.json")
	// THEN data should be nil (or other default value) and err not nil
	if data != nil && missingFileErr == nil {
		t.Error("ReadJSON unexpectedly succeeded")
	}

	// WHEN JSON file exists but underlying JSON is malformed
	_, badJsonErr := ReadJSON[any]("internal/util/test/bad.json")
	if badJsonErr == nil { // THEN should receive an error back
		t.Error("ReadJSON unexpectedly succeeded in unmarshaling JSON")
	}
	// WHEN a GraphQL query placed in a JSON file is input
	_, graphqlErr := ReadJSON[any]("internal/util/test/graphql_query.json")
	if graphqlErr == nil { // THEN should get an unexpected key char error
		t.Error("ReadJSON unexpectedly failed with graphql formatted JSON")
	}
	// WHEN any other file type is input
	_, fileTypeErr := ReadJSON[any]("internal/util/test/json.go")
	// THEN should get an incorrect file type error
	if fileTypeErr == nil || !strings.HasPrefix(fileTypeErr.Error(), "Incorrect File Type") {
		t.Error("ReadJSON unexpectedly read a non-JSON file")
	}

	// map[string][]map = parent jsonObj with key to an array of jsonObjs
	// WHEN JSON file exists with valid JSON, matching the generic type
	goodFile := "internal/util/test/good.json"
	jsonMap, err := ReadJSON[map[string][]map[string]any](goodFile)
	if err != nil { // THEN no err returned
		t.Error("ReadJSON unexpected failed to return a map")
	}
	jsonArr := jsonMap["objs"] // AND the returned type should be usable
	if jsonArr == nil || jsonArr[0] == nil || jsonArr[0]["foo"] != "bar" {
		t.Errorf("Json Map contains different values than expected = %v", jsonMap)
	}
}

func TestReadFileText(t *testing.T) {
	// WHEN unable to open file (since it doesn't exist)
	data, missingFileErr := ReadFileText("internal/util/test/unknown_file.json")
	// THEN data should be nil (or other default value) and err not nil
	if data != "" && missingFileErr == nil {
		t.Error("ReadFileText unexpectedly succeeded")
	}

	// WHEN file exists but underlying JSON is malformed
	badJsonStr, badJsonErr := ReadFileText("internal/util/test/bad.json")
	if badJsonStr == "" && badJsonErr != nil { // THEN should STILL get text back, no err
		t.Error("ReadFileText unexpected failed to read malformed JSON into string")
	}

	// WHEN a GraphQL query is packed into a json file
	graphqlText, graphqlErr := ReadFileText("internal/util/test/graphql_query.json")
	// THEN should get a perfectly valid query string back w/out error
	if graphqlText == "" && graphqlErr != nil {
		t.Error("ReadFileText unexpectedly failed to read basic GraphQL Query")
	}
}
