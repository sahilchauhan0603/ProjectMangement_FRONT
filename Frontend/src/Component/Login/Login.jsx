import React, { useEffect, useState } from "react";
import Image from "../../assets/image.png";
import Microsoft from "../../assets/Microsoft_logo.svg.png";
import "./Login.css";
import { FaEye } from "react-icons/fa6";
import { FaEyeSlash } from "react-icons/fa6";
import { Link } from "react-router-dom";
import Navbar from "../Navbar";

const Login = () => {
  const [showPassword, setShowPassword] = useState(false);

  return (
    <div className=" w-screen overflow-hidden">
      <Navbar/>
      <div className="login-main">
        <div className="login-left w-52">
          <img src={Image} alt="" className="rounded-lg" />
        </div>
        <div className="login-right">
          <div className="login-right-container">
            <div className="login-center ">
              <h1 className="md:text-[3rem] text-[3rem] text-white mb-5">Welcome back!</h1>
              <h1 className=" text-md md:text-lg text-zinc-500 mb-10">
                Please enter your details
              </h1>
              <form>
                <input
                  type="email"
                  placeholder="Email"
                  className="text-black text-xl input"
                />
                <div className="pass-input-div text-black">
                  <input
                    type={showPassword ? "text" : "password"}
                    placeholder="Password"
                    className="input"
                  />
                  {showPassword ? (
                    <FaEyeSlash
                      onClick={() => {
                        setShowPassword(!showPassword);
                      }}
                    />
                  ) : (
                    <FaEye
                      onClick={() => {
                        setShowPassword(!showPassword);
                      }}
                    />
                  )}
                </div>

                <div className="login-center-options">
                  <Link to="/" className="text-white text-sm">
                    Forgot password?
                  </Link>
                </div>
                <div className="login-center-buttons">
                  <Link to="/">
                    <button type="button">Log In</button>
                  </Link>
                  <button type="button">
                    <img src={Microsoft} alt=""  className="w-5"/>
                    Log In with Microsoft
                  </button>
                </div>
              </form>
            </div>

            <p className="login-bottom-p text-sm text-white">
              Don't have an account? <Link to="/signup">Sign Up</Link>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;
