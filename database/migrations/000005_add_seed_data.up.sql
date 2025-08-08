USE ordering_db;
    INSERT INTO User (user_id, first_name, last_name, email_id, mobile_number, password, role)
    VALUES 
        (1, 'admin', 'admin', 'admin.neptune@restaurant.com', '0000000000', '$2a$10$A3Zdq5MrRf6uO8R9jEzTseqRnPWNpIAEchoeNJbC2ZqBhMwMLiuhu', 'admin'),
        (2, 'John', 'Doe', 'johndoe@restaurant.com','0000000001', '$2a$10$xadrWAikpXNmFtOUhQo5rOw0ZIXvivGW04Ge1J3MU6Z8567CXH9Xm', 'chef'),
        (3, 'Jane', 'Doe', 'janedoe@restaurant.com','0000000002', '$2a$10$AvGZJkrrusI6hk6rhPWG7OOldZdMepOPNZFJQRfx2qF6QvOoviRAe', 'chef');

    INSERT INTO Categories (category_id, category_name)
    VALUES 
        (1, 'North Indian'),
        (2, 'South Indian'),
        (3, 'Italian'),
        (4, 'Mexican'),
        (5, 'Desserts'),
        (6, 'Beverages');
    
    INSERT INTO Items (item_id, item_name, category_id, price, description, item_image_url, is_veg, spice_level)
    VALUES
        (1, 'Paneer Tikka', 1, 180.00, 'Grilled paneer cubes with rich Indian spices', '/static/images/paneer_tikkaa.webp', TRUE, 3),  
        (2, 'Chicken Biryani', 1, 300.00, 'Classic Hyderabadi biryani', '/static/images/chicken_biryani.webp', FALSE, 2),
        (3, 'Malai Kofta', 1, 250.00, 'Soft paneer koftas in a rich, creamy spiced gravy', '/static/images/malai_kofta.webp', TRUE, 2),
        (4, 'Garlic Naan', 1, 100.00, 'Indian flatbread topped with fresh garlic and baked in a tandoor (4 pcs)', '/static/images/garlic_naan.webp', TRUE, 2),
        (5, 'Plain Dosa', 2, 150.00, 'Crispy South Indian crepe made from fermented rice and lentil batter', '/static/images/plain_dosa.webp', TRUE, 2),
        (6, 'Masala Uttapam', 2, 200.00,'Thick savory pancake topped with spiced onions, tomatoes, and chilies', '/static/images/masala_uttapam.webp', TRUE, 3),
        (7, 'Idli', 2, 110.00, 'Steamed rice cakes, light and fluffy, served with chutneys and sambar', '/static/images/idlie.webp', TRUE, 1),
        (8, 'Spaghetti Aglio Olio', 3, 300.00, 'Classic Italian pasta tossed with garlic, olive oil, and chili flakes', '/static/images/spaghetti_aglio_olio.webp', TRUE, 3),
        (9, 'Chicken Cacciatore', 3, 300.00, 'Rustic Italian stew of chicken braised with tomatoes, herbs, and bell peppers', '/static/images/chicken_cacciatore.webp', FALSE, 3),
        (10, 'Pizza Margherita', 3, 280.00, 'Typical Neapolitan pizza topped with tomato, mozzarella, and fresh basil', '/static/images/pizza_margherita.webp', TRUE, 2),
        (11, 'Tacos al Pastor', 4, 350.00,'Corn tortillas filled with marinated pork, pineapple, and fresh toppings','/static/images/tacos_al_pastor.webp', FALSE, 3),
        (12, 'Chiles Rellenos', 4, 300.00,'Roasted poblano peppers stuffed with cheese or meat, battered and fried','/static/images/chiles_rellenos.webp', TRUE, 2),
        (13, 'Nopales Salad', 4, 250.00,'Refreshing cactus salad with onions, tomatoes, and Mexican spices','/static/images/nopales_salad.webp', TRUE, 2),
        (14, 'Gulab Jamun', 5, 120.00, 'Soft milk dumplings soaked in fragrant rose-cardamom syrup', '/static/images/gulab_jamun.webp', TRUE, -5),
        (15, 'Churros', 5, 200.00, 'Golden fried dough sticks dusted with cinnamon sugar, served with chocolate', '/static/images/churros.webp', TRUE, -4),
        (16, 'New York Cheescake', 5, 250.00, 'Rich and creamy baked cheesecake with a classic graham cracker crust', '/static/images/cheesecake.webp', TRUE, -3),
        (17, 'Tiramisu', 5, 240.00, 'Layered Italian dessert with espresso-soaked ladyfingers and mascarpone cream', '/static/images/tiramisu.webp', TRUE, -2),
        (18, 'Chai', 6, 40.00, 'Spiced Indian tea brewed with milk, cardamom, ginger, and cloves', '/static/images/chai.webp', TRUE, -1),
        (19, 'Mocha Latte', 6, 120.00, 'Espresso blended with steamed milk and rich chocolate syrup', '/static/images/mocha_latte.webp', TRUE, -3),
        (20, 'Mint Mojito', 6, 100.00, 'Refreshing blend of mint, lime, and soda with a splash of sweetness', '/static/images/mint_mojito.webp', TRUE, -2),
        (21, 'Soft Drinks (Canned)', 6, 80.00, 'On Availability', '/static/images/soft_drinks.webp', TRUE, -4); 