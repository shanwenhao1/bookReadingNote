//go:generate mockgen -destination mock_matter/mock_matter.go bookReadingNote/project/mock/mockMatter IF
package mockMatter

type IF interface {
	SimpleMethod(i int)
	VariadicMethod(i ...int)
}
