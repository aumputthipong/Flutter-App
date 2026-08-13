import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

import 'package:flutter_app/main.dart';

void main() {
  testWidgets('กดปุ่ม + แล้วตัวนับเพิ่มขึ้น', (tester) async {
    await tester.pumpWidget(const MyApp());

    expect(find.text('0 ครั้ง'), findsOneWidget);

    // แตะไอคอน + แล้วสั่งวาดเฟรมใหม่ 1 ครั้ง
    await tester.tap(find.byIcon(Icons.add));
    await tester.pump();

    expect(find.text('1 ครั้ง'), findsOneWidget);
  });

  testWidgets('กดปุ่มรีเซ็ตแล้วตัวนับกลับเป็น 0', (tester) async {
    await tester.pumpWidget(const MyApp());

    await tester.tap(find.byIcon(Icons.add));
    await tester.pump();
    expect(find.text('1 ครั้ง'), findsOneWidget);

    await tester.tap(find.byIcon(Icons.refresh));
    await tester.pump();
    expect(find.text('0 ครั้ง'), findsOneWidget);
  });
}
