package products

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/external/products"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/add_to_cart"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/checkout"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/get_cart"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/remove_from_cart"
)

const ContentTypeHeaders = "Content-Type"
const JSONContentType = "application/json"

var _ add_to_cart.ProductsService = (*Client)(nil)
var _ remove_from_cart.ProductsService = (*Client)(nil)
var _ get_cart.ProductsService = (*Client)(nil)
var _ checkout.ProductsService = (*Client)(nil)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseUrl string) *Client {
	return &Client{
		baseURL:    baseUrl,
		httpClient: &http.Client{},
	}
}

type LockProductRequest struct {
	ProductID string `json:"product_id"`
	CartID    string `json:"cart_id"`
}

func (c *Client) LockProduct(productID string, cartID string) error {
	requestBody := LockProductRequest{
		ProductID: productID,
		CartID:    cartID,
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(requestBody); err != nil {
		return nil
	}
	headers := map[string]string{
		ContentTypeHeaders: JSONContentType,
	}
	resp, err := c.doRequest(http.MethodPost, "/api/v1/products/lock", headers, buf)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusNoContent {
		return errors.New("failed to lock product")
	}

	return nil
}

type UnlockProductRequest struct {
	ProductID string `json:"product_id"`
	CartID    string `json:"cart_id"`
}

func (c *Client) UnlockProduct(productID string, cartID string) error {
	requestBody := UnlockProductRequest{
		ProductID: productID,
		CartID:    cartID,
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(requestBody); err != nil {
		return nil
	}
	headers := map[string]string{
		ContentTypeHeaders: JSONContentType,
	}
	resp, err := c.doRequest(http.MethodPost, "/api/v1/products/unlock", headers, buf)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusNoContent {
		return errors.New("failed to unlock product")
	}

	return nil
}

type CheckoutProductsRequest struct {
	ProductIDs []string `json:"product_ids"`
	CartID     string   `json:"cart_id"`
}

func (c *Client) CheckoutProducts(productIDs []string, cartID string) error {
	requestBody := CheckoutProductsRequest{
		ProductIDs: productIDs,
		CartID:     cartID,
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(requestBody); err != nil {
		return nil
	}
	headers := map[string]string{
		ContentTypeHeaders: JSONContentType,
	}
	resp, err := c.doRequest(http.MethodPost, "/api/v1/products/checkout", headers, buf)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusNoContent {
		return errors.New("failed to checkout products")
	}

	return nil
}

type GetProductsByIDsRequest struct {
	ProductIDs []string `json:"product_ids"`
}

func (c *Client) GetProductsByIDs(productIDs []string) ([]products.Product, error) {
	requestBody := GetProductsByIDsRequest{
		ProductIDs: productIDs,
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(requestBody); err != nil {
		return nil, err
	}
	headers := map[string]string{
		ContentTypeHeaders: JSONContentType,
	}
	resp, err := c.doRequest(http.MethodPost, "/api/v1/products", headers, buf)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("RequestBody: %+v\n", requestBody)
		bodyString, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		fmt.Printf("ResponseBody: %s\n", bodyString)

		fmt.Printf("Code: %d\n", resp.StatusCode)
		return nil, errors.New("failed to get products")
	}

	if resp.Body == nil {
		return nil, errors.New("response body is nil")
	}

	var productsResponse []products.Product
	if err := json.NewDecoder(resp.Body).Decode(&productsResponse); err != nil {
		return nil, err
	}

	return productsResponse, nil
}

func (c *Client) doRequest(method, path string, headers map[string]string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, c.baseURL+path, body)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return c.httpClient.Do(req)
}
