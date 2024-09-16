package Util

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"time"
)

// GetCartFromCookie decodes the cookie value and returns a map of CartItems
func GetCartFromCookie(cookieValue string) map[string]BusinessObjects.Item {
	cart := make(map[string]BusinessObjects.Item)
	decodedBytes, _ := base64.StdEncoding.DecodeString(cookieValue)
	decodedString := string(decodedBytes)
	itemsList := strings.Split(decodedString, "|")

	for _, strItem := range itemsList {
		if strItem != "" {
			arrItemDetail := strings.Split(strItem, ",")
			ProductID := strings.TrimSpace(arrItemDetail[0])
			quantity, _ := strconv.Atoi(strings.TrimSpace(arrItemDetail[1]))

			item := BusinessObjects.Item{
				ProductID: ProductID,
				Quantity:  quantity,
			}
			cart[ProductID] = item
		}
	}

	return cart
}

// GetCookieByName retrieves a cookie by name from the request
func GetCookieByName(r *http.Request, cookieName string) *http.Cookie {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil
	}
	return cookie
}

// SaveCartToCookie saves the encoded cart string into a cookie
func SaveCartToCookie(w http.ResponseWriter, cartString string, userId string) {
	cookieName := "Cart_" + userId
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    cartString,
		MaxAge:   7200, // 120 minutes
		HttpOnly: true,
		Path:     "/",
	})
}

// DeleteCartToCookie deletes the cart cookie
func DeleteCartToCookie(w http.ResponseWriter, userId string) {
	cookieName := "Cart_" + userId
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Expires:  time.Now().Add(-24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})
}

// ConvertCartToString converts the list of CartItems to an encoded string
func ConvertCartToString(itemsList []BusinessObjects.Item) string {
	var strItemsInCart strings.Builder
	for _, item := range itemsList {
		strItemsInCart.WriteString(fmt.Sprintf("%s,%d|", item.ProductID, item.Quantity))
	}
	encodedString := base64.StdEncoding.EncodeToString([]byte(strItemsInCart.String()))
	return encodedString
}

// CookieNames returns a list of cookie names from the request
func CookieNames(r *http.Request) []string {
	var names []string
	for _, cookie := range r.Cookies() {
		names = append(names, cookie.Name)
	}
	return names
}
