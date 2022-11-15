package service

import (
	"context"
	"github.com/SamuelDevMobile/gRPC_GoLang/internal/database"
	"github.com/SamuelDevMobile/gRPC_GoLang/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	categoryDB, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id: 			categoryDB.ID,
		Name: 			categoryDB.Name,
		Description: 	categoryDB.Description,
	}

	return categoryResponse, nil
}