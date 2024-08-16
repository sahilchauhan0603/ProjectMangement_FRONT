import React from "react";
import Link from "./Routers/Links";
import axios from "axios";

axios.defaults.baseURL = "http://localhost:8000/api/v1";
axios.defaults.withCredentials = true;
axios.defaults.headers = {
  "Content-Type": "application/json",
  // "Authorization": `Bearer ${localStorage.getItem('token')}`,
};

const App = () => {
  return (
    <div className="w-screen h-screen">
      <Link />
    </div>
  );
};

export default App;
