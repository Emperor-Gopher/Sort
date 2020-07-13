package customsort

import (
	"sort"
	"time"
)

/*
SortIntSlice method is used to sort int slice
If nil Slice is provided, method returns nil
@author adit_chandel
@since 13th July 2020
*/
func SortIntSlice(slc []int) {

	if nil != slc {
		sort.Ints(slc)
	}
}

/*
SortFloat64Slice method is used to sort float slice
If nil Slice is provided, method returns nil
@author adit_chandel
@since 13th July 2020
*/
func SortFloat64Slice(slc []float64) {

	if nil != slc {
		sort.Float64s(slc)
	}
}

/*
SortStringSlice method is used to sort float slice
If nil Slice is provided, method returns nil
@author adit_chandel
@since 13th July 2020
*/
func SortStringSlice(slc []string) {

	if nil != slc {
		sort.Strings(slc)
	}
}

/*
SortStringSliceByLength method is used to sort slice on the basis of its element length
If nil Slice is provided, method returns nil
@author adit_chandel
@since 13th July 2020
*/
func SortStringSliceByLength(slc []string) {

	if nil != slc {
		sort.SliceStable(slc, func(i, j int) bool {
			if len(slc[i]) < len(slc[j]) {
				return true
			}
			return false
		})
	}
}

/*
ReverseSortStringSliceByLength  method is used to sort slice on the basis of its element length
If nil Slice is provided, method returns nil
@author adit_chandel
@since 13th July 2020
*/
func ReverseSortStringSliceByLength(slc []string) {

	if nil != slc {
		sort.SliceStable(slc, func(i, j int) bool {
			if len(slc[i]) > len(slc[j]) {
				return true
			}
			return false
		})
	}
}

/*
SortTimeSliceChronologically  method is used to sort slice time elements in the chronological order
If nil Slice is provided, method returns nil
@author adit_chandel
@since 13th July 2020
*/
func SortTimeSliceChronologically(slc []time.Time) {

	if nil != slc {
		sort.SliceStable(slc, func(i, j int) bool {
			if slc[i].Before(slc[j]) {
				return true
			}
			return false
		})
	}
}

/*
ReverseSortTimeSliceChronologically  method is used to sort slice time elements in the time reverse chronological order
If nil Slice is provided, method returns nil
@author adit_chandel
@since 13th July 2020
*/
func ReverseSortTimeSliceChronologically(slc []time.Time) {

	if nil != slc {
		sort.SliceStable(slc, func(i, j int) bool {
			if slc[j].Before(slc[i]) {
				return true
			}
			return false
		})
	}
}
