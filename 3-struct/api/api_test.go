package api_test

import (
	"3-struct/app/api"
	"3-struct/app/bins"
	"3-struct/app/storage"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

var fixtures = map[string]string{
	"create": "fixtures/create.json",
	"update": "fixtures/update.json",
}

type FakeStorage struct {
	BinList *bins.BinList
}

func newFakeStorage() *FakeStorage {
	return &FakeStorage{
		BinList: &bins.BinList{},
	}
}

func (fs *FakeStorage) ReadBins() (*bins.BinList, error) {
	return fs.BinList, nil
}

func (fs *FakeStorage) SaveBins(list storage.Serializable) (bool, error) {
	if bl, ok := list.(*bins.BinList); ok {
		fs.BinList = bl
		return true, nil
	}
	return false, fmt.Errorf("unsupported type %T", list)
}

type BrokenStorage struct{}

func (bs *BrokenStorage) ReadBins() (*bins.BinList, error) {
	return &bins.BinList{}, nil
}
func (bs *BrokenStorage) SaveBins(list storage.Serializable) (bool, error) {
	return false, fmt.Errorf("Ошибка сохранения")
}

func TestCreateBin(t *testing.T) {
	fakeStorage := newFakeStorage()

	data, err := api.CreateBin(fixtures["create"], "testBin", fakeStorage)

	if err != nil {
		t.Errorf("Возникла ошибка %v", err)
	}

	if len(data) == 0 {
		t.Errorf("Не пришли данные")
	}

	if len(fakeStorage.BinList.Bins) != 1 {
		t.Errorf("В хранилище должен быть 1 бин, фактически - %v", len(fakeStorage.BinList.Bins))
	}

	bin := fakeStorage.BinList.Bins[0]

	if bin.Name != "testBin" {
		t.Errorf("Ожидаемое имя testBin, фактическое %v", bin.Name)
	}
}

func TestCreateBinBadStorage(t *testing.T) {
	fakeStorage := &BrokenStorage{}
	_, err := api.CreateBin(fixtures["create"], "testBin", fakeStorage)

	if !errors.Is(err, api.Errors["storage"]) {
		t.Errorf("Ожидаемая ошибка %v, фактическая - %v", api.Errors["storage"], err)
	}
}

func TestCreateBinBadFile(t *testing.T) {
	fakeStorage := &BrokenStorage{}
	_, err := api.CreateBin("broken.txt", "testBin", fakeStorage)

	if !errors.Is(err, api.Errors["file"]) {
		t.Errorf("Ожидаемая ошибка %v, фактическая - %v", api.Errors["storage"], err)
	}
}

func TestUpdateBin(t *testing.T) {
	fakeStorage := newFakeStorage()

	api.CreateBin(fixtures["create"], "testBin", fakeStorage)
	bin := fakeStorage.BinList.Bins[0]

	data, err := api.UpdateBin(fixtures["update"], bin.Id)
	if err != nil {
		t.Errorf("Возникла ошибка %v", err)
	}

	if len(data) == 0 {
		t.Errorf("Не пришли данные")
	}

	var parsed map[string]any
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		t.Fatalf("Ошибка при разборе JSON: %v", err)
	}

	record, ok := parsed["record"].(map[string]any)

	if !ok {
		t.Fatalf("В ответе нет поля 'record'")
	}
	if val, ok := record["test"]; !ok {
		t.Errorf("В record нет поля 'test'")
	} else if v, ok := val.(float64); !ok || v != 1 {
		t.Errorf("Ожидаем test = 1, фактически %v", val)
	}
}

func TestUpdateBinIncorrectId(t *testing.T) {
	_, err := api.UpdateBin(fixtures["update"], "incorrect_id")

	if err != api.Errors["bad_response"] {
		t.Errorf("Ожидаемая ошибка %v, фактическая - %v", api.Errors["bad_response"], err)
	}
}

func TestUpdateBinBadFile(t *testing.T) {
	_, err := api.UpdateBin("broken.txt", "incorrect_id")

	if err != api.Errors["file"] {
		t.Errorf("Ожидаемая ошибка %v, фактическая - %v", api.Errors["file"], err)
	}
}

func TestGetBin(t *testing.T) {
	fakeStorage := newFakeStorage()

	api.CreateBin(fixtures["create"], "testBin", fakeStorage)
	bin := fakeStorage.BinList.Bins[0]

	data, err := api.GetBin(bin.Id)
	if err != nil {
		t.Errorf("Возникла ошибка %v", err)
	}

	if len(data) == 0 {
		t.Errorf("Не пришли данные")
	}
}

func TestGetBinIncorrectId(t *testing.T) {
	_, err := api.GetBin("no_id")

	if err != api.Errors["bad_response"] {
		t.Errorf("Ожидаемая ошибка %v, фактическая - %v", api.Errors["bad_response"], err)
	}
}

func TestDeleteBin(t *testing.T) {
	fakeStorage := newFakeStorage()

	api.CreateBin(fixtures["create"], "testBin", fakeStorage)
	bin := fakeStorage.BinList.Bins[0]

	data, err := api.DeleteBin(bin.Id, fakeStorage)
	if err != nil {
		t.Errorf("Возникла ошибка %v", err)
	}

	if len(data) == 0 {
		t.Errorf("Не пришли данные")
	}

	if len(fakeStorage.BinList.Bins) != 0 {
		t.Errorf("В хранилище должен быть 0 бин, фактически - %v", len(fakeStorage.BinList.Bins))
	}

}

func TestDeleteBinIncorrectId(t *testing.T) {
	fakeStorage := newFakeStorage()

	_, err := api.DeleteBin("no_id", fakeStorage)

	if err != api.Errors["bad_response"] {
		t.Errorf("Ожидаемая ошибка %v, фактическая - %v", api.Errors["bad_response"], err)
	}
}
