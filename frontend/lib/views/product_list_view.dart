// lib/views/product_list_view.dart

import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../services/api_service.dart';
import '../viewmodels/products_viewmodel.dart';
import 'product_details_view.dart';
import 'cart_view.dart';

class ProductListView extends StatefulWidget {
  const ProductListView({super.key});

  @override
  State<ProductListView> createState() => _ProductListViewState();
}

class _ProductListViewState extends State<ProductListView> {
  late TextEditingController _cartController;

  @override
  void initState() {
    super.initState();
    final viewModel = context.read<ProductsViewModel>();
    final initialCartId = viewModel.cartId.isNotEmpty ? viewModel.cartId : "";
    _cartController = TextEditingController(text: initialCartId);
    viewModel.setCartId(initialCartId);
  }

  @override
  void dispose() {
    _cartController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final viewModel = Provider.of<ProductsViewModel>(context);
    return Scaffold(
      appBar: AppBar(
        leading: Row(
                mainAxisSize: MainAxisSize.min,
                children: [
                  viewModel.isCheckedOut ? const Icon(Icons.lock, color: Colors.red) : const Icon(Icons.lock_open, color: Colors.green),
                  const SizedBox(width: 8),
                  Text(
                    viewModel.cartId,
                    style: const TextStyle(color: Colors.black),
                  ),
                ],
              ),
        title: const Text('Products'),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () async {
          await viewModel.loadCartItems();
          Navigator.push(
            context,
            MaterialPageRoute(builder: (_) => const CartView()),
          );
        },
        child: const Icon(Icons.shopping_cart),
      ),
      body: Column(
        children: [
          Padding(
            padding: const EdgeInsets.all(8.0),
            child: Row(
              children: [
                Expanded(
                  child: TextField(
                    controller: _cartController,
                    decoration: InputDecoration(
                      labelText: 'Cart ID',
                      border: const OutlineInputBorder(),
                      errorText: _cartController.text.isEmpty ? 'Required' : null,
                    ),
                  ),
                ),
                const SizedBox(width: 8),
                ElevatedButton(
                  onPressed: () async {
                    viewModel.setCartId(_cartController.text);
                    await viewModel.loadCartItems();
                  },
                  child: const Text('Confirm'),
                ),
              ],
            ),
          ),
          viewModel.isLoading
              ? const Expanded(
                  child: Center(child: CircularProgressIndicator()))
              : Expanded(
                  child: ListView.builder(
                    itemCount: viewModel.products.length,
                    itemBuilder: (context, index) {
                      final product = viewModel.products[index];
                      return ListTile(
                        leading: CircleAvatar(child: Text(product.id)),
                        title: Opacity(
                          opacity: product.available ? 1.0 : 0.5,
                          child: Text(product.name),
                        ),
                        subtitle: product.available
                            ? null
                            : const Text('Unavailable',
                                style: TextStyle(color: Colors.red)),
                        enabled: product.available,
                        onTap: product.available
                            ? () {
                                Navigator.push(
                                  context,
                                  MaterialPageRoute(
                                    builder: (context) =>
                                        ProductDetailsView(product: product),
                                  ),
                                );
                              }
                            : null,
                    trailing: !product.available
                        ? const Icon(Icons.lock, color: Colors.grey)
                        : (!viewModel.isCheckedOut
                            ? IconButton(
                                icon: const Icon(Icons.add_shopping_cart),
                                onPressed: () async {
                                  try {
                                    await viewModel.addProductToCart(product);
                                    ScaffoldMessenger.of(context).showSnackBar(
                                      const SnackBar(content: Text('Added to cart')),
                                    );
                                  } on ApiException catch (e) {
                                    ScaffoldMessenger.of(context).showSnackBar(
                                      SnackBar(content: Text(e.message)),
                                    );
                                  } catch (e) {
                                    ScaffoldMessenger.of(context).showSnackBar(
                                      SnackBar(content: Text('Error: $e')),
                                    );
                                  }
                                },
                              )
                            : null),
                      );
                    },
                  ),
                ),
        ],
      ),
    );
  }
}