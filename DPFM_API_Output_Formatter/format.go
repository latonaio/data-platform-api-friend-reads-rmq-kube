package dpfm_api_output_formatter

import (
	"data-platform-api-friend-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToGeneral(rows *sql.Rows) (*[]General, error) {
	defer rows.Close()
	general := make([]General, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.General{}

		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Friend,
			&pm.BPBusinessPartnerType,
			&pm.FriendBusinessPartnerType,
			&pm.CommunityRank,
			&pm.FriendIsBlocked,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &general, err
		}

		data := pm
		general = append(general, General{
			BusinessPartner:			data.BusinessPartner,
			Friend:						data.Friend,
			BPBusinessPartnerType:		data.BPBusinessPartnerType,
			FriendBusinessPartnerType:	data.FriendBusinessPartnerType,
			CommunityRank:				data.CommunityRank,
			FriendIsBlocked:			data.FriendIsBlocked,
			CreationDate:				data.CreationDate,
			CreationTime:				data.CreationTime,
			LastChangeDate:				data.LastChangeDate,
			LastChangeTime:				data.LastChangeTime,
			IsMarkedForDeletion:		data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &general, nil
	}

	return &general, nil
}
