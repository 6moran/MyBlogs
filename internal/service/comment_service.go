package service

import (
	"MyBlogs/internal/model/dto/response"
	"MyBlogs/internal/model/entity"
	"MyBlogs/internal/repository"
)

type commentService struct {
	CommentRepo repository.CommentRepo
	UserRepo    repository.UserRepo
	ArticleRepo repository.ArticleRepo
}

func NewCommentService(commentRepo repository.CommentRepo, userRepo repository.UserRepo, articleRepo repository.ArticleRepo) CommentService {
	return &commentService{
		CommentRepo: commentRepo,
		UserRepo:    userRepo,
		ArticleRepo: articleRepo,
	}
}

func (s *commentService) GetCommentsByArticleID(articleID, page, size int) (*response.CommentListResponse, error) {
	comments, total, err := s.CommentRepo.FindByArticleID(articleID, page, size)
	if err != nil {
		return nil, err
	}

	commentResponses := make([]response.CommentResponse, 0, len(comments))
	for _, c := range comments {
		user, _ := s.UserRepo.FindByID(c.UserID)
		cr := response.CommentResponse{
			ID:        c.ID,
			ParentID:  c.ParentID,
			RootID:    c.RootID,
			UserID:    c.UserID,
			Content:   c.Content,
			CreatedAt: c.CreatedAt,
		}
		if user != nil {
			cr.Username = user.Username
			cr.Nickname = user.Nickname
			cr.Avatar = user.Avatar
		}
		commentResponses = append(commentResponses, cr)
	}

	return &response.CommentListResponse{
		Comments: commentResponses,
		Total:    total,
		Page:     page,
		Size:     size,
	}, nil
}

func (s *commentService) CreateComment(articleID, userID int, content string, parentID int) error {
	comment := &entity.Comment{
		ArticleID: articleID,
		UserID:    userID,
		Content:   content,
		ParentID:  parentID,
		RootID:    0,
	}

	// 如果是回复评论，设置rootID
	if parentID > 0 {
		parent, err := s.CommentRepo.FindByID(parentID)
		if err != nil {
			return err
		}
		if parent.RootID > 0 {
			comment.RootID = parent.RootID
		} else {
			comment.RootID = parent.ID
		}
	}

	return s.CommentRepo.Create(comment)
}

func (s *commentService) DeleteComment(id int) error {
	comment, err := s.CommentRepo.FindByID(id)
	if err != nil {
		return err
	}

	// 如果是根评论，删除所有子评论（使用 rootID = id）
	if comment.ParentID == 0 {
		if err := s.CommentRepo.DeleteByRootID(id); err != nil {
			return err
		}
	}

	// 删除评论本身
	return s.CommentRepo.Delete(id)
}
