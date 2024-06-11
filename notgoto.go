package notgoto

import (
	"go/ast"
	"go/token"
	"regexp"

	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "notgoto",
		Doc:  "check if a file contains goto statement",
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if isGeneratedFile(file, pass.Fset) {
			continue
		}

		ast.Inspect(file, func(n ast.Node) bool {
			if stmt, ok := n.(*ast.BranchStmt); ok && stmt.Tok == token.GOTO {
				pass.Reportf(stmt.Pos(), "goto statement found")
			}
			return true
		})
	}
	return nil, nil
}

var generatedFilePattern = regexp.MustCompile(`^// Code generated .* DO NOT EDIT\.$`)

func isGeneratedFile(file *ast.File, fset *token.FileSet) bool {
	for _, commentGroup := range file.Comments {
		for _, comment := range commentGroup.List {
			if generatedFilePattern.MatchString(comment.Text) {
				return true
			}
		}
	}
	return false
}
