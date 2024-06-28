package utils

import "github.com/gofrs/uuid"

func NewUUID() uuid.UUID {
	guid, _ := uuid.NewV4()
	return guid
}

func UuidFromString(strUid string) (uuid.UUID, error) {
	uid, err := uuid.FromString(strUid)
	if err != nil {
		return uid, err
	}
	return uuid.Must(uid, err), nil
}
