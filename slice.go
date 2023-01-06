package tableprinter

import (
	"reflect"
)

type sliceParser struct {
	TagsOnly bool
}

var emptyStruct = struct{}{}

func (p *sliceParser) Parse(v reflect.Value, filters []RowFilter) ([]string, [][]string, []int) {
	headers, mapKeys := p.ParseHeaders(v)
	rows, nums := p.ParseRows(v, mapKeys, filters)
	return headers, rows, nums
}

func (p *sliceParser) ParseRows(v reflect.Value, mapKeys []reflect.Value, filters []RowFilter) (rows [][]string, nums []int) {
	for i, n := 0, v.Len(); i < n; i++ {
		item := indirectValue(v.Index(i))
		if !CanAcceptRow(item, filters) {
			continue
		}

		if item.Kind() == reflect.Struct {
			r, c := getRowFromStruct(item, p.TagsOnly)

			nums = append(nums, c...)

			rows = append(rows, r)
		} else if item.Kind() == reflect.Map {
			p := WhichParser(item.Type()).(*mapParser)
			rs, cs := p.ParseRows(item, mapKeys, filters)
			rows = append(rows, rs...)
			nums = append(nums, cs...)
		} else {
			// if not struct, don't search its fields, just put a row as it's.
			c, r := extractCells(i, emptyHeader, indirectValue(item), p.TagsOnly)
			rows = append(rows, r)
			nums = append(nums, c...)
		}
	}

	return
}

func (p *sliceParser) ParseHeaders(v reflect.Value) (headers []string, mapKeys []reflect.Value) {
	tmp := make(map[reflect.Type]struct{})

	for i, n := 0, v.Len(); i < n; i++ {
		item := indirectValue(v.Index(i))

		// no filters.
		itemTyp := item.Type()
		if _, ok := tmp[itemTyp]; !ok {
			// make headers once per type.
			tmp[itemTyp] = emptyStruct
			if itemTyp.Kind() == reflect.Struct {
				hs := extractHeadersFromStruct(itemTyp, p.TagsOnly)
				if len(hs) == 0 {
					continue
				}
				for _, h := range hs {
					headers = append(headers, h.Name)
				}
			} else if itemTyp.Kind() == reflect.Map {
				p := WhichParser(itemTyp).(*mapParser)
				mapKeys = p.Keys(item)
				hs := p.ParseHeaders(item, mapKeys)
				headers = append(headers, hs...)
			}
		}
	}

	return
}
