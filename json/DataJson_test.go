package json

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRealJsonFromFile(t *testing.T) {
	assert := assert.New(t)

	filePath := "./test.json"
	result := RealJsonFromFile(filePath)
	assert.NotNil(result)

	assert.Equal(2, len(result.Posts))
	assert.Equal(3, len(result.Comments))

	expectedTime, _ := time.Parse(time.RFC3339, "2018-09-22T12:42:31+07:00")

	assert.Equal(expectedTime, result.Posts[0].CreatedTime)
	assert.Equal(expectedTime, result.Posts[1].CreatedTime)

	assert.Equal(expectedTime, result.Comments[0].CreatedTime)
	assert.Equal(expectedTime, result.Comments[1].CreatedTime)
	assert.Equal(expectedTime, result.Comments[2].CreatedTime)
}

func TestSaveJsonToFile(t *testing.T) {
	assert := assert.New(t)

	// read data first
	filePath := "./test.json"
	result := RealJsonFromFile(filePath)
	assert.NotNil(result)

	// modify data
	result.NextCommentID = 5
	result.NextPostID = 1
	result.Posts[0].Title = "New title"
	result.Posts[0].Content = "New content"
	result.Posts[0].CoverImage = "New cover"
	result.Posts[0].UpdatedTime = time.Now()

	result.Comments[0].Username = "New user"
	result.Comments[0].Content = "New content"

	result.Comments[2].Username = "New user 12"
	result.Comments[2].Content = "New content 12"
	result.Comments[2].CreatedTime = time.Now()

	// save data back to test2.json
	err := SaveJsonToFile("./test2.json", result)
	log.Println(err)
	assert.Nil(err)

}
