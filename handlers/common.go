
// Возвращает страницу пользователю, обрабатывая заданный шаблон
func renderTemplate(w io.Writer, pagePath string, data interface{}) error {
	pageName := filepath.Base(pagePath)
	err := in.Templates.ExecuteTemplate(w, pageName, data)
	if err != nil {
		return err
	}
	return nil
}
