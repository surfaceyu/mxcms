package utils

func SliceRemoce(origin []interface{}) {
	for i := 0; i < len(origin); i++ {
		if origin[i] == 6 {
			origin = append(origin[:i], origin[i+1:]...)
			i-- // maintain the correct index
		}
	}
}
