package design

import "errors"

// File object
type File struct {
	name        string
	size        int
	category    int
	isDirectory bool
	children    []File
}

type Filter interface {
	apply(file File) bool
}

// TypeFilter object
type TypeFilter struct {
	category int
}

func NewTypeFilter(category int) TypeFilter {
	return TypeFilter{
		category: category,
	}
}

// 实现Filter接口的是 TypeFilter的指针类型
func (mf *TypeFilter) apply(file File) bool {
	return file.category == mf.category
}

// MinSizeFilter object
type MinSizeFilter struct {
	minSize int
}

func NewMinSizeFilter(minSize int) MinSizeFilter {
	return MinSizeFilter{
		minSize: minSize,
	}
}

// 实现Filter接口的是 MinSizeFilter的指针类型
func (ms *MinSizeFilter) apply(file File) bool {
	return file.size > ms.minSize
}

func FindWithFilters(directory File, filters []Filter) ([]File, error) {
	if !directory.isDirectory {
		return nil, errors.New("not a directory")
	}

	output := make([]File, 0)
	helper(directory, filters, &output)
	return output, nil
}

func helper(directory File, filters []Filter, output *[]File) {
	if directory.children == nil {
		return
	}
	for _, file := range directory.children {
		if file.isDirectory {
			helper(file, filters, output)
		} else {
			selectFile := true
			for _, filter := range filters {
				if !filter.apply(file) {
					selectFile = false
				}
			}
			if selectFile {
				*output = append(*output, file)
			}
		}
	}
}
