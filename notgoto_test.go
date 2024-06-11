package notgoto_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/matsuyoshi30/notgoto"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, notgoto.NewAnalyzer(), "a")
}
