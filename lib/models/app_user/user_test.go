package app_user_test

import (
	"testing"

	"github.com/shota-imoto/line-app/lib/models/app_user"
)

func TestBuildUserByIdToken(t *testing.T) {
	userName := "test"
	lineId := "line_id"

	parser := &app_user.ParseStruct{Parser: &app_user.DummyParser{Name: userName, LineId: lineId}}
	user, err := parser.BuildUserByIdToken()

	if err != nil {
		t.Error("raise error: ", err)
	}

	if user.Name != userName {
		t.Error("get wrong user name")
	}

	if user.LineId != lineId {
		t.Error("get wrong user name")
	}
}
