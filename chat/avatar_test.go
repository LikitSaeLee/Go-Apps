package main

import (
	"testing"
	"io/ioutil"
	"path"
	"os"
	gomniauthtest "github.com/stretchr/gomniauth/test"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatarURL)
	testChatUser := &chatUser{User: testUser}
	url, err := authAvatar.GetAvatarURL(testChatUser)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no value present")
	}
	// set a value
	testUrl := "http://url-to-gravatar/"
	testUser := &gomniauthtest.TestUser{}
	testChatUser.User = testUser
	url, err = authAvatar.GetAvatarURL(testChatUser)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should return no error when value present")
	} else {
		if url != testUrl {
			t.Error("AuthAvatar.GetAvatarURL should return correct url")
		}
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvitar GravatarAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := gravatarAvitar.GetAvatarURL(user)
	if err != nil {
		t.Error("GravatarAvitar.GetAvatarURL should not return an error")
	}
	if url != "//www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346" {
	t.Errorf("GravatarAvitar.GetAvatarURL wrongly return %s", url)
}
}

func TestFileSystemAvatar(t *testing.T) {
	// make a test avatar file
	filename := path.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer func() { os.Remove(filename) }()

	var fileSystemAvatar FileSystemAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("FileSystemAvatar.GetAvatarURL should not return an error")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.GetAvatarURL wrongly returned %s", url)
	}
}
