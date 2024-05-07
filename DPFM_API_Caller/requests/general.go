package requests

type General struct {
	BusinessPartner				int		`json:"BusinessPartner"`
	Friend						int		`json:"Friend"`
	BPBusinessPartnerType		string	`json:"BPBusinessPartnerType"`
	FriendBusinessPartnerType	string	`json:"FriendBusinessPartnerType"`
	CommunityRank				int		`json:"CommunityRank"`
	FriendIsBlocked				bool	`json:"FriendIsBlocked"`
	CreationDate				string	`json:"CreationDate"`
	CreationTime				string	`json:"CreationTime"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}
