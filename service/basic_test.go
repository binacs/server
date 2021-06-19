package service

import (
	"testing"
)

func TestBasicServiceImpl_ServeHome(t *testing.T) {
	tests := []struct {
		name string
		bs   *BasicServiceImpl
		want string
	}{
		{
			name: "normal",
			bs:   &BasicServiceImpl{},
			want: homeBody,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BasicServiceImpl{}
			if got := bs.ServeHome(); got != tt.want {
				t.Errorf("BasicServiceImpl.ServeHome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicServiceImpl_ServeToys(t *testing.T) {
	tests := []struct {
		name string
		bs   *BasicServiceImpl
		want string
	}{
		{
			name: "normal",
			bs:   &BasicServiceImpl{},
			want: toysBody,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BasicServiceImpl{}
			if got := bs.ServeToys(); got != tt.want {
				t.Errorf("BasicServiceImpl.ServeToys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicServiceImpl_ServeCrypto(t *testing.T) {
	tests := []struct {
		name string
		bs   *BasicServiceImpl
		want string
	}{
		{
			name: "normal",
			bs:   &BasicServiceImpl{},
			want: cryptoBody,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BasicServiceImpl{}
			if got := bs.ServeCrypto(); got != tt.want {
				t.Errorf("BasicServiceImpl.ServeCrypto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicServiceImpl_ServeTinyURL(t *testing.T) {
	tests := []struct {
		name string
		bs   *BasicServiceImpl
		want string
	}{
		{
			name: "normal",
			bs:   &BasicServiceImpl{},
			want: tinyurlBody,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BasicServiceImpl{}
			if got := bs.ServeTinyURL(); got != tt.want {
				t.Errorf("BasicServiceImpl.ServeTinyURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicServiceImpl_ServePastebin(t *testing.T) {
	tests := []struct {
		name string
		bs   *BasicServiceImpl
		want string
	}{
		{
			name: "normal",
			bs:   &BasicServiceImpl{},
			want: pastebinBody,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BasicServiceImpl{}
			if got := bs.ServePastebin(); got != tt.want {
				t.Errorf("BasicServiceImpl.ServePastebin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicServiceImpl_ServeStorage(t *testing.T) {
	tests := []struct {
		name string
		bs   *BasicServiceImpl
		want string
	}{
		{
			name: "normal",
			bs:   &BasicServiceImpl{},
			want: storageBody,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BasicServiceImpl{}
			if got := bs.ServeStorage(); got != tt.want {
				t.Errorf("BasicServiceImpl.ServeStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicServiceImpl_ServeAbout(t *testing.T) {
	tests := []struct {
		name string
		bs   *BasicServiceImpl
		want string
	}{
		{
			name: "normal",
			bs:   &BasicServiceImpl{},
			want: aboutBody,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BasicServiceImpl{}
			if got := bs.ServeAbout(); got != tt.want {
				t.Errorf("BasicServiceImpl.ServeAbout() = %v, want %v", got, tt.want)
			}
		})
	}
}
