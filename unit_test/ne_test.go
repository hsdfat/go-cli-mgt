package unit

import (
	"errors"
	"github.com/stretchr/testify/require"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/models/models_error"
	"go-cli-mgt/pkg/service/network_elements"
	"go-cli-mgt/pkg/service/utils/random"
	"testing"
)

func TestCreateAndDeleteNe(t *testing.T) {
	// Create neTest
	neTest := models_api.NeData{
		Name:             random.StringRandom(10),
		Url:              random.StringRandom(10),
		Type:             random.StringRandom(10),
		MasterIpConfig:   random.StringRandom(10),
		MasterPortConfig: int32(random.IntRandom(0, 9999)),
		SlaveIpConfig:    random.StringRandom(10),
		SlavePortConfig:  int32(random.IntRandom(0, 9999)),
		IpCommand:        random.StringRandom(10),
		PortCommand:      int32(random.IntRandom(0, 9999)),
		Namespace:        random.StringRandom(10),
	}

	// Create Ne
	err := network_elements.CreateNetworkElement(&neTest)
	require.NoError(t, err)

	// Get Ne
	neGetTest, err := network_elements.GetNetworkElement(neTest.Name, neTest.Namespace)
	require.NoError(t, err)
	require.NotEmpty(t, neGetTest)
	require.Equal(t, neTest.Name, neGetTest.Name)
	require.Equal(t, neTest.Url, neGetTest.Url)
	require.Equal(t, neTest.Type, neGetTest.Type)
	require.Equal(t, neTest.MasterIpConfig, neGetTest.MasterIpConfig)
	require.Equal(t, neTest.MasterPortConfig, neGetTest.MasterPortConfig)
	require.Equal(t, neTest.SlaveIpConfig, neGetTest.SlaveIpConfig)
	require.Equal(t, neTest.SlavePortConfig, neGetTest.SlavePortConfig)
	require.Equal(t, neTest.IpCommand, neGetTest.IpCommand)
	require.Equal(t, neTest.PortCommand, neGetTest.PortCommand)
	require.Equal(t, neTest.Namespace, neGetTest.Namespace)

	// Delete Ne
	err = network_elements.DeleteNetworkElement(neGetTest.Name, neGetTest.Namespace)
	require.NoError(t, err)

	// Get Ne again to test
	neGetTest, err = network_elements.GetNetworkElement(neTest.Name, neTest.Namespace)
	if errors.Is(err, models_error.ErrNotFoundUser) == false {
		require.Error(t, errors.New("delete user un success"))
	}
}

func TestGetNe(t *testing.T) {
	neTest := models_api.NeData{
		Name:      "phatlc-computer",
		Namespace: "Co Nhue, Ha Noi",
	}

	neGetTest, err := network_elements.GetNetworkElement(neTest.Name, neTest.Namespace)
	require.NoError(t, err)
	require.NotEmpty(t, neGetTest)
	require.Equal(t, neTest.Namespace, neGetTest.Namespace)
	require.Equal(t, neTest.Name, neGetTest.Name)
}
