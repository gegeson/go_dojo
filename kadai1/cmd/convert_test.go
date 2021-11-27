package convert_test

import (
	"convert"
	"strconv"
	"testing"
)


func TestConvert(t *testing.T) {
	testArgs := []struct {
		Path  string
		From string
		To   string
	} {
		// {Path: "../sample/test1.png", From: ".jpg", To: ".png"},
		{Path: "../sample/sample2/test3.png", From: ".png", To: ".jpg"},
	}

	for i, v := range testArgs {
		t.Run("test"+strconv.Itoa(i), func(t *testing.T) {
			HelperConvert(v.Path, v.From, v.To, t)
		})
		
	}

}

func HelperConvert(Path string, From string, To string, t *testing.T) {
	t.Helper()
	if err := convert.Convert(Path, From, To); err != nil {
		t.Errorf("failed test %#v", err)
	}
}