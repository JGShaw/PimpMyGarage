package services

import (
	"fmt"
	"time"
	"tinygo.org/x/bluetooth"
)

const scanTimeout = 10 * time.Second
const connectionTimeout = 5 * time.Second

type HrmService struct {
	adapter             *bluetooth.Adapter
	scanResults         map[string]bluetooth.ScanResult
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

func (h *HrmService) Scan() ([]bluetooth.ScanResult, error) {
	h.scanResults = map[string]bluetooth.ScanResult{}

	go func() {
		<-time.After(scanTimeout)
		h.adapter.StopScan()
	}()

	err := h.adapter.Scan(func(adapter *bluetooth.Adapter, scanResult bluetooth.ScanResult) {
		if scanResult.HasServiceUUID(bluetooth.ServiceUUIDHeartRate) {
			h.scanResults[scanResult.Address.String()] = scanResult
			return
		}
	})
	if err != nil {
		return nil, err
	}

	values := make([]bluetooth.ScanResult, len(h.scanResults))
	i := 0
	for _, v := range h.scanResults {
		values[i] = v
		i++
	}
	return values, nil
}

func (h *HrmService) Connect(address string) error {
	scanResult := h.scanResults[address]
	device, err := h.adapter.Connect(scanResult.Address, bluetooth.ConnectionParams{
		ConnectionTimeout: bluetooth.NewDuration(connectionTimeout),
	})
	if err != nil {
		fmt.Println("Here")
		return err
	}

	h.device = device
	h.connectedDeviceName = scanResult.LocalName()

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
