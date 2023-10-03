package model

type User struct {
	UserId    int64  `json:"userId"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	Token     string `json:"token"`
}

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

type UserGetRequest struct {
}

type UserGetResponse struct {
	Name string `json:"name"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}

type UserUpdateResponse struct {
	Name string `json:"name"`
}
