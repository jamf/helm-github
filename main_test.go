package main

import "testing"

func TestAssetMatches(t *testing.T) {
	tests := []struct {
		name  string
		asset string
		uri   string
		want  bool
	}{
		{"same repository", "dummy-app-0.1.1.tgz", "github.com/wandera/kubernetes-descriptors/releases/download/dummy-app-0.1.1/dummy-app-0.1.1.tgz", true},
		{"index keeps the old owner after a transfer", "dummy-app-0.1.1.tgz", "github.com/wandera/kubernetes-descriptors-test/releases/download/dummy-app-0.1.1/dummy-app-0.1.1.tgz", true},
		{"different file", "dummy-app-0.1.0.tgz", "github.com/jamf/kubernetes-descriptors/releases/download/dummy-app-0.1.1/dummy-app-0.1.1.tgz", false},
		{"prefix is not a match", "app-0.1.1.tgz", "github.com/jamf/kubernetes-descriptors/releases/download/dummy-app-0.1.1/dummy-app-0.1.1.tgz", false},
		{"empty asset name", "", "github.com/jamf/kubernetes-descriptors/releases/download/dummy-app-0.1.1/", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assetMatches(tt.asset, tt.uri); got != tt.want {
				t.Errorf("assetMatches(%q, %q) = %v, want %v", tt.asset, tt.uri, got, tt.want)
			}
		})
	}
}

func TestCanonicalOwnerRepository(t *testing.T) {
	tests := []struct {
		url, owner, repo string
	}{
		{"https://github.com/jamf/kubernetes-descriptors/releases/download/a-1.0.0/a-1.0.0.tgz", "jamf", "kubernetes-descriptors"},
		{"https://example.com/jamf/kubernetes-descriptors/releases/download/a-1.0.0/a-1.0.0.tgz", "", ""},
		{"https://github.com/jamf", "", ""},
		{"", "", ""},
	}
	for _, tt := range tests {
		if o, r := canonicalOwnerRepository(tt.url); o != tt.owner || r != tt.repo {
			t.Errorf("canonicalOwnerRepository(%q) = %q, %q; want %q, %q", tt.url, o, r, tt.owner, tt.repo)
		}
	}
}
