package handlers

import (
	"bytes"
	"errors"
	"fmt"
	db "github.com/richkule/prepareTestWeb/DBWorker"
	hf "github.com/richkule/prepareTestWeb/helpFun"
	in "github.com/richkule/prepareTestWeb/init"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
)

// Обрабатывает ошибку во время исполнения
func wrongFun(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	http.Error(w, "", http.StatusBadGateway)
}
