package rallyapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

func get[T any](ctx context.Context, url string, queryParams url.Values, headers http.Header) (T, error) {
	var m T
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return m, err
	}

	fetch, err := constructFetch(m)
	if err != nil {
		return m, err
	}

	req.Header = headers
	queryParams.Add("fetch", fetch)
	queryParams.Add("pagesize", "100")

	req.URL.RawQuery = queryParams.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return m, err
	}

	if res.StatusCode != 200 {
		return m, fmt.Errorf("Rally returns status code '%d'", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return m, err
	}

	return parseJSON[T](body)
}

func put[T RallyEntity](ctx context.Context, url string, headers http.Header, entity T) error {
	// Wrap the entity in a Request to get proper JSON structure
	request := Request[T]{Entity: entity}
	reqBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	requestBody := bytes.NewBuffer(reqBody)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, requestBody)
	if err != nil {
		return err
	}

	req.Header = headers

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("Status Code Returned: %d", res.StatusCode)
	}

	// Parse the response to check for errors and warnings
	var updateResponse UpdateResponse[T]
	err = json.Unmarshal(body, &updateResponse)
	if err != nil {
		log.Printf("Warning: Failed to parse update response: %v\n", err)
		return nil // Don't fail if we can't parse the response, as the update succeeded
	}

	// Check for errors and print warnings
	return CheckUpdateResponse(updateResponse)
}

func parseJSON[T any](body []byte) (T, error) {
	var res T
	err := json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func constructFetch(fetchType any) (string, error) {
	typedef := reflect.TypeOf(fetchType)

	// Find the innermost type that has rallyapi tags
	targetType := findTypeWithTags(typedef, "rallyapi")
	if targetType == nil {
		return "", fmt.Errorf("no type found with rallyapi tags")
	}

	// Collect all rallyapi tags from the target type
	stringBuilder := strings.Builder{}
	for i := 0; i < targetType.NumField(); i++ {
		fld := targetType.Field(i)
		if fetch := fld.Tag.Get("rallyapi"); fetch != "" {
			stringBuilder.WriteString(fetch)
			stringBuilder.WriteString(",")
		}
	}

	return strings.TrimRight(stringBuilder.String(), ","), nil
}

// findTypeWithTags recursively searches for a struct type that has fields with the specified tag
func findTypeWithTags(t reflect.Type, tagName string) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		return findTypeWithTags(t.Elem(), tagName)
	}

	// If it's not a struct, we can't search it
	if t.Kind() != reflect.Struct {
		return nil
	}

	// Check if this struct has any fields with the tag
	hasTag := false
	for i := 0; i < t.NumField(); i++ {
		fld := t.Field(i)
		if fld.Tag.Get(tagName) != "" {
			hasTag = true
			break
		}
	}

	// If this struct has the tag, return it
	if hasTag {
		return t
	}

	// Otherwise, recursively search all struct fields
	for i := 0; i < t.NumField(); i++ {
		fld := t.Field(i)
		if result := findTypeWithTags(fld.Type, tagName); result != nil {
			return result
		}
	}

	return nil
}
