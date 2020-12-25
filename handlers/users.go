package handlers

import (
	"errors"
	db "github.com/richkule/prepareTestWeb/DBWorker"
	hf "github.com/richkule/prepareTestWeb/helpFun"
	in "github.com/richkule/prepareTestWeb/init"
	"net/http"
)

// Выводит страницу выхода из системы /logout
func Logout(w http.ResponseWriter, req *http.Request, sessUs *in.SessUs) error {
	if sessUs.UsId == in.GuestUserId {
		http.Redirect(w, req, `/`, http.StatusFound)
		return nil
	}
	err := db.UpdateSessActivity(sessUs.SessId)
	if err != nil {
		err = errors.New("Ошибка деактивации сессии logout " + err.Error())
		return err
	}
	_, err = hf.CreateAndSetSess(w, req, nil, in.GuestUserId)
	if err != nil {
		err = errors.New("Ошибка генерации или установки сессии makeHandler " + err.Error())
		return err
	}
	http.Redirect(w, req, `/`, http.StatusFound)
	return nil
}
