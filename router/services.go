package router

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/smtnn-ks/go-test-2/algo"
	"github.com/smtnn-ks/go-test-2/db"
	"github.com/smtnn-ks/go-test-2/models"
)

func checkAuth(name, password string) error {
	usr, err := db.Client.FindOneFrom(models.UsrTable, "usr_name", name)
	if err != nil {
		log.Error(err)
		return fiber.ErrInternalServerError
	}
	if usr.(*models.Usr).Pass != password {
		return fiber.ErrInternalServerError
	}

	return nil
}

func register(dto registerDto_t) error {
	usr := &models.Usr{
		UsrName: dto.Name,
		Pass:    dto.Password,
	}

	if err := db.Client.Save(usr); err != nil {
		log.Error(err)
		return fiber.ErrInternalServerError
	}

	return nil
}

func create(dto newsDto_t) error {
	news := &models.News{
		Title: dto.Title,
		Cnt:   dto.Content,
	}

	if err := db.Client.Save(news); err != nil {
		log.Error(err)
		return fiber.ErrInternalServerError
	}

	for _, category := range dto.Categories {
		newsCategories := &models.NewsCategories{
			NewsID: news.ID,
		}
		newsCategories.CategoryID = category
		if err := db.Client.Save(newsCategories); err != nil {
			log.Error(err)
			return fiber.ErrInternalServerError
		}
	}

	return nil
}

func edit(newsId int64, newsDto newsDto_t) error {
	newsRecord, err := db.Client.FindByPrimaryKeyFrom(models.NewsTable, newsId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}
		log.Error(err)
		return fiber.ErrInternalServerError
	}

	news := newsRecord.(*models.News)
	if newsDto.Title != "" {
		news.Title = newsDto.Title
	}
	if newsDto.Content != "" {
		news.Cnt = newsDto.Content
	}

	if err := db.Client.Update(news); err != nil {
		log.Error(err)
		return fiber.ErrInternalServerError
	}

	rows, err := db.Client.FindAllFrom(models.NewsCategoriesTable, "news_id", newsId)
	if err != nil {
		log.Error(err)
		return fiber.ErrInternalServerError
	}

	var categories []int64

	for _, row := range rows {
		categories = append(categories, row.(*models.NewsCategories).CategoryID)
	}

	if newsDto.Categories != nil {
		toDelete, toInsert := algo.Exclude(newsDto.Categories, categories)

		for _, i := range toDelete {
			_, err := db.Client.DeleteFrom(
				models.NewsCategoriesTable,
				"WHERE news_id = $1 AND category_id = $2",
				newsId,
				i,
			)

			if err != nil {
				log.Error(err)
				return fiber.ErrInternalServerError
			}
		}

		for _, c := range toInsert {
			if err := db.Client.Insert(&models.NewsCategories{
				NewsID:     newsId,
				CategoryID: c,
			}); err != nil {
				log.Error(err)
				return fiber.ErrInternalServerError
			}
		}
	}
	return nil
}

func list(limit, offset int64) ([]list_t, error) {
	rows, err := db.Client.SelectAllFrom(models.NewsTable, "LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}
	list := make([]list_t, len(rows))
	for i, row := range rows {
		list[i].Id = row.(*models.News).ID
		list[i].Title = row.(*models.News).Title
		list[i].Content = row.(*models.News).Cnt

		newsCategories, err := db.Client.FindAllFrom(models.NewsCategoriesTable, "news_id", list[i].Id)
		if err != nil {
			log.Error(err)
			return nil, fiber.ErrInternalServerError
		}

		for _, c := range newsCategories {
			list[i].Categories = append(list[i].Categories, c.(*models.NewsCategories).CategoryID)
		}
	}

	return list, nil
}
