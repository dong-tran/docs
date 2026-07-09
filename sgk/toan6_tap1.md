# TÓM TẮT SÁCH TOÁN 6 - TẬP MỘT
### (Bộ sách "Kết nối tri thức với cuộc sống")

## Mục lục

- [Chương I. Tập hợp các số tự nhiên](#chương-i-tập-hợp-các-số-tự-nhiên)
- [Chương II. Tính chia hết trong tập hợp các số tự nhiên](#chương-ii-tính-chia-hết-trong-tập-hợp-các-số-tự-nhiên)
- [Chương III. Số nguyên](#chương-iii-số-nguyên)
- [Chương IV. Một số hình phẳng trong thực tiễn](#chương-iv-một-số-hình-phẳng-trong-thực-tiễn)
- [Chương V. Tính đối xứng của hình phẳng trong tự nhiên](#chương-v-tính-đối-xứng-của-hình-phẳng-trong-tự-nhiên)

---


*(Tóm tắt bằng lời văn riêng, giữ nguyên khái niệm và cấu trúc bài học của sách Toán 6 - tập 1)*

---

## Bài 1. Tập hợp

**Khái niệm:** Một **tập hợp** (gọi tắt là **tập**) gồm những đối tượng xác định nào đó. Mỗi đối tượng trong tập hợp gọi là một **phần tử** của tập hợp đó.

- Tên tập hợp thường viết bằng chữ **in hoa**: A, B, C, M, N...
- Nếu *x* là phần tử của tập *A*, ta viết $x \in A$ (đọc: *x thuộc A*).
- Nếu *y* không là phần tử của tập *A*, ta viết $y \notin A$ (đọc: *y không thuộc A*).

**Ví dụ minh hoạ:**
> Gọi A là tập hợp các bạn ngồi ở bàn đầu lớp em. Nếu Nam ngồi bàn đầu thì "Nam ∈ A"; nếu Lan ngồi bàn cuối thì "Lan ∉ A".

**Hai cách mô tả một tập hợp:**

| Cách | Nội dung | Ví dụ |
|---|---|---|
| **1. Liệt kê phần tử** | Viết tất cả phần tử trong dấu ngoặc nhọn { }, mỗi phần tử chỉ viết một lần, thứ tự tuỳ ý | P = {0; 1; 2; 3; 4} |
| **2. Nêu dấu hiệu đặc trưng** | Nêu tính chất chung mà mọi phần tử đều thoả mãn | $P = \{x \in \mathbb{N} \mid x < 5\}$ |

**Chú ý quan trọng:**
- Tập hợp số tự nhiên kí hiệu là **$\mathbb{N}$** = $\{0; 1; 2; 3; 4; \ldots\}$
- Tập hợp số tự nhiên khác 0 kí hiệu là **$\mathbb{N}^*$** = $\{1; 2; 3; \ldots\}$
- Viết "$n \in \mathbb{N}$" nghĩa là "n là một số tự nhiên".

**Tập hữu hạn và tập vô hạn:**
- Tập có số phần tử đếm được, có điểm dừng → **tập hữu hạn**. Ví dụ: X = {1; 2} có 2 phần tử.
- Tập có vô số phần tử → **tập vô hạn**. Ví dụ: N là tập vô hạn.

**Giao của hai tập hợp:** Tập hợp gồm các phần tử **chung** của hai tập A và B gọi là giao của A và B, kí hiệu $A \cap B$.
> Minh hoạ: A = {1; 2; 3; 4}, B = {3; 4; 5; 6} $\to A \cap B = \{3; 4\}$.

---

## Bài 2. Cách ghi số tự nhiên

**Hệ thập phân:** Mỗi số tự nhiên được viết bằng một dãy các **chữ số** lấy trong 10 chữ số: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9. Vị trí của mỗi chữ số trong dãy gọi là **hàng** (hàng đơn vị, hàng chục, hàng trăm...).

- Giá trị của một số phụ thuộc vào cả **chữ số** lẫn **vị trí (hàng)** của nó.
- Ví dụ: số 2035 gồm 2 nghìn, 0 trăm, 3 chục, 5 đơn vị:
  $2035 = 2 \times 1000 + 0 \times 100 + 3 \times 10 + 5$

**Số La Mã:** Dùng 7 kí hiệu cơ bản:

| I | V | X | L | C | D | M |
|---|---|---|---|---|---|---|
| 1 | 5 | 10 | 50 | 100 | 500 | 1000 |

> Ví dụ: IV = 4, IX = 9, XII = 12, XX = 20.
> Quy tắc: kí hiệu nhỏ đứng trước kí hiệu lớn thì trừ đi (IV = $5-1 = 4$); kí hiệu nhỏ đứng sau thì cộng thêm (VI = $5+1 = 6$).

---

## Bài 3. Thứ tự trong tập hợp các số tự nhiên

- Trên **tia số**, mỗi số tự nhiên ứng với một điểm gọi là **điểm biểu diễn** số đó; số nhỏ nằm bên trái, số lớn nằm bên phải.
- Với hai số tự nhiên a, b bất kì, luôn xảy ra một trong ba trường hợp: a = b, a < b hoặc a > b.
- **Tính chất bắc cầu:** nếu a < b và b < c thì a < c.
- **Số liền sau** của a là a + 1; **số liền trước** của a là a − 1 (với a > 0).
- Hai số tự nhiên liên tiếp hơn kém nhau 1 đơn vị. Ví dụ: 6 và 7 là hai số tự nhiên liên tiếp.

**Minh hoạ tia số:**
```
0 ---1---2---3---4---5---6---7--->
        ↑ số nhỏ hơn nằm bên trái, số lớn hơn nằm bên phải
```

---

## Bài 4. Phép cộng và phép trừ số tự nhiên

**Phép cộng:** $a + b = c$ (a, b: **số hạng**; c: **tổng**)

Tính chất:
- Giao hoán: $a + b = b + a$
- Kết hợp: $(a + b) + c = a + (b + c)$
- Cộng với 0: $a + 0 = 0 + a = a$

**Phép trừ:** $a - b = c$ chỉ thực hiện được khi **$a \ge b$** (a: **số bị trừ**, b: **số trừ**, c: **hiệu**)

> Ví dụ: $15 - 8 = 7$. Nhưng $5 - 9$ không thực hiện được trong tập số tự nhiên (vì $5 < 9$) — đây chính là lí do sau này cần đến số nguyên âm (sẽ học ở Chương III).

**Mẹo tính nhanh:** dùng tính chất giao hoán, kết hợp để nhóm các số "tròn chục, tròn trăm" lại với nhau.
> Ví dụ: $27 + 98 + 3 = (27 + 3) + 98 = 30 + 98 = 128$

---

## Bài 5. Phép nhân và phép chia số tự nhiên

**Phép nhân:** $a \times b = c$ (a, b: **thừa số**; c: **tích**)

Tính chất:
- Giao hoán: $a \times b = b \times a$
- Kết hợp: $(a \times b) \times c = a \times (b \times c)$
- Nhân với 1: $a \times 1 = a$
- **Phân phối của phép nhân với phép cộng:** $a \times (b + c) = a \times b + a \times c$

**Phép chia:**
- **Chia hết:** $a : b = c$ nếu $a = b \times c$ ($b \neq 0$). Ta nói a chia hết cho b.
- **Chia có dư:** $a = b \times q + r$ ($0 \le r < b$). q gọi là **thương**, r gọi là **số dư**.

> Ví dụ: $17 : 5 \to 17 = 5 \times 3 + 2$, vậy thương là 3, số dư là 2.

---

## Bài 6. Luỹ thừa với số mũ tự nhiên

**Định nghĩa:** Luỹ thừa bậc n của a (n là số tự nhiên, n > 1) là tích của n thừa số bằng nhau, mỗi thừa số bằng a:

**$a^n = a \times a \times \ldots \times a$** (n thừa số a)

- a gọi là **cơ số**, n gọi là **số mũ**.
- $a^2$ gọi là "a **bình phương**"; $a^3$ gọi là "a **lập phương**".

> Ví dụ: $2^3 = 2 \times 2 \times 2 = 8$; $5^2 = 5 \times 5 = 25$

**Nhân, chia hai luỹ thừa cùng cơ số:**
- $a^m \times a^n = a^{m+n}$
- $a^m : a^n = a^{m-n}$ ($m \ge n$, $a \neq 0$)

> Ví dụ: $2^3 \times 2^2 = 2^5 = 32$; $3^5 : 3^2 = 3^3 = 27$

---

## Bài 7. Thứ tự thực hiện các phép tính

**Quy tắc (rất quan trọng, cần nhớ theo thứ tự ưu tiên):**

1. Nếu có **dấu ngoặc**: ( ) → [ ] → { } — tính trong ngoặc trước, từ ngoặc trong cùng ra ngoài.
2. Nếu **không có ngoặc**: thực hiện **luỹ thừa** trước, rồi đến **nhân, chia** (trái sang phải), cuối cùng là **cộng, trừ** (trái sang phải).

> Ví dụ minh hoạ:
> $2 + 3 \times 2^2 = 2 + 3 \times 4 = 2 + 12 = 14$
> $(2 + 3) \times 2^2 = 5 \times 4 = 20$

*(Ghi nhớ dễ hiểu: "Ngoặc trước — Luỹ thừa sau — Nhân chia rồi mới Cộng trừ")*

---

### Luyện tập chung & Bài tập cuối chương I
Ôn lại toàn bộ 7 bài: đọc-viết tập hợp, ghi số tự nhiên, so sánh thứ tự, bốn phép tính, luỹ thừa và thứ tự thực hiện phép tính; dạng bài thường gặp là tính giá trị biểu thức, tìm x trong đẳng thức có phép tính hỗn hợp.

---

## Bài 8. Quan hệ chia hết và tính chất

**Định nghĩa:** Cho hai số tự nhiên a và b (b ≠ 0). Nếu có số tự nhiên q sao cho **$a = b \times q$** thì ta nói **a chia hết cho b**, kí hiệu **$a \vdots b$**. Khi đó b gọi là **ước** của a, còn a gọi là **bội** của b.

> Ví dụ: $15 = 3 \times 5 \to 15 \vdots 3$ và $15 \vdots 5$. Ta nói 3 và 5 là ước của 15; 15 là bội của 3 và của 5.

**Tính chất chia hết của một tổng, một hiệu:**
- Nếu $a \vdots m$ và $b \vdots m$ thì $(a + b) \vdots m$ và $(a - b) \vdots m$ ($a \ge b$)
- Nếu $a \vdots m$ và b không chia hết cho m thì $(a + b)$ không chia hết cho m

> Ví dụ: $12 \vdots 4$ và $8 \vdots 4 \to (12 + 8) = 20 \vdots 4$. Nhưng $12 \vdots 4$ và 7 không chia hết cho 4 → 19 không chia hết cho 4.

**Tính chất chia hết của một tích:** Nếu một thừa số trong tích chia hết cho m thì tích cũng chia hết cho m.
> Ví dụ: $6 \vdots 3 \to 6 \times 7 = 42$ cũng chia hết cho 3.

---

## Bài 9. Dấu hiệu chia hết

| Chia hết cho | Dấu hiệu | Ví dụ |
|---|---|---|
| **2** | Chữ số tận cùng là 0, 2, 4, 6, 8 | $128 \vdots 2$ |
| **5** | Chữ số tận cùng là 0 hoặc 5 | $435 \vdots 5$ |
| **3** | Tổng các chữ số chia hết cho 3 | 123: $1+2+3 = 6 \vdots 3$ |
| **9** | Tổng các chữ số chia hết cho 9 | 189: $1+8+9 = 18 \vdots 9$ |

**Minh hoạ dễ nhớ:** Để kiểm tra 2 và 5 → chỉ nhìn **chữ số cuối cùng**. Để kiểm tra 3 và 9 → **cộng tất cả chữ số lại**.

> Ví dụ tổng hợp: Số 720 — tận cùng là 0 nên chia hết cho cả 2 và 5; tổng chữ số $7+2+0 = 9$ nên chia hết cho cả 3 và 9.

---

## Bài 10. Số nguyên tố

**Định nghĩa:**
- **Số nguyên tố** là số tự nhiên lớn hơn 1, chỉ có đúng hai ước là 1 và chính nó.
- **Hợp số** là số tự nhiên lớn hơn 1, có nhiều hơn hai ước.

> Ví dụ: 2, 3, 5, 7, 11, 13... là số nguyên tố. 4 (ước: 1,2,4), 6 (ước: 1,2,3,6) là hợp số.
> Chú ý: số 0 và số 1 không phải là số nguyên tố cũng không phải là hợp số.

**Cách tìm số nguyên tố nhỏ hơn 100 — Sàng Eratosthenes:** liệt kê các số từ 2 đến 100, lần lượt gạch bỏ các bội của 2 (trừ 2), bội của 3 (trừ 3), bội của 5, bội của 7... Các số còn lại chưa bị gạch chính là số nguyên tố.

**Phân tích một số ra thừa số nguyên tố:** viết một hợp số thành tích các số nguyên tố.

> Ví dụ minh hoạ (sơ đồ cây):
```
        60
       /  \
      2    30
          /  \
         2    15
             /  \
            3    5
```
$\to 60 = 2 \times 2 \times 3 \times 5 = 2^2 \times 3 \times 5$

---

## Bài 11. Ước chung. Ước chung lớn nhất

- **Ước chung** của hai hay nhiều số là ước của tất cả các số đó, kí hiệu ƯC(a, b).
- **Ước chung lớn nhất (ƯCLN)** là số lớn nhất trong tập hợp các ước chung, kí hiệu ƯCLN(a, b).

**Cách tìm ƯCLN bằng phân tích ra thừa số nguyên tố:**
1. Phân tích mỗi số ra thừa số nguyên tố.
2. Chọn ra các thừa số nguyên tố **chung**.
3. Lập tích các thừa số đó, mỗi thừa số lấy với **số mũ nhỏ nhất**.

> Ví dụ: $36 = 2^2 \times 3^2$; $60 = 2^2 \times 3 \times 5 \to \text{ƯCLN}(36, 60) = 2^2 \times 3 = 12$

**Hai số nguyên tố cùng nhau:** là hai số có ƯCLN bằng 1.
> Ví dụ: $\text{ƯCLN}(8, 9) = 1 \to 8$ và 9 là hai số nguyên tố cùng nhau.

**Ứng dụng:** rút gọn phân số về phân số tối giản bằng cách chia cả tử và mẫu cho ƯCLN của chúng.
> Ví dụ: $18/24$, có $\text{ƯCLN}(18,24) = 6 \to 18/24 = 3/4$.

---

## Bài 12. Bội chung. Bội chung nhỏ nhất

- **Bội chung** của hai hay nhiều số là bội của tất cả các số đó, kí hiệu BC(a, b).
- **Bội chung nhỏ nhất (BCNN)** là số nhỏ nhất khác 0 trong tập hợp các bội chung, kí hiệu BCNN(a, b).

**Cách tìm BCNN bằng phân tích ra thừa số nguyên tố:**
1. Phân tích mỗi số ra thừa số nguyên tố.
2. Chọn ra **tất cả** các thừa số nguyên tố chung và riêng.
3. Lập tích các thừa số đó, mỗi thừa số lấy với **số mũ lớn nhất**.

> Ví dụ: $12 = 2^2 \times 3$; $18 = 2 \times 3^2 \to \text{BCNN}(12, 18) = 2^2 \times 3^2 = 36$

**Ứng dụng quan trọng: quy đồng mẫu số các phân số**, dùng BCNN của các mẫu số làm mẫu số chung.
> Ví dụ: quy đồng $1/12$ và $5/18$, mẫu chung là $\text{BCNN}(12,18) = 36 \to 1/12 = 3/36$; $5/18 = 10/36$.

---

### Luyện tập chung & Bài tập cuối chương II
Ôn tập quan hệ chia hết, dấu hiệu chia hết cho 2/3/5/9, số nguyên tố - hợp số, ƯCLN, BCNN. Dạng bài phổ biến: bài toán chia đều số lượng đồ vật (dùng ƯCLN) hoặc bài toán về các sự kiện lặp lại theo chu kì (dùng BCNN — ví dụ "hai bạn cùng trực nhật vào ngày nào").

---

## Bài 13. Tập hợp các số nguyên

**Vì sao cần số nguyên?** Trong tập số tự nhiên, phép trừ a − b chỉ thực hiện được khi a ≥ b. Để phép trừ luôn thực hiện được, và để biểu diễn các đại lượng "ngược chiều nhau" (nhiệt độ dưới 0°C, độ sâu dưới mực nước biển, nợ tiền...), người ta đưa thêm các **số nguyên âm**.

- **Số nguyên dương:** 1, 2, 3,... (chính là các số tự nhiên khác 0, có thể viết thêm dấu + phía trước: +1, +2...)
- **Số nguyên âm:** $-1, -2, -3, \ldots$
- **Z** = $\{\ldots; -3; -2; -1; 0; 1; 2; 3; \ldots\}$

**Trục số:** giống tia số nhưng kéo dài về bên trái điểm 0.
```
$<---(-3)--(-2)--(-1)---0---(1)---(2)--(3)--->$
   chiều âm            gốc         chiều dương
```

**Số đối:** hai số nguyên cách đều điểm 0 và nằm ở hai phía của trục số gọi là hai số đối nhau.
> Ví dụ: số đối của 5 là −5; số đối của −7 là 7.

**So sánh hai số nguyên:** trên trục số, số nào ở bên trái thì nhỏ hơn.
> Ví dụ: $-5 < -2 < 0 < 3$

---

## Bài 14. Phép cộng và phép trừ số nguyên

**Cộng hai số nguyên cùng dấu:** cộng hai giá trị tuyệt đối rồi giữ nguyên dấu chung.
> Ví dụ: $(-3) + (-5) = -(3+5) = -8$

**Cộng hai số nguyên khác dấu:** lấy số có giá trị tuyệt đối lớn hơn trừ số có giá trị tuyệt đối nhỏ hơn, rồi đặt dấu của số có giá trị tuyệt đối lớn hơn.
> Ví dụ: $(-7) + 4 = -(7-4) = -3$; $9 + (-3) = +(9-3) = 6$

**Tính chất phép cộng số nguyên:** giao hoán, kết hợp, cộng với 0, và **cộng với số đối bằng 0**: $a + (-a) = 0$

**Phép trừ:** $a - b = a + (-b)$ (trừ một số bằng cộng với số đối của nó)
> Ví dụ: $5 - 8 = 5 + (-8) = -3$

**Minh hoạ trực quan (nhiệt độ):** Buổi sáng −2°C, đến trưa tăng thêm 6°C → nhiệt độ trưa là (−2) + 6 = 4°C.

---

## Bài 15. Quy tắc dấu ngoặc

- Khi bỏ dấu ngoặc có **dấu "+" đằng trước**: giữ nguyên dấu các số hạng trong ngoặc.
- Khi bỏ dấu ngoặc có **dấu "−" đằng trước**: đổi dấu tất cả các số hạng trong ngoặc.

> Ví dụ:
> $5 + (3 - 7) = 5 + 3 - 7 = 1$
> $5 - (3 - 7) = 5 - 3 + 7 = 9$

*(Mẹo nhớ: dấu "+" trước ngoặc giống như "mở khoá tự do", dấu "−" trước ngoặc là "đảo ngược mọi thứ bên trong")*

---

## Bài 16. Phép nhân số nguyên

**Quy tắc dấu:**
- **Cùng dấu** → tích **dương**: $(+) \times (+) = +$; $(-) \times (-) = +$
- **Khác dấu** → tích **âm**: $(+) \times (-) = -$; $(-) \times (+) = -$

> Ví dụ: $(-4) \times (-5) = 20$; $(-4) \times 5 = -20$

Tính chất: giao hoán, kết hợp, nhân với 1, phân phối với phép cộng — giống hệt số tự nhiên nhưng cần chú ý thêm quy tắc dấu.

---

## Bài 17. Phép chia hết. Ước và bội của một số nguyên

**Định nghĩa:** Cho $a, b \in \mathbb{Z}$ ($b \neq 0$). Nếu có $q \in \mathbb{Z}$ sao cho $a = b \times q$ thì ta nói a chia hết cho b, kí hiệu $a \vdots b$. Khi đó b là **ước** của a, a là **bội** của b.

> Ví dụ: $(-12) : 4 = -3 \to -12 \vdots 4$; 4 là ước của -12, -12 là bội của 4.

**Chú ý:** Nếu a là bội của b thì −a cũng là bội của b; nếu b là ước của a thì −b cũng là ước của a.

---

### Luyện tập chung & Bài tập cuối chương III
Ôn lại tập hợp số nguyên, 4 phép tính cộng-trừ-nhân-chia trên số nguyên, quy tắc dấu ngoặc, ước-bội của số nguyên. Dạng bài phổ biến: tính giá trị biểu thức có dấu ngoặc lồng nhau, bài toán thực tế về nhiệt độ, độ cao/độ sâu, tài khoản (nợ - có).

---

## Bài 18. Hình tam giác đều. Hình vuông. Hình lục giác đều

| Hình | Đặc điểm nhận biết | Minh hoạ |
|---|---|---|
| **Tam giác đều** | 3 cạnh bằng nhau, 3 góc bằng nhau (mỗi góc 60°) | Biển báo giao thông hình tam giác, giá đỡ, mái nhà |
| **Hình vuông** | 4 cạnh bằng nhau, 4 góc vuông, 2 đường chéo bằng nhau và vuông góc | Viên gạch lát nền, ô bàn cờ |
| **Hình lục giác đều** | 6 cạnh bằng nhau, 6 góc bằng nhau | Tổ ong, viên bi-a, mặt cắt bút chì |

**Các yếu tố cơ bản:** cạnh, đỉnh, góc — mỗi hình đều được xác định bởi các yếu tố này bằng nhau theo quy tắc riêng.

> Ví dụ trực quan: gấp một tờ giấy hình vuông theo đường chéo sẽ được hai tam giác vuông cân bằng nhau — cho thấy rõ tính chất "hai đường chéo bằng nhau" của hình vuông.

---

## Bài 19. Hình chữ nhật. Hình thoi. Hình bình hành. Hình thang cân

| Hình | Cạnh | Góc | Đường chéo |
|---|---|---|---|
| **Hình chữ nhật** | 2 cặp cạnh đối song song và bằng nhau | 4 góc vuông | 2 đường chéo bằng nhau, cắt nhau tại trung điểm mỗi đường |
| **Hình thoi** | 4 cạnh bằng nhau | góc đối bằng nhau | 2 đường chéo vuông góc, cắt nhau tại trung điểm mỗi đường |
| **Hình bình hành** | 2 cặp cạnh đối song song và bằng nhau | góc đối bằng nhau | 2 đường chéo cắt nhau tại trung điểm mỗi đường |
| **Hình thang cân** | 1 cặp cạnh đối song song (2 đáy), 2 cạnh bên bằng nhau | 2 góc kề một đáy bằng nhau | 2 đường chéo bằng nhau |

**Cách nhớ dễ hiểu:**
- Hình chữ nhật = hình bình hành có 1 góc vuông.
- Hình thoi = hình bình hành có 2 cạnh kề bằng nhau.
- Hình vuông (đã học ở Bài 18) = vừa là hình chữ nhật, vừa là hình thoi.

> Ví dụ thực tế: cửa sổ, màn hình ti vi → hình chữ nhật; viên kim cương nhìn từ mặt trước, cánh diều → hình thoi; mặt bàn xếp, thanh ray cửa trượt → hình bình hành; mặt cắt của cái đê, chân bàn xếp → hình thang cân.

---

## Bài 20. Chu vi và diện tích của một số tứ giác đã học

| Hình | Chu vi | Diện tích |
|---|---|---|
| Hình chữ nhật (2 cạnh a, b) | $C = 2(a + b)$ | $S = a \times b$ |
| Hình vuông (cạnh a) | $C = 4a$ | $S = a^2$ |
| Hình thoi (cạnh a; hai đường chéo $d_1, d_2$) | $C = 4a$ | $S = (d_1 \times d_2) : 2$ |
| Hình bình hành (cạnh a; chiều cao h tương ứng) | $C = 2(a + b)$ | $S = a \times h$ |
| Hình thang (hai đáy a, b; chiều cao h) | C = tổng 4 cạnh | $S = (a + b) \times h : 2$ |

> Ví dụ minh hoạ: Một mảnh vườn hình chữ nhật có chiều dài 12 m, chiều rộng 8 m.
> - Chu vi: $2 \times (12 + 8) = 40$ (m)
> - Diện tích: $12 \times 8 = 96$ (m²)

---

### Luyện tập chung & Bài tập cuối chương IV
Ôn nhận diện hình, tính chất cạnh - góc - đường chéo, công thức chu vi và diện tích. Dạng bài phổ biến: tính diện tích mảnh đất/mảnh vườn, cắt ghép hình để tính diện tích hình phức tạp bằng cách chia thành các hình cơ bản đã học.

---

## Bài 21. Hình có trục đối xứng

**Khái niệm:** Một hình có **trục đối xứng** nếu có một đường thẳng chia hình đó thành hai phần mà khi "gấp" hình theo đường thẳng ấy, hai phần trùng khít lên nhau. Đường thẳng đó gọi là **trục đối xứng** của hình.

**Minh hoạ dễ hình dung:** giống như khi gấp đôi một tờ giấy cắt hình trái tim — nếp gấp chính là trục đối xứng, và hai nửa trái tim khớp hoàn toàn với nhau.

| Hình | Số trục đối xứng |
|---|---|
| Hình tam giác đều | 3 trục |
| Hình vuông | 4 trục |
| Hình chữ nhật | 2 trục |
| Hình thoi | 2 trục |
| Hình tròn | vô số trục (mọi đường kính đều là trục đối xứng) |

> Ví dụ thực tế: cánh quạt trần, chiếc lá, mặt người (gần đối xứng), chữ cái in hoa như A, H, M, T, U, V, W, X, Y đều có trục đối xứng.

---

## Bài 22. Hình có tâm đối xứng

**Khái niệm:** Một hình có **tâm đối xứng** nếu có một điểm O sao cho khi quay hình đó **$180°$** quanh điểm O thì hình thu được trùng khít với hình ban đầu. Điểm O gọi là **tâm đối xứng**.

**Minh hoạ dễ hình dung:** giống như con số 8 hoặc chữ cái "S" khi xoay ngược 180° vẫn giữ nguyên hình dạng.

| Hình | Tâm đối xứng |
|---|---|
| Hình bình hành | giao điểm hai đường chéo |
| Hình chữ nhật | giao điểm hai đường chéo |
| Hình thoi | giao điểm hai đường chéo |
| Hình vuông | giao điểm hai đường chéo |
| Hình tròn | tâm của đường tròn |
| Hình lục giác đều | giao điểm các đường chéo chính |

> Ví dụ thực tế: cánh quạt điện 2 cánh (đối xứng qua trục quay), chong chóng, viên xúc xắc nhìn theo một mặt.

**Phân biệt hai loại đối xứng (dễ nhầm):**
- **Trục đối xứng** → thao tác **gấp (lật)** qua một đường thẳng.
- **Tâm đối xứng** → thao tác **xoay 180°** quanh một điểm.
- Một hình có thể có cả hai loại (như hình vuông, hình tròn), chỉ một loại, hoặc không có loại nào.

---

### Luyện tập chung & Bài tập cuối chương V
Ôn cách nhận biết và tìm trục đối xứng, tâm đối xứng của các hình đã học ở Chương IV; liên hệ với các hình ảnh đối xứng có trong tự nhiên và đời sống (hoa, lá, kiến trúc, logo...).

---

## Hoạt động thực hành trải nghiệm (cuối sách)
- **Tấm thiệp và phòng học của em:** vận dụng kiến thức về hình phẳng và tính đối xứng để thiết kế, trang trí sản phẩm thực tế.
- **Vẽ hình đơn giản với phần mềm GeoGebra:** hướng dẫn dùng phần mềm để vẽ và kiểm tra tính chất của các hình đã học.
- **Sử dụng máy tính cầm tay:** hướng dẫn các phép tính cơ bản, luỹ thừa, ước - bội bằng máy tính cầm tay.
