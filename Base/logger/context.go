package logger

import "context"

type ctxMarker struct{}

var (
	ctxMarkerKey = &ctxMarker{}
	NoopTags     = &noopTags{}
)

type noopTags struct{}

func (t *noopTags) Set(key string, value interface{}) Tags {
	return t
}

func (t *noopTags) Has(key string) bool {
	return false
}

func (t *noopTags) Values() map[string]interface{} { return nil }

type Tags interface {
	// Set sets the given key in the metadata tags.
	Set(key string, value interface{}) Tags
	// Has checks if the given key exists.
	Has(key string) bool
	// Values returns a map of key to values.
	// Do not modify the underlying map, please use Set instead.
	Values() map[string]interface{}
}

func Extract(ctx context.Context) Tags {
	t, ok := ctx.Value(ctxMarkerKey).(Tags)
	if !ok {
		return NoopTags
	}
	return t
}
