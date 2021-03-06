package woogo

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestFetchProducts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	products := make([]map[string]interface{}, 0)

	// Mock GET request to products
	httpmock.RegisterResponder("GET", "https://test.com/products.json",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, products)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
}

func TestPostProduct(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	products := make([]map[string]interface{}, 0)

	// Mock POST request to products
	httpmock.RegisterResponder("POST", "https://api.mybiz.com/products.json",
		func(req *http.Request) (*http.Response, error) {
			product := make(map[string]interface{})
			if err := json.NewDecoder(req.Body).Decode(&product); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}

			products = append(products, product)

			resp, err := httpmock.NewJsonResponse(200, product)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
}

/*
data := []byte{
  ```{
	  "id": 794,
	  "name": "Premium Quality",
	  "slug": "premium-quality-19",
	  "permalink": "https://example.com/product/premium-quality-19/",
	  "date_created": "2017-03-23T17:01:14",
	  "date_created_gmt": "2017-03-23T20:01:14",
	  "date_modified": "2017-03-23T17:01:14",
	  "date_modified_gmt": "2017-03-23T20:01:14",
	  "type": "simple",
	  "status": "publish",
	  "featured": false,
	  "catalog_visibility": "visible",
	  "description": "<p>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.</p>\n",
	  "short_description": "<p>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.</p>\n",
	  "sku": "",
	  "price": "21.99",
	  "regular_price": "21.99",
	  "sale_price": "",
	  "date_on_sale_from": null,
	  "date_on_sale_from_gmt": null,
	  "date_on_sale_to": null,
	  "date_on_sale_to_gmt": null,
	  "price_html": "<span class=\"woocommerce-Price-amount amount\"><span class=\"woocommerce-Price-currencySymbol\">&#36;</span>21.99</span>",
	  "on_sale": false,
	  "purchasable": true,
	  "total_sales": 0,
	  "virtual": false,
	  "downloadable": false,
	  "downloads": [],
	  "download_limit": -1,
	  "download_expiry": -1,
	  "external_url": "",
	  "button_text": "",
	  "tax_status": "taxable",
	  "tax_class": "",
	  "manage_stock": false,
	  "stock_quantity": null,
	  "in_stock": true,
	  "backorders": "no",
	  "backorders_allowed": false,
	  "backordered": false,
	  "sold_individually": false,
	  "weight": "",
	  "dimensions": {
	    "length": "",
	    "width": "",
	    "height": ""
	  },
	  "shipping_required": true,
	  "shipping_taxable": true,
	  "shipping_class": "",
	  "shipping_class_id": 0,
	  "reviews_allowed": true,
	  "average_rating": "0.00",
	  "rating_count": 0,
	  "related_ids": [
	    53,
	    40,
	    56,
	    479,
	    99
	  ],
	  "upsell_ids": [],
	  "cross_sell_ids": [],
	  "parent_id": 0,
	  "purchase_note": "",
	  "categories": [
	    {
	      "id": 9,
	      "name": "Clothing",
	      "slug": "clothing"
	    },
	    {
	      "id": 14,
	      "name": "T-shirts",
	      "slug": "t-shirts"
	    }
	  ],
	  "tags": [],
	  "images": [
	    {
	      "id": 792,
	      "date_created": "2017-03-23T14:01:13",
	      "date_created_gmt": "2017-03-23T20:01:13",
	      "date_modified": "2017-03-23T14:01:13",
	      "date_modified_gmt": "2017-03-23T20:01:13",
	      "src": "https://example.com/wp-content/uploads/2017/03/T_2_front-4.jpg",
	      "name": "",
	      "alt": "",
	      "position": 0
	    },
	    {
	      "id": 793,
	      "date_created": "2017-03-23T14:01:14",
	      "date_created_gmt": "2017-03-23T20:01:14",
	      "date_modified": "2017-03-23T14:01:14",
	      "date_modified_gmt": "2017-03-23T20:01:14",
	      "src": "https://example.com/wp-content/uploads/2017/03/T_2_back-2.jpg",
	      "name": "",
	      "alt": "",
	      "position": 1
	    }
	  ],
	  "attributes": [],
	  "default_attributes": [],
	  "variations": [],
	  "grouped_products": [],
	  "menu_order": 0,
	  "meta_data": [],
	  "_links": {
	    "self": [
	      {
	        "href": "https://example.com/wp-json/wc/v2/products/794"
	      }
	    ],
	    "collection": [
	      {
	        "href": "https://example.com/wp-json/wc/v2/products"
	      }
	    ]
	  }
	}
```
}
*/
