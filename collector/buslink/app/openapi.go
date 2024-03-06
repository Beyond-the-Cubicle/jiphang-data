package app

type DocType string

const (
	Json = DocType("json")
	Xml  = DocType("xml")
)

type OpenAPIFailResponse struct {
	Result OpenAPIResultCode
}

type OpenAPIError struct {
	Url    string
	Result OpenAPIResultCode
}

type OpenAPIResultCode struct {
	Code    string
	Message string
}
