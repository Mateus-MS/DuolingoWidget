package routes

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/Mateus-MS/DuolingoWidget/utils"
)

func StreakRoute(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "https://cubonauta.com")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, X-CSRF-Token")
	w.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token")

	userID, err := utils.QueryFromURL[string]("id", r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := http.Get(fmt.Sprintf("https://www.duolingo.com/2017-06-30/users/%s", userID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	streak, err := queryStreak(body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(streak))

}

func queryStreak(responseBody []byte) (string, error) {
	regex := regexp.MustCompile(`"streak":\s*(\d+)`)
	match := regex.FindStringSubmatch(string(responseBody))

	if len(match) == 0 {
		return "", errors.New("Streak not finded in the body")
	}
	return match[1], nil
}
