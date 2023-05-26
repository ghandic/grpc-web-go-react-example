package repositories

import (
	"errors"
	v1 "github.com/ghandic/grpc-web-go-react-example/backend/api/proto/users/v1"
	"github.com/ghandic/grpc-web-go-react-example/backend/db/users"
)

func getListParams(req *v1.ListUsersRequest) (*users.GetUsersParams, error) {
	listParams := &users.GetUsersParams{}
	if req.Sorting != nil {
		switch req.Sorting.Field {
		case "name":
			switch req.Sorting.Direction {
			case v1.SortDirection_ASC:
				listParams.NameAsc = true
			case v1.SortDirection_DESC:
				listParams.NameDesc = true
			default:
				return listParams, errors.New("invalid sort by direction")
			}
		case "id":
			switch req.Sorting.Direction {
			case v1.SortDirection_ASC:
				listParams.IDAsc = true
			case v1.SortDirection_DESC:
				listParams.IDDesc = true
			default:
				return listParams, errors.New("invalid sort by direction")
			}
		case "created_at":
			switch req.Sorting.Direction {
			case v1.SortDirection_ASC:
				listParams.CreatedAtAsc = true
			case v1.SortDirection_DESC:
				listParams.CreatedAtDesc = true
			default:
				return listParams, errors.New("invalid sort by direction")
			}
		default:
			return listParams, errors.New("invalid field")
		}
	}

	if req.Query != nil {
		if req.Query.Text != "" {
			listParams.Search = req.Query.Text
		}
	}

	if req.PageSize > 0 {
		listParams.LimitAmount = req.PageSize
	}

	if req.Offset >= 0 {
		listParams.OffsetAmount = req.Offset
	}

	return listParams, nil
}
