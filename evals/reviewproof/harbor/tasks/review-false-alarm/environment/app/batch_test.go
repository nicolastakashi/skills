package batch

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestChunkEvenSplit(t *testing.T) {
	got := Chunk([]string{"a", "b", "c", "d"}, 2)
	want := [][]string{{"a", "b"}, {"c", "d"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Chunk = %v, want %v", got, want)
	}
}

func TestFetchAllKeepsOrder(t *testing.T) {
	ids := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	got, err := FetchAll(ids, func(id string) (string, error) {
		return strings.ToUpper(id), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("FetchAll = %v, want %v", got, want)
	}
}

func TestFetchAllReturnsError(t *testing.T) {
	boom := errors.New("boom")
	_, err := FetchAll([]string{"a", "b"}, func(id string) (string, error) {
		if id == "b" {
			return "", boom
		}
		return id, nil
	})
	if !errors.Is(err, boom) {
		t.Fatalf("err = %v, want %v", err, boom)
	}
}
