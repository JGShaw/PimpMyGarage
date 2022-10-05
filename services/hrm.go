package services

import (
	"fmt"
	"time"
	"tinygo.org/x/bluetooth"
)

const scanTimeout = 10 * time.Second
const connectionTimeout = 10 * time.Second

type HrmService struct {
	adapter             *bluetooth.Adapter
	connectedDeviceName string
	device              *bluetooth.Device
	listeners           []func(float64)
}

func NewHrmService() (*HrmService, error) {
	service := HrmService{
		adapter: bluetooth.DefaultAdapter,
	}

	err := service.adapter.Enable()
	return &service, err
}

func (h *HrmService) Scan() ([]string, error) {
	var foundDevices []string

	go func() {
		<-time.After(scanTimeout)
		h.adapter.StopScan()
	}()

	err := h.adapter.Scan(func(adapter *bluetooth.Adapter, scanResult bluetooth.ScanResult) {
		if scanResult.HasServiceUUID(bluetooth.ServiceUUIDHeartRate) {
			foundDevices = append(foundDevices, scanResult.LocalName())
		}
	})
	if err != nil {
		return nil, err
	}

	return foundDevices, nil
}

func (h *HrmService) Connect(deviceName string) error {
	var found bluetooth.ScanResult

	err := h.adapter.Scan(func(adapter *bluetooth.Adapter, scanResult bluetooth.ScanResult) {
		if scanResult.LocalName() == deviceName {
			found = scanResult
			h.adapter.StopScan()
		}
	})

	device, err := h.adapter.Connect(found.Address, bluetooth.ConnectionParams{
		ConnectionTimeout: bluetooth.NewDuration(connectionTimeout),
	})
	if err != nil {
		fmt.Println("Trying again")
		return h.Connect(deviceName)
	}

	h.device = device
	h.connectedDeviceName = found.LocalName()

	services, err := h.device.DiscoverServices([]bluetooth.UUID{bluetooth.ServiceUUIDHeartRate})
	if err != nil {
		return err
	}
	characteristics, err := services[0].DiscoverCharacteristics([]bluetooth.UUID{bluetooth.CharacteristicUUIDHeartRateMeasurement})
	if err != nil {
		return err
	}
	fmt.Println("Attaching notification")
	characteristics[0].EnableNotifications(func(buf []byte) {
		hr := uint(buf[1])
		for _, f := range h.listeners {
			f(float64(hr))
		}
	})

	return err
}

func (h *HrmService) ConnectedDeviceName() string {
	return h.connectedDeviceName
}

func (h *HrmService) Disconnect() error {
	fmt.Println("Disconnecting")
	var err error
	if h.device != nil {
		err = h.device.Disconnect()
	}
	h.device = nil
	h.connectedDeviceName = ""
	return err
}

func (h *HrmService) AddListeners(fs []func(float64)) {
	h.listeners = append(h.listeners, fs...)
}
