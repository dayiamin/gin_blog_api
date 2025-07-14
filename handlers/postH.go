package handlers

import (
	"net/http"

	db "github.com/dayiamin/gin_blog_api/database"
	"github.com/dayiamin/gin_blog_api/models"
	"github.com/gin-gonic/gin"
)

// @Summary Get all posts
// @Description Retrieve a list of all posts with their associated comments.
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]models.Post "posts: List of posts with comments"
// @Failure 500 {object} map[string]string "error: Failed to fetch posts"
// @Router /post [get]
func ShowPosts(c *gin.Context) {

	var posts []models.Post

	if err := db.DB.Preload("PostComment").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})

}

// @Summary Create a new post
// @Description Create a new post with image address, title, and caption for the authenticated user. Requires JWT authentication.
// @Tags posts
// @Accept json
// @Produce json
// @Security JWT
// @Param post body models.PostRegister true "Post creation details"
// @Success 201 {object} map[string]interface{} "message: Post created, post: Created post data"
// @Failure 400 {object} map[string]string "error: Invalid input"
// @Failure 401 {object} map[string]string "error: Unauthorized (missing or invalid JWT)"
// @Failure 500 {object} map[string]string "error: Internal server error (database issue)"
// @Router /post/register [post]
func RegisterPost(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input models.PostRegister
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := userIDVal.(uint)

	post := models.Post{
		PicAddres: input.PicAddres,
		Title:     input.Title,
		Caption:   input.Caption,
		UserID:    userID,
	}

	if err := db.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "post created", "post": post})

}

// @Summary Delete a post
// @Description Delete a post and its associated comments by post ID. Requires JWT authentication.
// @Tags posts
// @Accept json
// @Produce json
// @Security JWT
// @Param post_id path string true "Post ID"
// @Success 200 {object} map[string]string "message: Post deleted, Post ID: Deleted post ID"
// @Failure 400 {object} map[string]string "error: Failed to find the post"
// @Failure 401 {object} map[string]string "error: Unauthorized (missing or invalid JWT)"
// @Failure 500 {object} map[string]string "error: Internal server error (database issue)"
// @Router /post/{post_id} [delete]
func DeletePost(c *gin.Context) {
	postID := c.Param("post_id")

	var postDB models.Post
	if err := db.DB.Where("id =?", postID).First(&postDB).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find the post"})
		return
	}

	if err := db.DB.Preload("PostComment").Delete(&postDB).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to Delete the post"})
		return
	}
	var postComments models.PostComment
	if err := db.DB.Where("post_id=?",postID).Delete(&postComments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find the comments"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "post Deleted", "Post ID": postID})

}

// @Summary Create a comment
// @Description Add a comment to a specific post for the authenticated user. Requires JWT authentication.
// @Tags comments
// @Accept json
// @Produce json
// @Security JWT
// @Param post_id path string true "Post ID"
// @Param comment body models.PostCommentsRegister true "Comment details"
// @Success 201 {object} map[string]interface{} "message: Comment added successfully, comment: Created comment data"
// @Failure 400 {object} map[string]string "error: Invalid input"
// @Failure 401 {object} map[string]string "error: Unauthorized (missing or invalid JWT)"
// @Failure 500 {object} map[string]string "error: Internal server error (post not found or database issue)"
// @Router /post/{post_id}/comments [post]
func RegisterComment(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDVal.(uint)
	postID := c.Param("post_id")


	var input models.PostCommentsRegister
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	if err := db.DB.Where("id=?", postID).First(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "post not found"})
		return
	}
	comment := models.PostComment{
		Text:   input.Text,
		UserID: userID,
		PostID: post.ID,
	}

	if err := db.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment added successfully",
		"comment": comment,
	})
}

// @Summary Delete a comment
// @Description Delete a comment by its ID for a specific post. Requires JWT authentication.
// @Tags comments
// @Accept json
// @Produce json
// @Security JWT
// @Param post_id path string true "Post ID"
/// @Param comment_id path string true "Comment ID"
// @Success 200 {object} map[string]string "message: Comment deleted, comment ID: Deleted comment ID"
// @Failure 400 {object} map[string]string "error: Failed to find the comment"
// @Failure 401 {object} map[string]string "error: Unauthorized (missing or invalid JWT)"
// @Failure 500 {object} map[string]string "error: Internal server error (database issue)"
// @Router /post/{post_id}/comments/{comment_id} [delete]
func DeleteComment(c *gin.Context) {

	commentID := c.Param("comment_id")

	var commentDB models.PostComment
	if err := db.DB.Where("id =?", commentID).First(&commentDB).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find the comment"})
		return
	}

	if err := db.DB.Delete(&commentDB).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to Delete the Comment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "comment Deleted", "comment ID": commentID})
}
