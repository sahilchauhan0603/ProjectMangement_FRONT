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

  const handleSubmit = async (e) => {
      try {
        const ans = await Axios.post("/api/v1/upload", formData);
        console.log(ans);
        console.log("Uploader created successfully");
        fetchUploaders();
        setFormData({
          firstname: "",
          lastname: "",
          email: "",
          phonenumber: "",
          department: "",
          rollNumber: "",
          batchstart: "",
          batchend: "",
          Collagename:"",
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
        }); // Clear form data
      } catch (error) {
        console.log("Error creating uploader");
      }
  };
  console.log(uploaders)

  // useEffect(() => {
  //   fetchUploaders();
  // }, []);

  return (
    <UserForm.Provider
      value={{ formData, setFormData, handleSubmit, error }}
    >
      {props.children}
    </UserForm.Provider>
  );
};

export default Context;
