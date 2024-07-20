import React, { useEffect, useState } from "react";
import Image from "../../assets/image.png";
import Microsoft from "../../assets/Microsoft_logo.svg.png";
import Google from "../../assets/google.svg";
import { FaEye } from "react-icons/fa6";
import { FaEyeSlash } from "react-icons/fa6";
import { Link } from "react-router-dom";
import "./Signup.css";
const Signup = () => {
  const [showPassword, setShowPassword] = useState(false);

  const [formData, setFormData] = useState({
    firstName: "",
    lastName: "",
    email: "",
    password: "",
    confirmPassword: "",
  });

  const [errors, setErrors] = useState({});

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const validateForm = () => {
    let formErrors = {};
    if (!formData.firstName) formErrors.firstName = "First name is required";
    if (!formData.lastName) formErrors.lastName = "Last name is required";
    if (!formData.email) formErrors.email = "Email is required";
    else if (!/\S+@\S+\.\S+/.test(formData.email))
      formErrors.email = "Email is invalid";
    if (!formData.password) formErrors.password = "Password is required";
    else if (formData.password.length < 6)
      formErrors.password = "Password must be at least 6 characters";
    if (formData.password !== formData.confirmPassword)
      formErrors.confirmPassword = "Passwords do not match";

    setErrors(formErrors);
    return Object.keys(formErrors).length === 0;
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (validateForm()) {
      console.log("Form data submitted:", formData);
      // Handle form submission, e.g., send data to server
    }
  };

  const handleMicrosoftSignup = () => {
    // Handle Microsoft registration logic
    console.log("Register with Microsoft");
  };

  return (
    <div className="rounded-lg w-screen overflow-hidden">
      <div className="login-main">
        <div className="login-left w-52">
          <img src={Image} alt="" />
        </div>
        <div className="login-right">
          <div className="login-right-container">
            <div className="login-center text-white">
              <h1 className="text-[2rem] mb-3">
                Already have an account? <Link to="/login">Log in</Link>
              </h1>
              <div className="bg-red-100">
                <button className="">
                  <img src={Google} className="w-20" /> Sign up with Google
                </button>
                <button>
                  <img src={Microsoft} width="16px" /> Sign up with Microsoft
                </button>
              </div>
              <span>or</span>
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

                <div className="login-center-buttons">
                  <Link to="/">
                    <button type="button">Create Account</button>
                  </Link>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Signup;

{
  /* <h2>Get started.</h2>
<p>
  Already have an account? <Link href="#">Log in</Link>
</p>
<div class="sign-up-buttons">
  <button id="sign-up">
    <img src={Google} className="w-20" /> Sign up with Google
  </button>
  <button id="sign-up-facebook">
    <img src={Microsoft} width="16px" /> Sign up with
    Microsoft
  </button>
</div>
<p class="socials-divider">
  <span>or</span>
</p> */
}

{
  /* <div className="mb-4">
                <label
                  className="flex text-white text-xl font-bold mb-2"
                  htmlFor="firstName"
                >
                  First Name
                </label>
                <input
                  type="text"
                  id="firstName"
                  name="firstName"
                  value={formData.firstName}
                  onChange={handleChange}
                  className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                  placeholder="First Name"
                />
                {errors.firstName && (
                  <p className="text-red-500 text-xs italic">
                    {errors.firstName}
                  </p>
                )}
              </div>

              <div className="mb-4">
                <label
                  className="block text-gray-700 text-xl font-bold mb-2"
                  htmlFor="lastName"
                >
                  Last Name
                </label>
                <input
                  type="text"
                  id="lastName"
                  name="lastName"
                  value={formData.lastName}
                  onChange={handleChange}
                  className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                  placeholder="Last Name"
                />
                {errors.lastName && (
                  <p className="text-red-500 text-xs italic">
                    {errors.lastName}
                  </p>
                )}
              </div> */
}

// <div className="mb-6">
// <label
//   className="block text-gray-700 text-xl font-bold mb-2"
//   htmlFor="confirmPassword"
// >
//   Confirm Password
// </label>
// <input
//   type="password"
//   id="confirmPassword"
//   name="confirmPassword"
//   value={formData.confirmPassword}
//   onChange={handleChange}
//   className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
//   placeholder="Confirm Your Password"
// />
// {errors.confirmPassword && (
//   <p className="text-red-500 text-xs italic">
//     {errors.confirmPassword}
//   </p>
// )}
// </div>
