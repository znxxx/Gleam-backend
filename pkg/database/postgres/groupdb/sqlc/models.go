// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package groupdb

import (
	"database/sql"
)

type Group struct {
	GroupID        int32          `json:"group_id"`
	GroupName      string         `json:"group_name"`
	GroupCreatorID int32          `json:"group_creator_id"`
	PhotoUrl       sql.NullString `json:"photo_url"`
	CreatedAt      sql.NullTime   `json:"created_at"`
}

type GroupMember struct {
	GroupID   int32        `json:"group_id"`
	MemberID  int32        `json:"member_id"`
	Role      string       `json:"role"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Post struct {
	PostID      int32          `json:"post_id"`
	MemberID    int32          `json:"member_id"`
	GroupID     int32          `json:"group_id"`
	PhotoUrl    sql.NullString `json:"photo_url"`
	Description sql.NullString `json:"description"`
	CreatedAt   sql.NullTime   `json:"created_at"`
}

type PostComment struct {
	CommentID int32        `json:"comment_id"`
	PostID    int32        `json:"post_id"`
	Comment   string       `json:"comment"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type PostReaction struct {
	ReactionID int32        `json:"reaction_id"`
	PostID     int32        `json:"post_id"`
	Reaction   string       `json:"reaction"`
	CreatedAt  sql.NullTime `json:"created_at"`
}
