package models

type GoogleTokenResponseModel struct {
	UserData struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		PhotoURL  string `json:"photoUrl"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		AuthToken string `json:"authToken"`
		IDToken   string `json:"idToken"`
		Response  struct {
			OT string `json:"OT"`
			Ad string `json:"Ad"`
			SV string `json:"sV"`
			VT string `json:"vT"`
			IK string `json:"iK"`
			Du string `json:"du"`
		} `json:"response"`
		Provider string `json:"provider"`
	} `json:"userData"`
}
