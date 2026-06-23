package handlers

import (
	"github.com/edkariuki/url-shortener/codec"
	"github.com/edkariuki/url-shortener/store"
)

func Shorten(rawURL string) (string, error) {
    // 1. validate first, reject if not https
    if err := store.Validate(rawURL); err != nil {
        return "", err
    }

    // 2. check if already shortened
    id, err := store.Store.FindByURL(rawURL)
    if err != nil {
        // 3. not found, save it
        id, err = store.Save(rawURL)
        if err != nil {
            return "", err
        }
    }

    // 4. encode and return
    code := codec.Encode(uint64(id))
    return "https://short.ly/" + code, nil
}