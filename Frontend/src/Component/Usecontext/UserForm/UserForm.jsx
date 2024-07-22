import { createContext, useEffect, useState } from "react";
import Axios from "../../../Data/Axios";

export const UserForm = createContext();

const Context = (props) => {
  const [uploaders, setUploaders] = useState([]);

  const [formData, setFormData] = useState({
    firstname: "",
    lastname: "",
    email: "",
    phonenumber: "",
    department: "",
    rollNumber: "",
    batchstart: "",
    batchend: "",
    github: "",
    linkedin: "",
    faculty: "",
    student: "",
    research: "",
    projecturl: "",
    researchstart: "",
    researchend: "",
    aboutResearch: "",
    aboutapproach: "",
  });
  const [error, seterror] = useState({});

  const validateForm = () => {
    let formErrors = {};
    if (!formData.firstname) formErrors.firstname = "First name is required";
    if (!formData.lastname) formErrors.lastname = "Last name is required";
    if (!formData.email) formErrors.email = "Email is required";
    else if (!/\S+@\S+\.\S+/.test(formData.email))
      formErrors.email = "Email is invalid";
    // Add more validation rules as needed

    seterror(formErrors);
    return Object.keys(formErrors).length === 0;
  };

  // const fetchUploaders = async () => {
  //   try {
  //     const response = await Axios.get("api/v1/upload");
  //     console.log(response);
  //     setUploaders(response.data);
  //   } catch (error) {
  //     console.log("Error fetching uploaders");
  //   }
  // };
  const handleSubmit = async () => {

    try {
      const response = await fetch("http://localhost:8000/api/v1/upload", {
        method: "POST",
        body: formData,
        headers: {
          Authorization:
            "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE2MjI2MTksInN1YiI6IlhRalFWeDREX0dMRDZYWFJpb01lMVBKRTNQYy1GV1N5NkRZaS1XdTFPVEEifQ.UVDEtwd5PvX94Grn_EiaENdtCF9TghadZHf996DwRsU", // Replace with your actual token
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error("Error uploading file", error);
    }
  };

  // useEffect(() => {
  //   fetchUploaders();
  // }, []);

  return (
    <UserForm.Provider value={{ formData, setFormData, handleSubmit, error }}>
      {props.children}
    </UserForm.Provider>
  );
};

export default Context;
