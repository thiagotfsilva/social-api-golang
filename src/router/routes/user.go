package routes

import (
	"api-devbook/src/controllers"
	"net/http"
)

var UserRoute = []Route{
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.CreateUser,
		Authentication: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Function:       controllers.FindUsers,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodGet,
		Function:       controllers.FindUser,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodPut,
		Function:       controllers.UpdateUser,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodDelete,
		Function:       controllers.DeleteUser,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}/follow",
		Method:         http.MethodPost,
		Function:       controllers.FollowUser,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}/unfollow",
		Method:         http.MethodPost,
		Function:       controllers.UnfollowUser,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}/followers",
		Method:         http.MethodGet,
		Function:       controllers.GetFollowers,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}/following",
		Method:         http.MethodGet,
		Function:       controllers.GetFollowing,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}/update-password",
		Method:         http.MethodPost,
		Function:       controllers.UpdatePassword,
		Authentication: true,
	},
}
