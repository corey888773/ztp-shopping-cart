// lib/viewmodels/products_viewmodel.dart

import 'package:flutter/material.dart';

import '../models/product.dart';
import '../services/api_service.dart';

class ProductsViewModel extends ChangeNotifier {
  final ApiService _apiService = ApiService();

  List<Product> _products = [];
  List<Product> get products => _products;

  bool _isLoading = false;
  bool get isLoading => _isLoading;

  ProductsViewModel() {
    loadProducts();
  }

  Future<void> loadProducts() async {
    _isLoading = true;
    notifyListeners();

    _products = await _apiService.fetchProducts();

    _isLoading = false;
    notifyListeners();
  }

  // Cart management
  String _cartId = '';
  String get cartId => _cartId;
  void setCartId(String id) {
    _cartId = id;
    notifyListeners();
  }

  List<Product> _cartItems = [];
  List<Product> get cartItems => _cartItems;

  /// Add a product to cart and lock it locally
  Future<void> addProductToCart(Product product) async {
    if (_cartId.isEmpty) return;
    try {
      await _apiService.addToCart(product.id, _cartId);
      // Mark product unavailable locally
      final idx = _products.indexWhere((p) => p.id == product.id);
      if (idx != -1) {
        _products[idx] = Product(
          id: product.id,
          name: product.name,
          description: product.description,
          available: false,
        );
      }
      notifyListeners();
    } on ApiException catch (e) {
      // Handle locked error by marking product unavailable
      if (e.message.contains('failed to lock product')) {
        final idx = _products.indexWhere((p) => p.id == product.id);
        if (idx != -1) {
          _products[idx] = Product(
            id: product.id,
            name: product.name,
            description: product.description,
            available: false,
          );
          notifyListeners();
        }
      }
      rethrow;
    }
  }

  /// Remove a product from cart and unlock it locally
  Future<void> removeProductFromCart(Product product) async {
    if (_cartId.isEmpty) return;
    try {
      await _apiService.removeFromCart(product.id, _cartId);
      // Remove from local cart items
      _cartItems.removeWhere((p) => p.id == product.id);
      // Mark product available in product list
      final idx = _products.indexWhere((p) => p.id == product.id);
      if (idx != -1) {
        _products[idx] = Product(
          id: product.id,
          name: product.name,
          description: product.description,
          available: true,
        );
      }
      notifyListeners();
    } catch (e) {
      rethrow;
    }
  }

  bool _isCheckedOut = false;
  bool get isCheckedOut => _isCheckedOut;

  /// Load cart items and checkout status
  Future<void> loadCartItems() async {
    if (_cartId.isEmpty) return;
    _isLoading = true;
    notifyListeners();

    try {
      final cart = await _apiService.fetchCart(_cartId);
      _cartItems = cart.products;
      _isCheckedOut = cart.isCheckedOut;
    } on ApiException catch (e) {
      if (e.message.contains('no products found') || e.message.contains('invalid query')) {
        _cartItems = [];
      }
      // For other errors, rethrow to be handled by UI
      rethrow;
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }
/// Checkout the cart
  Future<void> checkoutCart() async {
    if (_cartId.isEmpty) return;
    _isLoading = true;
    notifyListeners();

    try {
      await _apiService.checkoutCart(_cartId);
      _isCheckedOut = true;
    } on ApiException catch (e) {
      // Handle already checked out error
      if (e.message.contains('already checked out')) {
        _isCheckedOut = true;
        notifyListeners();
      }
      rethrow;
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }
}
