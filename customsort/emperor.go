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

const (
	empty      = ""
	intType    = "int"
	stringType = "string"
	floatType  = "float64"
	boolType   = "bool"
	mixedType  = "mixed"
	invalid    = "invalid"
)

/*
CheckMapValuesHomogeneity  method is used to determine the homogeneity of the map values
It checks for int, string, float64, bool primitive values.
Method return dataType of map value and a bool determining whether all values
in map are of same data type.

If map contains values of same data type then it return true else return false
If input map is nil then it returns dataType as "" and true
Incase if map values has data type other than primitive then it returns dataType as
Invalid and false flag

@author adit_chandel
@since 13th July 2020
*/
func CheckMapValuesHomogeneity(m map[interface{}]interface{}) (string, bool) {

	var dataType = empty
	var flag = true

	if nil != m {

		for _, v := range m {

			switch v.(type) {

			case int:
				if dataType == empty {
					dataType = intType

				} else if dataType != intType {
					dataType = mixedType
					flag = false
				}

			case string:
				if dataType == empty {
					dataType = stringType

				} else if dataType != stringType {
					dataType = mixedType
					flag = false
				}

			case float64:
				if dataType == empty {
					dataType = floatType

				} else if dataType != floatType {
					dataType = mixedType
					flag = false
				}

			case bool:
				if dataType == empty {
					dataType = boolType

				} else if dataType != boolType {
					dataType = mixedType
					flag = false
				}

			default:

				dataType = invalid
				flag = false

			}

		}

	}
	return dataType, flag

}

/*
CheckMapKeyHomogeneity  method is used to determine the homogeneity of the map keys
It checks for int, string, float64, bool primitive values.
Method return dataType of map keys and a bool determining whether all keys
in map are of same data type.

If map contains keys of same data type then it return true else return false
If input map is nil then it returns dataType as "" and true

Incase if map keys has data type other than primitive then it returns dataType as
Invalid and false flag

@author adit_chandel
@since 13th July 2020
*/
func CheckMapKeyHomogeneity(m map[interface{}]interface{}) (string, bool) {

	var dataType = empty
	var flag = true

	if nil != m {

		for k := range m {

			switch k.(type) {

			case int:
				if dataType == empty {
					dataType = intType

				} else if dataType != intType {
					dataType = mixedType
					flag = false
				}

			case string:
				if dataType == empty {
					dataType = stringType

				} else if dataType != stringType {
					dataType = mixedType
					flag = false
				}

			case float64:
				if dataType == empty {
					dataType = floatType

				} else if dataType != floatType {
					dataType = mixedType
					flag = false
				}

			case bool:
				if dataType == empty {
					dataType = boolType

				} else if dataType != boolType {
					dataType = mixedType
					flag = false
				}

			default:

				dataType = invalid
				flag = false

			}

		}

	}
	return dataType, flag

}
