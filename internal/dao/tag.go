// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"blog/internal/dao/internal"
)

// internalTagDao is internal type for wrapping internal DAO implements.
type internalTagDao = *internal.TagDao

// tagDao is the data access object for table tag.
// You can define custom methods on it to extend its functionality as you wish.
type tagDao struct {
	internalTagDao
}

var (
	// Tag is globally public accessible object for table tag operations.
	Tag = tagDao{
		internal.NewTagDao(),
	}
)

// Fill with you ideas below.
