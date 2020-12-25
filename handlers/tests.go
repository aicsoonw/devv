package handlers

import (
	"errors"
	db "github.com/richkule/prepareTestWeb/DBWorker"
	in "github.com/richkule/prepareTestWeb/init"
	"net/http"
)

// Обрабатывает страницу create(Создать Тест)
func Create(w http.ResponseWriter, req *http.Request, sessUs *in.SessUs) error {
	_ = req // Переменная необходима для совместительства с типом HandlerIdFunc
	data := in.DataCreateTest{}
	var err error
	if data.Header, err = renderHeader(sessUs.UsId); err != nil {
		err = errors.New("Ошибка обработки шаблона шапки create " + err.Error())
		return err
	}
	data.Tests, err = db.GetTestsById(sessUs.UsId)
	if err != nil {
		err = errors.New("Ошибка получения тестов create " + err.Error())
		return err
	}
	err = renderTemplate(w, in.CreateTestPage, data)
	if err != nil {
		err = errors.New("Ошибка обработки шаблона create " + err.Error())
		return err
	}
	return nil
}

