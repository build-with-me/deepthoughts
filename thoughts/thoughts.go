package thoughts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://deepthoughts-api.azurewebsites.net/"

func Random() (DeepThought, error) {
    
    response, err := http.Get(url)

    if (err != nil) {
        return DeepThought{}, err
    }

    if (response.StatusCode == http.StatusOK) {
        defer response.Body.Close()

        bodyP, err := io.ReadAll(response.Body)

        if (err != nil) {
            return DeepThought{}, err
        }

        var thought DeepThought
        err = json.Unmarshal(bodyP, &thought)

        if (err != nil) {
            return DeepThought{}, err
        }

        return thought, nil
    } else {
        return DeepThought{}, nil
    }
}

func Vote(id int, isUpvote bool) (error) {

    putUrl := fmt.Sprintf("%s/%d?isUpvote=%t", url, id, isUpvote)

    client := &http.Client{}
    req, err := http.NewRequest(http.MethodPut, putUrl, nil)

    _, err = client.Do(req)

    if (err != nil) {
        return err
    }
    return nil
}

type DeepThought struct {
    Id int `json:"id"`
    Text string `json:"text"`
    Views int `json:"views"`
    Upvotes int `json:"upvotes"`
    Downvotes int `json:"downvotes"`
}
