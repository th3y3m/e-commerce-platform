package Util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"time"
)

// GetCartFromCookie decodes the cookie value and returns a map of CartItems, with error handling
func GetCartFromCookie(cookieValue string) (map[string]BusinessObjects.Item, error) {
	cart := make(map[string]BusinessObjects.Item)
	decodedBytes, err := base64.StdEncoding.DecodeString(cookieValue)
	if err != nil {
		return nil, fmt.Errorf("failed to decode cookie value: %w", err)
	}
	decodedString := string(decodedBytes)
	itemsList := strings.Split(decodedString, "|")

	for _, strItem := range itemsList {
		if strItem != "" {
			arrItemDetail := strings.Split(strItem, ",")
			if len(arrItemDetail) < 2 {
				return nil, errors.New("invalid cart item format")
			}
			ProductID := strings.TrimSpace(arrItemDetail[0])
			quantity, err := strconv.Atoi(strings.TrimSpace(arrItemDetail[1]))
			if err != nil {
				return nil, fmt.Errorf("failed to parse quantity: %w", err)
			}

			item := BusinessObjects.Item{
				ProductID: ProductID,
				Quantity:  quantity,
			}
			cart[ProductID] = item
		}
	}

	return cart, nil
}

// GetCookieByName retrieves a cookie by name from the request and returns an error if not found
func GetCookieByName(r *http.Request, cookieName string) (*http.Cookie, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, fmt.Errorf("cookie %s not found: %w", cookieName, err)
	}
	return cookie, nil
}

// SaveCartToCookie saves the encoded cart string into a cookie
func SaveCartToCookie(w http.ResponseWriter, cartString string, userId string) error {
	if cartString == "" {
		return errors.New("cart string is empty")
	}
	cookieName := "Cart_" + userId
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    cartString,
		MaxAge:   7200, // 120 minutes
		HttpOnly: true,
		Path:     "/",
	})
	return nil
}

// DeleteCartToCookie deletes the cart cookie
func DeleteCartToCookie(w http.ResponseWriter, userId string) error {
	if userId == "" {
		return errors.New("user ID is empty")
	}
	cookieName := "Cart_" + userId
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Expires:  time.Now().Add(-24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})
	return nil
}

// ConvertCartToString converts the list of CartItems to an encoded string
func ConvertCartToString(itemsList []BusinessObjects.Item) (string, error) {
	if len(itemsList) == 0 {
		return "", errors.New("cart is empty")
	}

	var strItemsInCart strings.Builder
	for _, item := range itemsList {
		if item.ProductID == "" || item.Quantity < 0 {
			return "", errors.New("invalid product ID or quantity in cart item")
		}
		strItemsInCart.WriteString(fmt.Sprintf("%s,%d|", item.ProductID, item.Quantity))
	}

	encodedString := base64.StdEncoding.EncodeToString([]byte(strItemsInCart.String()))
	return encodedString, nil
}

// CookieNames returns a list of cookie names from the request
func CookieNames(r *http.Request) ([]string, error) {
	if r == nil {
		return nil, errors.New("request is nil")
	}

	var names []string
	for _, cookie := range r.Cookies() {
		names = append(names, cookie.Name)
	}
	if len(names) == 0 {
		return nil, errors.New("no cookies found")
	}
	return names, nil
}
