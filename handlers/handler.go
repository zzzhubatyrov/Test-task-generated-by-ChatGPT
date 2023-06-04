package handlers

import (
	"MiddleTestTask/config"
	"MiddleTestTask/models"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetTasks(c *fiber.Ctx) error {
	db, err := config.Connection()
	if err != nil {
		return err
	}
	var data []models.Tasks
	if err := db.Find(&data).Error; err != nil {
		fmt.Println(err)
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

func CreateTasks(c *fiber.Ctx) error {
	db, err := config.Connection()
	if err != nil {
		return err
	}
	var data models.Tasks
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	createTasks := models.Tasks{
		Title:       data.Title,
		Description: data.Description,
	}
	if err := db.Create(&createTasks).Error; err != nil {
		fmt.Println(err)
	}
	return c.Status(fiber.StatusCreated).JSON(createTasks)
}

func GetTasksId(c *fiber.Ctx) error {
	db, err := config.Connection()
	if err != nil {
		return err
	}
	var data models.Tasks
	id := c.Params("id")
	if err := db.Where("id = ?", id).Find(&data, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": "Запись не найдена",
			})
		}
		return err
	}
	if data.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "Запись не найдена",
		})
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

func EditTasks(c *fiber.Ctx) error {
	db, err := config.Connection()
	if err != nil {
		return err
	}
	var data models.Tasks
	id := c.Params("id")
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	editTasks := models.Tasks{
		Title:       data.Title,
		Description: data.Description,
	}
	if err := db.Find(&data, id).Updates(&editTasks).Error; err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

func DeleteTasks(c *fiber.Ctx) error {
	db, err := config.Connection()
	if err != nil {
		return err
	}
	var data models.Tasks
	id := c.Params("id")
	if err := db.Find(&data, id).Delete(&data).Error; err != nil {
		return c.JSON(fiber.Map{
			"msg": err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Success",
	})
}
