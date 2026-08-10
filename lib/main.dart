import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

/// รากของแอป: ตั้งค่าธีมและกำหนดหน้าเริ่มต้น
class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'แอป 101 ของฉัน',
      theme: ThemeData(
        colorScheme: .fromSeed(seedColor: Colors.green),
      ),
      home: const MyHomePage(title: 'แอป 101 ของฉัน'),
    );
  }
}

/// หน้าหลัก — เป็น StatefulWidget เพราะมีค่าตัวนับที่เปลี่ยนได้
class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;

  // ทุกครั้งที่แก้ค่าที่แสดงบนจอ ต้องห่อใน setState เพื่อสั่งให้ build() วาดใหม่
  void _incrementCounter() {
    setState(() => _counter++);
  }

  void _resetCounter() {
    setState(() => _counter = 0);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(widget.title),
        actions: [
          IconButton(
            icon: const Icon(Icons.refresh),
            tooltip: 'รีเซ็ต',
            onPressed: _resetCounter,
          ),
        ],
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: .center,
          children: [
            const Text('คุณกดปุ่มไปทั้งหมด:'),
            Text(
              '$_counter ครั้ง',
              style: Theme.of(context).textTheme.headlineMedium,
            ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'เพิ่ม',
        child: const Icon(Icons.add),
      ),
    );
  }
}
