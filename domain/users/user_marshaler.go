package users

import "encoding/json"

// Struct for Public user JSON response
type PublicUser struct {
	Id 			int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status 		string `json:"status"`
}

// Struct for Private user JSON response
type PrivateUser struct {
	Id 			int64  `json:"id"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email 		string `json:"email"`
	DateCreated string `json:"date_created"`
	Status 		string `json:"status"`
}

// Marshal users found by status
func (users Users) Marshal(isPublic bool) []interface{}{
	result := make([]interface{}, len(users))
	for index, user := range users{
		result[index] = user.Marshal(isPublic)
	}

	return result
}

// JSON response for public or private user
func (user *User) Marshal(isPublic bool) interface{}{
	if isPublic{
		return PublicUser{
			Id : user.Id,
			DateCreated: user.DateCreated,
			Status: user.Status,
		}
	}

	userJson, _:= json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser

	//return PrivateUser{
	//	Id : user.Id,
	//	FirstName: user.FirstName,
	//	LastName: user.LastName,
	//	Email: user.Email,
	//	DateCreated: user.DateCreated,
	//	Status: user.Status,
	//}
}
