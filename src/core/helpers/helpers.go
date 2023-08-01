package helpers

import db "agora-vnf-manager/db"

type DAO[DataType any] interface {
	Insert(DataType) (DataType, int, error)
	Get(DataType) (DataType, bool, error)
	Find(DataType) ([]DataType, int, error)
	Update(DataType) (DataType, int, error)
	Remove(DataType) (int, error)
}

func GetEntry[DataType any](session db.IDbSession, ids DataType) (DataType, bool, error) {
	results := []DataType{}
	if err := session.FindOne(&results, &ids); err != nil {
		return ids, false, err
	} else if len(results) == 0 {
		return ids, false, nil
	}
	return results[0], true, nil
}

func ExistsEntry[DataType any](session db.IDbSession, ids DataType) (bool, error) {
	results := []DataType{}
	if err := session.FindOne(&results, &ids); err != nil {
		return false, err
	} else if len(results) == 0 {
		return false, nil
	}
	return true, nil
}

func FindEntry[DataType any](session db.IDbSession, ids DataType) ([]DataType, error) {
	results := []DataType{}
	if err := session.Find(&results, &ids); err != nil {
		return results, err
	} else {
		return results, nil
	}
}

func DeleteEntry[DataType any](session db.IDbSession, ids DataType) (DataType, error) {
	var placeholder DataType
	if err := session.Delete(&placeholder, &ids); err != nil {
		return ids, err
	} else {
		ids = placeholder
		return ids, nil
	}
}

func ReplaceEntry[DataType any](session db.IDbSession, newData DataType, ids ...any) (DataType, error) {
	if err := session.Update(&newData); err != nil {
		return newData, err
	} else {
		return newData, nil
	}
}

func CreateEntry[DataType any](session db.IDbSession, data DataType) (DataType, error) {
	if err := session.Create(&data); err != nil {
		return data, err
	} else {
		return data, nil
	}
}
