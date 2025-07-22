// lib/views/cart_view.dart

import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../viewmodels/products_viewmodel.dart';
import 'product_details_view.dart';
import '../services/api_service.dart';

class CartView extends StatelessWidget {
  const CartView({super.key});

  @override
  Widget build(BuildContext context) {
    final viewModel = Provider.of<ProductsViewModel>(context);
    return Scaffold(
      appBar: AppBar(
        leading: viewModel.isCheckedOut
            ? const Icon(Icons.lock, color: Colors.red)
            : null,
        title: const Text('Cart'),
      ),
      body: viewModel.isLoading
          ? const Center(child: CircularProgressIndicator())
          : Column(
              children: [
                if (viewModel.isCheckedOut)
                  Container(
                    color: Colors.red,
                    width: double.infinity,
                    padding: const EdgeInsets.all(8.0),
                    child: const Text(
                      'This cart has been checked out',
                      style: TextStyle(color: Colors.white),
                      textAlign: TextAlign.center,
                    ),
                  ),
                Expanded(
                  child: viewModel.cartItems.isEmpty
                      ? const Center(child: Text('Cart is empty'))
                      : ListView.builder(
                          itemCount: viewModel.cartItems.length,
                          itemBuilder: (context, index) {
                            final product = viewModel.cartItems[index];
                            return ListTile(
                              leading: CircleAvatar(child: Text(product.id)),
                              title: Text(product.name),
                              onTap: viewModel.isCheckedOut
                                  ? null
                                  : () {
                                      Navigator.push(
                                        context,
                                        MaterialPageRoute(builder: (_) => ProductDetailsView(product: product)),
                                      );
                                    },
                              trailing: viewModel.isCheckedOut
                                  ? null
                                  : IconButton(
                                      icon: const Icon(Icons.remove_circle),
                                      onPressed: () async {
                                        try {
                                          await viewModel.removeProductFromCart(product);
                                          ScaffoldMessenger.of(context).showSnackBar(
                                            const SnackBar(content: Text('Removed from cart')),
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
                                    ),
                            );
                          },
                        ),
                ),
                if (!viewModel.isCheckedOut && viewModel.cartItems.isNotEmpty)
                  Padding(
                    padding: const EdgeInsets.all(8.0),
                    child: ElevatedButton.icon(
                      icon: const Icon(Icons.check, color: Colors.black),
                      label: const Text('Checkout'),
                      style: ElevatedButton.styleFrom(minimumSize: const Size.fromHeight(48)),
                      onPressed: () async {
                        try {
                          await viewModel.checkoutCart();
                          ScaffoldMessenger.of(context).showSnackBar(
                            const SnackBar(content: Text('Cart checked out')),
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
                    ),
                  ),
              ],
            ),
    );
  }
}