package line_service

import (
	"github.com/shota-imoto/line-app/lib/db"
	"github.com/shota-imoto/line-app/lib/models/app_user"
)

func FindUserByIdToken(parser *app_user.ParseStruct) (app_user.User, error) {
	user := app_user.User{}
	claims, err := parser.GetJwtClaims()

	if err != nil {
		return user, nil
	}

	db.Db.Where("line_id = ?", claims["sub"].(string)).First(&user)

	return user, nil
}

// ex
// parser := &app_user.ParseStruct{Parser: &app_user.BasicParser{IdToken: tokenResponse.IdToken}}
// user, err := line_service.FindOrCreateUserByIdToken("test", parser)

func FindOrCreateUserByIdToken(parser *app_user.ParseStruct) (app_user.User, error) {
	user, err := FindUserByIdToken(parser)

	if err != nil {
		return app_user.User{}, err
	}

	if user.Id == 0 {

		user, err := parser.BuildUserByIdToken()

		if err != nil {
			return app_user.User{}, err
		}

		db.Db.Create(&user)
	}

	return user, nil
}
