// lib/services/api_service.dart

import 'dart:convert';
import 'package:http/http.dart' as http;
import '../models/product.dart';
import '../models/cart.dart';
import 'package:sqflite/sqflite.dart';
import 'package:path/path.dart';

import 'dart:io';

/// Exception thrown when an API call returns an error status code.
class ApiException implements Exception {
  final String message;
  final int statusCode;
  ApiException(this.message, this.statusCode);
  @override
  String toString() => 'ApiException($statusCode): $message';
}
  

class ApiService {
  String baseUrl = Platform.isAndroid ? 'http://10.0.2.2:8002/api/v1' : 'http://localhost:8002/api/v1';

  Future<Database>? _database;

  Future<Database> _getDatabase() async {
    if (_database != null) return _database!;
    String? dbDirectory;
    try {
      // Use the default database directory provided by sqflite
      dbDirectory = await getDatabasesPath();
    } catch (_) {
      dbDirectory = null;
    }
    final dbPath = dbDirectory != null
        ? join(dbDirectory, 'app.db')
        : inMemoryDatabasePath;
    _database = openDatabase(
      dbPath,
      version: 1,
      onCreate: (db, version) async {
        await db.execute(
          'CREATE TABLE IF NOT EXISTS products_cache(id INTEGER PRIMARY KEY, json TEXT, timestamp INTEGER)',
        );
      },
    );
    return _database!;
  }

  Future<List<Product>> fetchProducts() async {
    final db = await _getDatabase();
    final nowMs = DateTime.now().millisecondsSinceEpoch;
    final cached = await db.query(
      'products_cache',
      where: 'id = ?',
      whereArgs: [1],
    );
    if (cached.isNotEmpty) {
      final timestamp = cached.first['timestamp'] as int;
      if (nowMs - timestamp < 3 * 60 * 1000) {
        final jsonString = cached.first['json'] as String;
        final List<dynamic> productJson = json.decode(jsonString);
        print('list read from cache');
        return productJson.map((json) => Product.fromJson(json)).toList();
      }
    }
    final response = await http.get(
      Uri.parse('$baseUrl/products/all'),
    );
    if (response.statusCode == 200) {
      final jsonString = response.body;
      final List<dynamic> productJson = json.decode(jsonString);
      final products = productJson.map((json) => Product.fromJson(json)).toList();
      print('list read from api request');
      await db.insert(
        'products_cache',
        {'id': 1, 'json': jsonString, 'timestamp': nowMs},
        conflictAlgorithm: ConflictAlgorithm.replace,
      );
      return products;
    } else {
      final Map<String, dynamic> errorBody = json.decode(response.body);
      final errorMessage = errorBody['error'] ?? 'Failed to load products';
      throw ApiException(errorMessage, response.statusCode);
    }
  }

  Future<Cart> fetchCart(String cartId) async {
    final response = await http.get(
      Uri.parse('$baseUrl/carts/$cartId'),
    );
    if (response.statusCode == 200) {
      final Map<String, dynamic> jsonBody =
          json.decode(response.body) as Map<String, dynamic>;
      return Cart.fromJson(jsonBody);
    } else {
      final Map<String, dynamic> errorBody = json.decode(response.body);
      final errorMessage = errorBody['error'] ?? 'Failed to load cart';
      throw ApiException(errorMessage, response.statusCode);
    }
  }

  Future<void> addToCart(String productId, String cartId) async {
    final response = await http.post(
      Uri.parse('$baseUrl/carts/'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode({'product_id': productId, 'cart_id': cartId}),
    );
    if (response.statusCode != 204) {
      final Map<String, dynamic> errorBody = json.decode(response.body);
      final errorMessage = errorBody['error'] ?? 'Failed to add to cart';
      throw ApiException(errorMessage, response.statusCode);
    }
  }

  Future<void> removeFromCart(String productId, String cartId) async {
    final response = await http.delete(
      Uri.parse('$baseUrl/carts/'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode({'cart_id': cartId, 'product_id': productId}),
    );
    if (response.statusCode != 200 && response.statusCode != 204) {
      final Map<String, dynamic> errorBody = json.decode(response.body);
      final errorMessage = errorBody['error'] ?? 'Failed to remove from cart';
      throw ApiException(errorMessage, response.statusCode);
    }
  }

  Future<void> checkoutCart(String cartId) async {
    final response = await http.post(
      Uri.parse('$baseUrl/carts/checkout/$cartId'),
      headers: {'Content-Type': 'application/json'},
    );
    if (response.statusCode != 200 && response.statusCode != 204) {
      final Map<String, dynamic> errorBody = json.decode(response.body);
      final errorMessage = errorBody['error'] ?? 'Failed to checkout cart';
      throw ApiException(errorMessage, response.statusCode);
    }
  }
}