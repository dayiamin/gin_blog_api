package handlers

import (
	"fmt"

	"net/http"

	db "github.com/dayiamin/gin_blog_api/database"
	"github.com/dayiamin/gin_blog_api/models"
	"github.com/dayiamin/gin_blog_api/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


// @Summary Register a new user
// @Description Create a new user account with username, email, and password. Returns a JWT token upon successful registration.
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.RegisterUsers true "User registration details"
// @Success 201 {object} map[string]string "message: Success message, token: JWT token"
// @Failure 400 {object} map[string]string "error: Invalid input"
// @Failure 500 {object} map[string]string "error: Internal server error (hashing or database issue)"
// @Router /user/register [post]
func RegisterUser(c *gin.Context) {
	var body models.RegisterUsers
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error in hashing password"})
		return
	}

	user := models.User{
		UserName: body.UserName,
		Email:    body.Email,
		Password: string(hashedPassword),
	}

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user in db"})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	message := fmt.Sprintf("the User %v created succesfully", user.UserName)
	c.JSON(http.StatusCreated, gin.H{"message": message, "token": token})

}


// @Summary User login
// @Description Authenticate a user using username or email and password. Returns a JWT token upon successful login.
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body models.UserLoginRequest true "Login credentials (username or email and password)"
// @Success 200 {object} map[string]string "message: Success message, token: JWT token"
// @Failure 400 {object} map[string]string "error: Invalid input"
// @Failure 401 {object} map[string]string "error: Incorrect username/email or password"
// @Failure 500 {object} map[string]string "error: Internal server error (JWT generation)"
// @Router /user/login [post]
func Login(c *gin.Context) {
	var input models.UserLoginRequest

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	// checking input for username or email
	if err := db.DB.Where("email = ?", &input.Credential).First(&user).Error; err != nil {
		if err := db.DB.Where("user_name = ?", &input.Credential).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username or Email is incorecct"})
			return
		}
	}

	// decoding the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in generating jwt"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Loged in successfully",
		"token":   token,
	})

}

// @Summary Create or update user profile
// @Description Create a new user profile or update an existing one for the authenticated user. Requires JWT authentication.
// @Tags profiles
// @Accept json
// @Produce json
// @Security JWT
// @Param profile body models.UserProfileRegister true "User profile details"
// @Success 201 {object} map[string]interface{} "message: Profile created, profile: User profile data"
// @Success 200 {object} map[string]interface{} "message: Profile updated, profile: Updated user profile data"
// @Failure 400 {object} map[string]string "error: Invalid input"
// @Failure 401 {object} map[string]string "error: Unauthorized (missing or invalid JWT)"
// @Failure 500 {object} map[string]string "error: Internal server error (database issue)"
// @Router /user/profile [post]
func CreateProfile(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input models.UserProfileRegister

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userID := userIDVal.(uint)

	// checking if user profile already exist so we should use the update profile route
	var existingProfile models.UserProfile
	if err := db.DB.Where("user_id = ?", userID).First(&existingProfile).Error; err == nil {

		if helperUpdateProfile(input.FirstName) {
			existingProfile.FirstName = input.FirstName
		}
		if helperUpdateProfile(input.LastName) {
			existingProfile.LastName = input.LastName
		}
		if helperUpdateProfile(input.Bio) {
			existingProfile.Bio = input.Bio
		}
		if helperUpdateProfile(input.ProfilePic) {
			existingProfile.ProfilePic = input.ProfilePic
		}
		db.DB.Save(&existingProfile)
		c.JSON(http.StatusOK, gin.H{"message": "Profile updated", "profile": existingProfile})
		return
	}

	profile := models.UserProfile{
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Bio:        input.Bio,
		ProfilePic: input.ProfilePic,
		UserID:     userID,
	}

	if err := db.DB.Create(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create profile"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Profile created", "profile": profile})

}

func helperUpdateProfile(input string) bool {
	if len(input) == 0 || input == "" {
		return false
	}
	return true
}


// @Summary Get user profile
// @Description Retrieve the profile details of a user by their username. Requires JWT authentication.
// @Tags profiles
// @Accept json
// @Produce json
// @Security JWT
// @Param user_name path string true "Username of the user"
// @Success 200 {object} map[string]models.UserProfile "profile: User profile data"
// @Failure 400 {object} map[string]string "error: User or profile not found"
// @Failure 401 {object} map[string]string "error: Unauthorized (missing or invalid JWT)"
// @Router /user/profile/{user_name} [get]
func ShowProfile(c *gin.Context){
	userName := c.Param("user_name")

	var user models.User
	if err := db.DB.Where("user_name = ?", userName).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	var profile models.UserProfile
	if err := db.DB.Where("user_id = ?", user.ID).First(&profile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User profile not found"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"profile": profile})
}