// lib/models/product.dart

class Product {
   final String id;
   final String name;
   final String description;
   final bool available;

  Product({
    required this.id,
    required this.name,
    required this.description,
    required this.available,
  });

  factory Product.fromJson(Map<String, dynamic> json) {
    return Product(
      id: json['product_id'],
      name: json['name'],
      description: json['description'],
      available: (json['is_available'] is bool) ? json['is_available'] as bool : true,
    );
  }
}
