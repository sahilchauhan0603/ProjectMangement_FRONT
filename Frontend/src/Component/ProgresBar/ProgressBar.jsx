import { useContext, useState } from "react";
import "./Progress.css";
import Researchinfo from "../Researchinfo/Researchinfo";
import UserInfo from "../UserInfo/UserInfo";
import { UserForm } from "../Usecontext/UserForm/UserForm";

const CheckoutStepper = ({ stepsConfig = [] }) => {
  const [currentStep, setCurrentStep] = useState(1);
  const [isComplete, setIsComplete] = useState(false);
  const {handleSubmit, fetchUploaders} = useContext(UserForm);


  const handleNext = () => {
    setCurrentStep((prevStep) => {
      if (prevStep === stepsConfig.length) {
        setIsComplete(false);
        return prevStep - 1;
      } else {
        return prevStep + 1;
      }
    });
  };

  // Fetch all uploaders
  const handleSubmitAndFetch = async () => {
    await handleSubmit(); // Call the submit function to handle form data submission
    fetchUploaders(); // Fetch uploaders after submission
  };
  const handlePrevious = () => {
    setCurrentStep((prevStep) => (prevStep > 1 ? prevStep - 1 : prevStep));
  };


  const ActiveComponent = stepsConfig[currentStep - 1]?.Component;



  return (
    <>
      <div className="text-white w-screen px-1 md:px-10">
        <div className="flex md:flex-row flex-col mb-10 justify-center ">
          <div className="flex justify-center mb-5">
            {stepsConfig.map((step, index) => {
              return (
                <div
                  key={step.name}
                  className={`step flex w-[30%] md:w-[50%] ${
                    currentStep > index + 1 || isComplete ? "complete" : ""
                  } ${currentStep === index + 1 ? "active" : ""} `}
                >
                  <div className="step-number w-8 h-8 ">
                    {currentStep > index + 1 || isComplete ? (
                      <span>&#10003;</span>
                    ) : (
                      index + 1
                    )}
                  </div>
                  <div className="step-name ml-5 flex flex-col text-white mt-2 text-lg md:text-xl">
                    {step.name}
                  </div>
                  <div className="w-52 text-center mt-4 text-lg ">
                    {currentStep === index + 1 ? (
                      <ActiveComponent />
                    ) : null}
                  </div>
                </div>
              );
            })}
          </div>

          <div className="md:flex-row flex-col ">
            {!isComplete && (
              <div
                className=" md:gap-5 gap-10 flex px-24 md:items-center md:justify-center cursor-pointer"
                onClick={
                  currentStep === stepsConfig.length
                    ? handlePrevious
                    : handleNext
                }
              >
                <span className=" bg-blue-600  px-3 py-2 rounded-md w-22 md:w-24 h-fit md:text-center  ">
                  {currentStep === stepsConfig.length ? "Previous" : "Next"}
                </span>
                <span
                  className={` ${
                    currentStep === stepsConfig.length ? "bg-blue-600" : null
                  }  px-3 py-2 rounded-md w-15 md:w-24 h-fit text-center  `}
                >
                  {currentStep === stepsConfig.length ? (
                    <button
                      type="button"
                      onClick={handleSubmitAndFetch}                      className="text-sm"
                    >
                      Submit
                    </button>
                  ) : null}
                </span>
              </div>
            )}
          </div>
        </div>

        <div>
          {currentStep === 1 ? (
            <UserInfo />
          ) : (
            <Researchinfo />
          )}
        </div>
      </div>
    </>
  );
};

export default CheckoutStepper;
