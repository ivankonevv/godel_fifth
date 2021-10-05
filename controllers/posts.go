package controllers

import (
	"fifth_test/database"
	"fifth_test/models"
	"fifth_test/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"strings"
)

func GetAllPosts(c *fiber.Ctx) error {
	db := database.DB
	var posts []models.Post

	if c.Query("sort_by") != "" {
		db.Order(c.Query("sort_by")).Preload("Images").Find(&posts)
	} else {
		db.Preload("Images").Find(&posts)
	}
	if len(posts) == 0 {
		return c.Status(409).JSON(fiber.Map{
			"status": "error",
			"msg":    "No posts yet",
			"result": nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"msg":    "Posts found",
		"result": posts,
	})
}

func GetPost(c *fiber.Ctx) error {
	db := database.DB
	var post models.Post
	id := c.Params("postId")
	q := c.Query("fields")
	db.Preload("Images").First(&post, "post_id = ?", id)
	if post.PostID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No post present",
			"data":    nil,
		})
	}
	if strings.Contains(q, "images") && strings.Contains(q, "description") {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Post found",
			"data": fiber.Map{
				"title":       post.Title,
				"description": post.Description,
				"price":       post.Price,
				"images":      post.Images,
			},
		})
	} else if strings.Contains(q, "images") && !strings.Contains(q, "description") {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Post found",
			"data": fiber.Map{
				"title":  post.Title,
				"price":  post.Price,
				"images": post.Images,
			},
		})
	} else if !strings.Contains(q, "images") && strings.Contains(q, "description") {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Post found",
			"data": fiber.Map{
				"title":       post.Title,
				"description": post.Description,
				"price":       post.Price,
				"image":       post.Images[0],
			},
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Post found",
		"data": fiber.Map{
			"title": post.Title,
			"price": post.Price,
			"image": post.Images[0],
		},
	})
}

func CreatePost(c *fiber.Ctx) error {
	db := database.DB
	post := new(models.Post)

	err := c.BodyParser(post)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"msg":    "Incorrect input",
			"result": err,
		})
	}
	errors := validators.ValidateStruct(*post)
	if errors != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"msg":    "Validation error",
			"result": errors,
		})
	}
	post.PostID = uuid.New()

	err = db.Create(&post).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"msg":    "Could not create post",
			"result": err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"msg":    "Created post",
		"result": fiber.Map{
			"post_id": post.PostID,
		},
	})
}
