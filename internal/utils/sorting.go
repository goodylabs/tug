package utils

import (
	"sort"

	"github.com/goodylabs/tug/internal/constants"
)

func SortOptions(options []string) {
	sort.Slice(options, func(i, j int) bool {
		if options[i] == constants.ALL_OPTION {
			return false //
		}
		if options[j] == constants.ALL_OPTION {
			return true
		}
		return options[i] < options[j]
	})
}
