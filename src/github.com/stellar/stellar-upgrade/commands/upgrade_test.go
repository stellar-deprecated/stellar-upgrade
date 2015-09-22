package commands

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/go-stellar-base"
	"github.com/stellar/stellar-upgrade/api"
	"github.com/stellar/stellar-upgrade/commands"
	"github.com/stretchr/testify/mock"
)

type MockInput struct {
	mock.Mock
}
func (m *MockInput) GetOldNetworkSeedFromConsole() string {
	args := m.Called()
	return args.String(0)
}
func (m *MockInput) GetNewNetworkAddressFromConsole() string {
	args := m.Called()
	return args.String(0)
}
func (m *MockInput) GetConfirmationFromConsole(oldNetworkAddress, newNetworkAddress string) bool {
	args := m.Called(oldNetworkAddress, newNetworkAddress)
	return args.Bool(0)
}

type MockApi struct {
	UpgradeResponse *api.UpgradeResponse
	StatusResponse *api.StatusResponse
}
func (mockApi MockApi) SendUpgradeRequest(data api.MessageData, publicKey stellarbase.PublicKey, privateKey stellarbase.PrivateKey) (*api.UpgradeResponse, error) {
	return mockApi.UpgradeResponse, nil
}
func (mockApi MockApi) SendStatusRequest(address string) (*api.StatusResponse, error) {
	return mockApi.StatusResponse, nil
}

func TestUpgradeCommand(t *testing.T) {
	Convey("Given upgrade command", t, func() {
		mockInput := new(MockInput)
		mockApi := MockApi{}
		updateCommand := commands.UpdateCommand{
			Input: mockInput,
			ApiObject: &mockApi,
		}

		Convey("When old network seed is incorrect", func() {
			mockInput.On("GetOldNetworkSeedFromConsole").Return("wrong")

			Convey("it should panic with error", func() {
				displayedMessage := updateCommand.Run()
				So(displayedMessage, ShouldEqual, "Your old network account secret seed is incorrect.")
				mockInput.AssertExpectations(t)
			})
		})

		Convey("When old network seed is correct and user inputs new network address", func() {
			mockInput.On("GetOldNetworkSeedFromConsole").Return("s3tZPX5xE9obmKfR61vJwFVHHwVxG32DwCJb4XyMpC3Rtu4PsgG")
			mockInput.On("GetNewNetworkAddressFromConsole").Return("GD2EH5THFB4D575RHFKBCJBDNBEO53QUAETP7ZVH42RB2D3RRYCVPN6D")
			mockInput.On("GetConfirmationFromConsole", "gWRYUerEKuz53tstxEuR3NCkiQDcV4wzFHmvLnZmj7PUqxW2wn", "GD2EH5THFB4D575RHFKBCJBDNBEO53QUAETP7ZVH42RB2D3RRYCVPN6D").Return(false)

			Convey("it should ask for confirmation", func() {
				updateCommand.Run()
				mockInput.AssertExpectations(t)
			})
		})
	})
}
