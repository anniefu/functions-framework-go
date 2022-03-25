module example.com/conformance

go 1.13

replace example.com/function => ./function

require (
	example.com/function v0.0.0-00010101000000-000000000000
	github.com/GoogleCloudPlatform/functions-framework-go v1.5.3
)
