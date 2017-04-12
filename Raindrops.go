package raindrops

import "strconv"

const testVersion = 2

func Convert(num int) string {
	want := ""
	factors := []int{}
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}

	for i := 0; i < len(factors); i++ {
		switch factors[i] {
		case 3:
			want += "Pling"
		case 5:
			want += "Plang"
		case 7:
			want += "Plong"
		}
	}
	if len(want) == 0 {
		return strconv.Itoa(num)
	}
	return want

}
