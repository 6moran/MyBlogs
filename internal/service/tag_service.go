package service

import (
	"MyBlogs/internal/model/dto/request"
	"MyBlogs/internal/model/dto/response"
	"MyBlogs/internal/model/entity"
	"MyBlogs/internal/repository"
	bizerrors "MyBlogs/pkg/errors"
	"fmt"
)

type tagService struct {
	TagRepo repository.TagRepo
}

func NewTagService(repo repository.TagRepo) TagService {
	return &tagService{TagRepo: repo}
}

func (ts *tagService) CreateTag(req request.TagRequest) error {
	// 检查标签名是否已存在
	tags, err := ts.TagRepo.FindAll()
	if err != nil {
		return fmt.Errorf("查询标签失败: %w", err)
	}

	for _, tag := range tags {
		if tag.Name == req.Name {
			return bizerrors.New(bizerrors.CodeTagNameAlreadyExists, "标签名已存在")
		}
	}

	// 创建标签
	err = ts.TagRepo.CreateTag(&entity.Tag{
		Name: req.Name,
	})
	if err != nil {
		return fmt.Errorf("创建标签失败: %w", err)
	}

	return nil
}

func (ts *tagService) UpdateTag(id int, req request.TagRequest) error {
	// 检查标签是否存在
	existingTag, err := ts.TagRepo.FindByID(id)
	if err != nil {
		return bizerrors.New(bizerrors.CodeTagNotFound, "标签不存在")
	}

	// 检查标签名是否已被其他标签使用
	tags, err := ts.TagRepo.FindAll()
	if err != nil {
		return fmt.Errorf("查询标签失败: %w", err)
	}

	for _, tag := range tags {
		if tag.Name == req.Name && tag.ID != id {
			return bizerrors.New(bizerrors.CodeTagNameAlreadyExists, "标签名已存在")
		}
	}

	// 更新标签
	existingTag.Name = req.Name
	err = ts.TagRepo.EditTag(existingTag)
	if err != nil {
		return fmt.Errorf("更新标签失败: %w", err)
	}

	return nil
}

func (ts *tagService) DeleteTag(id int) error {
	// 检查标签是否存在
	_, err := ts.TagRepo.FindByID(id)
	if err != nil {
		return bizerrors.New(bizerrors.CodeTagNotFound, "标签不存在")
	}

	// 删除标签（级联删除文章-标签关联）
	err = ts.TagRepo.DeleteTag(id)
	if err != nil {
		return fmt.Errorf("删除标签失败: %w", err)
	}

	return nil
}

func (ts *tagService) GetTagByID(id int) (*response.TagResponse, error) {
	// 检查标签是否存在
	tag, err := ts.TagRepo.FindByID(id)
	if err != nil {
		return nil, bizerrors.New(bizerrors.CodeTagNotFound, "标签不存在")
	}

	// 转换为响应DTO
	resp := &response.TagResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}

	return resp, nil
}

func (ts *tagService) GetTagList(page, size int) ([]response.TagResponse, int64, error) {
	// 查询标签列表（分页）
	tags, err := ts.TagRepo.QueryTagLimit(repository.QueryLimit{
		Page: page,
		Size: size,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("查询标签列表失败: %w", err)
	}

	// 查询总数
	allTags, err := ts.TagRepo.FindAll()
	if err != nil {
		return nil, 0, fmt.Errorf("查询标签总数失败: %w", err)
	}
	total := int64(len(allTags))

	// 转换为响应DTO
	result := make([]response.TagResponse, len(tags))
	for i, tag := range tags {
		result[i] = response.TagResponse{
			ID:   tag.ID,
			Name: tag.Name,
		}
	}

	return result, total, nil
}

func (ts *tagService) GetAllTags() ([]response.TagResponse, error) {
	tags, err := ts.TagRepo.FindAll()
	if err != nil {
		return nil, err
	}

	result := make([]response.TagResponse, len(tags))
	for i, t := range tags {
		result[i] = response.TagResponse(t)
	}
	return result, nil
}
