package http

const (
	dataFileTestPath = "testdata"
)

var (
	expectedResponse string
)

func init() {
	expectedResponse = readFileFromPathToString(dataFileTestPath, "get_expected_response.json")
}
