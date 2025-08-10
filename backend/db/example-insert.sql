-- 1. contacts tablosu için örnek insertler
INSERT INTO contacts (name, email, subject, message) VALUES
('Ahmet Yılmaz', 'ahmet.yilmaz@example.com', 'Ürün Hakkında', 'Ürününüz hakkında bazı sorularım var.'),
('Elif Demir', 'elif.demir@example.com', 'Sipariş Durumu', 'Siparişim ne zaman kargolanacak?'),
('Mehmet Çelik', 'mehmet.celik@example.com', NULL, 'Web siteniz çok faydalı, teşekkürler!');

-- 2. home_contents tablosu için örnek insertler
INSERT INTO home_contents (title, description) VALUES
('Hoşgeldiniz', 'Sitemize hoşgeldiniz! En iyi ürünleri bulabilirsiniz.'),
('Kampanyalar', 'Bu ay tüm ürünlerde %20 indirim fırsatını kaçırmayın.'),
('Hakkımızda', 'Firmamız 10 yıldır kaliteli ürünler sunmaktadır.');

-- 3. photos tablosu için örnek insertler
INSERT INTO photos (url, category, title, description) VALUES
('https://example.com/images/photo1.jpg', 'doğa', 'Gün Batımı', 'Harika bir gün batımı manzarası.'),
('https://example.com/images/photo2.jpg', 'şehir', 'İstanbul', 'Boğaz köprüsünün gece görüntüsü.'),
('https://example.com/images/photo3.jpg', 'doğa', 'Orman', 'Yeşilin binbir tonu ile orman yürüyüşü.');

-- 4. products tablosu için örnek insertler
INSERT INTO products (name, description, price, stock, product_url) VALUES
('Bluetooth Kulaklık', 'Kablosuz, uzun pil ömürlü kulaklık.', 299.99, 50, 'https://example.com/products/bluetooth-kulaklik'),
('Laptop Çantası', 'Su geçirmez, kaliteli laptop çantası.', 149.90, 100, 'https://example.com/products/laptop-cantasi'),
('Gaming Mouse', 'Yüksek hassasiyetli oyuncu faresi.', 89.50, 30, 'https://example.com/products/gaming-mouse');
