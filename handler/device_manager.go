package handlers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"signaturesign/crypto"
	"signaturesign/domain"
	"sync"
)

var (
	devices = make(map[string]*domain.Device)
	mu      sync.Mutex
)

func NewDevice(id, algorithm, label string) (*domain.Device, error) {
	var signer crypto.Signer
	var err error

	switch algorithm {
	case "RSA":
		signer, err = crypto.NewRSASigner()
	case "ECC":
		signer, err = crypto.NewECCSigner()
	default:
		return nil, errors.New("unsupported algorithm")
	}

	if err != nil {
		return nil, err
	}

	device := &domain.Device{
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

func GetAllDevices() []*domain.Device {
	mu.Lock()
	defer mu.Unlock()

	var allDevices []*domain.Device
	for _, device := range devices {
		allDevices = append(allDevices, device)
	}
	return allDevices
}

func GetDeviceByID(id string) (*domain.Device, error) {
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
