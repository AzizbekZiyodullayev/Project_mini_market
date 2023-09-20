package testing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// func TestCreateBranch(t *testing.T) {
// 	response := &models.Branch{}

// 	request := &models.CreateBranch{
// 		Name:      faker.Name(),
// 		Address:   faker.MacAddress(),
// 		FoundedAt: faker.Date(),
// 	}

// 	resp, err := makeRequest(http.MethodPost, "/branch", request, response)

// 	assert.NoError(t, err)

// 	assert.NotNil(t, resp)

// 	if resp != nil {
// 		assert.Equal(t, resp.StatusCode, 201)
// 	}

// 	fmt.Println(response)

// }

// func TestUpdateBranch(t *testing.T) {
// 	response := &models.Branch{}

// 	request := &models.UpdateBranch{
// 		Name:    faker.Name(),
// 		Address: faker.MacAddress(),
// 	}

// 	resp, err := makeRequest(http.MethodPut, "/branch/:id", request, response)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, resp)

// 	if resp != nil {
// 		assert.Equal(t, resp.StatusCode, 200)
// 	}
// 	fmt.Println(response)
// }

func makeRequest(method, path string, req, res interface{}) (*http.Response, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	request, err := http.NewRequestWithContext(context.Background(),
		method,
		fmt.Sprintf("%s%s", "http://localhost:8080", path), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(resp_body, res)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
