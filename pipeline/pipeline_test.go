package pipeline

import (
	"CoolGoPkg/redis-go/client"
	"testing"
)

func init() {
	client.Init()
}

func TestGenBasicDataOfPipeline(t *testing.T) {
	GenBasicDataOfPipeline()
}

func TestGetDataByPipeline(t *testing.T) {
	GetDataByPipeline()
}
