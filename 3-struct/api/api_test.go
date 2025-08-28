package api_test

import (
	"3-struct/app/api"
	"3-struct/app/bins"
	"3-struct/app/storage"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	require.NoError(t, err)
	require.NotEmpty(t, data)
	require.Len(t, fakeStorage.BinList.Bins, 1)

	bin := fakeStorage.BinList.Bins[0]

	assert.Equal(t, "testBin", bin.Name)
}

func TestCreateBinBadStorage(t *testing.T) {
	fakeStorage := &BrokenStorage{}
	_, err := api.CreateBin(fixtures["create"], "testBin", fakeStorage)

	assert.ErrorIs(t, err, api.Errors["storage"])
}

func TestCreateBinBadFile(t *testing.T) {
	fakeStorage := &BrokenStorage{}
	_, err := api.CreateBin("broken.txt", "testBin", fakeStorage)

	assert.ErrorIs(t, err, api.Errors["file"])
}

func TestUpdateBin(t *testing.T) {
	fakeStorage := newFakeStorage()

	api.CreateBin(fixtures["create"], "testBin", fakeStorage)
	bin := fakeStorage.BinList.Bins[0]

	data, err := api.UpdateBin(fixtures["update"], bin.Id)

	require.NoError(t, err)
	require.NotEmpty(t, data)

	var parsed map[string]any
	err = json.Unmarshal(data, &parsed)

	require.NoError(t, err)

	record, ok := parsed["record"].(map[string]any)

	require.True(t, ok)
	val, ok := record["test"]

	require.True(t, ok)
	v, ok := val.(float64)
	require.True(t, ok)
	require.Equal(t, 1.0, v)
}

func TestUpdateBinIncorrectId(t *testing.T) {
	_, err := api.UpdateBin(fixtures["update"], "incorrect_id")

	assert.ErrorIs(t, err, api.Errors["bad_response"])
}

func TestUpdateBinBadFile(t *testing.T) {
	_, err := api.UpdateBin("broken.txt", "incorrect_id")

	assert.ErrorIs(t, err, api.Errors["file"])
}

func TestGetBin(t *testing.T) {
	fakeStorage := newFakeStorage()

	api.CreateBin(fixtures["create"], "testBin", fakeStorage)
	bin := fakeStorage.BinList.Bins[0]

	data, err := api.GetBin(bin.Id)

	require.NoError(t, err)
	require.NotEmpty(t, data)
}

func TestGetBinIncorrectId(t *testing.T) {
	_, err := api.GetBin("no_id")

	assert.ErrorIs(t, err, api.Errors["bad_response"])
}

func TestDeleteBin(t *testing.T) {
	fakeStorage := newFakeStorage()

	api.CreateBin(fixtures["create"], "testBin", fakeStorage)
	bin := fakeStorage.BinList.Bins[0]

	data, err := api.DeleteBin(bin.Id, fakeStorage)

	require.NoError(t, err)
	require.NotEmpty(t, data)
	require.Empty(t, fakeStorage.BinList.Bins)
}

func TestDeleteBinIncorrectId(t *testing.T) {
	fakeStorage := newFakeStorage()

	_, err := api.DeleteBin("no_id", fakeStorage)

	assert.ErrorIs(t, err, api.Errors["bad_response"])
}
