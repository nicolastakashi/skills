package config

import "testing"

func TestParsePortValid(t *testing.T) {
	p, err := ParsePort("8080")
	if err != nil || p != 8080 {
		t.Fatalf("ParsePort(8080) = %d, %v", p, err)
	}
}

func TestParsePortRejectsZero(t *testing.T) {
	if _, err := ParsePort("0"); err == nil {
		t.Fatal("ParsePort(0) should fail")
	}
}

func TestFromRawCopiesHostAndPort(t *testing.T) {
	cfg, err := FromRaw(Raw{Host: "localhost", Port: "8080"})
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Host != "localhost" || cfg.Port != 8080 {
		t.Fatalf("FromRaw = %+v", cfg)
	}
}
