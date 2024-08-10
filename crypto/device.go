package crypto

import (
	"encoding/base64"
	"errors"
	"fmt"
	"sync"
)

type Device struct {
	ID               string `json:"ID"`
	Algorithm        string `json:"Algorithm"`
	Label            string `json:"Label"`
	PublicKey        string `json:"PublicKey"`
	SignatureCounter int    `json:"SignatureCounter"`
	LastSignature    string `json:"LastSignature"`
	Signer           Signer `json:"Signer"`
}

var (
	devices = make(map[string]*Device)
	mu      sync.Mutex
)

func NewDevice(id, algorithm, label string) (*Device, error) {
	var signer Signer
	var err error

	switch algorithm {
	case "RSA":
		signer, err = NewRSASigner()
	case "ECC":
		signer, err = NewECCSigner()
	default:
		return nil, errors.New("unsupported algorithm")
	}

	if err != nil {
		return nil, err
	}

	device := &Device{
		ID:        id,
		Algorithm: algorithm,
		Label:     label,
		Signer:    signer,
		PublicKey: signer.PublicKey(),
	}

	mu.Lock()
	devices[id] = device
	mu.Unlock()

	return device, nil
}

func GetAllDevices() []*Device {
	mu.Lock()
	defer mu.Unlock()

	var allDevices []*Device
	for _, device := range devices {
		allDevices = append(allDevices, device)
	}
	return allDevices
}

func GetDeviceByID(id string) (*Device, error) {
	mu.Lock()
	defer mu.Unlock()

	device, exists := devices[id]
	if !exists {
		return nil, errors.New("device not found")
	}
	return device, nil
}
func SignTransaction(deviceID, data string) (string, string, error) {
	mu.Lock()
	defer mu.Unlock()

	device, exists := devices[deviceID]
	if !exists {
		return "", "", errors.New("device not found")
	}

	// Use the last signature if counter is not zero
	var lastSignature string
	if device.SignatureCounter == 0 {
		lastSignature = base64.StdEncoding.EncodeToString([]byte(deviceID))
	} else {
		lastSignature = device.LastSignature
	}

	// Create the data to be signed
	securedData := fmt.Sprintf("%d_%s_%s", device.SignatureCounter, data, lastSignature)
	signature, err := device.Signer.Sign(securedData)
	if err != nil {
		return "", "", err
	}

	// Update the device with the new signature and increment the counter
	device.LastSignature = signature
	device.SignatureCounter++

	return signature, securedData, nil
}
