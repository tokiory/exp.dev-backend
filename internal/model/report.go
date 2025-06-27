package model

type ReportAddReq struct {
	Person struct {
		Name string `json:"name"`
		Surname string `json:"surname"`
		Patronymic string `json:"patronymic"`
		Telegram string `json:"telegram"`
		Email string `json:"email"`
	} `json:"personal"`

	Skills map[string]uint8 `json:"skills"`

	Work struct {
		Position string `json:"position"`
		Grade string `json:"grade"`
		GrowthMessage string `json:"growthMessage"`
		TasksMessage string `json:"tasksMessage"`
	} `json:"work"`
}

type ReportAddRes struct {
	Id string `json:"id"`
}
