package tests

import (
	"swagger_to_test/services"
	"testing"
)

func TestAnalysisHtml(t *testing.T) {
	services.AnalysisHtml("http://127.0.0.1:8000/docs/")
}
