package analyzer

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUnusedFunctions(t *testing.T) {
	for _, tc := range []struct {
		name          string
		filePath      string
		expectedNames []string
	}{
		{name: "Public functions", filePath: "testdata/main.go", expectedNames: []string{"UnusedPublicFunction"}},
		{name: "Private functions", filePath: "testdata/main.go", expectedNames: []string{"unusedPrivateFunction"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			code, err := os.ReadFile(tc.filePath)
			assert.NoError(t, err)

			unusedFunctions, err := FindUnusedFunctions(string(code))
			unusedFunctionNames := make([]string, len(unusedFunctions))
			for i, fn := range unusedFunctions {
				unusedFunctionNames[i] = fn.Name
			}
			assert.NoError(t, err)
			assert.Len(t, unusedFunctions, len(tc.expectedNames))
			assert.ElementsMatch(t, tc.expectedNames, unusedFunctionNames)
		})
	}
}
