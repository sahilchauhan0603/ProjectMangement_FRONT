import React, { useState } from "react";
import Navbar from "./Navbar";
import CheckoutStepper from "./ProgresBar/ProgressBar";

const CHECKOUT_STEPS = [
  {
    name: "Student Info",
    Component: () => <div></div>,
  },
  {
    name: "Research  Info",
    Component: () => <div></div>,
  },
];

const AddResearch = () => {
  return (
    <div className="w-screen h-[130vh]  overflow-x-hidden md:overflow-hidden bg-zinc-800">
      <Navbar />
      <div className=" flex justify-center mt-10">
        <CheckoutStepper stepsConfig={CHECKOUT_STEPS} />
      </div>
    </div>
  );
};

export default AddResearch;

