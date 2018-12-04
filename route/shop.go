package route

import (
	"net/http"
	"github.com/QuanLab/go-service/service"
	"github.com/QuanLab/go-service/model"
	"encoding/json"
	"log"
)

func GetListId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var accessToken = r.Header.Get("access_token")
		user := service.ValidateToken(accessToken)

		if user.Role == model.GUEST {
			w.Write([]byte("{'message' : 'Current user is not authenticated.'}"))
		}

		if user.Role == model.BASIC {
			listUser := model.GetListShopToFollow(user.ID)
			result, _ := json.Marshal(model.Shop{Data:listUser})
			w.Write([]byte(result))
		}

		if user.Role == model.ADVANCE {
			listUser := model.GetListShopToFollow(user.ID)
			result, _ := json.Marshal(model.Shop{Data:listUser})
			w.Write([]byte(result))
		}

		if user.Role == model.PRO {
			listUser := model.GetListShopToFollow(user.ID)
			result, _ := json.Marshal(model.Shop{Data:listUser})
			w.Write([]byte(result))
		}

		if user.Role == model.ADMIN {
			result, err := json.Marshal(user)
			if err != nil {
				log.Panic(err)
			}
			w.Write(result)
		}
	default:
		w.Write([]byte("{'message', 'Method is not allowed here'}"))
	}
}
