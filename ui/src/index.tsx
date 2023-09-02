import React from "react";
import { createRoot } from "react-dom/client";
import * as sw from "./serviceWorker";
import "./index.css";
import "./App.css";
import "bootstrap/dist/css/bootstrap.min.css";
import App from "./App";
import Api from "./Api";
import { Routes, Route, BrowserRouter } from "react-router-dom";
sw.register();
const root = createRoot(document.getElementById("root")!);

root.render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/api" element={<Api />} />
      </Routes>
    </BrowserRouter>
  </React.StrictMode>,
);
