package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPagination_Valid(t *testing.T) {
	p := new(Pagination)
	assert.Equal(t, nil, p.Valid())
}

func TestPagination_GetTerm(t *testing.T) {
	p := &Pagination{
		Page:     1,
		PageSize: 10,
	}

	offset, limit := p.GetTerm()
	assert.Equal(t, 0, offset)
	assert.Equal(t, 10, limit)
}

func TestPagination_SelectAll(t *testing.T) {
	p := &Pagination{
		PageSize: -1,
	}

	assert.Equal(t, true, p.SelectAll())
}
