# 15 - JSON Basics

## Goal

Decode JSON into a typed Go struct and reject malformed or invalid input.

## Concepts

- Struct tags (`json:"title"`)
- `encoding/json` decoding with `json.Unmarshal`
- JSON-to-struct field mapping
- Decode errors as validation
- Validating content after decoding

## Syntax Primer

Struct tags map JSON keys to Go fields:

```go
type Task struct {
	Title string `json:"title"`
}
```

`json.Unmarshal` parses bytes into a pointer to the struct and returns an error for malformed JSON:

```go
var task Task
if err := json.Unmarshal(data, &task); err != nil {
	return Task{}, err
}
```

## Mental Model

JSON decoding is validation at the boundary between text and code: malformed text must never reach business logic. Decoding fills typed fields, and then you still validate the content — a well-formed payload like `{"title":"   "}` decodes fine but still holds a blank title.

## Annotated Examples

```go
type Config struct {
	Port int `json:"port"`
}

func loadConfig(data []byte) (Config, error) {
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
```

## Common Diagnostics

- `invalid character ... looking for beginning of value`: the input is not valid JSON; `Unmarshal` reports it instead of guessing.
- Empty or zero fields after decode: the JSON key does not match a struct tag (or the tag is missing). Check `json:"title"` spelling.
- `undefined: json`: add the `encoding/json` import when the starter does not provide it.
- `json: cannot unmarshal string into Go struct field`: the JSON value's type does not match the Go field's type.

## Exercise

Implement `DecodeTask` so it parses one task payload and validates the result.

## Acceptance Criteria

- `{"title":"ship API"}` decodes to a task with `Title` `"ship API"`.
- A blank title like `"   "` returns an error.
- Malformed JSON returns an error.

## Hints

- Unmarshal into the existing `Task` type and return `Task{}, err` on failure.
- After decoding, trim the title and reject it when empty.
- Return the task and `nil` only at the end.

## Verify

```bash
gofmt -w exercises/00-syntax-drills/15-json-basics
go test -tags exercise ./exercises/00-syntax-drills/15-json-basics/...
```

## Reflection Prompts

Why should decoding and validation happen before business logic? What is the difference between malformed JSON and valid JSON containing an empty title?