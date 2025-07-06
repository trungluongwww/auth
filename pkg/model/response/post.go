package response

import "time"

type PostResponse struct {
	ID           uint32       `json:"id"`
	Title        string       `json:"title"`
	Content      string       `json:"content"`
	ImageURL     *string      `json:"imageUrl"`
	IsPublic     bool         `json:"isPublic"`
	ViewCount    int          `json:"viewCount"`
	LikeCount    int          `json:"likeCount"`
	CommentCount int          `json:"commentCount"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	User         UserResponse `json:"user"`
	IsLiked      bool         `json:"isLiked"`
}

type CommentResponse struct {
	ID        uint32            `json:"id"`
	Content   string            `json:"content"`
	LikeCount int               `json:"likeCount"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
	User      UserResponse      `json:"user"`
	PostID    uint32            `json:"postId"`
	ParentID  *uint32           `json:"parentId"`
	IsLiked   bool              `json:"isLiked"`
	Replies   []CommentResponse `json:"replies,omitempty"`
}

type PostListResponse struct {
	Posts      []PostResponse `json:"posts"`
	TotalCount int            `json:"totalCount"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
}

type CommentListResponse struct {
	Comments   []CommentResponse `json:"comments"`
	TotalCount int               `json:"totalCount"`
	Page       int               `json:"page"`
	Limit      int               `json:"limit"`
}

type UserProfileResponse struct {
	UserResponse
	FollowerCount  int  `json:"followerCount"`
	FollowingCount int  `json:"followingCount"`
	PostCount      int  `json:"postCount"`
	IsFollowing    bool `json:"isFollowing"`
}

type FollowListResponse struct {
	Users      []UserResponse `json:"users"`
	TotalCount int            `json:"totalCount"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
}
