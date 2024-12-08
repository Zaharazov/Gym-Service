package services

import (
	"net/http"
	"path/filepath"
)

func DataPage(w http.ResponseWriter, r *http.Request) {
	// Путь к папке со статическими файлами
	staticDir := "./frontend"

	// Получение пути к запрашиваемому файлу
	requestedFile := r.URL.Path

	// Удаляем префикс /frontend/ из пути
	relativePath := requestedFile[len("/frontend/"):]

	// Генерируем полный путь к файлу
	fullPath := filepath.Join(staticDir, relativePath)

	// Отдаём файл
	http.ServeFile(w, r, fullPath)
}
