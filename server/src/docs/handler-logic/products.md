# Create Product

- Check user group (IMS User, Operations)
- Trim whitespaces in products fields
    - product_name, product_sku, product_colour, product_category, product_brand, size_name
- Products form validation
    - product_name
        - if empty, throw error to user that product name cannot be empty.
        - product_name cannot have more than 255 characters.
    - product_sku
        - if empty, throw error to user that product sku cannot be empty.
        - product_sku cannot have more than 50 characters.
    - brand_name
        - if empty, throw error to user that product brand cannot be empty.
        - brand_name cannot have more than 60 characters.
    - colour_name
        - colour_name cannot have more than 60 characters.
    - category_name
        - category_name cannot have more than 60 characters.
    - product_cost
        - check that product_cost has the format DECIMAL(10,2)
- Check size_name
    - size_name cannot have more than 5 characters.
    - size quantity cannot be negative or floating number.
    - this check should be done in a for loop together with trimspace
- Retrieve username from cookie
- Get the Organisation Name from the username (using SQL Query)
- Check if user belongs to organisation "InvenNexus" (duplicate product_sku not allowed)
    - If Yes:
        - check the product_sku of the product based on the **user_id** to see if it has already been taken.
    - If No:
        - check the product_sku of the product based on the **organisation_id** to see if it has already been taken.
        - to ensure that different users in the same organisation see the same products.
- Insert product to `products` table. (return product_id).
- Check if user has input for size_name and quantity
    - If yes:
        - Insert into `sizes` table.
        - Insert into `product_sizes_mapping` table.
- Insert into `product_user_organisation_mapping` table.
- Success message on inserting product.

