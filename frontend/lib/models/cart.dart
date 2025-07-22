// lib/models/cart.dart

import 'product.dart';

class Cart {
  final String cartId;
  final List<Product> products;
  final bool isCheckedOut;

  Cart({
    required this.cartId,
    required this.products,
    required this.isCheckedOut,
  });

  factory Cart.fromJson(Map<String, dynamic> json) {
    final productsJson = json['products'] as List<dynamic>? ?? [];
    return Cart(
      cartId: json['cart_id'] as String,
      products: productsJson
          .map((item) => Product.fromJson(item as Map<String, dynamic>))
          .toList(),
      isCheckedOut: json['is_checked_out'] as bool? ?? false,
    );
  }
}