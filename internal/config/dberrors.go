package config

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"

	"gorm.io/gorm"
)

// создаем универсальные ошибки, чтобы работать с ними в сервисе и хэндлерах
var (
	ErrNotFound            = errors.New("не найдено")
	ErrAlreadyExists       = errors.New("уже существует")
	ErrInvalidData         = errors.New("некорректные данные")
	ErrFieldRequired       = errors.New("отсутствует необходимое поле")
	ErrForeignKeyViolation = errors.New("нарушено условие внешнего ключа")
	ErrCheckViolation      = errors.New("нарушено условие поля")
	ErrInternal            = errors.New("ошибка базы данных")
)

// Map функция обрабатывает ошибки и ищет совпадения с ошибками gorm/pgError
func Map(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	if errors.Is(err, gorm.ErrInvalidData) {
		return ErrInvalidData
	}
	if errors.Is(err, gorm.ErrCheckConstraintViolated) {
		return ErrCheckViolation
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return ErrAlreadyExists

		case "23502":
			return ErrFieldRequired

		case "23503":
			return ErrForeignKeyViolation

		case "23514":
			return ErrCheckViolation

		case "22P02":
			return ErrInvalidData
		}
	}

	// если не нашли совпадение ошибки - возвращаем универсальный ответ
	return fmt.Errorf("%w: %v", ErrInternal, err)
}
