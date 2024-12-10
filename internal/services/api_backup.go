package services

import (
	"Gym-Service/internal/server/configs"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func CreateBackup(w http.ResponseWriter, r *http.Request) {
	backupFilePath := "./internal/database/data/backup.dump" // Путь для сохранения бэкапа
	dbName := configs.DBName
	username := configs.Username

	// Создание дампа в формате custom, который будет использоваться pg_restore
	cmd := exec.Command("pg_dump", "-U", username, "-F", "c", "-b", "-v", "-f", backupFilePath, dbName)
	cmd.Env = append(cmd.Env, "PGPASSWORD="+configs.Password)

	if err := cmd.Run(); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("Failed to create backup: %v", err), http.StatusInternalServerError)
		return
	}

	// Уведомление об успешном создании бэкапа
	http.Redirect(w, r, "/admin", http.StatusFound)
}

func RestoreBackup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Получаем файл из формы
	file, handler, err := r.FormFile("backup_file")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Сохраняем файл на сервере
	backupFilePath := fmt.Sprintf("./internal/database/data/%s", handler.Filename)
	dst, err := os.Create(backupFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save file: %v", err), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Копируем содержимое файла
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, fmt.Sprintf("Failed to save file: %v", err), http.StatusInternalServerError)
		return
	}

	// Восстанавливаем структуру базы данных (без данных)
	dbName := configs.DBName
	username := configs.Username

	cmd := exec.Command("pg_restore", "-U", username, "-d", dbName, "-v", "-no-data", backupFilePath)
	cmd.Env = append(os.Environ(), "PGPASSWORD="+configs.Password)

	if err := cmd.Run(); err != nil {
		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}

	// Восстанавливаем только данные
	cmd = exec.Command("pg_restore", "-U", username, "-d", dbName, "-v", "--data-only", "--disable-triggers", backupFilePath)
	cmd.Env = append(os.Environ(), "PGPASSWORD="+configs.Password)

	if err := cmd.Run(); err != nil {
		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}

	// Уведомляем об успешном восстановлении
	http.Redirect(w, r, "/admin", http.StatusFound)
}
