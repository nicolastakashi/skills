```go
func TestParseWebhookConfig(t *testing.T) {
	raw := `{"name":"alerts","url":"https://hooks.example.com/abc","headers":{"X-Token":"t1"},"retries":3}`
	cfg, err := ParseWebhookConfig(raw)
	require.NoError(t, err)
	assert.NotNil(t, cfg)
}
```
