import React, { useContext, useEffect, useRef, useState } from "react";
import { UserForm } from "../Usecontext/UserForm/UserForm";

const Researchinfo = () => {
  const { formData, setFormData, handleSubmit } = useContext(UserForm);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  return (
    <div className="w-[95vw] md:px-10 ml-2 md:ml-0">
      <form onSubmit={handleSubmit}>
        <div className="max-w-screen-xl h-auto p-8 flex flex-col bg-[#403e3e41] rounded-xl  items-center  mx-auto text-black">
          <div className=" w-full flex md:flex-row flex-col justify-between md:px-20 ">
            <div className="mb-5">
              <label className="flex text-md font-medium text-white">
                Faculty Name
              </label>
              <input
                type="text"
                name="faculty"
                value={formData.faculty}
                onChange={handleChange}
                placeholder="Faculty Name..."
                required
                className="mt-1 px-3 py-2 outline-none w-full  md:w-[16vw] rounded-md"
              />
            </div>
            <div className="mb-5">
              <label className="flex text-md font-medium text-white">
                Topic
              </label>
              <input
                type="text"
                name="Research"
                value={formData.Research}
                onChange={handleChange}
                placeholder="Research..."
                required
                className="mt-1 w-full  md:w-[16vw] outline-none rounded-md px-3 py-2 "
              />
            </div>
          </div>

          <div className="w-full  md:px-20 mb-5">
            <label className="flex text-md font-medium text-white">
              Project Url
            </label>
            <input
              type="url"
              name="projecturl"
              className="md:w-full mt-1 px-3 py-2 outline-none w-full   rounded-md"
              placeholder="Projecturl"
            />
          </div>

          <div className="w-full  md:px-20 mb-10 ">
            <label className="flex text-md font-medium text-white">Work</label>
            <div className=" flex md:flex-row flex-col justify-between ">
              <div className="mt-5">
                <h1 className="text-white">Start</h1>
                <input
                  type="Date"
                  name="researchstart"
                  value={formData.researchstart}
                  onChange={handleChange}
                  required
                  className="mt-1 w-full md:w-[10vw] outline-none rounded-md px-3 py-2 "
                />
              </div>
              <div className="mt-5">
                <h1 className="text-white">End</h1>
                <input
                  type="Date"
                  name="researchend"
                  value={formData.researchend}
                  onChange={handleChange}
                  required
                  className="mt-1 w-full md:w-[10vw] outline-none rounded-md px-3 py-2 "
                />
              </div>
            </div>
          </div>

          <div className="w-full  md:px-20 mb-5 ">
            <label className="flex text-md font-medium text-white">
              About Research
            </label>
            <textarea
              name="aboutResearch"
              value={formData.aboutResearch}
              onChange={handleChange}
              rows="7"
              className="w-full outline-none px-5 py-4 overflow-scroll  text-black"
            ></textarea>
          </div>
          <div className="w-full md:px-20 ">
            <label className="flex text-md font-medium text-white">
              About Approach
            </label>
            <textarea
              name="aboutapproach"
              value={formData.aboutapproach}
              onChange={handleChange}
              rows="7"
              className="w-full  outline-none px-5 py-4 overflow-scroll text-black"
            ></textarea>
          </div>
        </div>
      </form>
    </div>
  );
};

export default Researchinfo;
