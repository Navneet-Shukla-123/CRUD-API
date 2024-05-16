package models

// Pet store the data of individual animal
type Pet struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	PhotoUrls []string `json:"photo_urls"`
	Tags      []string `json:"tags"`
	Status    int    `json:"status"`
	/*
		1-- available
		2-- pending
		3-- sold
	*/
}
