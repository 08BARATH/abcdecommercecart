import React, { useState, useEffect } from "react";
import API from "./api";

function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [token, setToken] = useState("");
  const [items, setItems] = useState([]);

  const login = async () => {
    try {
      const res = await API.post("/users/login", { username, password });
      setToken(res.data.token);
      fetchItems(res.data.token);
    } catch {
      alert("Invalid username or password");
    }
  };

  const fetchItems = async (tkn) => {
    const res = await API.get("/items", {
      headers: { Authorization: tkn },
    });
    setItems(res.data);
  };

  const addToCart = async (itemId) => {
    await API.post(
      "/carts",
      { itemID: itemId },
      { headers: { Authorization: token } }
    );
    alert("Item added to cart!");
  };

  const checkout = async () => {
    const res = await API.post(
      "/orders",
      {},
      { headers: { Authorization: token } }
    );
    alert("Order placed! Order ID: " + res.data.order_id);
    fetchItems(token); 
  };

  const viewCart = async () => {
    const res = await API.get("/carts", {
      headers: { Authorization: token },
    });
    const items = res.data.map((cart) =>
      cart.Items.map((item) => `Cart ID: ${cart.ID}, Item ID: ${item.ItemID}`)
    );
    alert("Cart Items:\n" + items.flat().join("\n"));
  };

  const viewOrders = async () => {
    const res = await API.get("/orders", {
      headers: { Authorization: token },
    });
    const orders = res.data.map((o) => "Order ID: " + o.ID);
    alert("Your Orders:\n" + orders.join("\n"));
  };

  if (!token) {
    return (
      <div style={{ padding: 50 }}>
        <h2>Login</h2>
        <input
          placeholder="username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <br />
        <input
          placeholder="password"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <br />
        <button onClick={login}>Login</button>
      </div>
    );
  }

  return (
    <div style={{ padding: 50 }}>
      <h2>Item List</h2>
      <button onClick={checkout}>Checkout</button>
      <button onClick={viewCart}>View Cart</button>
      <button onClick={viewOrders}>Order History</button>
      <ul>
        {items.map((item) => (
          <li key={item.ID}>
            {item.Name} - ₹{item.Price}{" "}
            <button onClick={() => addToCart(item.ID)}>Add to Cart</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
