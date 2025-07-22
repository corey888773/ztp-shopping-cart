// lib/main.dart

import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'viewmodels/products_viewmodel.dart';
import 'views/product_list_view.dart';

void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    // The ChangeNotifierProvider creates and provides the ViewModel to all
    // widgets below it in the tree. This is the core of MVVM with Provider.
    return ChangeNotifierProvider(
      create: (context) => ProductsViewModel(),
      child: MaterialApp(
        title: 'Product Shop',
        theme: ThemeData(
          primarySwatch: Colors.blue,
          visualDensity: VisualDensity.adaptivePlatformDensity,
        ),
        home: const ProductListView(),
      ),
    );
  }
}