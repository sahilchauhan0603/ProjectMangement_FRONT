import React, { useContext, useEffect, useRef, useState } from "react";
import { UserForm } from "../Usecontext/UserForm/UserForm";

const UserInfo = () => {
  const {formData, setFormData , handleSubmit , error } = useContext(UserForm);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };


 
  return (
    <div className="w-[95vw]  ml-2 md:ml-0 md:px-10 ">
      <form onSubmit={handleSubmit}>
        <div className=" md:max-w-screen-lg md:mx-auto  p-8 flex flex-col bg-[#605e5e2e] rounded-xl  md:items-center h-auto  text-black">
          <div className="flex flex-col md:flex-row gap-[10vw] mb-10">
            <div>
              <label className="flex text-md font-medium text-white">
                First Name
              </label>
              <input
                type="text"
                name="firstname"
                value={formData.firstname}
                onChange={handleChange}
                placeholder="Username..."
                className="mt-1 px-3 py-2 outline-none w-full  md:w-[16vw] rounded-md"
              />
              {error.firstName && (
                <p className="text-red-500 text-xs italic">{error.firstName}</p>
              )}
            </div>
            <div>
              <label className="flex text-md font-medium text-white">
                Last Name
              </label>
              <input
                type="text"
                name="Lastname"
                value={formData.Lastname}
                onChange={handleChange}
                placeholder="Lastname..."
                className="mt-1 w-full  md:w-[16vw] outline-none rounded-md px-3 py-2 "
              />
              {error.Lastname && (
                <p className="text-red-500 text-xs italic">{error.Lastname}</p>
              )}
            </div>
          </div>
          <div className=" flex md:flex-row flex-col gap-[10vw] mb-10">
            <div>
              <label className="flex text-md font-medium text-white">
                Email Name
              </label>
              <input
                type="email"
                name="email"
                value={formData.email}
                onChange={handleChange}
                placeholder="Email..."
                className="mt-1 w-full md:w-[16vw]  rounded-md px-3 py-2 outline-none"
              />
              {error.email && (
                <p className="text-red-500 text-xs italic">{error.email}</p>
              )}
            </div>
            <div>
              <label className="flex  text-md font-medium text-white">
                Phone Number
              </label>
              <input
                type="text"
                name="phonenumber"
                value={formData.phonenumber}
                onChange={handleChange}
                placeholder="Phone no."
                className="mt-1  w-full md:w-[16vw] outline-none rounded-md px-3 py-2 "
              />
            </div>
            {error.phonenumber && (
              <p className="text-red-500 text-xs italic">{error.phonenumber}</p>
            )}
          </div>

          <div className="flex md:flex-row flex-col gap-[10vw] mb-10">
            <div>
              <label className=" flex text-md font-medium text-white">
                Department
              </label>
              <input
                type="text"
                name="department"
                value={formData.department}
                onChange={handleChange}
                placeholder="Department..."
                className="mt-1  w-full md:w-[16vw] outline-none rounded-md px-3 py-2 "
              />
              {error.department && (
                <p className="text-red-500 text-xs italic">
                  {error.department}
                </p>
              )}
            </div>
            <div>
              <label className="flex text-md font-medium text-white">
                Roll Number
              </label>
              <input
                type="text"
                name="rollNumber"
                value={formData.rollNumber}
                onChange={handleChange}
                placeholder="Roll No."
                className="mt-1  w-full md:w-[16vw] outline-none rounded-md px-3 py-2 "
              />
              {error.rollNumber && (
                <p className="text-red-500 text-xs italic">
                  {error.rollNumber}
                </p>
              )}
            </div>
          </div>
          <div className="w-full  md:px-20 mb-10 ">
            <label className="flex text-md font-medium text-white">
              Year of Graduation
            </label>
            <div className="flex md:flex-row flex-col   md:justify-between ">
              <div className="mt-5">
                <h1 className="text-white">Start</h1>
                <input
                  type="Date"
                  name="batchstart"
                  value={formData.batchstart}
                  onChange={handleChange}
                  className="mt-1 w-full md:w-[10vw] outline-none rounded-md px-3 py-2 "
                />
                {error.batchstart && (
                  <p className="text-red-500 text-xs italic">
                    {error.batchstart}
                  </p>
                )}
              </div>
              <div className="mt-5">
                <h1 className="text-white">End</h1>
                <input
                  type="Date"
                  name="batchend"
                  value={formData.batchend}
                  onChange={handleChange}
                  className="mt-1 w-full md:w-[10vw] outline-none rounded-md px-3 py-2 "
                />
                {error.batchend && (
                  <p className="text-red-500 text-xs italic">
                    {error.batchend}
                  </p>
                )}
              </div>
            </div>
          </div>

          <div className=" flex flex-col md:flex-row gap-[10vw]">
            <div>
              <label className="flex text-md font-medium text-white">
                Github
              </label>
              <input
                type="url"
                name="github"
                value={formData.github}
                onChange={handleChange}
                placeholder="Link..."
                className="mt-1  w-full md:w-[16vw] outline-none  rounded-md px-3 py-2 "
              />
              {error.github && (
                <p className="text-red-500 text-xs italic">{error.github}</p>
              )}
            </div>
            <div>
              <label className="flex text-md font-medium text-white">
                Linkedin
              </label>
              <input
                type="url"
                name="linkedin"
                value={formData.linkedin}
                onChange={handleChange}
                placeholder="Link..."
                className="mt-1  w-full md:w-[16vw] outline-none rounded-md px-3 py-2 "
              />
            </div>
            {error.linkedin && (
              <p className="text-red-500 text-xs italic">{error.linkedin}</p>
            )}
          </div>

          <div className="flex justify-between fle-col w-full px-20 mt-16">
            <button
              type="reset"
              className="w-fit h-fit px-4 py-3 cursor-pointer mt-8 text-white font-bold rounded-md bg-gray-400"
              onClick={() => setFormData("")}
            >
              Reset
            </button>
          </div>
        </div>
      </form>
    </div>
  );
};

export default UserInfo;
