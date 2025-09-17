
/**
 * API Client — stubs only. No mock/phantom data here.
 * Replace URL paths with your real backend endpoints when they are ready.
 */
window.API = {
  // Catalog & search
  searchProducts: async (query, filters, sort, page) => {
    throw new Error("searchProducts: требуется бэкенд. Подключите эндпоинт /api/products/search");
  },
  getProductById: async (id) => {
    throw new Error("getProductById: требуется бэкенд. Подключите эндпоинт /api/products/:id");
  },
  // Auth
  login: async (email, password) => {
    throw new Error("login: требуется бэкенд. Подключите эндпоинт /api/auth/login");
  },
  signup: async (payload) => {
    throw new Error("signup: требуется бэкенд. Подключите эндпоинт /api/auth/signup");
  },
  // Cart
  getCart: async () => {
    throw new Error("getCart: требуется бэкенд. Подключите эндпоинт /api/cart");
  },
  addToCart: async (productId, qty = 1) => {
    throw new Error("addToCart: требуется бэкенд. Подключите эндпоинт /api/cart/items");
  },
  updateCartItem: async (itemId, qty) => {
    throw new Error("updateCartItem: требуется бэкенд. Подключите эндпоинт /api/cart/items/:id");
  },
  removeCartItem: async (itemId) => {
    throw new Error("removeCartItem: требуется бэкенд. Подключите эндпоинт DELETE /api/cart/items/:id");
  },
  // Orders
  createOrder: async (payload) => {
    throw new Error("createOrder: требуется бэкенд. Подключите эндпоинт POST /api/orders");
  },
  getOrders: async () => {
    throw new Error("getOrders: требуется бэкенд. Подключите эндпоинт GET /api/orders");
  },
  // Favorites
  getFavorites: async () => {
    throw new Error("getFavorites: требуется бэкенд. Подключите эндпоинт GET /api/favorites");
  },
  toggleFavorite: async (productId) => {
    throw new Error("toggleFavorite: требуется бэкенд. Подключите эндпоинт POST /api/favorites/:id");
  },
  // Compare
  getCompare: async () => {
    throw new Error("getCompare: требуется бэкенд. Подключите эндпоинт GET /api/compare");
  },
  toggleCompare: async (productId) => {
    throw new Error("toggleCompare: требуется бэкенд. Подключите эндпоинт POST /api/compare/:id");
  },
};
